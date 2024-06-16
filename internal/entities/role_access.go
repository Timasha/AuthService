package entities

import "github.com/Timasha/AuthService/pkg/errlist"

type RoleAccess []byte

func (r *RoleAccess) Scan(src interface{}) error {
	var ok bool

	*r, ok = src.([]byte)
	if !ok {
		return errlist.ErrCantScanRoleID
	}

	return nil
}

func (r *RoleAccess) HaveAccess(required RoleAccess) bool {
	if *r == nil {
		return true
	}

	accessCopy := make([]byte, 0, len(required))

	if len(required) <= len(*r) {
		accessCopy = append(accessCopy, (*r)[:len(required)]...)
	} else if len(required) > len(*r) {
		accessCopy = append(accessCopy, *r...)
		accessCopy = accessCopy[:cap(accessCopy)]
	}

	for i, part := range accessCopy {
		if (required[i] & part) != required[i] {
			return false
		}
	}

	return true
}
