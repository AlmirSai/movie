package controllers

import (
	"log"
	"movie/internal/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	services "movie/internal/services"
)

type MovieController struct {
	MovieService *services.MovieService
}

func NewMovieController(ms *services.MovieService) *MovieController {
	return &MovieController{
		MovieService: ms,
	}
}

type CreateMovieRequest struct {
	Title       string `json:"title" binding:"required"`
	Director    string `json:"director" binding:"required"`
	ReleaseDate string `json:"releaseDate" binding:"required"`
	Duration    int    `json:"duration" binding:"required"`
	Description string `json:"description" binding:"required"`
	Genre       string `json:"genre" binding:"required"`
	PosterURL   string `json:"posterURL" binding:"required"`
}

func (this *MovieController) CreateMovie(c *gin.Context) {
	var req CreateMovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		log.Printf("Error parsing request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	releaseDate, err := time.Parse("2006-01-02", req.ReleaseDate)
	if err != nil {
		log.Printf("Error parsing releaseDate: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	movie := models.Movie{
		ID:          uuid.New(),
		Title:       req.Title,
		Director:    req.Director,
		ReleaseDate: releaseDate,
		Duration:    req.Duration,
		Description: req.Description,
		Genre:       models.Genre(req.Genre),
		PosterImage: req.PosterURL,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	createMovie, err := this.MovieService.CreateMovie(&movie)
	if err != nil {
		log.Printf("Error creating movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createMovie)
}

func (this *MovieController) GetMovies(c *gin.Context) {
	movies, err := this.MovieService.GetMovies()
	if err != nil {
		log.Printf("Error getting movies: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, movies)
}

func (this *MovieController) UpdateMovie(c *gin.Context) {
	var movie models.Movie
	if err := c.ShouldBindJSON(&movie); err != nil {
		log.Printf("Error parsing request: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updateMovie, err := this.MovieService.UpdateMovie(movie)
	if err != nil {
		log.Printf("Error updating movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updateMovie)
}

func (this *MovieController) DeleteMovie(c *gin.Context) {
	movieId := c.Param("movieId")
	if err := this.MovieService.DeleteMovie(movieId); err != nil {
		log.Printf("Error deleting movie: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "movie deleted"})
}
