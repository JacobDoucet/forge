package actor_role

import (
	"github.com/JacobDoucet/forge/example/generated/go/enum_role"
)

type HTTPRecord struct {
	Role *enum_role.Value `json:"role,omitempty"`
}

func (r *HTTPRecord) ToModel() (Model, error) {
	m := Model{}
	if r.Role != nil {
		elemrole0 := r.Role
		m.Role = *elemrole0
	}
	return m, nil
}

func (r *HTTPRecord) ToProjection() (Projection, error) {
	p := Projection{}
	if r.Role != nil {
		p.Role = true
	}
	return p, nil
}

type HTTPWhereClause struct {
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
}

func (o HTTPWhereClause) ToWhereClause() (WhereClause, error) {
	to := WhereClause{}
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
	return to, nil
}

type HTTPSortParams struct {
}

func (s HTTPSortParams) ToSortParams() SortParams {
	to := SortParams{}
	return to
}
