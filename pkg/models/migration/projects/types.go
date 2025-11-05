package projects

import "time"

// LegacyProject represents the payload returned by the legacy PWA v1 project endpoints.
// It includes just enough information to map and validate between PWA v1 and PWA v3 project models.
// This is not meant to be used outside of migration paths between PWA v1 and PWA v3.
type LegacyProject struct {
	// The legacy unique identifier used in PWA V1 for a live environment.
	LiveProjectID string `json:"live_project_id"`
	// The legacy unique identifier used in PWA V1 for a test environment.
	TestProjectID string `json:"test_project_id"`
	// The PWA V3 identifier (slug) for the live environment.
	LiveEnvironmentSlug string `json:"live_project_slug"`
	// The PWA V3 identifier (slug) for the test environment.
	TestEnvironmentSlug string `json:"test_project_slug"`
	// The PWA V3 identifier (slug) for the project.
	ProjectSlug string `json:"project_slug"`

	// The following can be used to validate that the correct project is being referenced during migration operations.
	// The name of the project.
	Name string `json:"name"`
	// The vertical associated with the project.
	Vertical string `json:"vertical"`
	// The timestamp when the project was created.
	CreatedAt time.Time `json:"created_at"`
}

// GetProjectRequest represents the request payload for retrieving a legacy project by ID.
type GetProjectRequest struct {
	// The legacy unique identifier used in PWA V1 for projects.
	ProjectID string
}

// GetProjectResponse represents the response payload for retrieving project information by legacy ID.
type GetProjectResponse struct {
	// The unique request ID associated with the API call.
	RequestID string `json:"request_id"`
	// The project identifier details.
	Project LegacyProject `json:"project"`
}

// GetProjectsRequest represents the request payload for retrieving all projects.
type GetProjectsRequest struct{}

// GetProjectsResponse represents the response payload for retrieving all projects.
type GetProjectsResponse struct {
	// The unique request ID associated with the API call.
	RequestID string `json:"request_id"`
	// The list of all projects, with their PWA V1 and PWA V3 identifiers.
	Projects []LegacyProject `json:"projects"`
}
