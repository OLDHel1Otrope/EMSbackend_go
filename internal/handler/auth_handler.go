package handler

import (
	"encoding/json"
	"net/http"

	"server.go/internal/model"
	"server.go/internal/service"
)

type AuthHandler struct {
	sessionService *service.SessionService
	userService    *service.UserService
}

func NewAuthHandler(sessionService *service.SessionService, userService *service.UserService) *AuthHandler {
	return &AuthHandler{
		sessionService: sessionService,
		userService:    userService,
	}
}

func (h *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var creds struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	user, err := h.userService.ValidateUser(creds.Email, creds.Password)
	if err != nil {
		http.Error(w, "Invalid credentials", http.StatusUnauthorized)
		return
	}

	session, err := h.sessionService.Login(user)
	if err != nil {
		http.Error(w, "Could not create session", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(session)
}

func (h *AuthHandler) Signup(w http.ResponseWriter, r *http.Request) {
	var user model.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userService.CreateUser(&user)
	if err != nil {
		http.Error(w, "Could not create user", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdUser)
}
