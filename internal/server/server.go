package server

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/masonschafercodes/go-fivem-api/internal/routes/health"
)

type FiveMServer struct {
	Router *chi.Mux
	Port   int
}

func NewServer(router *chi.Mux, port int) *FiveMServer {
	return &FiveMServer{
		Router: router,
		Port:   port,
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
	s.Router.Get("/", health.Handler)
}
