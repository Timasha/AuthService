package api

import (
	"AuthService/internal/cases"
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/errsutil"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type AuthorizeUserRequest struct {
	AccessToken    string        `json:"accessToken" validate:"required"`
	RequiredRoleId models.RoleId `json:"requiredRoleId" validate:"required"`
}

type AuthorizeUserResponses struct {
	BaseResponse

	UserId string `json:"userId"`
}

func (a *Auth) GetAuthorizeUserHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			req  AuthorizeUserRequest
			resp AuthorizeUserResponses
		)

		unmarshErr := a.bodySerializer.Unmarshal(c.Body(), &req)

		if unmarshErr != nil {
			resp.Err = ErrInvalidInput.Error() + unmarshErr.Error()
			resp.ErrCode = ErrInvalidInput.ErrorCode()

			data, _ := a.bodySerializer.Marshal(resp)

			c.Status(400).Write(data)
			return nil
		}

		var args cases.AuthorizeUserArgs = cases.AuthorizeUserArgs{
			Ctx:            a.ctx,
			AccessToken:    req.AccessToken,
			RequiredRoleId: req.RequiredRoleId,
		}

		returned := a.casesProvider.AuthorizeUser(args)

		var errWithCode errsutil.AuthErr

		errors.As(returned.Err, &errWithCode)

		if returned.Err != nil {
			if errWithCode == cases.ErrServiceInternal || errWithCode == cases.ErrServiceNotAvaliable {
				c.Status(500)
			} else {
				c.Status(400)
			}
			resp.Err = errWithCode.Error()
			resp.ErrCode = errWithCode.ErrorCode()

			data, _ := a.bodySerializer.Marshal(resp)

			c.Write(data)
			return nil
		}

		resp.UserId = returned.UserId
		resp.ErrCode = errsutil.SuccessCode

		data, _ := a.bodySerializer.Marshal(resp)

		c.Status(200).Write(data)
		return nil
	}
}
