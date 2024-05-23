package convert

import (
	"github.com/Timasha/AuthService/internal/usecase"
	"github.com/Timasha/AuthService/pkg/api"
)

func AuthorizeUserRequestFromProto(m *api.AuthorizeRequest) (args usecase.AuthorizeUserRequest) {
	args = usecase.AuthorizeUserRequest{
		AccessToken:        m.GetAccessToken(),
		RequiredRoleAccess: m.GetRequiredRoleAccess(),
	}

	return args
}

func AuthorizeUserResponseToProto(m usecase.AuthorizeUserResponse) (resp *api.AuthorizeResponse) {
	resp = &api.AuthorizeResponse{
		UserID: m.UserID,
	}

	return resp
}
