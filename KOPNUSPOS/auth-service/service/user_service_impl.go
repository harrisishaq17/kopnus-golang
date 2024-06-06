package service

import (
	"auth-service/config"
	"auth-service/entity"
	"auth-service/helpers"
	"auth-service/model"
	"auth-service/repository"
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserService struct {
	DB       *gorm.DB
	Log      *logrus.Logger
	RepoUser repository.UserRepository
}

func NewUserService(db *gorm.DB, log *logrus.Logger, repoUser repository.UserRepository) *UserService {
	return &UserService{
		DB:       db,
		Log:      log,
		RepoUser: repoUser,
	}
}

func (svc *UserService) GenerateTokenAndSession(dataUser entity.User) (map[string]interface{}, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = dataUser.UserID
	formatDate := time.Now().Add(time.Minute * 50)
	claims["exp"] = formatDate.Unix() // Expires in 24 hours
	signedToken, err := token.SignedString([]byte(config.AppConfig.JWTSecret))
	if err != nil {
		log.Println("Error when generate JWT Token, cause: ", err)
		return nil, model.NewError("500", "Internal server error.")
	}

	return map[string]interface{}{
		"token":     signedToken,
		"expiresIn": formatDate.Format("2006/01/02 15:04:05"),
	}, nil
}

func (svc *UserService) Login(ctx context.Context, req *model.LoginUserRequest) (*model.LoginUserResponse, error) {
	tx := svc.DB.WithContext(ctx).Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			svc.Log.WithFields(logrus.Fields{
				"error": r,
			}).Error("[UserService:Login] Rollback due to panic")
		}
	}()

	dataUser, err := svc.RepoUser.GetByUserID(tx, req.UserID)
	if err != nil {
		svc.Log.WithFields(logrus.Fields{
			"userID": req.UserID,
			"error":  err,
		}).Warnf("[UserService:Login] Error while get data, cause: %+v", err)
		return nil, model.NewError("500", "Internal server error.")
	} else if dataUser == nil {
		return nil, model.NewError("404", "Data tidak ditemukan.")
	} else if dataUser.Counter == "4" || dataUser.Flag == "2" {
		// Validasi untuk cek user terblokir
		return nil, model.NewError("401", "User terblokir, Silahkan menghubungi IT Ops untuk membuka blokir user anda.")
	}

	var dict = helpers.HashDictionaries()
	var password string

	// Buat hash password dari data request berdasarkan data dictionaries
	for i := 0; i < len(req.Password); i++ {
		password = fmt.Sprintf("%s%s", password, dict[string(req.Password[i])])
	}

	log.Printf("Password User: %s, Hash Result: %s", dataUser.Password, password)

	if password != dataUser.Password {
		counter, _ := strconv.Atoi(dataUser.Counter)
		counter += 1
		dataUser.Counter = strconv.Itoa(counter)
		log.Printf("Counter: %d", counter)

		// Jika counter mencapai 4 maka nilai flag diubah menjadi 2
		if counter >= 4 {
			dataUser.Flag = "2"
		}

		err = svc.RepoUser.Update(tx, dataUser)
		if err != nil {
			svc.Log.WithFields(logrus.Fields{
				"userID": req.UserID,
				"error":  err,
			}).Warnf("[UserService:Login] Error while update user data, cause: %+v", err)
			return nil, model.NewError("500", "Internal server error.")
		}

		if err := tx.Commit().Error; err != nil {
			svc.Log.WithFields(logrus.Fields{
				"error": err,
			}).Warnf("[UserService:Login] Failed commit transaction : %+v", err)
			return nil, model.NewError("500", "Internal server error.")
		}

		svc.Log.WithFields(logrus.Fields{
			"error": "Success Test",
		}).Warnf("[UserService:Login] Success Test")

		if counter >= 4 {
			return nil, model.NewError("401", "User's Blocked")
		}

		return nil, model.NewError("401", "Wrong Password")
	}

	//Generate Token JWT User
	respToken, err := svc.GenerateTokenAndSession(*dataUser)
	if err != nil {
		return nil, err
	}

	// Reset Counter
	dataUser.Counter = "0"
	err = svc.RepoUser.Update(tx, dataUser)
	if err != nil {
		svc.Log.WithFields(logrus.Fields{
			"userID": req.UserID,
			"error":  err,
		}).Warnf("[UserService:Login] Error while update user data, cause: %+v", err)
		return nil, model.NewError("500", "Internal server error.")
	}
	log.Println("Counter Reset")

	// commit code for all success flow login
	if err := tx.Commit().Error; err != nil {
		svc.Log.WithFields(logrus.Fields{
			"error": err,
		}).Warnf("[UserService:Login] Failed commit transaction : %+v", err)
		return nil, model.NewError("500", "Internal server error.")
	}

	svc.Log.WithFields(logrus.Fields{
		"error": "Success Test",
	}).Warnf("[UserService:Login] Success Test")

	return &model.LoginUserResponse{
		Token: &model.Token{
			AccessToken: respToken["token"].(string),
			Type:        "bearer",
			ExpiresIn:   respToken["expiresIn"].(string),
		},
		User: &model.User{
			UserID:       strings.TrimSpace(dataUser.UserID),
			Nama:         strings.TrimSpace(dataUser.Nama),
			KdCabang:     strings.TrimSpace(dataUser.KdCabang),
			StatusKantor: strings.TrimSpace(dataUser.StatusKantor),
			FlagBudget:   strings.TrimSpace(dataUser.FlagBudget),
		},
	}, nil
}

func (svc *UserService) GetDataUser(ctx context.Context, req *model.GetDataUserRequest) (*entity.User, error) {
	tx := svc.DB.WithContext(ctx).Begin()
	dataUser, err := svc.RepoUser.GetByUserID(tx, req.UserID)
	if err != nil {
		svc.Log.WithFields(logrus.Fields{
			"userID": req.UserID,
			"error":  err,
		}).Warnf("[UserService:Login] Error while get user data, cause: %+v", err)
		return nil, model.NewError("500", "Internal server error.")
	} else if dataUser == nil {
		return nil, model.NewError("404", "Data not found.")
	}

	return dataUser, nil
}

func (svc *UserService) ChangePassword(ctx context.Context, req *model.ChangePasswordRequest) {

}

func (svc *UserService) SendOTP(ctx context.Context, req *model.SendOTPRequest) error {
	tx := svc.DB.WithContext(ctx).Begin()
	dataUser, err := svc.RepoUser.GetByUserID(tx, req.UserID)
	if err != nil {
		log.Println("Error while get data, cause: ", err)
		return model.NewError("500", "Internal server error.")
	} else if dataUser == nil {
		return model.NewError("404", "User not found.")
	}

	return nil
}
