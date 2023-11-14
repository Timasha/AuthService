package api

import (
	"auth/internal/cases"
	"auth/internal/utils/errsutil"
	"encoding/json"
	"errors"

	"github.com/gofiber/fiber/v2"
)

type RefreshTokensRequest struct {
	RefreshToken string `json:"refreshToken"`
	AccessToken  string `json:"accessToken"`
}
type RefreshTokensResponse struct {
	BaseResponse

	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

func (a *Auth) GetRefreshTokensHandler() fiber.Handler {
	return func(c *fiber.Ctx) error {
		var (
			req  RefreshTokensRequest
			resp RefreshTokensResponse
		)

		unmarshalErr := json.Unmarshal(c.Body(), &req)

		if unmarshalErr != nil {
			resp.Err = "Unmarshal json error: " + unmarshalErr.Error()
			resp.ErrCode = errsutil.ErrInputCode

			data, _ := json.Marshal(resp)

			c.Status(400).Write(data)
			return nil
		}

		var args cases.RefreshTokensArgs = cases.RefreshTokensArgs{
			Ctx:          a.ctx,
			AccessToken:  req.AccessToken,
			RefreshToken: req.RefreshToken,
		}

		returned := a.casesProvider.RefreshTokens(args)

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

		resp.AccessToken = returned.AccessToken
		resp.RefreshToken = returned.RefreshToken
		resp.ErrCode = errsutil.SuccessCode

		data, _ := json.Marshal(resp)

		c.Status(200).Write(data)

		return nil
	}
}