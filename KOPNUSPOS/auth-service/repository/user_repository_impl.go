package repository

import (
	"auth-service/entity"
	"errors"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository struct {
	Repository[entity.User]
	Log *logrus.Logger
}

func NewUserRepository(log *logrus.Logger) *UserRepository {
	return &UserRepository{
		Log: log,
	}
}

func (repo *UserRepository) GetByUserID(db *gorm.DB, user_id string) (*entity.User, error) {
	var result entity.User
	qResult := db.First(&result, "user_id = ?", user_id)
	if errors.Is(qResult.Error, gorm.ErrRecordNotFound) {
		repo.Log.Warnf("[UserRepository:Get] Request data user id %s not found", user_id)
		return nil, nil
	} else if qResult.Error != nil {
		repo.Log.Warnf("[UserRepository:Get] Failed to get data user id %s cause:%+v", user_id, qResult.Error)
		return nil, qResult.Error
	}
	return &result, nil
}
