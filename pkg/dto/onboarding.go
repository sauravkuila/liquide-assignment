package dto

import (
	"database/sql"
	"time"
)

type UserSignupRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	UserType string `json:"type" binding:"required,oneof=USER ADMIN"`
	Email    string `json:"email" binding:"required"`
	Mobile   string `json:"mobile" binding:"required"`
}

func (obj *UserSignupRequest) ToUserDetails() UserDetail {
	return UserDetail{
		UserName:     obj.UserName,
		UserPassword: obj.Password,
		UserType:     obj.UserType,
		Email:        obj.Email,
		Mobile:       obj.Mobile,
	}
}

type UserSignupResponse struct {
	Data *UserSignup `json:"data,omitempty"`
	CommonResponse
}

type UserSignup struct {
	UserName string `json:"username"`
	UserId   int64  `json:"userId"`
}

type UserLoginRequest struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserLoginResponse struct {
	Data *UserLogin `json:"data,omitempty"`
	CommonResponse
}

type UserLogin struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
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
