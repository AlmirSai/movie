package controllers

import (
	"log"
	"movie/internal/models"
	"movie/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type ReservationController struct {
	ReservationService *services.ReservationService
	ShowtimeService    *services.ShowtimeService
}

func NewReservationService(
	reservationService *services.ReservationService,
	showtimeService *services.ShowtimeService,
) *ReservationController {
	return &ReservationController{
		ReservationService: reservationService,
		ShowtimeService:    showtimeService,
	}
}

func (this *ReservationController) GetAvailableSeats(c *gin.Context) {
	showtimeID := c.Param("showtimeId")
	log.Printf("Fetching seats for showtime ID: %s\n", showtimeID)

	seats, err := this.ShowtimeService.GetAvailableSeats(showtimeID)
	if err != nil {
		log.Printf("Error fetching seats: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Available seats: %+v\n", seats)

	c.JSON(http.StatusOK, seats)
}

func (this *ReservationController) CreateReservation(c *gin.Context) {
	var reservation models.Reservation
	if err := c.ShouldBindJSON(&reservation); err != nil {
		log.Printf("Error parsing request: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, ok := c.Get("userId")
	if !ok {
		log.Printf("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userUUID, err := uuid.Parse(userID.(string))
	if err != nil {
		log.Printf("Error parsing user UUID: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	reservation.UserID = userUUID

	newReservation, err := this.ReservationService.CreateReservation(reservation)
	if err != nil {
		log.Printf("Error creating reservation: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, newReservation)
}

func (this *ReservationController) GetUserReservations(c *gin.Context) {
	userID, ok := c.Get("userId")
	if !ok {
		log.Printf("User not authenticated")
		c.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	userIDString, ok := userID.(string)
	if !ok {
		log.Printf("Invalid user ID format")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid user ID format"})
		return
	}

	reservations, err := this.ReservationService.GetUserReservations(userIDString)
	if err != nil {
		log.Printf("Error fetching reservations: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}

func (this *ReservationController) CancelReservation(c *gin.Context) {
	reservationID := c.Param("reservationId")
	if err := this.ReservationService.CancelReservation(reservationID); err != nil {
		log.Printf("Error cancelling reservation: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Reservation cancelled successfully"})
}

func (this *ReservationController) GetAllReservations(c *gin.Context) {
	reservations, err := this.ReservationService.GetAllReservations()
	if err != nil {
		log.Printf("Error fetching reservations: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reservations)
}
