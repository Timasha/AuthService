package entities

type User struct {
	ID       int64  `db:"id"`
	Login    string `db:"login"`
	Password string `db:"password"`

	OtpEnabled bool   `db:"otp_enabled"`
	OtpKey     string `db:"otp_key"`

	Role
}

func (u *User) Check(haveUser *User) bool {
	return (u.Login == haveUser.Login) && (u.Password == haveUser.Password)
}

func (u *User) HaveAccess(required RoleAccess) bool {
	return u.Role.HaveAccess(required)
}
