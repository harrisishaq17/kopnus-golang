package repository

import "auth-service/entity"

type (
	OTPRepository interface {
		Create(model entity.OTP) error
		Get(id string) (*entity.OTP, error)
		Update(model *entity.OTP) error
	}
)
