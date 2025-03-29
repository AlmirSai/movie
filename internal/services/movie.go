package services

import (
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"movie/internal/models"
)

type MovieService struct {
	DB *gorm.DB
}

func NewMovieService(db *gorm.DB) *MovieService {
	return &MovieService{
		DB: db,
	}
}

func (this *MovieService) CreateMovie(movie *models.Movie) (*models.Movie, error) {
	movie.ID = uuid.New()
	if err := this.DB.Create(&movie).Error; err != nil {
		return nil, fmt.Errorf("error creating movie: %v", err)
	}
	return movie, nil
}

func (this *MovieService) GetMovies() ([]*models.Movie, error) {
	var movies []*models.Movie
	if err := this.DB.Find(&movies).Error; err != nil {
		return nil, err
	}
	return movies, nil
}

func (this *MovieService) UpdateMovie(movie models.Movie) (*models.Movie, error) {
	if err := this.DB.Save(&movie).Error; err != nil {
		return nil, err
	}
	return &movie, nil
}

func (this *MovieService) DeleteMovie(movieID string) error {
	if err := this.DB.Where("id = ?", movieID).Delete(&models.Movie{}).Error; err != nil {
		return err
	}
	return nil
}
