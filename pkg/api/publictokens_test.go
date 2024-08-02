package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/publictokens"
)

func TestPublicTokensClient_CreatePublicToken(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()

	// Act
	_, err := client.PublicTokens.Create(ctx, publictokens.CreatePublicTokenRequest{
		ProjectID: project.LiveProject.ID,
	})

	// Assert
	assert.NoError(t, err)
}

func TestPublicTokensClient_GetPublicTokens(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp, err := client.PublicTokens.Create(ctx, publictokens.CreatePublicTokenRequest{
		ProjectID: project.LiveProject.ID,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.PublicTokens.GetAll(ctx, publictokens.GetAllPublicTokensRequest{
		ProjectID: project.LiveProject.ID,
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
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp, err := client.PublicTokens.Create(ctx, publictokens.CreatePublicTokenRequest{
		ProjectID: project.LiveProject.ID,
	})
	require.NoError(t, err)

	// Act
	_, err = client.PublicTokens.Delete(ctx, publictokens.DeletePublicTokenRequest{
		ProjectID:   project.LiveProject.ID,
		PublicToken: createResp.PublicToken,
	})

	// Assert
	assert.NoError(t, err)
}
