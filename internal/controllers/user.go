package controllers

import (
	"context"
	"fmt"
	"myblogapi/internal/db"
	"myblogapi/internal/models"

	"github.com/jackc/pgx/v5"
)

func GetUserByEmail(ctx context.Context, email string) (models.User, error) {
	var user models.User

	err := db.DB.QueryRow(ctx, "SELECT email FROM users WHERE email = $1", email).Scan(&user.Email)
	fmt.Println("hey error", err)
	if err != nil {
		fmt.Println("error is here", err.Error())
		if err == pgx.ErrNoRows {
			return models.User{}, nil
		}
		return models.User{}, fmt.Errorf("unable to retrieve user: %v", err)
	}
	return user, nil
}
