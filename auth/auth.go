package auth

import (
	"database/sql"
	"encoding/base64"
	"net/http"
	"strings"
)

// adminAuthentication проверяет аутентификацию администратора на основе учетных данных из базы данных.
func adminAuthentication(r *http.Request, db *sql.DB) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	if !strings.HasPrefix(authHeader, "Basic ") {
		return false
	}

	token := strings.TrimPrefix(authHeader, "Basic ")

	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false
	}

	credentials := strings.SplitN(string(decodedToken), ":", 2)
	if len(credentials) != 2 {
		return false
	}
	username := credentials[0]
	password := credentials[1]

	row := db.QueryRow("SELECT username, password FROM admin WHERE username = $1", username)
	var dbUsername, dbPassword string
	err = row.Scan(&dbUsername, &dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		return false
	}

	return username == dbUsername && password == dbPassword
}

// AdminAuthMiddleware предоставляет промежуточное ПО для аутентификации администратора.
func AdminAuthMiddleware(next http.Handler, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !adminAuthentication(r, db) {
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func userAuthentication(r *http.Request, db *sql.DB) bool {
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		return false
	}

	if !strings.HasPrefix(authHeader, "Basic ") {
		return false
	}

	token := strings.TrimPrefix(authHeader, "Basic ")

	decodedToken, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		return false
	}

	credentials := strings.SplitN(string(decodedToken), ":", 2)
	if len(credentials) != 2 {
		return false
	}
	username := credentials[0]
	password := credentials[1]

	row := db.QueryRow("SELECT username, password FROM users WHERE username = $1", username)
	var dbUsername, dbPassword string
	err = row.Scan(&dbUsername, &dbPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		// Если произошла другая ошибка, возвращаем false
		return false
	}

	// Сравниваем полученные учетные данные с учетными данными пользователя из базы данных
	return username == dbUsername && password == dbPassword
}

// UserAuthMiddleware предоставляет промежуточное ПО для аутентификации пользователя.
func UserAuthMiddleware(next http.Handler, db *sql.DB) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Проверяем аутентификацию пользователя
		if !userAuthentication(r, db) {
			// Возвращаем ошибку, если аутентификация не удалась
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte("Unauthorized"))
			return
		}

		next.ServeHTTP(w, r)
	})
}
