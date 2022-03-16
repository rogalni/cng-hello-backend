package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/rogalni/cng-hello-backend/internal/adapter/db/postgres/model"
	"github.com/rogalni/cng-hello-backend/internal/adapter/rest/v1/dto"
	"github.com/rogalni/cng-hello-backend/internal/service"
	"github.com/rogalni/cng-hello-backend/pkg/gin/errors"
)

func SetupMessages(g *gin.RouterGroup) {
	messages := g.Group("/messages")
	{
		messages.GET("/", getMessages)
		messages.GET("/:id", getMessage)
		messages.POST("/", createMessage)
		messages.DELETE("/:id", deleteMessage)
	}
}

func getMessages(c *gin.Context) {
	ms := service.NewMessageService()
	m, err := ms.GetMessages(c.Request.Context())
	if err != nil {
		errors.Handle(c, err)
		return
	}
	c.IndentedJSON(200, toDtos(m))
}

func getMessage(c *gin.Context) {
	ms := service.NewMessageService()
	id := c.Param("id")
	m, err := ms.GetMessage(c.Request.Context(), id)
	if err != nil {
		errors.Handle(c, err)
		return
	}

	c.IndentedJSON(200, &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	})
}

func createMessage(c *gin.Context) {
	ms := service.NewMessageService()
	var m *model.Message
	c.Bind(&m)
	err := ms.CreateMessage(c.Request.Context(), m)
	if err != nil {
		errors.Handle(c, err)
		return
	}
	c.Status(204)
}

func deleteMessage(c *gin.Context) {
	ms := service.NewMessageService()
	id := c.Param("id")
	err := ms.DeleteMessage(c.Request.Context(), id)
	if err != nil {
		errors.Handle(c, err)
		return
	}
	c.Status(204)
}

func toDto(m *model.Message) *dto.Message {
	return &dto.Message{
		Id:   m.Id,
		Code: m.Code,
		Text: m.Text,
	}
}

func toDtos(mm []*model.Message) (dtoM []dto.Message) {
	dtoM = make([]dto.Message, 0)
	for _, m := range mm {
		dtoM = append(dtoM, *toDto(m))
	}
	return
}
