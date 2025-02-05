package service

import (
	"github.com/gin-gonic/gin"
)

type service struct {
}

type ServiceGroupLayer interface {
	Health(*gin.Context)
}

func NewServiceGroupObject() ServiceGroupLayer {
	return &service{}
}
