package errsutil

type AuthErr interface {
	error
	ErrCode() int
}
