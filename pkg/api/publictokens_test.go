package api_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/publictokens"
	"os"
	"testing"
)

func TestPublicTokensClient_CreatePublicToken(t *testing.T) {
	// Arrange
	projectID := os.Getenv("STYTCH_PROJECT_ID")
	if projectID == "" {
		t.Skip("STYTCH_PROJECT_ID environment variable is required for this test")
	}

	// Act
	client := NewTestClient(t)
	ctx := context.Background()
	_, err := client.PublicTokens.CreatePublicToken(ctx, publictokens.CreatePublicTokenRequest{
		ProjectID: projectID,
	})

	// Assert
	assert.NoError(t, err)
}

func TestPublicTokensClient_GetPublicTokens(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	ctx := context.Background()
	projectID := os.Getenv("STYTCH_PROJECT_ID")
	if projectID == "" {
		t.Skip("STYTCH_PROJECT_ID environment variable is required for this test")
	}
	createResp, err := client.PublicTokens.CreatePublicToken(ctx, publictokens.CreatePublicTokenRequest{
		ProjectID: projectID,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.PublicTokens.GetPublicTokens(ctx, publictokens.GetPublicTokensRequest{
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
	client := NewTestClient(t)
	ctx := context.Background()
	projectID := os.Getenv("STYTCH_PROJECT_ID")
	if projectID == "" {
		t.Skip("STYTCH_PROJECT_ID environment variable is required for this test")
	}
	createResp, err := client.PublicTokens.CreatePublicToken(ctx, publictokens.CreatePublicTokenRequest{
		ProjectID: projectID,
	})
	require.NoError(t, err)

	// Act
	_, err = client.PublicTokens.DeletePublicToken(ctx, publictokens.DeletePublicTokenRequest{
		ProjectID:     projectID,
		PublicTokenID: createResp.PublicToken,
	})

	// Assert
	assert.NoError(t, err)
}
