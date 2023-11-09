package handlers

import (
	"auth/internal/cases/errs"
	"auth/internal/cases/iomodels"
	"auth/internal/utils/errsutil"
	"auth/internal/utils/iomodels/requests"
	"auth/internal/utils/iomodels/responses"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (h *FiberHandlersProvider) AuthorizeUserHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Request().Body()
		var (
			req  requests.AuthorizeUserRequest
			resp responses.AuthorizeUserResponses
		)
		unmarshErr := json.Unmarshal(body, &req)
		if unmarshErr != nil {
			resp.Err = "unmarshal request error: " + unmarshErr.Error()
			resp.ErrCode = errsutil.ErrInputCode
			c.Status(400).JSON(resp)
			return nil
		}

		var args iomodels.AuthorizeUserArgs = iomodels.AuthorizeUserArgs{
			Ctx: h.ctx,
			AccessToken: req.AccessToken,
			Login: req.Login,
		}

		returned := h.casesProvider.AuthorizeUser(args)

		var authErrWithCode errsutil.AuthErr

		errors.As(returned.Err, &authErrWithCode)

		if authErrWithCode == (errs.ErrServiceInternal{}) || authErrWithCode == (errs.ErrServiceNotAvaliable{}) {
			resp.Err = authErrWithCode.Error()
			resp.ErrCode = authErrWithCode.ErrCode()
			c.Status(500).JSON(resp)
			return nil
		} else if returned.Err != nil {
			resp.Err = authErrWithCode.Error()
			resp.ErrCode = authErrWithCode.ErrCode()
			c.Status(400).JSON(resp)
			return nil
		}

		resp.Uuid = returned.UserId
		resp.ErrCode = errsutil.SuccessCode
		c.Status(200).JSON(resp)
		return nil
	}
}
