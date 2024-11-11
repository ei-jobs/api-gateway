package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/aidosgal/ei-jobs-core/internal/http/handler"
	"github.com/aidosgal/ei-jobs-core/internal/repository"
	"github.com/aidosgal/ei-jobs-core/internal/service"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type APIServer struct {
	addr string
	db   *sql.DB
}

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db:   db,
	}
}

func (s *APIServer) Run() error {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)

	userRepository := repository.NewUserRepository(s.db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	router.Route("/api/v1", func(router chi.Router) {
		router.Route("/user", func(router chi.Router) {
			router.Get("/", userHandler.GetAllUsers)
		})
	})

	log.Print("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
