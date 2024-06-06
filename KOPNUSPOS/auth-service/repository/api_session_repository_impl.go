package repository

import (
	"auth-service/entity"

	"github.com/sirupsen/logrus"
)

type ApiSessionRepository struct {
	Repository[entity.ApiSession]
	Log *logrus.Logger
}

func NewApiSessionRepository(log *logrus.Logger) *ApiSessionRepository {
	return &ApiSessionRepository{
		Log: log,
	}
}
