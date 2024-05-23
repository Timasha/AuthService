package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserValidateTestCase struct {
	name      string
	user      *User
	givenUser *User
	want      bool
}

func TestUserValidate(t *testing.T) {
	t.Parallel()
	testCases := []UserValidateTestCase{
		{
			name: "Valid user",
			user: &User{
				Login:    "login",
				Password: "password",
			},
			givenUser: &User{
				Login:    "login",
				Password: "password",
			},
			want: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, tc.user.Check(tc.givenUser))
		})
	}
}
