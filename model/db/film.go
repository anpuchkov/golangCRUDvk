package db

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"vkFilmoteka/domain"
	"vkFilmoteka/server/logs"
)

func UpdateMovie(db *pgxpool.Pool, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAdmin(r) {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}
		var movie domain.Movie

		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			logs.Error("Invalid JSON", zap.Error(err))
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		_, err = db.Exec(r.Context(), `
            UPDATE Movies 
            SET title = $1, description = $2, release_date = $3, rating = $4
            WHERE id = $5
        `, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating, movie.ID)
		if err != nil {
			logs.Error("Failed to update movie in database", zap.Error(err))
			http.Error(w, "Failed to update movie in database", http.StatusInternalServerError)
			return
		}

		logs.Info("Movie updated", zap.Int("id", movie.ID))

		w.WriteHeader(http.StatusOK)
	}
}

func DeleteMovie(db *pgxpool.Pool, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAdmin(r) {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}
		movieID := r.URL.Query().Get("id")
		if movieID == "" {
			http.Error(w, "Movie ID is required", http.StatusBadRequest)
			return
		}

		_, err := db.Exec(r.Context(), `
            DELETE FROM movies
            WHERE id = $1
        `, movieID)
		if err != nil {
			logs.Error("Failed to delete movie from database", zap.Error(err))
			http.Error(w, "Failed to delete movie from database", http.StatusInternalServerError)
			return
		}

		logs.Info("Movie deleted", zap.String("id", movieID))

		w.WriteHeader(http.StatusOK)

	}
}

func GetMovies(db *pgxpool.Pool, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(r.Context(), `
            SELECT id, title, description, release_date, rating
            FROM movies ORDER BY rating DESC
        `)
		if err != nil {
			logs.Error("Failed to fetch movies from database", zap.Error(err))
			http.Error(w, "Failed to fetch movies from database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var movies []domain.Movie
		for rows.Next() {
			var movie domain.Movie
			err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
			if err != nil {
				logs.Error("Failed to scan movie row", zap.Error(err))
				http.Error(w, "Failed to fetch movies from database", http.StatusInternalServerError)
				return
			}
			movies = append(movies, movie)
		}

		err = json.NewEncoder(w).Encode(movies)
		if err != nil {
			logs.Error("Failed to encode movies to JSON", zap.Error(err))
			http.Error(w, "Failed to encode movies to JSON", http.StatusInternalServerError)
			return
		}
		logs.Info("sent all movies")
	}
}

func GetMoviesByPartOfTitle(db *pgxpool.Pool, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		partOfTitle := r.URL.Query().Get("partOfTitle")
		if partOfTitle == "" {
			http.Error(w, "Part of title is required", http.StatusBadRequest)
			return
		}

		rows, err := db.Query(r.Context(), `
            SELECT id, title, description, release_date, rating
            FROM Movies
            WHERE title LIKE '%' || $1 || '%'
        `, partOfTitle)
		if err != nil {
			logs.Error("Failed to fetch movies from database", zap.Error(err))
			http.Error(w, "Failed to fetch movies from database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var movies []domain.Movie
		for rows.Next() {
			var movie domain.Movie
			err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
			if err != nil {
				logs.Error("Failed to scan movie row", zap.Error(err))
				http.Error(w, "Failed to fetch movies from database", http.StatusInternalServerError)
				return
			}
			movies = append(movies, movie)
		}

		err = json.NewEncoder(w).Encode(movies)
		if err != nil {
			logs.Error("Failed to encode movies to JSON", zap.Error(err))
			http.Error(w, "Failed to encode movies to JSON", http.StatusInternalServerError)
			return
		}
		logs.Info("Sent movies", zap.String("partOfTitle", partOfTitle))
	}
}

func GetMoviesWithSort(db *pgxpool.Pool, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sortType := r.URL.Query().Get("sortType")
		if sortType == "" {
			sortType = "rating"
		}

		validSortTypes := map[string]bool{"date": true, "title": true, "rating": true}
		if !validSortTypes[sortType] {
			http.Error(w, "Invalid sort type", http.StatusBadRequest)
			return
		}

		var rows pgx.Rows
		var err error
		switch sortType {
		case "date":
			rows, err = db.Query(r.Context(), `
                SELECT id, title, description, release_date, rating
                FROM Movies
                ORDER BY release_date DESC
            `)
		case "title":
			rows, err = db.Query(r.Context(), `
                SELECT id, title, description, release_date, rating
                FROM Movies
                ORDER BY title
            `)
		case "rating":
			rows, err = db.Query(r.Context(), `
                SELECT id, title, description, release_date, rating
                FROM Movies
                ORDER BY rating DESC
            `)
		}
		if err != nil {
			logs.Error("Failed to fetch movies from database", zap.Error(err))
			http.Error(w, "Failed to fetch movies from database", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var movies []domain.Movie
		for rows.Next() {
			var movie domain.Movie
			err := rows.Scan(&movie.ID, &movie.Title, &movie.Description, &movie.ReleaseDate, &movie.Rating)
			if err != nil {
				logs.Error("Failed to scan movie row", zap.Error(err))
				http.Error(w, "Failed to fetch movies from database", http.StatusInternalServerError)
				return
			}
			movies = append(movies, movie)
		}

		err = json.NewEncoder(w).Encode(movies)
		if err != nil {
			logs.Error("Failed to encode movies to JSON", zap.Error(err))
			http.Error(w, "Failed to encode movies to JSON", http.StatusInternalServerError)
			return
		}
		logs.Info("GetMoviesWithSort", zap.String("sortType", sortType))
	}
}

func AddMovie(db *pgxpool.Pool, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAdmin(r) {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}
		var movie domain.Movie

		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			logs.Error("Invalid JSON", zap.Error(err))
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		movie.ReleaseDate = pgtype.Date{
			Time:  time.Now(),
			Valid: true,
		}

		var movieID int
		err = db.QueryRow(r.Context(), `
            INSERT INTO Movies (title, description, release_date, rating)
            VALUES ($1, $2, $3, $4)
            RETURNING id
        `, movie.Title, movie.Description, movie.ReleaseDate, movie.Rating).Scan(&movieID)
		if err != nil {
			logs.Error("Failed to add movie to database", zap.Error(err))
			http.Error(w, "Failed to add movie to database", http.StatusInternalServerError)
			return
		}
		logs.Info("Movie added successfully", zap.Int("movie_id", movieID))
		w.WriteHeader(http.StatusCreated)
	}
}

func isAdmin(r *http.Request) bool {
	// Здесь должна быть реализация проверки, является ли пользователь администратором
	// Например, проверка роли пользователя или наличие аутентификационного токена с правами администратора
	// В данном примере всегда возвращаем true для демонстрации
	return true
}
