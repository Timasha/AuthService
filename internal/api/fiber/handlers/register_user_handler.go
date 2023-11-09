package handlers

import (
	"auth/internal/cases/errs"
	"auth/internal/cases/iomodels"
	"auth/internal/logic/models"
	"auth/internal/utils/errsutil"
	"auth/internal/utils/iomodels/requests"
	"auth/internal/utils/iomodels/responses"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (h *FiberHandlersProvider) RegisterUserHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Request().Body()
		var (
			req  requests.RegisterUserRequest
			resp responses.RegisterUserResponses
		)
		unmarshErr := json.Unmarshal(body, &req)
		if unmarshErr != nil {
			resp.Err = "unmarshal request error: " + unmarshErr.Error()
			resp.ErrCode = errsutil.ErrInputCode
			c.Status(400).JSON(resp)
			return nil
		}

		var args iomodels.RegisterUserArgs = iomodels.RegisterUserArgs{
			Ctx: h.ctx,
			User: models.User{
				Login: req.Login,
				Password: req.Password,
			},
		}
		returned := h.casesProvider.RegisterUser(args)

		var regErrWithCode errsutil.AuthErr

		errors.As(returned.Err, &regErrWithCode)

		if returned.Err == (errs.ErrServiceInternal{}) || returned.Err == (errs.ErrServiceNotAvaliable{}) {
			resp.Err = regErrWithCode.Error()
			resp.ErrCode = regErrWithCode.ErrCode()
			c.Status(500).JSON(resp)
			return nil
		} else if returned.Err != nil {
			resp.Err = regErrWithCode.Error()
			resp.ErrCode = regErrWithCode.ErrCode()
			c.Status(400).JSON(resp)
			return nil
		}
		resp.ErrCode = errsutil.SuccessCode

		c.Status(200).JSON(resp)
		return nil
	}
}
