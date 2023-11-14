package models

type User struct {
	UserID   string
	Login    string
	Password string

	OtpEnabled bool
	OtpKey     string

	Role Role
}
