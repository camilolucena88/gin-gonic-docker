package services

import (
	"errors"
	"github.com/camilolucena88/gin-gonic-docker/entities"
)

type LevelService interface {
	Save(level entities.Level) entities.Level
	FindOne(id uint64) (entities.Level, error)
	FindAll() []entities.Level
	Update(id uint64, level entities.Level) (entities.Level, error)
	Delete(id uint64) ([]entities.Level, error)
}

type levelService struct {
	levels []entities.Level
}

func New() LevelService {
	return &levelService{}
}

func (services *levelService) Save(level entities.Level) entities.Level {
	services.levels = append(services.levels, level)
	return level
}

func (services *levelService) FindAll() []entities.Level {
	return services.levels
}

func (services *levelService) FindOne(id uint64) (entities.Level, error) {
	for i, value := range services.levels {
		if id == value.Id {
			return services.levels[i], nil
		}
	}
	var level entities.Level
	return level, errors.New("not found")
}

func (services *levelService) Update(id uint64, level entities.Level) (entities.Level, error) {
	for i, value := range services.levels {
		if id == value.Id {
			services.levels[i].Level = level.Level
			return services.levels[i], nil
		}
	}
	return level, errors.New("not found")
}

func remove(levels []entities.Level, index uint64) []entities.Level {
	return append(levels[:index], levels[index+1:]...)
}

func (services *levelService) Delete(id uint64) ([]entities.Level, error) {
	for _, value := range services.levels {
		if id == value.Id {
			return remove(services.levels, id), nil
		}
	}
	return remove(services.levels, id), nil
}
