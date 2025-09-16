package environmentmetrics

// Metrics contains various metrics for an environment, such as the count of active users or
// organizations.
type Metrics struct {
	// UserCount is the number of active users in the environment (only relevant for Consumer
	// projects).
	UserCount uint32 `json:"user_count"`
	// OrganizationCount is the number of active organizations in the environment (only relevant for
	// B2B projects).
	OrganizationCount uint32 `json:"organization_count"`
	// MemberCount is the number of active members in the environment (only relevant for B2B
	// projects).
	MemberCount uint32 `json:"member_count"`
	// M2MClientCount is the number of active M2M clients in the environment.
	M2MClientCount uint32 `json:"m2m_client_count"`
}

type GetRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve metrics.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to retrieve metrics.
	EnvironmentSlug string `json:"-"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// Metrics contains various metrics for the environment.
	Metrics Metrics `json:"metrics"`
}
