package rest

import (
	"fmt"
	"log"
	"strconv"

	"github.com/NeoJay0705/go-server-template/internal/service"
	"github.com/gin-gonic/gin"
)

type UserHttpHandlerImpl struct {
	svc service.UserService
}

func NewUserHttpHandler(svc service.UserService) UserHttpHandler {
	return &UserHttpHandlerImpl{svc: svc}
}

func (g *UserHttpHandlerImpl) GetUser(c *gin.Context) {
	//將 middleware 傳入的 i18n 進行轉換
	id, err := strconv.ParseInt(c.Params.ByName("id"), 10, 64)
	if err != nil {
		// return g.i18nMsg.GetError(err)
		log.Println(fmt.Errorf("error parsing id"))
	}
	if result, err := g.svc.GetUser(id); err != nil {
		// return g.i18nMsg.GetError(err)
		log.Println(fmt.Errorf("error getting user"))
	} else {
		// c.JSON(g.i18nMsg.GetSuccess(result))
		c.JSON(200, result)
	}
}
