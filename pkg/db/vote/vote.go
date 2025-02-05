package vote

import "gorm.io/gorm"

type voteDb struct {
	dbObj *gorm.DB
}

type DbVoteInterface interface {
}

func NewVoteDbObject(db *gorm.DB) DbVoteInterface {
	return &voteDb{
		dbObj: db,
	}
}
