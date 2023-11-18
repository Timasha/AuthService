package api

import (
	"AuthService/internal/cases"
	"AuthService/internal/utils/errsutil"
	"errors"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func (a *Auth) GetAuthorizeMiddleware() fiber.Handler {
	return func(c *fiber.Ctx) error {

		var (
			resp BaseResponse
		)

		authHeader, ok := c.GetReqHeaders()["Authorization"]
		if !ok{
			err := ErrWrongAuthorizationMethod
			resp.Err = err.Error()
			resp.ErrCode = err.ErrorCode()

			data, _ := a.bodySerializer.Marshal(resp)

			c.Status(403).Write(data)

			return nil 
		}

		authorizeHeader := strings.Split(authHeader, " ")
		if len(authorizeHeader) != 2 || authorizeHeader[0] != "Bearer" {
			err := ErrWrongAuthorizationMethod
			resp.Err = err.Error()
			resp.ErrCode = err.ErrorCode()

			data, _ := a.bodySerializer.Marshal(resp)

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

			data, _ := a.bodySerializer.Marshal(resp)

			c.Write(data)
			return nil
		}
		c.Locals("userId", returned.UserId)
		return c.Next()
	}
}
