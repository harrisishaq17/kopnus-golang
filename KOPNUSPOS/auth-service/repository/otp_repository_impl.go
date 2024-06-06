package repository

import (
	"auth-service/entity"
	"errors"
	"log"

	"gorm.io/gorm"
)

type (
	otpRepository struct {
		DB *gorm.DB
	}
)

func NewOTPRepository(conn *gorm.DB) OTPRepository {
	return &otpRepository{
		DB: conn,
	}
}

func (repo *otpRepository) Create(model entity.OTP) error {
	db := repo.DB.Create(&model)
	err := db.Error
	if err != nil {
		log.Println("error cause: ", err)
		return err
	}
	return nil
}

func (repo *otpRepository) Get(id string) (*entity.OTP, error) {
	var result entity.OTP
	qResult := repo.DB.First(&result, "user_id = ?", id)
	if errors.Is(qResult.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if qResult.Error != nil {
		return nil, qResult.Error
	}
	return &result, nil
}

func (repo *otpRepository) Update(model *entity.OTP) error {
	qResult := repo.DB.Select("*").Where("user_id = ?", model.UserID).Updates(model)
	if errors.Is(qResult.Error, gorm.ErrRecordNotFound) {
		return nil
	} else if qResult.Error != nil {
		return qResult.Error
	}
	return nil
}
