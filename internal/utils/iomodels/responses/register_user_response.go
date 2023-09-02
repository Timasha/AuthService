package responses

type RegisterUserResponses struct {
	Err     string `json:"error"`
	ErrCode int    `json:"errorCode"`
}
