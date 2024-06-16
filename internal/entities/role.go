package entities

type Role struct {
	ID     int64      `json:"roleID" db:"role_id"`
	Access RoleAccess `json:"roleAccess" db:"role_access"`
	Name   string     `json:"roleName" db:"role_name"`
}

func (r *Role) HaveAccess(required RoleAccess) bool {
	return r.Access.HaveAccess(required)
}
