package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/qw-scripts/go-fivem-api/internal/models"
	"github.com/qw-scripts/go-fivem-api/internal/server"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	r := chi.NewRouter()

	db, err := gorm.Open(sqlite.Open("../../test.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&models.Player{})
	if err != nil {
		log.Panicf("failed to migrate model: %s", err.Error())
	}

	err = db.AutoMigrate(&models.Character{})
	if err != nil {
		log.Panicf("failed to migrate model: %s", err.Error())
	}

	err = db.AutoMigrate(&models.CharacterLocation{})
	if err != nil {
		log.Panicf("failed to migrate model: %s", err.Error())
	}

	portFromEnv := os.Getenv("PORT")

	if portFromEnv == "" {
		portFromEnv = "3030"
	}

	port, err := strconv.Atoi(portFromEnv)
	if err != nil {
		log.Panic("unable to parse server port")
	}

	svr := server.NewServer(r, port, db)

	svr.Router.Use(middleware.RequestID)
	svr.Router.Use(middleware.RealIP)
	svr.Router.Use(middleware.Logger)
	svr.Router.Use(middleware.Recoverer)
	svr.Router.Use(middleware.Timeout(60 * time.Second))

	svr.CreateRoutes()

	err = svr.StartServer()

	if err != nil {
		log.Fatalf("unable to start server: %s", err.Error())
	}
}
