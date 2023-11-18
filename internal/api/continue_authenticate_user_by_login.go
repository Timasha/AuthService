package api

import (
	"AuthService/internal/cases"
	"AuthService/internal/utils/errsutil"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type ContinueAuthenticateOtpUserByLoginRequest struct {
	IntermediateToken string `json:"intermediateToken"`
	OtpCode           string `json:"otpCode"`
}

type ContinueAuthenticateOtpUserByLoginResponse struct {
	BaseResponse
	AuthInfo struct {
		AccessToken  string `json:"accessToken"`
		RefreshToken string `json:"refreshToken"`
	} `json:"authInfo"`
}

func (a *Auth)GetContinueAuthenticateOtpUserByLoginHandler()fiber.Handler{
	return func(c *fiber.Ctx) error {
		var(
			req ContinueAuthenticateOtpUserByLoginRequest
			resp ContinueAuthenticateOtpUserByLoginResponse
		)


		unmarshErr := a.bodySerializer.Unmarshal(c.Body(),&req)

		if unmarshErr != nil{
			resp.Err = ErrInvalidInput.Error() + unmarshErr.Error()
			resp.ErrCode = ErrInvalidInput.ErrorCode()

			data, _ := a.bodySerializer.Marshal(resp)

			c.Status(400).Write(data)
			return nil
		}

		args := cases.ContinueAuthenticateOtpUserByLoginArgs{
			Ctx: a.ctx,
			IntermediateToken: req.IntermediateToken,
			OtpCode: req.OtpCode,
		}

		returned := a.casesProvider.ContinueAuthenticateOtpUserByLogin(args)

		var errWithCode errsutil.AuthErr

		errors.As(returned.Err,&errWithCode)

		if returned.Err != nil{
			if returned.Err == cases.ErrServiceInternal{
				c.Status(500)
			}else{
				c.Status(400)
			}

			resp.Err = errWithCode.Error()
			resp.ErrCode = errWithCode.ErrorCode()

			data, _ := a.bodySerializer.Marshal(resp)

			c.Write(data)
			return nil
		}

		resp.AuthInfo.AccessToken = returned.AuthInfo.AccessToken
		resp.AuthInfo.RefreshToken = returned.AuthInfo.RefreshToken
		resp.ErrCode = errsutil.SuccessCode

		data, _ := a.bodySerializer.Marshal(resp)

		c.Status(200).Write(data)

		return nil
	}
}