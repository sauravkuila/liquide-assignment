package dto

import (
	e "liquide-assignment/pkg/errors"
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
	Data    *UserSignup `json:"data,omitempty"`
	Status  bool        `json:"success"`
	Errors  []e.Error   `json:"errors,omitempty"`
	Message string      `json:"message,omitempty"`
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
	Data    *UserLogin `json:"data,omitempty"`
	Status  bool       `json:"success"`
	Errors  []e.Error  `json:"errors,omitempty"`
	Message string     `json:"message,omitempty"`
}

type UserLogin struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Expiry       string `json:"expiry"`
}
