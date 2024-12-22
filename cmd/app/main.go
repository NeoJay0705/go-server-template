package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/NeoJay0705/go-server-template/internal/controller"
	"github.com/NeoJay0705/go-server-template/pkg/logger"
	"github.com/NeoJay0705/go-server-template/pkg/middleware"
	"github.com/gin-gonic/gin"
	"gopkg.in/olahol/melody.v1"
)

type IndexData struct {
	Title   string
	Content string
}

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一個首頁"
	c.HTML(http.StatusOK, "index.html", data)
}
func main() {
	server := gin.Default()
	server.LoadHTMLGlob("../../web/template/**/*.html")
	// /assets in client url is a mapping to directory ../../web/template/assets
	server.Static("/assets", "../../web/template/assets")

	logger.InitLogger()
	defer logger.Logger.Sync() // Flush logs before exiting
	server.Use(middleware.RequestLogger())

	server.GET("/", test)
	server.GET("/login", controller.LoginPage)
	server.POST("/login", controller.LoginAuth)

	// userRepo := repository.NewUserTemplate(db.Conn.Pgdb)
	// userService := service.NewUserService(userRepo)
	// userHandler := rest.NewUserHttpHandler(userService)
	// server.GET("/v1/user/:id", userHandler.GetUser)

	server.GET("/chat", func(c *gin.Context) {
		c.HTML(http.StatusOK, "chat.html", nil)
	})
	m := melody.New()
	server.GET("/ws/chat", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	m.HandleConnect(func(session *melody.Session) {
		id := session.Request.URL.Query().Get("id")
		m.Broadcast(NewMessage("other", id, "加入聊天室").GetByteMessage())
	})

	m.HandleClose(func(session *melody.Session, i int, s string) error {
		id := session.Request.URL.Query().Get("id")
		m.Broadcast(NewMessage("other", id, "離開聊天室").GetByteMessage())
		return nil
	})

	server.Run(":8888")
}

type Message struct {
	Event     string    `json:"event"`
	Name      string    `json:"name"`
	Content   string    `json:"content"`
	TimeStamp time.Time `json:"timestamp"`
}

func NewMessage(event, name, content string) *Message {
	return &Message{
		Event:     event,
		Name:      name,
		Content:   content,
		TimeStamp: time.Now(),
	}
}

func (m *Message) GetByteMessage() []byte {
	result, _ := json.Marshal(m)
	return result
}
