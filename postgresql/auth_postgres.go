package postgresql

import (
	"bookmyshow/models"
	"fmt"

	"github.com/google/uuid"
)

func GetUserDetailsQuery(username string) (models.User, error) {
	var user models.User

	smt := `SELECT * FROM users WHERE username = $1`
	err := db.QueryRow(smt, username).Scan(&user.Id, &user.Username, &user.Password, &user.RefreshToken)
	if err != nil {
		return user, fmt.Errorf("failed to fetch user details: %w", err)
	}

	return user, nil
}

func GetIdPasswordQuery(username string) (uuid.UUID, string, error) {
	var password string
	var id uuid.UUID
	smt := `SELECT id,password FROM users WHERE username = $1`
	err := db.QueryRow(smt, username).Scan(&id, &password)
	if err != nil {
		return id, "", err
	}
	return id, password, nil
}

func RegisterNewUserQuery(user models.User) (uuid.UUID, error) {
	var id uuid.UUID
	smt := `INSERT INTO users(name,username,password) VALUES($1,$2,$3) RETURNING id`
	err := db.QueryRow(smt, user.Name, user.Username, user.Password).Scan(&id)
	return id, err
}

func UpdateUserTokensQuery(refreshToken string, id uuid.UUID) error {
	smt := `UPDATE users SET refresh_token = $1 WHERE id = $2`
	_, err := db.Exec(smt, refreshToken, id)
	return err
}

func CompareRefreshToken(refreshToken string, id uuid.UUID) bool {
	var refresh_token string
	smt := `SELECT refresh_token from users WHERE id = $1 AND refresh_token = $2`
	err := db.QueryRow(smt, id, refreshToken).Scan(&refresh_token)
	return err == nil
}
