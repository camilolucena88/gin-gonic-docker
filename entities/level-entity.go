package entities

type Level struct {
	Id    uint64    `gorm:"primary_key;auto_increment"`
	Level [][]uint8 `sql:"type:json" json:"level" binding:"gt=0,dive,gt=0,dive,gte=0,lte=2"`
}
