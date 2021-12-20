package main

import (
	"github.com/camilolucena88/gin-gonic-docker/controllers"
	"github.com/camilolucena88/gin-gonic-docker/repositories"
	"github.com/camilolucena88/gin-gonic-docker/services"
	"github.com/gin-gonic/gin"
)

var (
	levelRepository repositories.LevelRepository = repositories.New()
	levelService    services.LevelService        = services.New(levelRepository)
	levelController controllers.LevelController  = controllers.New(levelService)
)

func setupRouter() *gin.Engine {
	router := gin.Default()
	levels := router.Group("/levels")
	{
		levels.POST("", func(ctx *gin.Context) {
			levelController.Create(ctx)
		})
		levels.GET("", func(ctx *gin.Context) {
			levelController.FindAll(ctx)
		})
		levels.GET("/:id", func(ctx *gin.Context) {
			levelController.FindOne(ctx)
		})

		levels.PUT("/:id", func(ctx *gin.Context) {
			levelController.Update(ctx)
		})
		levels.DELETE("/:id", func(ctx *gin.Context) {
			levelController.Delete(ctx)
		})
	}
	return router
}

func main() {
	router := setupRouter()
	router.Run()
}
