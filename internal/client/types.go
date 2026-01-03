package client

type CreateSessionRequest struct {
	Region string `json:"region"`
}

type Region struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type RegionsResponse struct {
	Regions []Region `json:"regions"`
}

type CreateSessionResponse struct {
	SessionID       string `json:"session_id"`
	WireGuardConfig string `json:"wireguard_config"`
}

type UsageResponse struct {
	SessionID         string  `json:"session_id"`
	Cost              float64 `json:"cost"`
	DurationInSeconds float64 `json:"duration"`
}

type DeleteSessionResponse struct {
	SessionID         string  `json:"session_id"`
	DurationInSeconds float64 `json:"duration"`
	Cost              float64 `json:"cost"`
}
