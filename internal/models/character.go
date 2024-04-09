package models

import "time"

type Character struct {
	ModelDefaults
	PlayerID          uint              `json:"player_id"`
	StateID           string            `json:"state_id"`
	FirstName         string            `json:"first_name"`
	LastName          string            `json:"last_name"`
	Gender            string            `json:"gender"`
	DateOfBirth       time.Time         `json:"date_of_birth"`
	LastPlayed        time.Time         `json:"last_played"`
	IsDead            uint8             `json:"is_dead"`
	CharacterLocation CharacterLocation `json:"character_location"`
}

type CharacterLocation struct {
	ModelDefaults
	CharacterID uint    `json:"character_id"`
	X           float32 `json:"x"`
	Y           float32 `json:"y"`
	Z           float32 `json:"z"`
	Heading     float32 `json:"heading"`
}
