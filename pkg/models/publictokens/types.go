package publictokens

import "time"

// PublicToken represents a public token for an environment. This token can be used for SDK authentication and OAuth integrations.
type PublicToken struct {
	// PublicToken is the public token value. This is a unique ID which is also the identifier for the token.
	PublicToken string `json:"public_token"`
	// CreatedAt is the ISO-8601 timestamp for when the object was created
	CreatedAt time.Time `json:"created_at"`
}

type GetAllRequest struct {
	// Project is the project to retrieve the public tokens for
	Project string `json:"-"`
	// Environment is the environment to retrieve the public tokens for
	Environment string `json:"-"`
}

type GetAllResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// PublicTokens is all the public tokens that belong to the environment
	PublicTokens []PublicToken `json:"public_tokens"`
}

type CreateRequest struct {
	// Project is the project to create the public token for
	Project string `json:"-"`
	// Environment is the environment to create the public token for
	Environment string `json:"-"`
}

type CreateResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// PublicToken is the public token that was created
	PublicToken PublicToken `json:"public_token"`
}

type DeleteRequest struct {
	// Project is the project where the public token is located
	Project string `json:"-"`
	// Environment is the environment where the public token is located
	Environment string `json:"-"`
	// PublicToken is the public token to delete
	PublicToken string `json:"public_token"`
}

type DeleteResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging
	RequestID string `json:"request_id"`
}
