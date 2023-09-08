package responses

type AuthorizeUserResponses struct {
	Uuid string `json:"uuid"`

	Err     string `json:"error"`
	ErrCode int    `json:"errorCode"`
}
