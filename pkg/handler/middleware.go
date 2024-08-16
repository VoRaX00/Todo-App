package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	userCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.Request.Header.Get(authorizationHeader)

	if header == "" {
		NewErrorResponse(c, http.StatusUnauthorized, "No authorization header")
		return
	}

	headerParts := strings.SplitN(header, " ", 2)
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		NewErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
		return
	}

	userId, err := h.service.Authorization.ParseToken(headerParts[1])
	if err != nil {
		NewErrorResponse(c, http.StatusUnauthorized, "Invalid authorization header")
	}

	c.Set(userCtx, userId)
}

func getUserId(c *gin.Context) (int, error) {
	id, ok := c.Get(userCtx)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user is not found")
		return 0, errors.New("user is not found")
	}

	idInt, ok := id.(int)
	if !ok {
		NewErrorResponse(c, http.StatusInternalServerError, "user id is not int")
		return 0, errors.New("user id is not int")
	}

	return idInt, nil
}
