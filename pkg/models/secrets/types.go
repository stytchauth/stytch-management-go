package secrets

import "time"

type Secret struct {
	SecretID  string    `json:"secret_id"`
	ProjectID string    `json:"project_id"`
	LastFour  string    `json:"last_four"`
	CreatedAt time.Time `json:"created_at"`
	UsedAt    time.Time `json:"used_at"`
}

type GetAllSecretsRequest struct {
	ProjectID string `json:"project_id"`
}

type GetAllSecretsResponse struct {
	StatusCode int      `json:"status_code"`
	RequestID  string   `json:"request_id"`
	Secrets    []Secret `json:"secrets"`
}

type CreateSecretRequest struct {
	ProjectID string `json:"project_id"`
}

type CreateSecretResponse struct {
	StatusCode int       `json:"status_code"`
	RequestID  string    `json:"request_id"`
	SecretID   string    `json:"secret_id"`
	ProjectID  string    `json:"project_id"`
	Secret     string    `json:"secret"`
	CreatedAt  time.Time `json:"created_at"`
}

type DeleteSecretRequest struct {
	ProjectID string `json:"project_id"`
	SecretID  string `json:"secret_id"`
}

type DeleteSecretResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id"`
	SecretID   string `json:"secret_id"`
}
