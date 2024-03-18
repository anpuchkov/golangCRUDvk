package db

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"net/http"
)

// SetupRoutes настраивает маршруты для обработки запросов к фильмам.
func SetupRoutes(db *pgxpool.Pool, logger *zap.Logger) {
	http.HandleFunc("/movies", GetMovies(db, logger))
	http.HandleFunc("/movies/add", AddMovie(db, logger))
	http.HandleFunc("/movies/update", UpdateMovie(db, logger))
	http.HandleFunc("/movies/delete", DeleteMovie(db, logger))
	http.HandleFunc("/movies/search", GetMoviesByPartOfTitle(db, logger))
	http.HandleFunc("/movies/sort", GetMoviesWithSort(db, logger))
}
