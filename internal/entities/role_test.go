package entities_test

import (
	"testing"

	"github.com/Timasha/AuthService/internal/entities"
	"github.com/Timasha/AuthService/utils/consts"
	"github.com/stretchr/testify/assert"
)

type roleTestCase struct {
	name       string
	role       entities.Role
	required   entities.RoleAccess
	haveAccess bool
}

func TestRoleHaveAccess(t *testing.T) {
	t.Parallel()
	testCases := []roleTestCase{
		{
			"Exact access",
			entities.Role{
				Access: entities.RoleAccess{1, 2, 3},
				Name:   "",
			},
			entities.RoleAccess{1, 2, 3},
			true,
		},
		{
			"Partial access",
			entities.Role{
				Access: entities.RoleAccess{1, 2, 3},
				Name:   "",
			},
			entities.RoleAccess{1, 2},
			true,
		},
		{
			"No access",
			entities.Role{
				Access: entities.RoleAccess{1, 2, 3},
				Name:   "",
			},
			entities.RoleAccess{1, 2, 4},
			false,
		},
		{
			"Partial access with different length",
			entities.Role{
				Access: entities.RoleAccess{1, 2, 3},
				Name:   "",
			},
			entities.RoleAccess{1, 2, 3, 4},
			false,
		},
		{
			"No access. One access is missing",
			entities.Role{
				Access: entities.RoleAccess{44},
				Name:   "",
			},
			entities.RoleAccess{45},
			false,
		},
		{
			"Have access. One access more",
			entities.Role{
				Access: entities.RoleAccess{108},
				Name:   "",
			},
			entities.RoleAccess{44},
			true,
		},
		{
			"No access. Have number is less than required",
			entities.Role{
				Access: entities.RoleAccess{128},
				Name:   "",
			},
			entities.RoleAccess{44},
			false,
		},
		{
			"Check root role",
			consts.RootRole,
			entities.RoleAccess{255, 255, 255, 255},
			true,
		},
		{
			"Check default role",
			consts.DefaultRole,
			entities.RoleAccess{1},
			false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.haveAccess, tc.role.HaveAccess(tc.required))
		})
	}
}
