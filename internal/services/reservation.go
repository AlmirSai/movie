package services

import (
	"fmt"
	"movie/internal/models"
	"strconv"
	"strings"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ReservationService struct {
	DB *gorm.DB
}

func NewReservationService(db *gorm.DB) *ReservationService {
	return &ReservationService{
		DB: db,
	}
}

func (this *ReservationService) GetAvailableSeats(showtimeID string) ([]int, error) {
	var showtime models.Showtime
	if err := this.DB.Where("id = ?", showtimeID).First(&showtime).Error; err != nil {
		return nil, fmt.Errorf("failed to get showtime: %w", err)
	}
	reservations := make([]models.Reservation, 0)
	if err := this.DB.Where("showtime_id = ?", showtimeID).Find(&reservations).Error; err != nil {
		return nil, fmt.Errorf("failed to get reservations: %w", err)
	}

	var bookedSeats []int
	for _, reservation := range reservations {
		seatNumbers := strings.Split(reservation.SeatNumbers, ",")
		for _, seatString := range seatNumbers {
			seat, err := strconv.Atoi(strings.TrimSpace(seatString))
			if err != nil {
				return nil, fmt.Errorf("failed to parse seat number: %w", err)
			}
			bookedSeats = append(bookedSeats, seat)
		}
	}

	availableSeats := make([]int, showtime.AvailableSeats)
	for i := 0; i < showtime.AvailableSeats; i++ {
		availableSeats[i] = i + 1
	}

	for _, bookedSeat := range bookedSeats {
		for i, seat := range availableSeats {
			if seat == bookedSeat {
				availableSeats = append(availableSeats[:i], availableSeats[i+1:]...)
				break
			}
		}
	}

	return availableSeats, nil
}

func (this *ReservationService) GetUserReservations(userID string) ([]*models.Reservation, error) {
	var reservations []*models.Reservation
	if err := this.DB.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}

func (this *ReservationService) CreateReservation(reservation models.Reservation) (*models.Reservation, error) {
	reservation.ID = uuid.New()
	if err := this.DB.Create(&reservation).Error; err != nil {
		return nil, fmt.Errorf("failed to create reservation: %w", err)
	}
	return &reservation, nil
}

func (this *ReservationService) GetReservationsByUserID(userID string) ([]models.Reservation, error) {
	var reservations []models.Reservation
	if err := this.DB.Where("user_id = ?", userID).Find(&reservations).Error; err != nil {
		return nil, fmt.Errorf("failed to get reservations: %w", err)
	}
	return reservations, nil
}

func (this *ReservationService) CancelReservation(reservationID string) error {
	if err := this.DB.Where("id = ?", reservationID).Delete(&models.Reservation{}).Error; err != nil {
		return err
	}
	return nil
}

func (this *ReservationService) GetAllReservations() ([]*models.Reservation, error) {
	var reservations []*models.Reservation
	if err := this.DB.Find(&reservations).Error; err != nil {
		return nil, err
	}
	return reservations, nil
}
