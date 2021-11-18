package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/joao14/api-sports.git/model"
	"github.com/joao14/api-sports.git/services"
)

type PlayerController interface {
	Save(ctx *gin.Context) error
	FindAll() []model.Player
}

type playerController struct {
	service services.PlayerService
}

func New(services services.PlayerService) PlayerController {
	return &playerController{
		service: services,
	}
}

func (c *playerController) Save(ctx *gin.Context) error {
	var player model.Player
	error := ctx.ShouldBindJSON(&player)
	if error != nil {
		return error
	}
	c.service.Save(player)
	return nil
}

func (c *playerController) FindAll() []model.Player {
	return c.service.FindAll()
}
