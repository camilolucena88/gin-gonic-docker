package main

import (
	"github.com/camilolucena88/docker-go-gin/controllers"
	"github.com/camilolucena88/docker-go-gin/services"
	"github.com/gin-gonic/gin"
)

var (
	levelService    services.LevelService       = services.New()
	videoController controllers.LevelController = controllers.New(levelService)
)

func main() {
	router := gin.Default()
	router.GET("/levels", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.FindOne(ctx))
	})
	router.POST("/levels", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Save(ctx))
	})
	router.PUT("/levels", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Update(ctx))
	})
	router.DELETE("/levels", func(ctx *gin.Context) {
		ctx.JSON(200, videoController.Delete(ctx))
	})
	router.Run()
}
