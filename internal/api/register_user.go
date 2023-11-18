package api

import (
	"AuthService/internal/cases"
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/errsutil"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type RegisterUserRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type RegisterUserResponses struct {
	BaseResponse
}

func (a *Auth) GetRegisterUserHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			req  RegisterUserRequest
			resp RegisterUserResponses
		)
		unmarshErr := a.bodySerializer.Unmarshal(c.Body(), &req)
		if unmarshErr != nil {
			resp.Err = ErrInvalidInput.Error() + unmarshErr.Error()
			resp.ErrCode = ErrInvalidInput.ErrorCode()

			data, _ := a.bodySerializer.Marshal(resp)

			c.Status(400).Write(data)
			return nil
		}

		var args cases.RegisterUserArgs = cases.RegisterUserArgs{
			Ctx: a.ctx,
			User: models.User{
				Login:    req.Login,
				Password: req.Password,
			},
		}
		returned := a.casesProvider.RegisterUser(args)

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
		resp.ErrCode = errsutil.SuccessCode

		data, _ := a.bodySerializer.Marshal(resp)

		c.Status(200).Write(data)
		return nil
	}
}
