package controllers

import (
	"github.com/camilolucena88/docker-go-gin/entities"
	"github.com/camilolucena88/docker-go-gin/services"
	"github.com/gin-gonic/gin"
)

type LevelController interface {
	FindOne(ctx *gin.Context) entities.Level
	Save(ctx *gin.Context) entities.Level
	Update(ctx *gin.Context) entities.Level
	Delete(ctx *gin.Context) []entities.Level
}

type controller struct {
	services services.LevelService
}

func New(services services.LevelService) LevelController {
	return &controller{
		services: services,
	}
}

func (c *controller) FindOne(ctx *gin.Context) entities.Level {
	var level entities.Level
	err := ctx.BindJSON(&level)
	if err != nil {
		return entities.Level{}
	}
	return c.services.FindOne(level)
}
func (c *controller) Save(ctx *gin.Context) entities.Level {
	var level entities.Level
	err := ctx.BindJSON(&level)
	if err != nil {
		return entities.Level{}
	}
	return c.services.Save(level)
}
func (c *controller) Update(ctx *gin.Context) entities.Level {
	var level entities.Level
	err := ctx.BindJSON(&level)
	if err != nil {
		return entities.Level{}
	}
	return c.services.Update(level)
}
func (c *controller) Delete(ctx *gin.Context) []entities.Level {
	var level entities.Level
	return c.services.Delete(level)
}
