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
	// Project is the project for which to create the public token.
	Project string `json:"-"`
	// Environment is the environment for which to create the public token.
	Environment string `json:"-"`
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
	// Project is the project for which to retrieve the public token.
	Project string `json:"-"`
	// Environment is the environment for which to retrieve the public token.
	Environment string `json:"-"`
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
	// Project is the project for which to retrieve the public tokens.
	Project string `json:"-"`
	// Environment is the environment for which to retrieve the public tokens.
	Environment string `json:"-"`
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
	// Project is the project where the public token is located.
	Project string `json:"-"`
	// Environment is the environment where the public token is located.
	Environment string `json:"-"`
	// PublicToken is the public token to delete.
	PublicToken string `json:"public_token"`
}

type DeleteResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging.
	RequestID string `json:"request_id"`
}
