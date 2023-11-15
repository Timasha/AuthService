package api

import (
	"AuthService/internal/cases"
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/errsutil"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type AuthorizeUserRequest struct {
	AccessToken    string        `json:"accessToken" validate:"required"`
	RequiredRoleId models.RoleId `json:"requiredRoleId" validate:"required"`
}

type AuthorizeUserResponses struct {
	BaseResponse

	Uuid string `json:"uuid"`
}

func (a *Auth) GetAuthorizeUserHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			req  AuthorizeUserRequest
			resp AuthorizeUserResponses
		)

		unmarshErr := json.Unmarshal(c.Body(), &req)

		if unmarshErr != nil {
			resp.Err = "unmarshal request error: " + unmarshErr.Error()
			resp.ErrCode = errsutil.ErrInputCode

			data, _ := json.Marshal(resp)

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

			data, _ := json.Marshal(resp)

			c.Write(data)
			return nil
		}

		resp.Uuid = returned.UserId
		resp.ErrCode = errsutil.SuccessCode

		data, _ := json.Marshal(resp)

		c.Status(200).Write(data)
		return nil
	}
}
