package repositories

import (
	"context"
	"github.com/camilolucena88/gin-gonic-docker/entities"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type LevelRepository interface {
	Create(level entities.Level) (primitive.ObjectID, error)
	Update(id primitive.ObjectID, level entities.Level) (entities.Level, error)
	Delete(id primitive.ObjectID) error
	FindOne(id primitive.ObjectID) (entities.Level, error)
	FindAll() ([]entities.Level, error)
}

type database struct {
	connection *mongo.Client
	context    context.Context
}

const uri = "mongodb://root:example@db:27017/?maxPoolSize=20&w=majority"

func New() LevelRepository {
	ctx, _ := context.WithTimeout(context.Background(), 100*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}
	if err != nil {
		log.Printf("Could not create Task: %v", err)
	}
	return &database{
		connection: client,
		context:    ctx,
	}
}

func (db *database) Create(level entities.Level) (primitive.ObjectID, error) {
	level.Id = primitive.NewObjectID()
	coll := db.connection.Database("levels").Collection("levels")
	result, err := coll.InsertOne(db.context, level)
	if err != nil {
		log.Printf("Could not create Task: %v", err)
		return primitive.NilObjectID, err
	}
	oid := result.InsertedID.(primitive.ObjectID)
	return oid, nil
}

func (db *database) Update(id primitive.ObjectID, level entities.Level) (entities.Level, error) {
	filter := bson.D{{"_id", id}}
	coll := db.connection.Database("levels").Collection("levels")
	update := bson.D{{"$set", bson.D{{"level", level.Level}}}}
	_, err := coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return level, nil
	}
	return level, nil
}

func (db *database) Delete(id primitive.ObjectID) error {
	coll := db.connection.Database("levels").Collection("levels")
	filter := bson.D{{"_id", id}}
	_, err := coll.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (db *database) FindOne(id primitive.ObjectID) (entities.Level, error) {
	var level entities.Level
	filter := bson.D{{"_id", id}}
	cursor := db.connection.Database("levels").Collection("levels").FindOne(db.context, filter)
	err := cursor.Decode(&level)
	defer db.connection.Disconnect(db.context)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return entities.Level{}, err
		}
		return entities.Level{}, err
	}
	return level, nil
}

func (db *database) FindAll() ([]entities.Level, error) {
	var levels []entities.Level
	filter := bson.D{}
	cursor, err := db.connection.Database("levels").Collection("levels").Find(db.context, filter)
	if err != nil {
		return levels, err
	}
	defer cursor.Close(db.context)
	for cursor.Next(db.context) {
		var level entities.Level
		cursor.Decode(&level)
		levels = append(levels, level)
	}
	if cursor.Err(); err != nil {
		return levels, err
	}
	return levels, err
}
