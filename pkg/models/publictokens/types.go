package publictokens

import "time"

// PublicToken represents a public token for an environment. This token can be used for SDK
// authentication and OAuth integrations.
type PublicToken struct {
	// PublicToken is the public token value as well as the unique identifier for the token.
	PublicToken string `json:"public_token"`
	// CreatedAt is the ISO-8601 timestamp for when the object was created.
	CreatedAt time.Time `json:"created_at"`
}

type CreateRequest struct {
	// ProjectSlug is the slug of the project for which to create the public token.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to create the public token.
	EnvironmentSlug string `json:"-"`
}

type CreateResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// PublicToken is the public token that was created.
	PublicToken PublicToken `json:"public_token"`
}

type GetRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve the public token.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to retrieve the public token.
	EnvironmentSlug string `json:"-"`
	// PublicToken is the identifier of the public token to retrieve.
	PublicToken string `json:"public_token"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// PublicToken is the requested public token.
	PublicToken PublicToken `json:"public_token"`
}

type GetAllRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve the public tokens.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to retrieve the public tokens.
	EnvironmentSlug string `json:"-"`
}

type GetAllResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// PublicTokens is all the public tokens that belong to the environment.
	PublicTokens []PublicToken `json:"public_tokens"`
}

type DeleteRequest struct {
	// ProjectSlug is the slug of the project for which to delete the public token.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to delete the public token.
	EnvironmentSlug string `json:"-"`
	// PublicToken is the public token to delete.
	PublicToken string `json:"public_token"`
}

type DeleteResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging.
	RequestID string `json:"request_id"`
}
