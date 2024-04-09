package models

type Player struct {
	ModelDefaults
	Username string `json:"username"`
	License  string `json:"license"`
	Steam    string `json:"steam_id"`
	FiveM    string `json:"fivem_id"`
	Discord  string `json:"discord_id"`
}
