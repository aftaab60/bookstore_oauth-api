package access_token

import (
	"fmt"
	"github.com/aftaab60/bookstore_oauth-api/src/utils/crypto_utils"
	"github.com/aftaab60/bookstore_oauth-api/src/utils/errors"
	"strings"
	"time"
)

const (
	expirationTime             = 24
	grantTypePassword          = "password"
	grandTypeClientCredentials = "client_credentials"
)

type AccessTokenRequest struct {
	GrantType string `json:"grant_type"`
	Scope string `json:"scope"`

	//email, password based authentication
	Email string `json:"email"`
	Password string `json:"password"`

	//clientId and secret based authentication
	ClientId string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
}

func (atRequest *AccessTokenRequest) Validate() *errors.RestErr{
	if atRequest.GrantType == grantTypePassword || atRequest.GrantType == grandTypeClientCredentials {
		return nil
	}
	return errors.NewBadRequestError("invalid grand type in request")
}

type AccessToken struct {
	AccessToken string `json:"access_token"`
	UserId      int64  `json:"user_id"`
	ClientId    int64  `json:"client_id"`
	Expires     int64  `json:"expires"`
}

func (at *AccessToken) Validate() *errors.RestErr{
	at.AccessToken = strings.TrimSpace(at.AccessToken)
	if at.AccessToken == "" {
		return errors.NewBadRequestError("Invalid token id")
	}
	if at.UserId <=0 {
		return errors.NewBadRequestError("Invalid user id")
	}
	if at.ClientId <= 0 {
		return errors.NewBadRequestError("Invalid client id")
	}
	if at.Expires <= 0 {
		return errors.NewBadRequestError("Invalid expiration time")
	}
	return nil
}

func GetNewAccessToken(id int64) AccessToken {
	return AccessToken{
		UserId: id,
		Expires: time.Now().UTC().Add(expirationTime * time.Hour).Unix(),
	}
}

func (at AccessToken) IsExpired() bool {
	return time.Unix(at.Expires, 0).Before(time.Now().UTC())
}

func (at *AccessToken) Generate() {
	at.AccessToken = crypto_utils.GetMd5(fmt.Sprintf("at-%d-%d-ran", at.UserId, at.Expires))
}
