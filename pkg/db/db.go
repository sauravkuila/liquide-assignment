package db

import "gorm.io/gorm"

type dbService struct {
}

type DBLayer interface {
}

func NewDBObject(psqlDB *gorm.DB) DBLayer {
	temp := &dbService{}
	return temp
}
