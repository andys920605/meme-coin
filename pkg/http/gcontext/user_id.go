package gcontext

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetUserIdInt64(c *gin.Context) (int64, bool) {
	userId, ok := GetUserIdString(c)
	if !ok {
		return 0, false
	}

	parsedUserId, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		return 0, false
	}

	return parsedUserId, true
}

func GetUserIdString(c *gin.Context) (string, bool) {
	userId, ok := c.Get(ContextKeyUserId)
	if !ok {
		return "", false
	}

	_userId, ok := userId.(string)
	if !ok || _userId == "" {
		return "", false
	}

	return _userId, true
}

func SetUserId(c *gin.Context, userId string) {
	c.Set(ContextKeyUserId, userId)
}
