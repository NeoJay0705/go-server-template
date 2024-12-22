package rest

import "github.com/gin-gonic/gin"

type UserHttpHandler interface {
	GetUser(c *gin.Context)
}
