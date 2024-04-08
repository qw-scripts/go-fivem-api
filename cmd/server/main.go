package main

import (
	"log"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/masonschafercodes/go-fivem-api/internal/server"
)

func main() {
	r := chi.NewRouter()

	svr := server.NewServer(r, 3030)

	svr.Router.Use(middleware.RequestID)
	svr.Router.Use(middleware.RealIP)
	svr.Router.Use(middleware.Logger)
	svr.Router.Use(middleware.Recoverer)
	svr.Router.Use(middleware.Timeout(60 * time.Second))

	svr.CreateRoutes()

	err := svr.StartServer()

	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}
