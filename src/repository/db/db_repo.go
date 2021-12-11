package db

import (
	"github.com/aftaab60/bookstore_oauth-api/src/clients/cassandra"
	"github.com/aftaab60/bookstore_oauth-api/src/domain/access_token"
	"github.com/aftaab60/bookstore_oauth-api/src/utils/errors"
	"github.com/gocql/gocql"
)

var (
	queryGetAccessToken = "SELECT access_token, user_id, client_id, expires from access_tokens where access_token=?;"
	queryCreateAccessToken = "INSERT into access_tokens(access_token, user_id, client_id, expires) VALUES (?,?,?,?);"
	queryUpdateExpirationTime = "UPDATE access_tokens set expires=? where access_token=?;"
)

func NewRepository() DbRepository{
	return &dbRepository{}
}

type DbRepository interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessToken) *errors.RestErr
	UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr
}

type dbRepository struct {
}

func (r *dbRepository) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	var result access_token.AccessToken
	if err := cassandra.GetSession().Query(queryGetAccessToken, accessTokenId).Scan(&result.AccessToken, &result.UserId, &result.ClientId, &result.Expires); err != nil {
		if err == gocql.ErrNotFound {
			return nil, errors.NewNotFoundError("No access token with given token id")
		}
	}
	return &result, nil
}

func (r *dbRepository) Create(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryCreateAccessToken, at.AccessToken, at.UserId, at.ClientId, at.Expires).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

func (r *dbRepository) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	if err := cassandra.GetSession().Query(queryUpdateExpirationTime, at.Expires, at.AccessToken).Exec(); err != nil {
		return errors.NewInternalServerError(err.Error())
	}
	return nil
}

