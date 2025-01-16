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

func createUser(ctx context.Context, user models.User, w http.ResponseWriter) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Something went wrong while hashing the password"})
		return err
	}
	query := `INSERT INTO users (name, email, password, country) VALUES ($1, $2, $3, $4)`
	_, err = db.DB.Exec(ctx, query, user.Name, user.Email, hashedPassword, user.Country)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Internal server error"})
		return err
	}
	return nil
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
	err = createUser(ctx, user, w)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}
	emailService, err := services.NewEmailService()
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Error creating email service"})
		return
	}
	err = emailService.SendEmail(user.Email, "Welcome to MyBlog", "Thank you for signing up with MyBlog")
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, map[string]string{"message": "Error sending email"})
		return
	}
	utils.JSONResponse(w, http.StatusCreated, map[string]string{"message": "User created successfully"})

}
