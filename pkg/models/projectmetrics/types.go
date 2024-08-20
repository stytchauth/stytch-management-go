package projectmetrics

// Metrics contains various metrics for a project, such as the count of active users or organizations.
type Metrics struct {
	// UserCount is the number of active users in the project (only relevant for consumer projects)
	UserCount uint32 `json:"user_count"`
	// OrganizationCount is the number of active organizations in the project (only relevant for B2B projects)
	OrganizationCount uint32 `json:"organization_count"`
	// MemberCount is the number of active members in the project (only relevant for B2B projects)
	MemberCount uint32 `json:"member_count"`
	// M2MClientCount is the number of active M2M clients in the project
	M2MClientCount uint32 `json:"m2m_client_count"`
}

type GetRequest struct {
	// ProjectID is the unique ID of the project for which to retrieve metrics
	ProjectID string `json:"project_id"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// ProjectID is the unique ID of the project for which the metrics were retrieved
	ProjectID string `json:"project_id"`
	// Metrics contains various metrics for the project
	Metrics Metrics `json:"metrics"`
}
