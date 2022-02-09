package response

type HealthCheckResponse struct {
	Mongo      string `json:"mongo"`
	ServerPort string `json:"serverPort"`
	Database   string `json:"database"`
}

func NewHealthCheckResponse(status, database, port string) HealthCheckResponse {
	return HealthCheckResponse{
		Mongo:      status,
		ServerPort: port,
		Database:   database,
	}
}
