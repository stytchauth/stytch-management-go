package publictokens

import "time"

type PublicToken struct {
	ProjectID   string    `json:"project_id"`
	PublicToken string    `json:"public_token"`
	CreatedAt   time.Time `json:"created_at"`
}

type GetAllPublicTokensRequest struct {
	ProjectID string `json:"project_id"`
}

type GetAllPublicTokensResponse struct {
	ProjectID    string        `json:"project_id"`
	PublicTokens []PublicToken `json:"public_tokens"`
	RequestID    string        `json:"request_id"`
	StatusCode   int           `json:"status_code"`
}

type CreatePublicTokenRequest struct {
	ProjectID string `json:"project_id"`
}

type CreatePublicTokenResponse struct {
	ProjectID   string    `json:"project_id"`
	PublicToken string    `json:"public_token"`
	RequestID   string    `json:"request_id"`
	CreatedAt   time.Time `json:"created_at"`
	StatusCode  int       `json:"status_code"`
}

type DeletePublicTokenRequest struct {
	ProjectID     string `json:"project_id"`
	PublicTokenID string `json:"public_token_id"`
}

type DeletePublicTokenResponse struct {
	RequestID  string `json:"request_id"`
	StatusCode int    `json:"status_code"`
}
