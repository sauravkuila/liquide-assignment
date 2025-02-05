package comment

import (
	"gorm.io/gorm"
)

type commentDb struct {
	dbObj *gorm.DB
}

type DbCommentInterface interface {
}

func NewCommentDbObject(db *gorm.DB) DbCommentInterface {
	return &commentDb{
		dbObj: db,
	}
}
