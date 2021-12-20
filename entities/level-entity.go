package entities

import "go.mongodb.org/mongo-driver/bson/primitive"

type Level struct {
	Id    primitive.ObjectID
	Level [][]int `json:"level" binding:"gt=0,dive,gt=0,dive,gte=0,lte=2"`
}
