package user

import (
	"errors"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_role"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_role"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Model struct {
	Id            string
	ActorRoles    []actor_role.Model
	Created       actor_trace.Model
	Email         string
	FirstName     string
	Language      string
	LastName      string
	Role          enum_role.Value
	Updated       actor_trace.Model
	UpdatedByUser actor_trace.Model
}

func (m *Model) ToMongoRecord(projection Projection) (MongoRecord, error) {
	r := MongoRecord{}
	if m.Id != "" {
		elemid0, err := primitive.ObjectIDFromHex(m.Id)
		if err != nil {
			return r, errors.Join(errors.New("invalid m.Id"), err)
		}
		r.Id = &elemid0
	}
	if projection.ActorRoles {
		elemactorRoles0 := make([]actor_role.MongoRecord, 0)
		for _, mactorRoles0 := range m.ActorRoles {
			elemactorRoles1, err := mactorRoles0.ToMongoRecord(projection.ActorRolesFields)
			if err != nil {
				return r, err
			}
			elemactorRoles0 = append(elemactorRoles0, elemactorRoles1)
		}
		r.ActorRoles = &elemactorRoles0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToMongoRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Email {
		elememail0 := m.Email
		r.Email = &elememail0
	}
	if projection.FirstName {
		elemfirstName0 := m.FirstName
		r.FirstName = &elemfirstName0
	}
	if projection.Language {
		elemlanguage0 := m.Language
		r.Language = &elemlanguage0
	}
	if projection.LastName {
		elemlastName0 := m.LastName
		r.LastName = &elemlastName0
	}
	if projection.Role {
		elemrole0 := m.Role
		r.Role = &elemrole0
	}
	if projection.Updated {
		elemupdated0, err := m.Updated.ToMongoRecord(projection.UpdatedFields)
		if err != nil {
			return r, err
		}
		r.Updated = &elemupdated0
	}
	if projection.UpdatedByUser {
		elemupdatedByUser0, err := m.UpdatedByUser.ToMongoRecord(projection.UpdatedByUserFields)
		if err != nil {
			return r, err
		}
		r.UpdatedByUser = &elemupdatedByUser0
	}
	return r, nil
}

func (m *Model) ToHTTPRecord(projection Projection) (HTTPRecord, error) {
	r := HTTPRecord{}
	if m.Id != "" {
		elemid0 := m.Id
		r.Id = &elemid0
	}
	if projection.ActorRoles {
		elemactorRoles0 := make([]actor_role.HTTPRecord, 0)
		for _, mactorRoles0 := range m.ActorRoles {
			elemactorRoles1, err := mactorRoles0.ToHTTPRecord(projection.ActorRolesFields)
			if err != nil {
				return r, err
			}
			elemactorRoles0 = append(elemactorRoles0, elemactorRoles1)
		}
		r.ActorRoles = &elemactorRoles0
	}
	if projection.Created {
		elemcreated0, err := m.Created.ToHTTPRecord(projection.CreatedFields)
		if err != nil {
			return r, err
		}
		r.Created = &elemcreated0
	}
	if projection.Email {
		elememail0 := m.Email
		r.Email = &elememail0
	}
	if projection.FirstName {
		elemfirstName0 := m.FirstName
		r.FirstName = &elemfirstName0
	}
	if projection.Language {
		elemlanguage0 := m.Language
		r.Language = &elemlanguage0
	}
	if projection.LastName {
		elemlastName0 := m.LastName
		r.LastName = &elemlastName0
	}
	if projection.Role {
		elemrole0 := m.Role
		r.Role = &elemrole0
	}
	if projection.Updated {
		elemupdated0, err := m.Updated.ToHTTPRecord(projection.UpdatedFields)
		if err != nil {
			return r, err
		}
		r.Updated = &elemupdated0
	}
	if projection.UpdatedByUser {
		elemupdatedByUser0, err := m.UpdatedByUser.ToHTTPRecord(projection.UpdatedByUserFields)
		if err != nil {
			return r, err
		}
		r.UpdatedByUser = &elemupdatedByUser0
	}
	return r, nil
}

type SelectByIdQuery struct {
	Id string
}
type SelectByEmailIdxQuery struct {
	Email string
}

type WhereClause struct {
	// id (Ref<User>) search options
	IdEq     *string
	IdIn     *[]string
	IdNin    *[]string
	IdExists *bool
	// actorRoles (List<ActorRole>) search options
	ActorRoles      *actor_role.WhereClause
	ActorRolesEmpty *bool
	// created (ActorTrace) search options
	Created *actor_trace.WhereClause
	// email (string) search options
	EmailEq     *string
	EmailNe     *string
	EmailGt     *string
	EmailGte    *string
	EmailLt     *string
	EmailLte    *string
	EmailIn     *[]string
	EmailNin    *[]string
	EmailExists *bool
	EmailLike   *string
	EmailNlike  *string
	// firstName (string) search options
	FirstNameEq     *string
	FirstNameNe     *string
	FirstNameGt     *string
	FirstNameGte    *string
	FirstNameLt     *string
	FirstNameLte    *string
	FirstNameIn     *[]string
	FirstNameNin    *[]string
	FirstNameExists *bool
	FirstNameLike   *string
	FirstNameNlike  *string
	// language (string) search options
	LanguageEq     *string
	LanguageNe     *string
	LanguageGt     *string
	LanguageGte    *string
	LanguageLt     *string
	LanguageLte    *string
	LanguageIn     *[]string
	LanguageNin    *[]string
	LanguageExists *bool
	LanguageLike   *string
	LanguageNlike  *string
	// lastName (string) search options
	LastNameEq     *string
	LastNameNe     *string
	LastNameGt     *string
	LastNameGte    *string
	LastNameLt     *string
	LastNameLte    *string
	LastNameIn     *[]string
	LastNameNin    *[]string
	LastNameExists *bool
	LastNameLike   *string
	LastNameNlike  *string
	// role (Role) search options
	RoleEq     *enum_role.Value
	RoleNe     *enum_role.Value
	RoleGt     *enum_role.Value
	RoleGte    *enum_role.Value
	RoleLt     *enum_role.Value
	RoleLte    *enum_role.Value
	RoleIn     *[]enum_role.Value
	RoleNin    *[]enum_role.Value
	RoleExists *bool
	// updated (ActorTrace) search options
	Updated *actor_trace.WhereClause
	// updatedByUser (ActorTrace) search options
	UpdatedByUser *actor_trace.WhereClause
}

func (o SelectByIdQuery) ToMongoSelectByIdQuery() (MongoSelectByIdQuery, error) {
	to := MongoSelectByIdQuery{}
	elemid0, err := primitive.ObjectIDFromHex(o.Id)
	if err != nil {
		return to, errors.Join(errors.New("invalid o.Id"), err)
	}
	to.Id = elemid0
	return to, nil
}
func (o SelectByEmailIdxQuery) ToMongoSelectByEmailIdxQuery() (MongoSelectByEmailIdxQuery, error) {
	to := MongoSelectByEmailIdxQuery{}
	elememail0 := o.Email
	to.Email = elememail0
	return to, nil
}

func (o WhereClause) ToMongoWhereClause() (MongoWhereClause, error) {
	to := MongoWhereClause{}
	if o.IdEq != nil {
		elemidEq0, err := primitive.ObjectIDFromHex(*o.IdEq)
		if err != nil {
			return to, errors.Join(errors.New("invalid o.IdEq"), err)
		}
		to.IdEq = &elemidEq0
	}
	if o.IdIn != nil {
		elemidIn0 := make([]primitive.ObjectID, 0)
		for _, oidIn0 := range *o.IdIn {
			elemidIn1, err := primitive.ObjectIDFromHex(oidIn0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oidIn0"), err)
			}
			elemidIn0 = append(elemidIn0, elemidIn1)
		}
		to.IdIn = &elemidIn0
	}
	if o.IdNin != nil {
		elemidNin0 := make([]primitive.ObjectID, 0)
		for _, oidNin0 := range *o.IdNin {
			elemidNin1, err := primitive.ObjectIDFromHex(oidNin0)
			if err != nil {
				return to, errors.Join(errors.New("invalid oidNin0"), err)
			}
			elemidNin0 = append(elemidNin0, elemidNin1)
		}
		to.IdNin = &elemidNin0
	}
	if o.IdExists != nil {
		elemidExists0 := o.IdExists
		to.IdExists = elemidExists0
	}
	if o.ActorRoles != nil {
		elemactorRoles0, err := o.ActorRoles.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.ActorRoles = &elemactorRoles0
	}
	if o.ActorRolesEmpty != nil {
		elemactorRolesEmpty0 := o.ActorRolesEmpty
		to.ActorRolesEmpty = elemactorRolesEmpty0
	}
	if o.Created != nil {
		elemcreated0, err := o.Created.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Created = &elemcreated0
	}
	if o.EmailEq != nil {
		elememailEq0 := o.EmailEq
		to.EmailEq = elememailEq0
	}
	if o.EmailNe != nil {
		elememailNe0 := o.EmailNe
		to.EmailNe = elememailNe0
	}
	if o.EmailGt != nil {
		elememailGt0 := o.EmailGt
		to.EmailGt = elememailGt0
	}
	if o.EmailGte != nil {
		elememailGte0 := o.EmailGte
		to.EmailGte = elememailGte0
	}
	if o.EmailLt != nil {
		elememailLt0 := o.EmailLt
		to.EmailLt = elememailLt0
	}
	if o.EmailLte != nil {
		elememailLte0 := o.EmailLte
		to.EmailLte = elememailLte0
	}
	if o.EmailIn != nil {
		elememailIn0 := make([]string, 0)
		for _, oemailIn0 := range *o.EmailIn {
			elememailIn1 := oemailIn0
			elememailIn0 = append(elememailIn0, elememailIn1)
		}
		to.EmailIn = &elememailIn0
	}
	if o.EmailNin != nil {
		elememailNin0 := make([]string, 0)
		for _, oemailNin0 := range *o.EmailNin {
			elememailNin1 := oemailNin0
			elememailNin0 = append(elememailNin0, elememailNin1)
		}
		to.EmailNin = &elememailNin0
	}
	if o.EmailExists != nil {
		elememailExists0 := o.EmailExists
		to.EmailExists = elememailExists0
	}
	if o.EmailLike != nil {
		elememailLike0 := o.EmailLike
		to.EmailLike = elememailLike0
	}
	if o.EmailNlike != nil {
		elememailNlike0 := o.EmailNlike
		to.EmailNlike = elememailNlike0
	}
	if o.FirstNameEq != nil {
		elemfirstNameEq0 := o.FirstNameEq
		to.FirstNameEq = elemfirstNameEq0
	}
	if o.FirstNameNe != nil {
		elemfirstNameNe0 := o.FirstNameNe
		to.FirstNameNe = elemfirstNameNe0
	}
	if o.FirstNameGt != nil {
		elemfirstNameGt0 := o.FirstNameGt
		to.FirstNameGt = elemfirstNameGt0
	}
	if o.FirstNameGte != nil {
		elemfirstNameGte0 := o.FirstNameGte
		to.FirstNameGte = elemfirstNameGte0
	}
	if o.FirstNameLt != nil {
		elemfirstNameLt0 := o.FirstNameLt
		to.FirstNameLt = elemfirstNameLt0
	}
	if o.FirstNameLte != nil {
		elemfirstNameLte0 := o.FirstNameLte
		to.FirstNameLte = elemfirstNameLte0
	}
	if o.FirstNameIn != nil {
		elemfirstNameIn0 := make([]string, 0)
		for _, ofirstNameIn0 := range *o.FirstNameIn {
			elemfirstNameIn1 := ofirstNameIn0
			elemfirstNameIn0 = append(elemfirstNameIn0, elemfirstNameIn1)
		}
		to.FirstNameIn = &elemfirstNameIn0
	}
	if o.FirstNameNin != nil {
		elemfirstNameNin0 := make([]string, 0)
		for _, ofirstNameNin0 := range *o.FirstNameNin {
			elemfirstNameNin1 := ofirstNameNin0
			elemfirstNameNin0 = append(elemfirstNameNin0, elemfirstNameNin1)
		}
		to.FirstNameNin = &elemfirstNameNin0
	}
	if o.FirstNameExists != nil {
		elemfirstNameExists0 := o.FirstNameExists
		to.FirstNameExists = elemfirstNameExists0
	}
	if o.FirstNameLike != nil {
		elemfirstNameLike0 := o.FirstNameLike
		to.FirstNameLike = elemfirstNameLike0
	}
	if o.FirstNameNlike != nil {
		elemfirstNameNlike0 := o.FirstNameNlike
		to.FirstNameNlike = elemfirstNameNlike0
	}
	if o.LanguageEq != nil {
		elemlanguageEq0 := o.LanguageEq
		to.LanguageEq = elemlanguageEq0
	}
	if o.LanguageNe != nil {
		elemlanguageNe0 := o.LanguageNe
		to.LanguageNe = elemlanguageNe0
	}
	if o.LanguageGt != nil {
		elemlanguageGt0 := o.LanguageGt
		to.LanguageGt = elemlanguageGt0
	}
	if o.LanguageGte != nil {
		elemlanguageGte0 := o.LanguageGte
		to.LanguageGte = elemlanguageGte0
	}
	if o.LanguageLt != nil {
		elemlanguageLt0 := o.LanguageLt
		to.LanguageLt = elemlanguageLt0
	}
	if o.LanguageLte != nil {
		elemlanguageLte0 := o.LanguageLte
		to.LanguageLte = elemlanguageLte0
	}
	if o.LanguageIn != nil {
		elemlanguageIn0 := make([]string, 0)
		for _, olanguageIn0 := range *o.LanguageIn {
			elemlanguageIn1 := olanguageIn0
			elemlanguageIn0 = append(elemlanguageIn0, elemlanguageIn1)
		}
		to.LanguageIn = &elemlanguageIn0
	}
	if o.LanguageNin != nil {
		elemlanguageNin0 := make([]string, 0)
		for _, olanguageNin0 := range *o.LanguageNin {
			elemlanguageNin1 := olanguageNin0
			elemlanguageNin0 = append(elemlanguageNin0, elemlanguageNin1)
		}
		to.LanguageNin = &elemlanguageNin0
	}
	if o.LanguageExists != nil {
		elemlanguageExists0 := o.LanguageExists
		to.LanguageExists = elemlanguageExists0
	}
	if o.LanguageLike != nil {
		elemlanguageLike0 := o.LanguageLike
		to.LanguageLike = elemlanguageLike0
	}
	if o.LanguageNlike != nil {
		elemlanguageNlike0 := o.LanguageNlike
		to.LanguageNlike = elemlanguageNlike0
	}
	if o.LastNameEq != nil {
		elemlastNameEq0 := o.LastNameEq
		to.LastNameEq = elemlastNameEq0
	}
	if o.LastNameNe != nil {
		elemlastNameNe0 := o.LastNameNe
		to.LastNameNe = elemlastNameNe0
	}
	if o.LastNameGt != nil {
		elemlastNameGt0 := o.LastNameGt
		to.LastNameGt = elemlastNameGt0
	}
	if o.LastNameGte != nil {
		elemlastNameGte0 := o.LastNameGte
		to.LastNameGte = elemlastNameGte0
	}
	if o.LastNameLt != nil {
		elemlastNameLt0 := o.LastNameLt
		to.LastNameLt = elemlastNameLt0
	}
	if o.LastNameLte != nil {
		elemlastNameLte0 := o.LastNameLte
		to.LastNameLte = elemlastNameLte0
	}
	if o.LastNameIn != nil {
		elemlastNameIn0 := make([]string, 0)
		for _, olastNameIn0 := range *o.LastNameIn {
			elemlastNameIn1 := olastNameIn0
			elemlastNameIn0 = append(elemlastNameIn0, elemlastNameIn1)
		}
		to.LastNameIn = &elemlastNameIn0
	}
	if o.LastNameNin != nil {
		elemlastNameNin0 := make([]string, 0)
		for _, olastNameNin0 := range *o.LastNameNin {
			elemlastNameNin1 := olastNameNin0
			elemlastNameNin0 = append(elemlastNameNin0, elemlastNameNin1)
		}
		to.LastNameNin = &elemlastNameNin0
	}
	if o.LastNameExists != nil {
		elemlastNameExists0 := o.LastNameExists
		to.LastNameExists = elemlastNameExists0
	}
	if o.LastNameLike != nil {
		elemlastNameLike0 := o.LastNameLike
		to.LastNameLike = elemlastNameLike0
	}
	if o.LastNameNlike != nil {
		elemlastNameNlike0 := o.LastNameNlike
		to.LastNameNlike = elemlastNameNlike0
	}
	if o.RoleEq != nil {
		elemroleEq0 := o.RoleEq
		to.RoleEq = elemroleEq0
	}
	if o.RoleNe != nil {
		elemroleNe0 := o.RoleNe
		to.RoleNe = elemroleNe0
	}
	if o.RoleGt != nil {
		elemroleGt0 := o.RoleGt
		to.RoleGt = elemroleGt0
	}
	if o.RoleGte != nil {
		elemroleGte0 := o.RoleGte
		to.RoleGte = elemroleGte0
	}
	if o.RoleLt != nil {
		elemroleLt0 := o.RoleLt
		to.RoleLt = elemroleLt0
	}
	if o.RoleLte != nil {
		elemroleLte0 := o.RoleLte
		to.RoleLte = elemroleLte0
	}
	if o.RoleIn != nil {
		elemroleIn0 := make([]enum_role.Value, 0)
		for _, oroleIn0 := range *o.RoleIn {
			elemroleIn1 := oroleIn0
			elemroleIn0 = append(elemroleIn0, elemroleIn1)
		}
		to.RoleIn = &elemroleIn0
	}
	if o.RoleNin != nil {
		elemroleNin0 := make([]enum_role.Value, 0)
		for _, oroleNin0 := range *o.RoleNin {
			elemroleNin1 := oroleNin0
			elemroleNin0 = append(elemroleNin0, elemroleNin1)
		}
		to.RoleNin = &elemroleNin0
	}
	if o.RoleExists != nil {
		elemroleExists0 := o.RoleExists
		to.RoleExists = elemroleExists0
	}
	if o.Updated != nil {
		elemupdated0, err := o.Updated.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.Updated = &elemupdated0
	}
	if o.UpdatedByUser != nil {
		elemupdatedByUser0, err := o.UpdatedByUser.ToMongoWhereClause()
		if err != nil {
			return to, err
		}
		to.UpdatedByUser = &elemupdatedByUser0
	}
	return to, nil
}

type SortParams struct {
	CreatedAt int8
	Email     int8
	UpdatedAt int8
}

func (s SortParams) ToMongoSortParams() MongoSortParams {
	to := MongoSortParams{}
	to.CreatedAt = s.CreatedAt
	to.Email = s.Email
	to.UpdatedAt = s.UpdatedAt
	return to
}
