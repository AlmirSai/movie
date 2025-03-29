package services

import (
	"fmt"
	"movie/internal/models"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ShowtimeService struct {
	DB *gorm.DB
}

func NewShowtimeService(db *gorm.DB) *ShowtimeService {
	return &ShowtimeService{
		DB: db,
	}
}

func (this *ShowtimeService) CreateShowtime(showtime models.Showtime) (*models.Showtime, error) {
	showtime.ID = uuid.New()
	if err := this.DB.Create(&showtime).Error; err != nil {
		return nil, fmt.Errorf("error creating showtime: %v", err)
	}
	return &showtime, nil
}

func (this *ShowtimeService) GetShowtimes(movieID string) ([]*models.Showtime, error) {
	var showtimes []*models.Showtime
	if err := this.DB.Where("movie_id = ?", movieID).Find(&showtimes).Error; err != nil {
		return nil, fmt.Errorf("error fetching showtimes: %v", err)
	}

	return showtimes, nil
}

func (this *ShowtimeService) UpdateShowtime(showtime models.Showtime) (*models.Showtime, error) {
	if err := this.DB.Save(&showtime).Error; err != nil {
		return nil, fmt.Errorf("error updating showtime: %v", err)
	}
	return &showtime, nil
}

func (this *ShowtimeService) DeleteShowtime(showtimeID string) error {
	if err := this.DB.Where("id = ?", showtimeID).Delete(&models.Showtime{}).Error; err != nil {
		return fmt.Errorf("error deleting showtime: %v", err)
	}
	return nil
}

func (this *ShowtimeService) GetAvailableSeats(showtimeID string) ([]string, error) {
	var seats string
	if err := this.DB.Table("showtimes").Select("available_seats").Where(
		"id = ?", showtimeID).Scan(&seats).Error; err != nil {
		return nil, fmt.Errorf("error fetching available seats: %v", err)
	}

	seatsList := strings.Split(seats, ",")
	for i, seat := range seatsList {
		seatsList[i] = strings.TrimSpace(seat)
	}

	return seatsList, nil
}
