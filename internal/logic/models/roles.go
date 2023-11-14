package models

type RoleId uint

type Role struct {
	RoleId   RoleId `json:"roleId"`
	RoleName string `json:"roleName"`
}

func (r RoleId) HavaAccess(required RoleId) bool {
	return r >= required
}
