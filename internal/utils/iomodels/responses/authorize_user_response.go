package responses

type AuthorizeUserResponses struct {
	Err     string `json:"error"`
	ErrCode int    `json:"errorCode"`
}
