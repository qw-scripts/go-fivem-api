package server

import (
	"log"
	"net/http"
	"time"

	"github.com/masonschafercodes/go-fivem-api/internal/models"
	"github.com/masonschafercodes/go-fivem-api/pkg/utils"
)

func (s *FiveMServer) GetCharactersHandler(w http.ResponseWriter, r *http.Request) {
	var characters []models.Character

	err := s.DB.Model(&models.Character{}).Preload("CharacterLocation").Find(&characters).Error

	if err != nil {
		log.Printf("error while querying characters: %s", err.Error())
		utils.RespondWithJSON(w, 500, models.ApiError{
			Message: http.StatusText(500),
			Error:   err.Error(),
		})
		return
	}

	utils.RespondWithJSON(w, 200, &characters)
}

func (s *FiveMServer) CreateNewCharacterHandler(w http.ResponseWriter, r *http.Request) {
	newCharacter := models.Character{
		PlayerID:    1,
		StateID:     "testing_id",
		FirstName:   "Qwade",
		LastName:    "Bot",
		Gender:      "Male",
		DateOfBirth: time.Now(),
		LastPlayed:  time.Now(),
		IsDead:      0,
	}

	result := s.DB.Create(&newCharacter)

	if result.Error != nil {
		log.Printf("error while creating a character: %s", result.Error.Error())
		utils.RespondWithJSON(w, 500, models.ApiError{
			Message: http.StatusText(500),
			Error:   result.Error.Error(),
		})
		return
	}

	newCharacterLocation := models.CharacterLocation{
		CharacterID: newCharacter.ID,
		X:           0.0,
		Y:           0.0,
		Z:           0.0,
		Heading:     180.0,
	}

	result = s.DB.Create(&newCharacterLocation)

	if result.Error != nil {
		log.Printf("error while creating character location: %s", result.Error.Error())
		utils.RespondWithJSON(w, 500, models.ApiError{
			Message: http.StatusText(500),
			Error:   result.Error.Error(),
		})
		return
	}

	newCharacter.CharacterLocation = newCharacterLocation
	result = s.DB.Save(&newCharacter)

	if result.Error != nil {
		log.Printf("error while updating character: %s", result.Error.Error())
		utils.RespondWithJSON(w, 500, models.ApiError{
			Message: http.StatusText(500),
			Error:   result.Error.Error(),
		})
		return
	}

	utils.RespondWithJSON(w, 200, &newCharacter)
}
