package db

import (
	"github.com/jackc/pgx/v5"
	"net/http"
)

func AddActor(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAdmin(r) {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}
	}
}

func DeleteActor(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !isAdmin(r) {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}
	}
}

func UpdateActor(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if !isAdmin(r) {
			http.Error(w, "Access forbidden", http.StatusForbidden)
			return
		}
	}
}

func GetActor(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}

func GetActorsAndFilms(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
	}
}
