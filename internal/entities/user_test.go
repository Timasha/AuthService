package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type userCheckTestCase struct {
	name      string
	user      *User
	givenUser *User
	want      bool
}

func TestUserCheck(t *testing.T) {
	t.Parallel()
	testCases := []userCheckTestCase{
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
		{
			name: "invalid user",
			user: &User{
				Login:    "login",
				Password: "password",
			},
			givenUser: &User{
				Login:    "login",
				Password: "password1",
			},
			want: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			assert.Equal(t, tc.want, tc.user.Check(tc.givenUser))
		})
	}
}
