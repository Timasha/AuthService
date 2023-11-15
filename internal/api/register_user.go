package api

import (
	"AuthService/internal/cases"
	"AuthService/internal/logic/models"
	"AuthService/internal/utils/errsutil"
	"encoding/json"
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
		unmarshErr := json.Unmarshal(c.Body(), &req)
		if unmarshErr != nil {
			resp.Err = "unmarshal request error: " + unmarshErr.Error()
			resp.ErrCode = errsutil.ErrInputCode

			data, _ := json.Marshal(resp)

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

			data, _ := json.Marshal(resp)

			c.Write(data)
			return nil
		}
		resp.ErrCode = errsutil.SuccessCode

		c.Status(200).JSON(resp)
		return nil
	}
}
