# Movie Booking System

A Go-based RESTful API service for movie ticket booking and management.

## Features

- User Authentication and Authorization
- Movie Management
- Showtime Scheduling
- Seat Reservation System
- Admin Controls

## Project Structure

```
├── cmd/
│   └── server/          # Main application entry point
├── config/              # Configuration files
├── docker/              # Docker configuration
├── internal/
│   ├── config/         # Internal configuration
│   ├── controllers/    # HTTP request handlers
│   ├── models/         # Data models
│   ├── routes/         # API route definitions
│   └── services/       # Business logic
├── pkg/                # Reusable packages
└── tests/              # Test files
```

## Prerequisites

- Go 1.x
- Docker and Docker Compose (for containerized deployment)
- PostgreSQL (or your preferred database)

## Getting Started

1. Clone the repository

```bash
git clone <repository-url>
cd movie
```

2. Set up environment variables

```bash
cp config/.env.example config/.env
# Edit .env with your configuration
```

3. Run locally

```bash
go mod download
go run cmd/server/main.go
```

4. Run with Docker

```bash
docker-compose -f docker/docker-compose.yml up --build
```

## API Endpoints

### Authentication
- POST /api/auth/register - Register new user
- POST /api/auth/login - User login

### Movies
- GET /api/movies - List all movies
- GET /api/movies/{id} - Get movie details
- POST /api/movies - Add new movie (Admin)
- PUT /api/movies/{id} - Update movie (Admin)
- DELETE /api/movies/{id} - Delete movie (Admin)

### Showtimes
- GET /api/showtimes - List all showtimes
- GET /api/showtimes/{id} - Get showtime details
- POST /api/showtimes - Add new showtime (Admin)
- PUT /api/showtimes/{id} - Update showtime (Admin)
- DELETE /api/showtimes/{id} - Delete showtime (Admin)

### Reservations
- GET /api/reservations - List user reservations
- POST /api/reservations - Create new reservation
- GET /api/reservations/{id} - Get reservation details
- DELETE /api/reservations/{id} - Cancel reservation

## Testing

API endpoints can be tested using the provided HTTP test files:

```bash
# Using the test.http file in tests/http/
# Requires REST Client extension in VS Code or similar tool
```

Test data is available in `tests/json/test.json`

## Contributing

1. Fork the repository
2. Create your feature branch
3. Commit your changes
4. Push to the branch
5. Create a new Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details