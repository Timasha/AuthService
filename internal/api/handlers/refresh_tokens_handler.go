package handlers

import (
	casesErrs "auth/internal/cases/errs"
	"auth/internal/utils/errsutil"
	"auth/internal/utils/iomodels/requests"
	"auth/internal/utils/iomodels/responses"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

func (h *FiberHandlersProvider) RefreshTokensHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			req  requests.RefreshTokenRequest
			resp responses.RefreshTokensResponse
		)

		unmarshalErr := json.Unmarshal(c.Body(), &req)

		if unmarshalErr != nil {
			resp.Err = "Unmarshal json error: " + unmarshalErr.Error()
			resp.ErrCode = errsutil.ErrInputCode
			c.Status(400).JSON(resp)
			return nil
		}

		accessToken, refreshToken, err := h.casesProvider.RefreshTokens(h.ctx, req.RefreshToken, req.AccessToken, req.Login)

		var errWithCode errsutil.AuthErr

		errors.As(err, &errWithCode)

		if errWithCode == (casesErrs.ErrServiceInternal{}) || errWithCode == (casesErrs.ErrServiceNotAvaliable{}) {
			resp.Err = errWithCode.Error()
			resp.ErrCode = errWithCode.ErrCode()
			c.Status(500).JSON(resp)
			return nil
		} else if err != nil {

			resp.Err = err.Error()
			resp.ErrCode = errWithCode.ErrCode()
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
