package services

import "github.com/camilolucena88/gin-gonic-docker/entities"

type LevelService interface {
	Save(level entities.Level) entities.Level
	FindOne(level entities.Level) entities.Level
	Update(level entities.Level) entities.Level
	Delete(level entities.Level) []entities.Level
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

func (services *levelService) FindOne(level entities.Level) entities.Level {
	return services.levels[0]
}

func (services *levelService) Update(level entities.Level) entities.Level {
	return services.levels[0]
}

func (services *levelService) Delete(level entities.Level) []entities.Level {
	return services.levels
}
