package jwt

type TokenType int

const (
	AccessTokenType TokenType = iota
	IntermediateTokenType
	RefreshTokenType
)
