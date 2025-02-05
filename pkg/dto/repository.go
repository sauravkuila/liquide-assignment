package dto

import (
	"database/sql"
	"time"
)

type DbUserDetail struct {
	UserId       sql.NullInt64
	UserName     sql.NullString
	UserPassword sql.NullString
	UserType     sql.NullString
	Email        sql.NullString
	Mobile       sql.NullString
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}

func (u *DbUserDetail) ToUserDetail() UserDetail {
	return UserDetail{
		UserId:       u.UserId.Int64,
		UserName:     u.UserName.String,
		UserPassword: u.UserPassword.String,
		UserType:     u.UserType.String,
		Email:        u.Email.String,
		Mobile:       u.Mobile.String,
		CreatedAt:    u.CreatedAt.Time,
		UpdatedAt:    u.UpdatedAt.Time,
	}
}

type UserDetail struct {
	UserId       int64
	UserName     string
	UserPassword string
	UserType     string
	Email        string
	Mobile       string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

func (u *UserDetail) ToDbUserDetail() DbUserDetail {
	return DbUserDetail{
		UserId:       sql.NullInt64{Int64: u.UserId, Valid: true},
		UserName:     sql.NullString{String: u.UserName, Valid: true},
		UserPassword: sql.NullString{String: u.UserPassword, Valid: true},
		UserType:     sql.NullString{String: u.UserType, Valid: true},
		Email:        sql.NullString{String: u.Email, Valid: true},
		Mobile:       sql.NullString{String: u.Mobile, Valid: true},
		CreatedAt:    sql.NullTime{Time: u.CreatedAt, Valid: true},
		UpdatedAt:    sql.NullTime{Time: u.UpdatedAt, Valid: true},
	}
}
