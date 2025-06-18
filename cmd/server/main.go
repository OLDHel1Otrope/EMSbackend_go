package server

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"server.go/database"
	"server.go/internal/handler"
	"server.go/internal/middlewares"
	"server.go/internal/repository"
	"server.go/internal/service"
)

func StartServer() {
	if err := database.ConnectAndMigrate(os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		database.GetSSLMode()); err != nil {
		logrus.WithError(err).Panic("Failed to initialize and migrate database")
	}

	userRepo := repository.NewUserRepository(database.NoteDB)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	r := mux.NewRouter()
	sessionRepo := repository.NewSessionRepository(database.NoteDB)
	sessionService := service.NewSessionService(sessionRepo)

	authHandler := handler.NewAuthHandler(sessionService, userService)

	r.HandleFunc("/login", authHandler.Login).Methods("POST")
	r.HandleFunc("/signup", authHandler.Signup).Methods("POST")

	authMiddleware := &middlewares.AuthenticationMiddleware{
		SessionRepo: sessionRepo,
	}

	r.Use(authMiddleware.AuthMiddleware)
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST")
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", userHandler.CreateUser).Methods("PATCH")

	r.HandleFunc("/users/{id}/notes").Methods("GET")
	r.HandleFunc("/users/note").Methods("POST")
	r.HandleFunc("/users/{id}/notes/{note_id}").Methods("PATCH", "DELETE")

	log.Fatal(http.ListenAndServe(":8080", r))
}
