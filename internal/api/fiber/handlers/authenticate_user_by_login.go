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

		var args iomodels.AuthenticateUserByLoginArgs = iomodels.AuthenticateUserByLoginArgs{
			Ctx: h.ctx,
			Login: req.Login,
			Password: req.Password,
		}

		returned := h.casesProvider.AuthenticateUserByLogin(args)

		var authErrWithCode errsutil.AuthErr

		errors.As(returned.Err, &authErrWithCode)

		if returned.Err == (errs.ErrServiceInternal{}) || returned.Err == (errs.ErrServiceNotAvaliable{}) {
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

		if (returned.OtpEnabled){
			resp.OtpEnabled = returned.OtpEnabled
			resp.IntermediateToken = returned.IntermediateToken
		}else{
			resp.AuthInfo.AccessToken = returned.AuthInfo.AccessToken
			resp.AuthInfo.RefreshToken = returned.AuthInfo.RefreshToken
		}

		resp.ErrCode = errsutil.SuccessCode

		c.Status(200).JSON(resp)

		return nil
	}
}
