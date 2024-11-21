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

	resumeRepository := repository.NewResumeRepository(s.db)
	resumeService := service.NewResumeService(resumeRepository)
	resumeHandler := handler.NewResumeHandler(resumeService)

	vacanyRepository := repository.NewVacancyRepository(s.db)
	vacancyService := service.NewVacancyService(vacanyRepository)
	vacancyHandler := handler.NewVacancyHandler(vacancyService)

    assistanceRepository := repository.NewAssistanceRepository(s.db)
    assistanceService := service.NewAssistanceService(assistanceRepository)
    assistanceHandler := handler.NewAssistanceHandler(assistanceService)

	router.Route("/api/v1", func(router chi.Router) {
		router.Route("/user", func(router chi.Router) {
			router.Post("/login", userHandler.HandleLogin)
			router.Post("/register", userHandler.HandleRegister)
		})
		router.Route("/resume", func(router chi.Router) {
			router.Get("/{userID}", resumeHandler.GetResumesByUserID)
			router.Post("/", resumeHandler.CreateResume)
			router.Put("/{resumeID}", resumeHandler.UpdateResume)
			router.Delete("/{resumeID}", resumeHandler.DeleteResume)
		})
		router.Route("/vacancy", func(router chi.Router) {
			router.Get("/", vacancyHandler.GetAllVacancies)
			router.Get("/{id}", vacancyHandler.GetVacancy)
			router.Post("/", vacancyHandler.CreateVacancy)
			router.Put("/{id}", vacancyHandler.UpdateVacancy)
			router.Delete("/{id}", vacancyHandler.DeleteVacancy)
		})
        router.Route("/assitance", func(router chi.Router) {
			router.Get("/{userId}", assistanceHandler.GetAssistancesByUserId)
		})
	})

	log.Print("Listening on", s.addr)
	return http.ListenAndServe(s.addr, router)
}
