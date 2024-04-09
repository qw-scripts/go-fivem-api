package server

import (
	"fmt"
	"net/http"
	"runtime"
	"time"

	"github.com/masonschafercodes/go-fivem-api/internal/models"
	"github.com/masonschafercodes/go-fivem-api/pkg/utils"
)

func (s *FiveMServer) GetHealthHandler(w http.ResponseWriter, r *http.Request) {
	var mem runtime.MemStats
	runtime.ReadMemStats(&mem)
	allocatedHeapMB := float64(mem.HeapAlloc) / 1024 / 1024
	lastGC := time.Unix(0, int64(mem.LastGC)).String()

	r.Header.Set("Cache-Control", "no-cache")

	utils.RespondWithJSON(w, http.StatusOK, models.HealthResponse{
		Status:    http.StatusText(http.StatusOK),
		StartedAt: s.StartTime.String(),
		Uptime:    time.Since(s.StartTime).String(),
		Checks: models.HealthChecks{
			AllocatedHeap: fmt.Sprintf("%f MB", allocatedHeapMB),
			LastGC:        lastGC,
		},
	})
}
