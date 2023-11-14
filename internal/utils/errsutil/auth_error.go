package errsutil

type AuthErr struct {
	Msg     string
	ErrCode AuthErrCode
}

func (a AuthErr) Error() string {
	return a.Msg
}

func (a AuthErr) ErrorCode() AuthErrCode {
	return a.ErrCode
}
