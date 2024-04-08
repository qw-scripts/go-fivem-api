package health

import (
	"net/http"

	"github.com/masonschafercodes/go-fivem-api/internal/utils"
)

func GetHealthHandler(w http.ResponseWriter, r *http.Request) {
	utils.RespondWithJSON(w, 200, HealthResponse{
		Status: http.StatusText(200),
	})
}
