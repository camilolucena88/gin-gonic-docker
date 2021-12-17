package controllers

import (
	"github.com/camilolucena88/gin-gonic-docker/entities"
	"github.com/camilolucena88/gin-gonic-docker/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type LevelController interface {
	FindOne(ctx *gin.Context)
	FindAll(ctx *gin.Context)
	Save(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	services services.LevelService
}

func New(services services.LevelService) LevelController {
	return &controller{
		services: services,
	}
}

func (c *controller) FindOne(ctx *gin.Context) {
	var level entities.Level
	var id uint64
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	level, err = c.services.FindOne(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, level)
	}
}
func (c *controller) Save(ctx *gin.Context) {
	var level entities.Level
	err := ctx.ShouldBindJSON(&level)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusCreated, c.services.Save(level))
	}
}
func (c *controller) Update(ctx *gin.Context) {
	var level entities.Level
	var id uint64
	err := ctx.BindJSON(&level)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id, err = strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	level, err = c.services.Update(id, level)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusCreated, level)
	}
}

func (c *controller) Delete(ctx *gin.Context) {
	var levels []entities.Level
	var id uint64
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	levels, err = c.services.Delete(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, levels)
	}
}

func (c *controller) FindAll(ctx *gin.Context) {
	ctx.IndentedJSON(200, c.services.FindAll())
}
