package controllers

import (
	"context"
	"encoding/json"
	"myblogapi/internal/db"
	"myblogapi/internal/models"
	"myblogapi/internal/services"
	"myblogapi/pkg/utils"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

func createUser(ctx context.Context, user models.User) (models.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, err
	}
	query := `INSERT INTO users (name, email, password, country) VALUES ($1, $2, $3, $4) RETURNING id, role`
	var newUser models.User
	err = db.DB.QueryRow(ctx, query, user.Name, user.Email, hashedPassword, user.Country).Scan(
		&newUser.ID, &newUser.Role,
	)
	if err != nil {
		return models.User{}, err
	}
	return newUser, nil
}

func SignUp(w http.ResponseWriter, r *http.Request) {

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	ctx := r.Context()

	existingUser, err := GetUserByEmail(ctx, user.Email)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
		return
	}
	if existingUser.Email != "" {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "User already exists"})
		return
	}
	if user.Name == "" || user.Email == "" || user.Password == "" || user.Country == "" {
		utils.JSONResponse(w, http.StatusBadRequest, map[string]string{"message": "All fields are required"})
		return
	}
	newUser, err := createUser(ctx, user)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
		return
	}
	emailService, err := services.NewEmailService()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Error creating email service"})
		return
	}
	err = emailService.SendEmail(newUser.ID, newUser.Role, user.Email, "Welcome to MyBlog", "Thank you for signing up with MyBlog")
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Error sending email"})
		return
	}
	utils.JSONResponse(w, http.StatusCreated, map[string]string{"message": "User created successfully"})

}
