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

func (h *FiberHandlersProvider) AuthenticateUserByLoginHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		body := c.Request().Body()
		var (
			req  requests.AuthenticateUserByLoginRequest
			resp responses.AuthenticateUserByLoginResponse
		)

		unmarshErr := json.Unmarshal(body, &req)

		if unmarshErr != nil {
			resp.Err = "unmarshal request error: " + unmarshErr.Error()
			resp.ErrCode = errsutil.ErrInputCode
			c.Status(400).JSON(resp)
			return nil
		}

		accessToken, refreshToken, authErr := h.casesProvider.AuthenticateUserByLogin(h.ctx, req.Login, req.Password)

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

		resp.AccessToken = accessToken
		resp.RefreshToken = refreshToken
		resp.ErrCode = errsutil.SuccessCode

		c.Status(200).JSON(resp)

		return nil
	}
}
