package services

import (
	"database/sql"
	"errors"
	"personal-dashboard-backend/api/user/model"
	"personal-dashboard-backend/db"
)

func GetUserService(identifier string) (*model.User, error) {

	var user model.User

	query := "SELECT id, username, email, created_at FROM users WHERE email = $1 OR username = $1"
	row := db.DB.QueryRow(query, identifier)

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	
	return &user, nil
	
}