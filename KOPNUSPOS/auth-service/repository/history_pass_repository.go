package repository

import "auth-service/entity"

type (
	HistoryPassRepository interface {
		Create(model entity.HistoryPass) (string, error)
		Get(id string) (*entity.HistoryPass, error)
		List(limit, offset int, filters map[string]interface{}) ([]entity.HistoryPass, int64, error)
		Update(model *entity.HistoryPass) error
	}
)
