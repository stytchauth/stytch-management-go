package public_tokens

type PublicToken struct {
	ProjectId   string `json:"project_id"`
	PublicToken string `json:"public_token"`
	CreatedAt   string `json:"created_at"`
}

type GetPublicTokensRequest struct {
	ProjectId string `json:"project_id"`
}

type GetPublicTokensResponse struct {
	ProjectId    string        `json:"project_id"`
	PublicTokens []PublicToken `json:"public_tokens"`
	RequestId    string        `json:"request_id"`
	StatusCode   int           `json:"status_code"`
}

type CreatePublicTokenRequest struct {
	ProjectId string `json:"project_id"`
}

type CreatePublicTokenResponse struct {
	ProjectId   string `json:"project_id"`
	PublicToken string `json:"public_token"`
	RequestId   string `json:"request_id"`
	CreatedAt   string `json:"created_at"`
	StatusCode  int    `json:"status_code"`
}

type DeletePublicTokenRequest struct {
	ProjectId string `json:"project_id"`
	// This is the public token string
	PublicTokenId string `json:"public_token_id"`
}

type DeletePublicTokenResponse struct {
	RequestId  string `json:"request_id"`
	StatusCode int    `json:"status_code"`
}
