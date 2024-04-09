package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"gorm.io/gorm"
)

type FiveMServer struct {
	Router    *chi.Mux
	Port      int
	StartTime time.Time
	DB        *gorm.DB
}

func NewServer(router *chi.Mux, port int, db *gorm.DB) *FiveMServer {
	startTime := time.Now()
	return &FiveMServer{
		Router:    router,
		Port:      port,
		StartTime: startTime,
		DB:        db,
	}
}

func (s *FiveMServer) StartServer() error {
	srv := &http.Server{
		Addr:         ":" + strconv.Itoa(s.Port),
		Handler:      s.Router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	err := srv.ListenAndServe()

	return err
}

func (s *FiveMServer) CreateRoutes() {
	s.Router.Get("/", s.GetHealthHandler)
	s.Router.Get("/players", s.GetPlayersHandler)
	s.Router.Post("/players", s.CreateNewPlayer)
}
