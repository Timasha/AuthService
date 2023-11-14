package api

import (
	"auth/internal/cases"
	"auth/internal/utils/errsutil"
	"encoding/json"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) GetAuthorizeMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			resp BaseResponse
		)

		authorizeHeader := strings.Split(c.GetReqHeaders()["Authorize"], " ")
		if len(authorizeHeader) != 2 || authorizeHeader[0] != "Bearer" {
			err := ErrWrongAuthMethod
			resp.Err = err.Error()
			resp.ErrCode = err.ErrorCode()

			data, _ := json.Marshal(resp)

			c.Status(403).Write(data)

			return nil
		}
		var args cases.AuthorizeUserArgs = cases.AuthorizeUserArgs{
			Ctx:            a.ctx,
			AccessToken:    authorizeHeader[1],
			RequiredRoleId: 0,
		}
		returned := a.casesProvider.AuthorizeUser(args)

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
		return c.Next()
	}
}
