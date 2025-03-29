package routes

import (
	"movie/internal/controllers"
	"movie/internal/services"
	"movie/pkg/middleware"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	router := gin.Default()

	authService := services.NewAuthService(db)
	movieService := services.NewMovieService(db)
	reservationService := services.NewReservationService(db)
	showtimeService := services.NewShowtimeService(db)

	authController := controllers.NewAuthController(authService)
	movieController := controllers.NewMovieController(movieService)
	reservationController := controllers.NewReservationService(reservationService, showtimeService)
	showtimeController := controllers.NewShowtimeController(showtimeService)

	public := router.Group("/api")
	{
		public.POST("/singup", authController.SingUp)
		public.POST("/login", authController.Login)
		public.POST("/movies", movieController.CreateMovie)
	}

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/user/reservations", reservationController.GetUserReservations)
		protected.POST("/reservations", reservationController.CreateReservation)
		protected.DELETE("/reservations/:reservationId", reservationController.CancelReservation)
		protected.GET("/showtimes/:showtimeId/seats", reservationController.GetAvailableSeats)
		protected.GET("/movies/:movieID/showtimes", showtimeController.GetShowtimes)
	}

	admin := router.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware(), middleware.AdminMiddleware())
	{
		admin.POST("/movies", movieController.CreateMovie)
		admin.PUT("/movies/:movieId", movieController.UpdateMovie)
		admin.DELETE("/movies/:movieId", movieController.DeleteMovie)
		admin.GET("/reservations", reservationController.GetAllReservations)
		admin.POST("/users/:userId/promote", authController.PromoteToAdmin)
		admin.POST("/showtimes", showtimeController.CreateShowtime)
		admin.PUT("/showtimes/:showtimeId", showtimeController.UpdateShowtime)
		admin.DELETE("/showtimes/:showtimeId", showtimeController.DeleteShowtime)
	}

	return router
}
