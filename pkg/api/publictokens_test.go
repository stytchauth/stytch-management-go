package api_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/publictokens"
	"testing"
)

func TestPublicTokensClient_CreatePublicToken(t *testing.T) {
	// Arrange
	projectID := GetProjectID(t)
	client := NewTestClient(t)
	ctx := context.Background()

	// Act
	_, err := client.PublicTokens.Create(ctx, publictokens.CreatePublicTokenRequest{
		ProjectID: projectID,
	})

	// Assert
	assert.NoError(t, err)
}

func TestPublicTokensClient_GetPublicTokens(t *testing.T) {
	// Arrange
	projectID := GetProjectID(t)
	client := NewTestClient(t)
	ctx := context.Background()
	createResp, err := client.PublicTokens.Create(ctx, publictokens.CreatePublicTokenRequest{
		ProjectID: projectID,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.PublicTokens.GetAll(ctx, publictokens.GetAllPublicTokensRequest{
		ProjectID: projectID,
	})
	var tokens []string
	for _, token := range resp.PublicTokens {
		tokens = append(tokens, token.PublicToken)
	}

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, tokens, createResp.PublicToken)
}

func TestPublicTokensClient_DeletePublicToken(t *testing.T) {
	// Arrange
	projectID := GetProjectID(t)
	client := NewTestClient(t)
	ctx := context.Background()
	createResp, err := client.PublicTokens.Create(ctx, publictokens.CreatePublicTokenRequest{
		ProjectID: projectID,
	})
	require.NoError(t, err)

	// Act
	_, err = client.PublicTokens.Delete(ctx, publictokens.DeletePublicTokenRequest{
		ProjectID:     projectID,
		PublicTokenID: createResp.PublicToken,
	})

	// Assert
	assert.NoError(t, err)
}
