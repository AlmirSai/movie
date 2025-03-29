package config

import (
	"fmt"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"movie/internal/models"
)

type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
	JWTSecret  string
}

func LoadConfig() (*Config, error) {
	viper.SetConfigFile("config/.env")
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Error reading config file: %v", err)
	}

	config := &Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetInt("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
	}

	return config, nil
}

func NewDB(config *Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("Error connecting to database: %v", err)
	}

	if err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";").Error; err != nil {
		return nil, fmt.Errorf("Error creating extension: %v", err)
	}

	if err := db.AutoMigrate(
		&models.Movie{},
		&models.User{},
		&models.Showtime{},
		&models.Reservation{},
	); err != nil {
		return nil, fmt.Errorf("Error initializing database: %v", err)
	}

	return db, nil
}
