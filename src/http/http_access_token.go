package http

import (
	atDomain "github.com/aftaab60/bookstore_oauth-api/src/domain/access_token"
	"github.com/aftaab60/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
	Create(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	accessToken, err := handler.service.GetById(strings.TrimSpace(c.Param("access_token_id")))
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusOK, accessToken)
}

func (handler *accessTokenHandler) Create(c *gin.Context) {
	var accessTokenRequest atDomain.AccessTokenRequest
	if err := c.ShouldBindJSON(&accessTokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}
	accessToken, err := handler.service.Create(accessTokenRequest)
	if err != nil {
		c.JSON(err.Status, err)
		return
	}
	c.JSON(http.StatusCreated, accessToken)
}

func NewAccessTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}