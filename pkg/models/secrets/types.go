package secrets

import "time"

// Secret represents a secret used for a project. This secret can be used for interaction with the Stytch API
type Secret struct {
	// SecretID is the unique ID of the secret in the project
	SecretID string `json:"secret_id"`
	// ProjectID is the unique ID of the project to which the secret belongs
	ProjectID string `json:"project_id"`
	// LastFour is the last four characters of the secret
	LastFour string `json:"last_four"`
	// CreatedAt is the ISO-8601 timestamp for when the object was created
	CreatedAt time.Time `json:"created_at"`
	// UsedAt is the ISO-8601 timestamp for when the secret was last used
	UsedAt time.Time `json:"used_at"`
}

// CreatedSecret represents a Secret that has just been created
type CreatedSecret struct {
	// SecretID is the unique ID of the secret in the project
	SecretID string `json:"secret_id"`
	// ProjectID is the unique ID of the project to which the secret belongs
	ProjectID string `json:"project_id"`
	// Secret is the secret value. This is only visible once upon secret creation
	Secret string `json:"secret"`
	// CreatedAt is the ISO-8601 timestamp for when the object was created
	CreatedAt time.Time `json:"created_at"`
	// LastFour is the last four characters of the secret
	LastFour string `json:"last_four"`
}

type GetAllSecretsRequest struct {
	// ProjectID is the project to retrieve the secrets for
	ProjectID string `json:"project_id"`
}

type GetAllSecretsResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Secrets is all the secrets that belong to the project
	Secrets []Secret `json:"secrets"`
}

type GetSecretRequest struct {
	// ProjectID is the project to retrieve the secret for
	ProjectID string `json:"project_id"`
	// SecretID is the ID of the secret to retrieve
	SecretID string `json:"secret_id"`
}

type GetSecretResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Secret is the secret that was retrieved
	Secret Secret `json:"secret"`
}

type CreateSecretRequest struct {
	// ProjectID is the project to create the secret for
	ProjectID string `json:"project_id"`
}

type CreateSecretResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Secret is the newly created secret. The value of this secret is only visible in this response.
	Secret CreatedSecret `json:"secret"`
}

type DeleteSecretRequest struct {
	// ProjectID is the project where the secret is located
	ProjectID string `json:"project_id"`
	// SecretID is the ID of the secret to delete
	SecretID string `json:"secret_id"`
}

type DeleteSecretResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging
	RequestID string `json:"request_id"`
}
