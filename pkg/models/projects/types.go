package projects

import "time"

type Vertical string

const (
	VerticalConsumer Vertical = "CONSUMER"
	VerticalB2B      Vertical = "B2B"
)

type Project struct {
	LiveProjectID       string    `json:"live_project_id"`
	TestProjectID       string    `json:"test_project_id"`
	Name                string    `json:"name"`
	LiveOAuthCallbackID string    `json:"live_oauth_callback_id"`
	TestOAuthCallbackID string    `json:"test_oauth_callback_id"`
	Vertical            Vertical  `json:"vertical"`
	CreatedAt           time.Time `json:"created_at"`
}

type CreateRequest struct {
	ProjectName string   `json:"project_name"`
	Vertical    Vertical `json:"vertical"`
}

type CreateResponse struct {
	StatusCode int     `json:"status_code"`
	RequestID  string  `json:"request_id"`
	Project    Project `json:"project"`
}

type GetRequest struct {
	ProjectID string `json:"project_id"`
}

type GetResponse struct {
	StatusCode int     `json:"status_code"`
	RequestID  string  `json:"request_id"`
	Project    Project `json:"project"`
}

type GetAllRequest struct{}

type GetAllResponse struct {
	StatusCode int       `json:"status_code"`
	RequestID  string    `json:"request_id"`
	Projects   []Project `json:"projects"`
}

type DeleteRequest struct {
	ProjectID string `json:"project_id"`
}

type DeleteResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id"`
}
