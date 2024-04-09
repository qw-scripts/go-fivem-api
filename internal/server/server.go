package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
)

type FiveMServer struct {
	Router    *chi.Mux
	Port      int
	StartTime time.Time
}

func NewServer(router *chi.Mux, port int) *FiveMServer {
	startTime := time.Now()
	return &FiveMServer{
		Router:    router,
		Port:      port,
		StartTime: startTime,
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
}
