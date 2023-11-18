package api

import (
	"AuthService/internal/cases"
	"AuthService/internal/utils/errsutil"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type EnableOtpAuthenticationForUserRequest struct{

}

type EnableOtpAuthenticationForUserResponse struct{
	BaseResponse
	OtpKey string `json:"otpKey"`
	OtpUrl string `json:"otpUrl"`
}

func (a *Auth) GetEnableOtpAuthenticationForUserHandler()fiber.Handler{
	return func(c *fiber.Ctx) error {
		var (
			resp EnableOtpAuthenticationForUserResponse
		)

		args := cases.EnableOtpAuthenticationForUserArgs{
			Ctx: a.ctx,
			UserId: c.Locals("userId").(string),
		}

		returned := a.casesProvider.EnableOtpAuthenticationForUser(args)

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

			data, _ := a.bodySerializer.Marshal(resp)

			c.Write(data)
			return nil
		}

		resp.OtpKey = returned.OtpKey
		resp.OtpUrl = returned.OtpUrl
		resp.ErrCode = errsutil.SuccessCode

		data, _ := a.bodySerializer.Marshal(resp)

		c.Status(200).Write(data)
		return nil
	}
}