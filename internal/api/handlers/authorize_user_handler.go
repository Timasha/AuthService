package handlers

import (
	"auth/internal/cases/errs"
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

		uuid, authErr := h.casesProvider.AuthorizeUser(h.ctx, req.AccessToken, req.Login)
		var authErrWithCode errsutil.AuthErr

		errors.As(authErr, &authErrWithCode)

		if authErr == (errs.ErrServiceInternal{}) || authErr == (errs.ErrServiceNotAvaliable{}) {
			resp.Err = authErr.Error()
			resp.ErrCode = authErrWithCode.ErrCode()
			c.Status(500).JSON(resp)
			return nil
		} else if authErr != nil {
			resp.Err = authErr.Error()
			resp.ErrCode = authErrWithCode.ErrCode()
			c.Status(400).JSON(resp)
			return nil
		}
		resp.Uuid = uuid
		resp.ErrCode = errsutil.SuccessCode
		c.Status(200).JSON(resp)
		return nil
	}
}
