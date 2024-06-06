package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type EnvGormConfig struct {
	DBHost     string `mapstructure:"DB_HOST" validate:"required"`
	DBPort     string `mapstructure:"DB_PORT" validate:"required"`
	DBUser     string `mapstructure:"DB_USER" validate:"required"`
	DBPassword string `mapstructure:"DB_PASSWORD" validate:"required"`
	DBName     string `mapstructure:"DB_NAME" validate:"required"`
}

var (
	GormConfig EnvGormConfig
)

type (
	DB struct {
		DB *gorm.DB
	}
)

func InitDBConnection(host, username, password, name, port string) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, username, password, name, port)
	log.Println("dsn:", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Failed to connect to database, error cause:", err)
		return nil
	}

	log.Println("Successfully connected to DB!")

	return db
}
