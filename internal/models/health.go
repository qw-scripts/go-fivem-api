package models

type HealthChecks struct {
	AllocatedHeap string `json:"allocated_heap"`
	LastGC        string `json:"last_gc"`
}

type HealthResponse struct {
	Status    string       `json:"status"`
	StartedAt string       `json:"started_at"`
	Uptime    string       `json:"uptime"`
	Checks    HealthChecks `json:"checks"`
}
