package service

import (
	"liquide-assignment/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

type service struct {
}

type ServiceGroupLayer interface {
	Health(*gin.Context)
}

func NewServiceGroupObject(db db.DBLayer) ServiceGroupLayer {
	return &service{}
}

func (s *service) Health(c *gin.Context) {
	c.String(http.StatusOK, "I am Healthy")
}
