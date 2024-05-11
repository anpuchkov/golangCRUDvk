package server

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"net/http"
	db2 "vkFilmoteka/model/db"
)

// SetupRoutes настраивает маршруты для обработки запросов к фильмам.
func SetupRoutes(db *pgxpool.Pool, logger *zap.Logger) {
	http.HandleFunc("/login", db2.AuthenticateUser(db, logger))
	http.HandleFunc("/movies/update", db2.UpdateMovie(db, logger))
	http.HandleFunc("/register", db2.RegisterUser(db, logger))
	http.HandleFunc("/movies", db2.GetMovies(db, logger))
	http.HandleFunc("/movies/search", db2.GetMoviesByPartOfTitle(db, logger))
	http.HandleFunc("/movies/add", db2.AddMovie(db, logger))
	http.HandleFunc("/movies/delete", db2.DeleteMovie(db, logger))
	http.HandleFunc("/movies/sort", db2.GetMoviesWithSort(db, logger))
	http.HandleFunc("/movies", db2.GetMovies(db, logger))
	http.HandleFunc("/movies/search", db2.GetMoviesByPartOfTitle(db, logger))
}
