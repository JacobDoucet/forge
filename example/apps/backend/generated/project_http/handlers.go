package project_http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/coded_error"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/permissions"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/project"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/project_api"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/utils"
	"github.com/rs/zerolog/log"
	"io"
	"net/http"
)

type GetMetadataOnSearchHook func(ctx context.Context, actor permissions.Actor, queryResult project_api.HTTPQueryResult) (map[string]any, error)
type GetMetadataOnCreateHook func(ctx context.Context, actor permissions.Actor, mutationResult project_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnUpdateHook func(ctx context.Context, actor permissions.Actor, mutationResult project_api.HTTPMutationResult) (map[string]any, error)
type GetMetadataOnDeleteHook func(ctx context.Context, actor permissions.Actor, deleteResult project_api.HTTPDeleteResult) (map[string]any, error)

type MetadataHooks struct {
	OnSearch GetMetadataOnSearchHook
	OnCreate GetMetadataOnCreateHook
	OnUpdate GetMetadataOnUpdateHook
	OnDelete GetMetadataOnDeleteHook
}

type HandlerProps struct {
	Api           project_api.Client
	ResolveActor  func(r *http.Request) (permissions.Actor, error)
	MetadataHooks []MetadataHooks
	OnError       func(handler string, e error)
}

func (p *HandlerProps) Validate() error {
	var err error
	if p.Api == nil {
		err = errors.Join(err, errors.New("api is required"))
	}
	if p.ResolveActor == nil {
		err = errors.Join(err, errors.New("resolveActor is required"))
	}
	return err
}

func (p *HandlerProps) onError(endpoint string, e error) {
	if p.OnError != nil {
		p.OnError(endpoint+"<Project>", e)
	}
}

type SearchRequest struct {
	Query      project.HTTPWhereClause `json:"query"`
	Sort       project.HTTPSortParams  `json:"sort"`
	Projection *project_api.Projection `json:"projection,omitempty"`
	Limit      int                     `json:"limit,omitempty"`
	Skip       int                     `json:"skip,omitempty"`
}

func (sr *SearchRequest) ResolveProjection() *project_api.Projection {
	if sr.Projection != nil {
		return sr.Projection
	}
	projection := project_api.NewProjection(true)

	return &projection
}

func GetSearchHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		searchRequest, err := resolveSearchRequest(r)
		if err != nil {
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusBadRequest)
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		if searchRequest.Limit > 1000 {
			props.onError("Search", errors.New("search limit exceeded"))
			log.Debug().Msg("limit exceeded")
			http.Error(w, "MAX_LIMIT_1000", http.StatusBadRequest)
			return
		}

		searchQuery, err := searchRequest.Query.ToWhereClause()
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to resolve search query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		queryResult, projectionResult, err := props.Api.Search(ctx, actor, searchQuery, project_api.QueryOptions{
			Projection: searchRequest.ResolveProjection(),
			Sort:       searchRequest.Sort.ToSortParams(),
			Limit:      searchRequest.Limit,
			Skip:       searchRequest.Skip,
		})
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to execute search query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := project_api.ToHTTPQueryResult(queryResult, projectionResult)
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to convert query result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		for i, metadataHook := range props.MetadataHooks {
			if metadataHook.OnSearch != nil {
				metadata, err := metadataHook.OnSearch(ctx, actor, response)
				if err != nil {
					log.Debug().Err(err).Int("idx", i).Msg("failed to execute metadata hook")
					http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
					return
				}
				response.Metadata = metadata
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("Search", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusInternalServerError)
			return
		}
	}, nil
}

type SelectRequest struct {
	Projection *project_api.Projection `json:"projection,omitempty"`
}

func (sr *SelectRequest) ResolveProjection() project_api.Projection {
	if sr.Projection != nil {
		return *sr.Projection
	}
	return project_api.NewProjection(true)
}

func resolveSelectRequest(r *http.Request) (SelectRequest, error) {
	switch r.Method {
	case http.MethodPost:
		selectRequest := SelectRequest{}
		err := json.NewDecoder(r.Body).Decode(&selectRequest)
		if err != nil {
			return SelectRequest{}, coded_error.NewInvalidRequestError()
		}
		return selectRequest, nil
	case http.MethodGet:
		return SelectRequest{}, nil
	default:
		return SelectRequest{}, coded_error.NewMethodNotAllowedError(r.Method)
	}
}

type SelectByIdRequest struct {
	Query project.HTTPSelectByIdQuery `json:"query"`
}

func GetSelectByIdHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		selectRequest, err := resolveSelectRequest(r)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to resolve select request")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		id := r.PathValue("id")
		if id == "" {
			props.onError("SelectById", errors.New("missing path parameter: id"))
			log.Debug().Err(err).Msg("missing path parameter")
			http.Error(w, "missing query parameter: id", http.StatusBadRequest)
			return
		}

		selectByIdRequest := SelectByIdRequest{
			Query: project.HTTPSelectByIdQuery{
				Id: id,
			},
		}

		selectByIdQuery, err := selectByIdRequest.Query.ToSelectByIdQuery()
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to resolve select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		projection := selectRequest.ResolveProjection()

		getByIdResult, projectionResult, err := props.Api.SelectById(ctx, actor, selectByIdQuery, projection)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to execute select query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := project_api.ToHTTPModel(getByIdResult, projectionResult)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to convert query result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("SelectById", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

type SaveRequest struct {
	Data project.HTTPRecord `json:"data"`
}

func GetCreateHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			props.onError("Create", errors.New("method not allowed"))
			log.Debug().Msg("method not allowed")
			http.Error(w, "METHOD_NOT_ALLOWED", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to read request body")
			http.Error(w, "INVALID_REQUEST", http.StatusBadRequest)
			return
		}

		projection := project.NewProjection(true)
		projection.Id = false

		var createRequest SaveRequest
		err = json.Unmarshal(body, &createRequest)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to unmarshal request body")
			http.Error(w, "INVALID_REQUEST", http.StatusBadRequest)
			return
		}

		create, err := createRequest.Data.ToModel()
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to convert request body to model")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusBadRequest)
			return
		}

		createResult, projectionResult, err := props.Api.Create(ctx, actor, create, projection)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to execute create query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response, err := project_api.ToHTTPMutationResult(createResult, projectionResult)
		if err != nil {
			props.onError("Create", err)
			log.Debug().Err(err).Msg("failed to convert create result to HTTP response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		for i, metadataHook := range props.MetadataHooks {
			if metadataHook.OnCreate != nil {
				metadata, err := metadataHook.OnCreate(ctx, actor, response)
				if err != nil {
					props.onError("Create", err)
					log.Debug().Err(err).Int("idx", i).Msg("failed to execute metadata hook")
					http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
					return
				}
				response.Metadata = metadata
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("Create", err)
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

func GetDeleteHandler(props HandlerProps) (http.HandlerFunc, error) {
	if err := props.Validate(); err != nil {
		return nil, err
	}
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodDelete {
			props.onError("Delete", errors.New("method not allowed"))
			log.Debug().Msg("method not allowed")
			http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
			return
		}

		ctx := r.Context()

		actor, err := props.ResolveActor(r)
		if err != nil {
			props.onError("Delete", err)
			log.Debug().Err(err).Msg("failed to resolve actor")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), http.StatusUnauthorized)
			return
		}

		id := r.PathValue("id")
		if id == "" {
			props.onError("Delete", errors.New("missing path parameter: id"))
			log.Debug().Msg("missing path parameter")
			http.Error(w, "missing path parameter: id", http.StatusBadRequest)
			return
		}

		err = props.Api.Delete(ctx, actor, id)
		if err != nil {
			props.onError("Delete", err)
			log.Debug().Err(err).Msg("failed to execute delete query")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}

		response := project_api.ToHTTPDeleteResult(id)

		for i, metadataHook := range props.MetadataHooks {
			if metadataHook.OnDelete != nil {
				metadata, err := metadataHook.OnDelete(ctx, actor, response)
				if err != nil {
					props.onError("Delete", err)
					log.Debug().Err(err).Int("idx", i).Msg("failed to execute metadata hook")
					http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
					return
				}
				response.Metadata = metadata
			}
		}

		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(response)
		if err != nil {
			props.onError("Delete", err)
			log.Debug().Err(err).Msg("failed to encode response")
			http.Error(w, coded_error.ResolveErrorCodeAsString(err), coded_error.ResolveHTTPStatus(err))
			return
		}
	}, nil
}

func resolveSearchRequest(r *http.Request) (SearchRequest, error) {
	switch r.Method {
	case http.MethodPost:
		searchRequest := SearchRequest{
			Limit: 100,
		}
		err := json.NewDecoder(r.Body).Decode(&searchRequest)
		if err != nil {
			return SearchRequest{}, err
		}
		return searchRequest, nil
	case http.MethodGet:
		searchRequest := SearchRequest{
			Query: project.HTTPWhereClause{},
		}
		for key, values := range r.URL.Query() {
			switch key {
			case "Limit":
				limit := utils.StringSliceToIntPtr(values)
				if limit == nil {
					return SearchRequest{}, coded_error.NewInvalidRequestError("invalid limit")
				}
				searchRequest.Limit = *limit
				continue
			case "Skip":
				skip := utils.StringSliceToIntPtr(values)
				if skip == nil {
					return SearchRequest{}, coded_error.NewInvalidRequestError("invalid skip")
				}
				searchRequest.Skip = *skip
				continue
			case "id":
				if len(values) == 1 {
					searchRequest.Query.IdEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.IdIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "description":
				if len(values) == 1 {
					searchRequest.Query.DescriptionEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.DescriptionIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "name":
				if len(values) == 1 {
					searchRequest.Query.NameEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.NameIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			case "ownerId":
				if len(values) == 1 {
					searchRequest.Query.OwnerIdEq = utils.StringSliceToStringPtr(values)
					continue
				}
				if len(values) > 1 {
					searchRequest.Query.OwnerIdIn = utils.StringSliceToStringSlicePtr(values)
					continue
				}
			}
		}
		return searchRequest, nil
	default:
		return SearchRequest{}, coded_error.NewMethodNotAllowedError(r.Method)
	}
}
