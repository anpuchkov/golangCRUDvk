package db

import (
	"encoding/json"
	"net/http"
	"vkFilmoteka/auth"
	"vkFilmoteka/server/logs"

	"github.com/gorilla/sessions"
	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
)

var store *sessions.CookieStore

func RegisterUser(db *pgxpool.Pool, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logs.Error("Method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var credentials struct {
			Username string `json:"username" required:"true"`
			Password string `json:"password" required:"true"`
		}

		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			logs.Error("Invalid JSON", zap.Error(err))
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		if credentials.Username == "" || credentials.Password == "" {
			logs.Error("Username and password are required")
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}

		err = auth.RegisterUser(db, credentials.Username, credentials.Password)
		if err != nil {
			logs.Error("Failed to register user", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		session, _ := store.Get(r, "session-name")
		session.Values["authenticated"] = true
		session.Save(r, w)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(struct{ Status string }{Status: "Success"})
	}
}

func AuthenticateUser(db *pgxpool.Pool, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			logs.Error("Method not allowed")
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		var credentials struct {
			Username string `json:"username" required:"true"`
			Password string `json:"password" required:"true"`
		}

		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			logs.Error("Invalid JSON", zap.Error(err))
			http.Error(w, "Invalid JSON", http.StatusBadRequest)
			return
		}

		_, err = auth.AuthenticateUser(db, credentials.Username, credentials.Password)
		if err != nil {
			logs.Error("Failed to authenticate user", zap.Error(err))
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		session, _ := store.Get(r, "session-name")
		session.Values["authenticated"] = true
		session.Save(r, w)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(struct{ Status string }{Status: "Success"})
	}
}

func isUser(r *http.Request) bool {
	session, _ := store.Get(r, "session-name")
	val := session.Values["authenticated"]
	authenticated, ok := val.(bool)
	return ok && authenticated
}

func InitSessionStore() {
	store = sessions.NewCookieStore([]byte("your-secret-key"))
}
