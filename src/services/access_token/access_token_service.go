package access_token

import (
	"github.com/aftaab60/bookstore_oauth-api/src/domain/access_token"
	"github.com/aftaab60/bookstore_oauth-api/src/repository/db"
	"github.com/aftaab60/bookstore_oauth-api/src/repository/rest"
	"github.com/aftaab60/bookstore_oauth-api/src/utils/errors"
)

type Service interface {
	GetById(string) (*access_token.AccessToken, *errors.RestErr)
	Create(access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr)
	UpdateExpirationTime(access_token.AccessToken) *errors.RestErr
}

type service struct {
	restUserRepo rest.RestUserRepository
	dbRepo db.DbRepository
}

func NewService(userRepo rest.RestUserRepository, dbRepo db.DbRepository) Service {
	return &service{
		restUserRepo: userRepo,
		dbRepo: dbRepo,
	}
}

func (s *service) GetById(accessTokenId string) (*access_token.AccessToken, *errors.RestErr) {
	accessToken, err := s.dbRepo.GetById(accessTokenId)
	if err != nil {
		return nil, err
	}
	return accessToken, nil
}

func (s *service) Create(request access_token.AccessTokenRequest) (*access_token.AccessToken, *errors.RestErr) {
	if err := request.Validate(); err != nil {
		return nil, err
	}
	//TODO: support all grandtypes login.
	user, err := s.restUserRepo.LoginUser(request.Email, request.Password)
	if err != nil {
		return nil, err
	}

	at := access_token.GetNewAccessToken(user.Id)
	at.Generate()
	if err := s.dbRepo.Create(at); err != nil {
		return nil, errors.NewInternalServerError(err.Error)
	}
	return &at, nil
}

func (s *service) UpdateExpirationTime(at access_token.AccessToken) *errors.RestErr {
	return nil
}