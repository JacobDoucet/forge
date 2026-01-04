package user

import (
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_role"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/actor_trace"
	"github.com/JacobDoucet/forge/example/apps/backend/generated/enum_role"
)

type HTTPRecord struct {
	Id            *string                  `json:"id,omitempty"`
	ActorRoles    *[]actor_role.HTTPRecord `json:"actorRoles,omitempty"`
	Created       *actor_trace.HTTPRecord  `json:"created,omitempty"`
	Email         *string                  `json:"email,omitempty"`
	FirstName     *string                  `json:"firstName,omitempty"`
	Language      *string                  `json:"language,omitempty"`
	LastName      *string                  `json:"lastName,omitempty"`
	Role          *enum_role.Value         `json:"role,omitempty"`
	Updated       *actor_trace.HTTPRecord  `json:"updated,omitempty"`
	UpdatedByUser *actor_trace.HTTPRecord  `json:"updatedByUser,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Id != nil {
		elemid0 := r.Id
		m.Id = *elemid0
	}
	if r.ActorRoles != nil {
		elemactorRoles0 := make([]actor_role.Model, 0)
		for _, ractorRoles0 := range *r.ActorRoles {
			elemactorRoles1, err := ractorRoles0.ToModel()
			if err != nil {
				return m, err
			}
			elemactorRoles0 = append(elemactorRoles0, elemactorRoles1)
		}
		m.ActorRoles = elemactorRoles0
	}
	if r.Created != nil {
		elemcreated0, err := r.Created.ToModel()
		if err != nil {
			return m, err
		}
		m.Created = elemcreated0
	}
	if r.Email != nil {
		elememail0 := r.Email
		m.Email = *elememail0
	}
	if r.FirstName != nil {
		elemfirstName0 := r.FirstName
		m.FirstName = *elemfirstName0
	}
	if r.Language != nil {
		elemlanguage0 := r.Language
		m.Language = *elemlanguage0
	}
	if r.LastName != nil {
		elemlastName0 := r.LastName
		m.LastName = *elemlastName0
	}
	if r.Role != nil {
		elemrole0 := r.Role
		m.Role = *elemrole0
	}
	if r.Updated != nil {
		elemupdated0, err := r.Updated.ToModel()
		if err != nil {
			return m, err
		}
		m.Updated = elemupdated0
	}
	if r.UpdatedByUser != nil {
		elemupdatedByUser0, err := r.UpdatedByUser.ToModel()
		if err != nil {
			return m, err
		}
		m.UpdatedByUser = elemupdatedByUser0
	}
	return m, nil
}

func (r *HTTPRecord) ToProjection() (Projection, error) {
	p := Projection{}
	if r.Id != nil {
		p.Id = true
	}
	if r.ActorRoles != nil {
		p.ActorRoles = true
		p.ActorRolesFields = actor_role.NewProjection(true)
	}
	if r.Created != nil {
		p.Created = true
		p.CreatedFields = actor_trace.NewProjection(true)
	}
	if r.Email != nil {
		p.Email = true
	}
	if r.FirstName != nil {
		p.FirstName = true
	}
	if r.Language != nil {
		p.Language = true
	}
	if r.LastName != nil {
		p.LastName = true
	}
	if r.Role != nil {
		p.Role = true
	}
	if r.Updated != nil {
		p.Updated = true
		p.UpdatedFields = actor_trace.NewProjection(true)
	}
	if r.UpdatedByUser != nil {
		p.UpdatedByUser = true
		p.UpdatedByUserFields = actor_trace.NewProjection(true)
	}
	return p, nil
}

type HTTPSelectByIdQuery struct {
	Id string `json:"id"`
}
type HTTPSelectByEmailIdxQuery struct {
	Email string `json:"email"`
}

type HTTPWhereClause struct {
	// id (Ref<User>) search options
	IdEq     *string   `json:"idEq,omitempty"`
	IdIn     *[]string `json:"idIn,omitempty"`
	IdNin    *[]string `json:"idNin,omitempty"`
	IdExists *bool     `json:"idExists,omitempty"`
	// actorRoles (List<ActorRole>) search options
	ActorRoles      *actor_role.HTTPWhereClause `json:"actorRoles,omitempty"`
	ActorRolesEmpty *bool                       `json:"actorRolesEmpty,omitempty"`
	// created (ActorTrace) search options
	Created *actor_trace.HTTPWhereClause `json:"created,omitempty"`
	// email (string) search options
	EmailEq     *string   `json:"emailEq,omitempty"`
	EmailNe     *string   `json:"emailNe,omitempty"`
	EmailGt     *string   `json:"emailGt,omitempty"`
	EmailGte    *string   `json:"emailGte,omitempty"`
	EmailLt     *string   `json:"emailLt,omitempty"`
	EmailLte    *string   `json:"emailLte,omitempty"`
	EmailIn     *[]string `json:"emailIn,omitempty"`
	EmailNin    *[]string `json:"emailNin,omitempty"`
	EmailExists *bool     `json:"emailExists,omitempty"`
	EmailLike   *string   `json:"emailLike,omitempty"`
	EmailNlike  *string   `json:"emailNlike,omitempty"`
	// firstName (string) search options
	FirstNameEq     *string   `json:"firstNameEq,omitempty"`
	FirstNameNe     *string   `json:"firstNameNe,omitempty"`
	FirstNameGt     *string   `json:"firstNameGt,omitempty"`
	FirstNameGte    *string   `json:"firstNameGte,omitempty"`
	FirstNameLt     *string   `json:"firstNameLt,omitempty"`
	FirstNameLte    *string   `json:"firstNameLte,omitempty"`
	FirstNameIn     *[]string `json:"firstNameIn,omitempty"`
	FirstNameNin    *[]string `json:"firstNameNin,omitempty"`
	FirstNameExists *bool     `json:"firstNameExists,omitempty"`
	FirstNameLike   *string   `json:"firstNameLike,omitempty"`
	FirstNameNlike  *string   `json:"firstNameNlike,omitempty"`
	// language (string) search options
	LanguageEq     *string   `json:"languageEq,omitempty"`
	LanguageNe     *string   `json:"languageNe,omitempty"`
	LanguageGt     *string   `json:"languageGt,omitempty"`
	LanguageGte    *string   `json:"languageGte,omitempty"`
	LanguageLt     *string   `json:"languageLt,omitempty"`
	LanguageLte    *string   `json:"languageLte,omitempty"`
	LanguageIn     *[]string `json:"languageIn,omitempty"`
	LanguageNin    *[]string `json:"languageNin,omitempty"`
	LanguageExists *bool     `json:"languageExists,omitempty"`
	LanguageLike   *string   `json:"languageLike,omitempty"`
	LanguageNlike  *string   `json:"languageNlike,omitempty"`
	// lastName (string) search options
	LastNameEq     *string   `json:"lastNameEq,omitempty"`
	LastNameNe     *string   `json:"lastNameNe,omitempty"`
	LastNameGt     *string   `json:"lastNameGt,omitempty"`
	LastNameGte    *string   `json:"lastNameGte,omitempty"`
	LastNameLt     *string   `json:"lastNameLt,omitempty"`
	LastNameLte    *string   `json:"lastNameLte,omitempty"`
	LastNameIn     *[]string `json:"lastNameIn,omitempty"`
	LastNameNin    *[]string `json:"lastNameNin,omitempty"`
	LastNameExists *bool     `json:"lastNameExists,omitempty"`
	LastNameLike   *string   `json:"lastNameLike,omitempty"`
	LastNameNlike  *string   `json:"lastNameNlike,omitempty"`
	// role (Role) search options
	RoleEq     *enum_role.Value   `json:"roleEq,omitempty"`
	RoleNe     *enum_role.Value   `json:"roleNe,omitempty"`
	RoleGt     *enum_role.Value   `json:"roleGt,omitempty"`
	RoleGte    *enum_role.Value   `json:"roleGte,omitempty"`
	RoleLt     *enum_role.Value   `json:"roleLt,omitempty"`
	RoleLte    *enum_role.Value   `json:"roleLte,omitempty"`
	RoleIn     *[]enum_role.Value `json:"roleIn,omitempty"`
	RoleNin    *[]enum_role.Value `json:"roleNin,omitempty"`
	RoleExists *bool              `json:"roleExists,omitempty"`
	// updated (ActorTrace) search options
	Updated *actor_trace.HTTPWhereClause `json:"updated,omitempty"`
	// updatedByUser (ActorTrace) search options
	UpdatedByUser *actor_trace.HTTPWhereClause `json:"updatedByUser,omitempty"`
}

func (o HTTPSelectByIdQuery) ToSelectByIdQuery() (SelectByIdQuery, error) {
	to := SelectByIdQuery{}
	elemid0 := o.Id
	to.Id = elemid0
	return to, nil
}
func (o HTTPSelectByEmailIdxQuery) ToSelectByEmailIdxQuery() (SelectByEmailIdxQuery, error) {
	to := SelectByEmailIdxQuery{}
	elememail0 := o.Email
	to.Email = elememail0
	return to, nil
}

func (o HTTPWhereClause) ToWhereClause() (WhereClause, error) {
	to := WhereClause{}
	if o.IdEq != nil {
		elemidEq0 := o.IdEq
		to.IdEq = elemidEq0
	}
	if o.IdIn != nil {
		elemidIn0 := make([]string, 0)
		for _, oidIn0 := range *o.IdIn {
			elemidIn1 := oidIn0
			elemidIn0 = append(elemidIn0, elemidIn1)
		}
		to.IdIn = &elemidIn0
	}
	if o.IdNin != nil {
		elemidNin0 := make([]string, 0)
		for _, oidNin0 := range *o.IdNin {
			elemidNin1 := oidNin0
			elemidNin0 = append(elemidNin0, elemidNin1)
		}
		to.IdNin = &elemidNin0
	}
	if o.IdExists != nil {
		elemidExists0 := o.IdExists
		to.IdExists = elemidExists0
	}
	if o.ActorRoles != nil {
		elemactorRoles0, err := o.ActorRoles.ToWhereClause()
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
		elemcreated0, err := o.Created.ToWhereClause()
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
		elemupdated0, err := o.Updated.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.Updated = &elemupdated0
	}
	if o.UpdatedByUser != nil {
		elemupdatedByUser0, err := o.UpdatedByUser.ToWhereClause()
		if err != nil {
			return to, err
		}
		to.UpdatedByUser = &elemupdatedByUser0
	}
	return to, nil
}

type HTTPSortParams struct {
	CreatedAt *int8 `json:"createdAt,omitempty"`
	Email     *int8 `json:"email,omitempty"`
	UpdatedAt *int8 `json:"updatedAt,omitempty"`
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	if s.CreatedAt != nil {
		to.CreatedAt = *s.CreatedAt
	}
	if s.Email != nil {
		to.Email = *s.Email
	}
	if s.UpdatedAt != nil {
		to.UpdatedAt = *s.UpdatedAt
	}
	return to
}
