package users

import (
	"github.com/aftaab60/bookstore_oauth-api/src/utils/errors"
	"strings"
)

type User struct {
	Id int64 `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
}

type UserLoginRequest struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func (userLoginRequest *UserLoginRequest) Validate() *errors.RestErr {
	if strings.TrimSpace(userLoginRequest.Email) == "" {
		return errors.NewBadRequestError("invalid user email id")
	}
	if strings.TrimSpace(userLoginRequest.Password) == "" {
		return errors.NewBadRequestError("invalid user password")
	}
	return nil
}


