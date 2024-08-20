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
	StatusCode   int           `json:"status_code"`
	RequestID    string        `json:"request_id"`
	PublicTokens []PublicToken `json:"public_tokens"`
}

type CreatePublicTokenRequest struct {
	ProjectID string `json:"project_id"`
}

type CreatePublicTokenResponse struct {
	StatusCode  int         `json:"status_code"`
	RequestID   string      `json:"request_id"`
	PublicToken PublicToken `json:"public_token"`
}

type DeletePublicTokenRequest struct {
	ProjectID   string `json:"project_id"`
	PublicToken string `json:"public_token"`
}

type DeletePublicTokenResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id"`
}
