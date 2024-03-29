package domain

import (
	"github.com/jackc/pgx/v5/pgtype"
	"time"
)

// Actor Структура для актеров
type Actor struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Gender    string    `json:"gender"`
	Birthdate time.Time `json:"birthdate"`
}

// Movie Структура для фильмов
type Movie struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ReleaseDate pgtype.Date `json:"release_date"`
	Rating      float64     `json:"rating"`
}

// User Структура для пользователей
type User struct {
	ID       int
	Username string
	Password string
}

// MovieActor Структура для фильмов и актеров
type MovieActor struct {
	ID          int         `json:"id"`
	Title       string      `json:"title"`
	Description string      `json:"description"`
	ReleaseDate pgtype.Date `json:"release_date"`
	Rating      float64     `json:"rating"`
	Actors      []Actor     `json:"actors"`
}
