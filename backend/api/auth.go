package api

import (
	"encoding/json"
	"net/http"

	"POS/backend/database/models"
	"POS/backend/database/sqlite"

	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

type AuthResponse struct {
	Token string      `json:"token"`
	User  UserResponse `json:"user"`
}

type UserResponse struct {
	ID    uint   `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var req RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	if req.Email == "" || req.Password == "" || req.Name == "" {
		errorResponse(w, http.StatusBadRequest, "email, password y nombre requeridos")
		return
	}

	var existing models.User
	if sqlite.DB.Where("email = ?", req.Email).First(&existing).RowsAffected > 0 {
		errorResponse(w, http.StatusConflict, "el email ya esta registrado")
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "error al registrar")
		return
	}

	user := models.User{Email: req.Email, Password: string(hash), Name: req.Name}
	if err := sqlite.DB.Create(&user).Error; err != nil {
		errorResponse(w, http.StatusInternalServerError, "error al crear usuario")
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "error al generar token")
		return
	}

	jsonResponse(w, http.StatusCreated, AuthResponse{
		Token: token,
		User:  UserResponse{ID: user.ID, Email: user.Email, Name: user.Name},
	})
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var req LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		errorResponse(w, http.StatusBadRequest, "cuerpo invalido")
		return
	}
	if req.Email == "" || req.Password == "" {
		errorResponse(w, http.StatusBadRequest, "email y password requeridos")
		return
	}

	var user models.User
	if sqlite.DB.Where("email = ?", req.Email).First(&user).Error != nil {
		errorResponse(w, http.StatusUnauthorized, "credenciales invalidas")
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		errorResponse(w, http.StatusUnauthorized, "credenciales invalidas")
		return
	}

	token, err := generateToken(user.ID)
	if err != nil {
		errorResponse(w, http.StatusInternalServerError, "error al generar token")
		return
	}

	jsonResponse(w, http.StatusOK, AuthResponse{
		Token: token,
		User:  UserResponse{ID: user.ID, Email: user.Email, Name: user.Name},
	})
}
