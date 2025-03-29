package controllers

import (
	"movie/internal/models"
	services "movie/internal/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ShowtimeController struct {
	showtimeService *services.ShowtimeService
}

func NewShowtimeController(
	showtimeService *services.ShowtimeService,
) *ShowtimeController {
	return &ShowtimeController{
		showtimeService: showtimeService,
	}
}

func (this *ShowtimeController) CreateShowtime(c *gin.Context) {
	var showtime models.Showtime
	if err := c.ShouldBindJSON(&showtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdShowtime, err := this.showtimeService.CreateShowtime(showtime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdShowtime)
}

func (this *ShowtimeController) GetShowtimes(c *gin.Context) {
	movieID := c.Param("movieID")
	showtimes, err := this.showtimeService.GetShowtimes(movieID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, showtimes)
}

func (this *ShowtimeController) UpdateShowtime(c *gin.Context) {
	var showtime models.Showtime
	if err := c.ShouldBindJSON(&showtime); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedShowtime, err := this.showtimeService.UpdateShowtime(showtime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedShowtime)
}

func (this *ShowtimeController) DeleteShowtime(c *gin.Context) {
	showtimeID := c.Param("showtimeID")
	err := this.showtimeService.DeleteShowtime(showtimeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Showtime deleted successfully"})
}

func (this *ShowtimeController) GetAvailableSeats(c *gin.Context) {
	showtimeID := c.Param("showtimeID")
	availableSeats, err := this.showtimeService.GetAvailableSeats(showtimeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"available_seats": availableSeats})
}
