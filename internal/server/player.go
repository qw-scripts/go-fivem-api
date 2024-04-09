package server

import (
	"log"
	"net/http"

	"github.com/masonschafercodes/go-fivem-api/internal/models"
	"github.com/masonschafercodes/go-fivem-api/pkg/utils"
)

func (s *FiveMServer) GetPlayersHandler(w http.ResponseWriter, r *http.Request) {
	var players []models.Player

	result := s.DB.Find(&players)

	if result.Error != nil {
		log.Printf("error while querying players: %s", result.Error.Error())
		utils.RespondWithJSON(w, 500, models.ApiError{
			Message: http.StatusText(500),
			Error:   result.Error.Error(),
		})
		return
	}

	utils.RespondWithJSON(w, 200, players)
}

func (s *FiveMServer) CreateNewPlayerHandler(w http.ResponseWriter, r *http.Request) {
	newPlayer := models.Player{
		Username: "qwade",
		License:  "qwadebot",
	}
	result := s.DB.Create(&newPlayer)

	if result.Error != nil {
		log.Printf("error while creating a player: %s", result.Error.Error())
		utils.RespondWithJSON(w, 500, models.ApiError{
			Message: http.StatusText(500),
			Error:   result.Error.Error(),
		})
		return
	}

	utils.RespondWithJSON(w, 200, &newPlayer)
}
