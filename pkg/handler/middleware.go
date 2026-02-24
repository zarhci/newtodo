package handler

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

const (
	AuthorizationHeader = "Authorization"
	UserCtx             = "userId"
)

func (h *Handler) userIdentity(c *gin.Context) {
	header := c.GetHeader(AuthorizationHeader)
	if header == "" {
		newErrorResponse(c, http.StatusUnauthorized, "empty suth header")
		return
	}
	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 {
		newErrorResponse(c, http.StatusUnauthorized, "invalid auth header")
		return
	}

	userId, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		newErrorResponse(c, http.StatusUnauthorized, err.Error())
		return
	}
	c.Set(UserCtx, userId)
}

func getUserId(c *gin.Context) (int, bool) {
	id, ok := c.Get(UserCtx)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user не найден")
		return 0, false
	}

	idInt, ok := id.(int)
	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id не является числом")
		return 0, false
	}
	return idInt, true
}
