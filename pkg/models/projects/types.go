package projects

import "time"

type Vertical string

const (
	VerticalConsumer Vertical = "CONSUMER"
	VerticalB2B      Vertical = "B2B"
)

func Verticals() []Vertical {
	return []Vertical{
		VerticalConsumer,
		VerticalB2B,
	}
}

// Project represents a Stytch Project.
type Project struct {
	// ProjectSlug is the immutable unique identifier (slug) of the project.
	ProjectSlug string `json:"project_slug"`
	// Name is the project's name.
	Name string `json:"name"`
	// Vertical is the project's vertical.
	Vertical Vertical `json:"vertical"`
	// CreatedAt is the ISO-8601 timestamp for when the project was created.
	CreatedAt time.Time `json:"created_at"`
}

type CreateRequest struct {
	// Name is the name of the project.
	Name string `json:"name"`
	// Vertical is the project's vertical.
	Vertical Vertical `json:"vertical"`
	// ProjectSlug is the immutable unique identifier (slug) of the project, and cannot be changed
	// after the project is created. If not provided, a slug will be generated.
	ProjectSlug string `json:"project_slug"`
}

type CreateResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// Project contains the details for the newly created project.
	Project Project `json:"project"`
}

type GetRequest struct {
	// ProjectSlug is the slug of the project to retrieve.
	ProjectSlug string `json:"-"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// Project contains the details for the requested project.
	Project Project `json:"project"`
}

type GetAllRequest struct{}

type GetAllResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// Projects is a list of all active projects in the workspace.
	Projects []Project `json:"projects"`
}

type UpdateRequest struct {
	// ProjectSlug is the slug of the project to update.
	ProjectSlug string `json:"-"`
	// Name is the new name for the project.
	Name *string `json:"name,omitempty"`
}

type UpdateResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// Project contains the updated project details.
	Project Project `json:"project"`
}

type DeleteRequest struct {
	// ProjectSlug is the slug of the project to delete.
	ProjectSlug string `json:"-"`
}

type DeleteResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
}
