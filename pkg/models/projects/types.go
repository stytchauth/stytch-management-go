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

// Project encompasses the relevant fields for a live and test project
type Project struct {
	// LiveProjectID is the unique identifier for the live project
	LiveProjectID string `json:"live_project_id"`
	// TestProjectID is the unique identifier for the test project
	TestProjectID string `json:"test_project_id"`
	// Name is the project's name
	Name string `json:"name"`
	// LiveOAuthCallbackID is the callback ID used in OAuth requests for the live project
	LiveOAuthCallbackID string `json:"live_oauth_callback_id"`
	// TestOAuthCallbackID is the callback ID used in OAuth requests for the test project
	TestOAuthCallbackID string `json:"test_oauth_callback_id"`
	// Vertical is the project's vertical
	Vertical Vertical `json:"vertical"`
	// CreatedAt is the ISO-8601 timestamp for when the project was created
	CreatedAt time.Time `json:"created_at"`
	// TestUserImpersonationEnabled is a boolean indicating whether user impersonation is enabled for the test project
	TestUserImpersonationEnabled bool `json:"test_user_impersonation_enabled"`
	// LiveUserImpersonationEnabled is a boolean indicating whether user impersonation is enabled for the live project
	LiveUserImpersonationEnabled bool `json:"live_user_impersonation_enabled"`
	// TestCrossOrgPasswordsEnabled is a boolean indicating whether the test project uses cross-org passwords
	TestCrossOrgPasswordsEnabled bool `json:"test_cross_org_password_enabled"`
	// LiveCrossOrgPasswordsEnabled is a boolean indicating whether the live project uses cross-org passwords
	LiveCrossOrgPasswordsEnabled bool `json:"live_cross_org_password_enabled"`

	LiveIDPAuthorizationURL                 string `json:"live_idp_authorization_url"`
	TestIDPAuthorizationURL                 string `json:"test_idp_authorization_url"`
	LiveIDPDynamicClientRegistrationEnabled bool   `json:"live_idp_dynamic_client_registration_enabled"`
	TestIDPDynamicClientRegistrationEnabled bool   `json:"test_idp_dynamic_client_registration_enabled"`
}

type CreateRequest struct {
	// ProjectName is the name of the project
	ProjectName string `json:"project_name"`
	// Vertical is the project's vertical
	Vertical Vertical `json:"vertical"`
	// TestUserImpersonationEnabled is a boolean indicating whether user impersonation is enabled for the test project
	TestUserImpersonationEnabled bool `json:"test_user_impersonation_enabled"`
	// LiveUserImpersonationEnabled is a boolean indicating whether user impersonation is enabled for the live project
	LiveUserImpersonationEnabled bool `json:"live_user_impersonation_enabled"`
	// TestCrossOrgPasswordsEnabled is a boolean indicating whether cross org passwords are enabled for the test project
	TestCrossOrgPasswordsEnabled bool `json:"test_cross_org_password_enabled"`
	// LiveCrossOrgPasswordsEnabled is a boolean indicating whether cross org passwords are enabled for the live project
	LiveCrossOrgPasswordsEnabled bool `json:"live_cross_org_password_enabled"`
}

type CreateResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Project contains the project details for the newly created live and test project
	Project Project `json:"project"`
}

type GetRequest struct {
	// ProjectID is the unique identifier for the project to retrieve
	ProjectID string `json:"project_id"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Project contains the project details for the requested project
	Project Project `json:"project"`
}

type GetAllRequest struct{}

type GetAllResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Projects is a list of all active projects in the workspace
	Projects []Project `json:"projects"`
}

type DeleteRequest struct {
	// ProjectID is the live project ID of the project to delete
	ProjectID string `json:"project_id"`
}

type DeleteResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
}

type UpdateRequest struct {
	// ProjectID is the unique id for the live project to update
	ProjectID string `json:"project_id"`
	// Name is the new name for the project
	Name string `json:"name"`
	// LiveUserImpersonationEnabled is a boolean indicating whether user impersonation is enabled for the project
	LiveUserImpersonationEnabled *bool `json:"live_user_impersonation_enabled"`
	// TestUserImpersonationEnabled is a boolean indicating whether user impersonation is enabled for the test project
	TestUserImpersonationEnabled *bool `json:"test_user_impersonation_enabled"`
	// LiveUseCrossOrgPasswords is a boolean indicating whether the project uses cross-org passwords
	LiveUseCrossOrgPasswords *bool `json:"live_use_cross_org_passwords"`
	// TestUseCrossOrgPasswords is a boolean indicating whether the test project uses cross-org passwords
	TestUseCrossOrgPasswords *bool `json:"test_use_cross_org_passwords"`

	LiveIdpDynamicClientRegistrationEnabled *bool `json:"live_idp_dynamic_client_registration_enabled,omitempty"`
	TestIdpDynamicClientRegistrationEnabled *bool `json:"test_idp_dynamic_client_registration_enabled,omitempty"`
}

type UpdateResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Project contains the updated project details
	Project Project `json:"project"`
}
