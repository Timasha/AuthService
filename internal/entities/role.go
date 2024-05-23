package entities

import (
	"github.com/Timasha/AuthService/pkg/errlist"
)

type RoleAccess []byte

func (r *RoleAccess) Scan(src interface{}) error {
	var ok bool

	*r, ok = src.([]byte)
	if !ok {
		return errlist.ErrCantScanRoleID
	}

	return nil
}

type Role struct {
	ID     int64      `json:"roleID" db:"role_id"`
	Access RoleAccess `json:"roleAccess" db:"role_access"`
	Name   string     `json:"roleName" db:"role_name"`
}

func (r *Role) HaveAccess(required RoleAccess) bool {
	if r.Access == nil {
		return true
	}

	accessCopy := make([]byte, 0, len(required))

	if len(required) <= len(r.Access) {
		accessCopy = append(accessCopy, r.Access[:len(required)]...)
	} else if len(required) > len(r.Access) {
		accessCopy = append(accessCopy, r.Access...)
		accessCopy = accessCopy[:cap(accessCopy)]
	}

	for i, part := range accessCopy {
		if (required[i] & part) != required[i] {
			return false
		}
	}

	return true
}
