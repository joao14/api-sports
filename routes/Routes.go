package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joao14/api-sports.git/controllers"
	"github.com/joao14/api-sports.git/services"
)

var (
	playerService    services.PlayerService       = services.New()
	playerController controllers.PlayerController = controllers.New(playerService)
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Logger())
	grp1 := r.Group("api/v1/players")
	{
		grp1.POST("/add", func(ctx *gin.Context) {
			err := playerController.Save(ctx)
			if err != nil {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Player created successfully"})
			}
		})

		grp1.POST("/players", func(ctx *gin.Context) {
			ctx.JSON(200, playerController.FindAll())
		})
	}
	return r
}
