package db

import (
	"github.com/jackc/pgx/v5"
	"net/http"
)

func FragmentedSearch(db *pgx.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Обрабатываем GET-запрос
		// Здесь ваш код для поиска фрагментов в базе данных
	}
}
