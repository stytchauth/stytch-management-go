package api

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/public_tokens"
)

type PublicTokensClient struct {
	client *internal.Client
}

func NewPublicTokensClient(c *internal.Client) *PublicTokensClient {
	return &PublicTokensClient{client: c}
}

func (c *PublicTokensClient) GetPublicTokens(ctx context.Context, body public_tokens.GetPublicTokensRequest) (*public_tokens.GetPublicTokensResponse, error) {

	var resp public_tokens.GetPublicTokensResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/web/v1/projects/%s/public_tokens", body.ProjectId),
		nil,
		nil,
		&resp)

	return &resp, err
}

func (c *PublicTokensClient) CreatePublicToken(ctx context.Context, body public_tokens.CreatePublicTokenRequest) (*public_tokens.CreatePublicTokenResponse, error) {
	jsonBody, err := json.Marshal(body)

	var res public_tokens.CreatePublicTokenResponse
	err = c.client.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/web/v1/projects/%s/public_tokens", body.ProjectId),
		nil,
		jsonBody,
		&res)
	return &res, err
}

func (c *PublicTokensClient) DeletePublicToken(ctx context.Context, body public_tokens.DeletePublicTokenRequest) error {
	return c.client.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/web/v1/projects/%s/public_tokens/%s", body.ProjectId, body.PublicTokenId),
		nil,
		nil,
		nil)
}
