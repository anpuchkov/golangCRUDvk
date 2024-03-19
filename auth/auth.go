package db

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/crypto/bcrypt"
)

// User represents a user in the system
type User struct {
	ID       int
	Username string
	Password string
}

// RegisterUser registers a new user in the system
func RegisterUser(db *pgxpool.Pool, username, password string) error {
	var count int
	err := db.QueryRow(context.TODO(), "SELECT COUNT(id) FROM users WHERE username = $1", username).Scan(&count)
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("username already exists")
	}

	hashedPassword := hashPassword(password)

	_, err = db.Exec(context.TODO(), "INSERT INTO users (username, password) VALUES ($1, $2)", username, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}

// AuthenticateUser аутентифицирует пользователя по имени пользователя и паролю
func AuthenticateUser(db *pgxpool.Pool, username, password string) (*User, error) {
	var user User
	err := db.QueryRow(context.TODO(), "SELECT id, username, password FROM users WHERE username = $1", username).Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}

	if !checkPassword(password, user.Password) {
		return nil, errors.New("incorrect password")
	}

	return &user, nil
}

// hashPassword хэширует пароль
func hashPassword(password string) string {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	return string(hashedPassword)
}

// checkPassword проверяет совпадение хэшированного пароля с исходным паролем
func checkPassword(password, hashedPassword string) bool {
	bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return password == hashedPassword
}
