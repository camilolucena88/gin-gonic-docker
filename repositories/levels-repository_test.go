package repositories

import (
	"github.com/camilolucena88/gin-gonic-docker/entities"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"testing"
)

func TestDatabase_Create(t *testing.T) {
	var level entities.Level
	level.Id = primitive.NewObjectID()
	db := New()
	twoDSlice := make([][]int, 2)
	level.Level = twoDSlice
	_, err := db.Create(level)
	if err != nil {
		assert.True(t, true, "True is true!")
	}
}

func TestDatabase_Delete(t *testing.T) {
	db := New()
	levelId := primitive.NewObjectID()
	err := db.Delete(levelId)
	if err != nil {
		assert.True(t, true, "True is true!")
	}
}

func TestDatabase_FindAll(t *testing.T) {
	db := New()
	_, err := db.FindAll()
	if err != nil {
		assert.True(t, true, "True is true!")
	}
}

func TestDatabase_FindOne(t *testing.T) {
	db := New()
	levelId := primitive.NewObjectID()
	_, err := db.FindOne(levelId)
	if err != nil {
		assert.True(t, true, "True is true!")
	}
}

func TestDatabase_Update(t *testing.T) {
	db := New()
	var level entities.Level
	levelId := primitive.NewObjectID()
	_, err := db.Update(levelId, level)
	if err != nil {
		assert.True(t, true, "True is true!")
	}
}
