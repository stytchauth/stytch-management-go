package publictokens

import "time"

// PublicToken represents a public token for a project. This token can be used for SDK authentication and OAuth integrations.
type PublicToken struct {
	// ProjectID is the unique ID of the project to which the token belongs
	ProjectID string `json:"project_id"`
	// PublicToken is the public token value. This is a unique ID which is also the identifier for the token.
	PublicToken string `json:"public_token"`
	// CreatedAt is the ISO-8601 timestamp for when the object was created
	CreatedAt time.Time `json:"created_at"`
}

type GetAllPublicTokensRequest struct {
	// ProjectID is the project to retrieve the public tokens for
	ProjectID string `json:"project_id"`
}

type GetAllPublicTokensResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// PublicTokens is all the public tokens that belong to the project
	PublicTokens []PublicToken `json:"public_tokens"`
}

type CreatePublicTokenRequest struct {
	// ProjectID is the project to create the public token for
	ProjectID string `json:"project_id"`
}

type CreatePublicTokenResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// PublicToken is the public token that was created
	PublicToken PublicToken `json:"public_token"`
}

type DeletePublicTokenRequest struct {
	// ProjectID is the project where the public token is located
	ProjectID string `json:"project_id"`
	// PublicToken is the public token to delete
	PublicToken string `json:"public_token"`
}

type DeletePublicTokenResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging
	RequestID string `json:"request_id"`
}
