package services

import (
	"github.com/camilolucena88/gin-gonic-docker/entities"
	"github.com/camilolucena88/gin-gonic-docker/repositories"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type LevelService interface {
	Create(level entities.Level) (primitive.ObjectID, error)
	FindOne(id primitive.ObjectID) (entities.Level, error)
	FindAll() ([]entities.Level, error)
	Update(id primitive.ObjectID, level entities.Level) (entities.Level, error)
	Delete(id primitive.ObjectID) error
}
type levelService struct {
	levelRepository repositories.LevelRepository
}

func New(repo repositories.LevelRepository) LevelService {
	return &levelService{
		levelRepository: repo,
	}
}

func (services *levelService) Create(level entities.Level) (primitive.ObjectID, error) {
	var levelId primitive.ObjectID
	levelId, err := services.levelRepository.Create(level)
	if err != nil {
		return levelId, err
	}
	return levelId, err
}

func (services *levelService) FindAll() ([]entities.Level, error) {
	return services.levelRepository.FindAll()
}

func (services *levelService) FindOne(id primitive.ObjectID) (entities.Level, error) {
	return services.levelRepository.FindOne(id)
}

func (services *levelService) Update(id primitive.ObjectID, level entities.Level) (entities.Level, error) {
	return services.levelRepository.Update(id, level)
}

func (services *levelService) Delete(id primitive.ObjectID) error {
	return services.levelRepository.Delete(id)
}
