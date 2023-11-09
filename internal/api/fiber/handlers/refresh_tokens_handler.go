package handlers

import (
	casesErrs "auth/internal/cases/errs"
	"auth/internal/cases/iomodels"
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

		var args iomodels.RefreshTokensArgs = iomodels.RefreshTokensArgs{
			Ctx: h.ctx,
			AccessToken: req.AccessToken,
			RefreshToken: req.RefreshToken,
			Login: req.Login,
		}

		returned := h.casesProvider.RefreshTokens(args)

		var errWithCode errsutil.AuthErr

		errors.As(returned.Err, &errWithCode)

		if errWithCode == (casesErrs.ErrServiceInternal{}) || errWithCode == (casesErrs.ErrServiceNotAvaliable{}) {
			resp.Err = errWithCode.Error()
			resp.ErrCode = errWithCode.ErrCode()
			c.Status(500).JSON(resp)
			return nil
		} else if returned.Err != nil {

			resp.Err = errWithCode.Error()
			resp.ErrCode = errWithCode.ErrCode()
			c.Status(400).JSON(resp)
			return nil
		}
		resp.AccessToken = returned.AccessToken
		resp.RefreshToken = returned.RefreshToken
		resp.ErrCode = errsutil.SuccessCode
		c.Status(200).JSON(resp)

		return nil
	}
}
