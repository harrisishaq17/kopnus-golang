package repository

import (
	"auth-service/entity"
	"errors"
	"fmt"
	"log"
	"reflect"
	"strings"

	"gorm.io/gorm"
)

type (
	historyPassRepository struct {
		DB *gorm.DB
	}
)

func NewHistoryPassRepository(conn *gorm.DB) HistoryPassRepository {
	return &historyPassRepository{
		DB: conn,
	}
}

func (repo *historyPassRepository) Create(model entity.HistoryPass) (string, error) {
	db := repo.DB.Create(&model)
	err := db.Error
	if err != nil {
		log.Println("error cause: ", err)
		return "", err
	}
	return model.UserID, nil
}

func (repo *historyPassRepository) Get(id string) (*entity.HistoryPass, error) {
	var result entity.HistoryPass
	qResult := repo.DB.First(&result, "user_id = ?", id)
	if errors.Is(qResult.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	} else if qResult.Error != nil {
		return nil, qResult.Error
	}
	return &result, nil
}

func (repo *historyPassRepository) List(limit, offset int, filters map[string]interface{}) ([]entity.HistoryPass, int64, error) {
	var results []entity.HistoryPass
	var tx = repo.DB.Model(&entity.HistoryPass{})
	tx.Order("id asc")

	if limit > 0 {
		tx.Limit(limit)
		tx.Offset(offset)
	}

	if len(filters) > 0 {
		for field, value := range filters {
			// avoid sql injection in column name
			var checkField = strings.Split(field, " ")
			if len(checkField) > 1 {
				field = checkField[0]
			}

			if field == "last_login_date" {
				var values = value.([]interface{})
				tx.Where("last_login_date between ? AND ?", values[0], values[1])
				continue
			}

			switch reflect.TypeOf(value).Kind() {
			case reflect.Float64, reflect.Int, reflect.Int64:
				tx.Where(fmt.Sprintf("%s = ?", field), value)
			case reflect.Slice:
				tx.Where(fmt.Sprintf("%s IN (?)", field), value)
			default:
				tx.Where(fmt.Sprintf("%s LIKE ?", field), "%"+value.(string)+"%")
			}
		}
	}

	err := tx.Find(&results).Error
	if err != nil {
		return make([]entity.HistoryPass, 0), 0, err
	}

	total, err := repo.listCount(&entity.User{}, filters)
	if err != nil {
		return make([]entity.HistoryPass, 0), 0, err
	}

	return results, total, nil
}

func (repo *historyPassRepository) listCount(model interface{}, filters map[string]interface{}, condition ...interface{}) (int64, error) {
	var results []entity.HistoryPass
	var tx = repo.DB.Model(&model)

	if len(filters) > 0 {
		for field, value := range filters {
			// avoid sql injection in column name
			var checkField = strings.Split(field, " ")
			if len(checkField) > 1 {
				field = checkField[0]
			}

			if field == "last_login_date" {
				var values = value.([]interface{})
				tx.Where("last_login_date between ? AND ?", values[0], values[1])
				continue
			}

			switch reflect.TypeOf(value).Kind() {
			case reflect.Float64, reflect.Int, reflect.Int64:
				tx.Where(fmt.Sprintf("%s = ?", field), value)
			case reflect.Slice:
				tx.Where(fmt.Sprintf("%s IN (?)", field), value)
			default:
				tx.Where(fmt.Sprintf("%s LIKE ?", field), "%"+value.(string)+"%")
			}
		}
	}

	if len(condition) > 0 {
		tx.Where(condition[0], condition[1:]...)
	}

	err := tx.Find(&results).Error
	if err != nil {
		return 0, err
	}

	var total int64
	err = tx.Count(&total).Error
	if err != nil {
		return 0, err
	}

	return total, nil
}

func (repo *historyPassRepository) Update(model *entity.HistoryPass) error {
	qResult := repo.DB.Select("*").Where("user_id = ?", model.UserID).Updates(model)
	if errors.Is(qResult.Error, gorm.ErrRecordNotFound) {
		return nil
	} else if qResult.Error != nil {
		return qResult.Error
	}
	return nil
}
