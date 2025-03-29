package services

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"movie/internal/models"
	jwt "movie/pkg/jwt"
	password "movie/pkg/password"
)

type AuthService struct {
	DB *gorm.DB
}

func NewAuthService(db *gorm.DB) *AuthService {
	return &AuthService{
		DB: db,
	}
}

func (this *AuthService) SingUp(user models.User) (*models.User, error) {
	user.ID = uuid.New()
	user.Role = models.RegularUserRole
	hashedPassword, err := password.HashPassword(user.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %w", err)
	}
	user.Password = hashedPassword

	if err := this.DB.Create(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to create user: %w", err)
	}

	return &user, nil
}

func (this *AuthService) Login(credits models.Credentials) (*models.User, string, error) {
	var user models.User
	if err := this.DB.Where("email = ?", credits.Email).First(&user).Error; err != nil {
		return nil, "", fmt.Errorf("failed to login: %w", err)
	}

	if !password.CheckPasswordHash(credits.Password, user.Password) {
		return nil, "", fmt.Errorf("invalid credentials")
	}

	token, err := jwt.GenerateJWTToken(user.ID.String(), string(rune(user.Role)))
	if err != nil {
		return nil, "", fmt.Errorf("failed to generate token: %w", err)
	}

	return &user, token, nil
}

func (this *AuthService) PromoteToAdmin(userID string) error {
	var user models.User
	if err := this.DB.Where("id = ?", userID).First(&user).Error; err != nil {
		return fmt.Errorf("failed to get user: %w", err)
	}

	user.Role = models.AdminRole
	if err := this.DB.Save(&user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	return nil
}
