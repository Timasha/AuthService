package api

import (
	"auth/internal/cases"
	"auth/internal/utils/errsutil"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type AuthenticateUserByLoginRequest struct {
	Login    string `json:"login" validate:"required"`
	Password string `json:"password" validate:"required"`
}
type AuthenticateUserByLoginResponse struct {
	BaseResponse

	OtpEnabled bool `json:"otpEnabled"`

	IntermediateToken string `json:"IntermediateToken"`

	AuthInfo struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshtoken"`
	} `json:"authInfo"`
}

func (a *Auth) GetAuthenticateUserByLoginHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			req  AuthenticateUserByLoginRequest
			resp AuthenticateUserByLoginResponse
		)

		unmarshErr := json.Unmarshal(c.Request().Body(), &req)

		if unmarshErr != nil {
			resp.Err = "unmarshal request error: " + unmarshErr.Error()
			resp.ErrCode = errsutil.ErrInputCode

			data, _ := json.Marshal(resp)

			c.Status(400).Write(data)
			return nil
		}

		var args cases.AuthenticateUserByLoginArgs = cases.AuthenticateUserByLoginArgs{
			Ctx:      a.ctx,
			Login:    req.Login,
			Password: req.Password,
		}

		returned := a.casesProvider.AuthenticateUserByLogin(args)

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

		if returned.OtpEnabled {
			resp.OtpEnabled = returned.OtpEnabled
			resp.IntermediateToken = returned.IntermediateToken
		} else {
			resp.AuthInfo.AccessToken = returned.AuthInfo.AccessToken
			resp.AuthInfo.RefreshToken = returned.AuthInfo.RefreshToken
		}

		resp.ErrCode = errsutil.SuccessCode

		data, _ := json.Marshal(resp)

		c.Status(200).Write(data)

		return nil
	}
}
