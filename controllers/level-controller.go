package controllers

import (
	"github.com/camilolucena88/gin-gonic-docker/entities"
	"github.com/camilolucena88/gin-gonic-docker/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type LevelController interface {
	Create(ctx *gin.Context)
	FindOne(ctx *gin.Context)
	FindAll(ctx *gin.Context)
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

func (c *controller) Create(ctx *gin.Context) {
	var level entities.Level
	err := ctx.ShouldBindJSON(&level)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	} else {
		id, _ := c.services.Create(level)
		ctx.IndentedJSON(http.StatusCreated, gin.H{"id": id})
	}
}

func (c *controller) Update(ctx *gin.Context) {
	var level entities.Level
	var id primitive.ObjectID
	err := ctx.BindJSON(&level)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	id, err = primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	level.Id = id
	level, err = c.services.Update(id, level)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, level)
	}
}

func (c *controller) Delete(ctx *gin.Context) {
	var id primitive.ObjectID
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	err = c.services.Delete(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	}
	ctx.IndentedJSON(http.StatusAccepted, gin.H{"deleted": true})
}

func (c *controller) FindAll(ctx *gin.Context) {
	levels, err := c.services.FindAll()
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, levels)
	}
}

func (c *controller) FindOne(ctx *gin.Context) {
	var level entities.Level
	err := ctx.BindQuery(&level)
	var id primitive.ObjectID
	id, err = primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	level, err = c.services.FindOne(id)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusOK, level)
	}
}
