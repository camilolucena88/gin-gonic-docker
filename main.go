package main

import (
	"github.com/camilolucena88/gin-gonic-docker/controllers"
	"github.com/camilolucena88/gin-gonic-docker/services"
	"github.com/gin-gonic/gin"
)

var (
	levelService    services.LevelService       = services.New()
	levelController controllers.LevelController = controllers.New(levelService)
)

func main() {
	router := gin.Default()

	levels := router.Group("/levels")
	{
		levels.GET("", func(ctx *gin.Context) {
			levelController.FindAll(ctx)
		})
		levels.GET("/:id", func(ctx *gin.Context) {
			levelController.FindOne(ctx)
		})
		levels.POST("", func(ctx *gin.Context) {
			levelController.Save(ctx)
		})
		levels.PUT("/:id", func(ctx *gin.Context) {
			levelController.Update(ctx)
		})
		levels.DELETE("/:id", func(ctx *gin.Context) {
			levelController.Delete(ctx)
		})
	}
	router.Run()
}
