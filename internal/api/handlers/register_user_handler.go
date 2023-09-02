package handlers

import (
	"auth/internal/cases/errs"
	"auth/internal/logic/models"
	"auth/internal/utils/errsutil"
	"auth/internal/utils/iomodels/requests"
	"auth/internal/utils/iomodels/responses"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (h *HandlersProvider) RegisterUserHandler() fiber.Handler {
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

		var user models.User = models.User{
			Login:    req.Login,
			Password: req.Password,
		}
		regErr := h.casesProvider.RegisterUser(h.ctx, user)

		var regErrWithCode errsutil.AuthErr

		errors.As(regErr, &regErrWithCode)

		if regErr == (errs.ErrServiceInternal{}) || regErr == (errs.ErrServiceNotAvaliable{}) {
			resp.Err = regErr.Error()
			resp.ErrCode = regErrWithCode.ErrCode()
			c.Status(500).JSON(resp)
			return nil
		} else if regErr != nil {
			resp.Err = regErr.Error()
			resp.ErrCode = regErrWithCode.ErrCode()
			c.Status(400).JSON(resp)
			return nil
		}
		resp.ErrCode = errsutil.SuccessCode

		c.Status(200).JSON(resp)
		return nil
	}
}
