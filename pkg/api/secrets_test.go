package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/secrets"
)

func TestSecretsClient_Create(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()

	// Act
	_, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
		ProjectID: project.LiveProjectID,
	})

	// Assert
	assert.NoError(t, err)
}

func TestSecretsClient_GetAll(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
		ProjectID: project.LiveProjectID,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.Secrets.GetAll(ctx, secrets.GetAllSecretsRequest{
		ProjectID: project.LiveProjectID,
	})
	var secretIDs []string
	for _, secret := range resp.Secrets {
		secretIDs = append(secretIDs, secret.SecretID)
	}

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, secretIDs, createResp.CreatedSecret.SecretID)
}

func TestSecretsClient_Get(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
		ProjectID: project.LiveProjectID,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.Secrets.Get(ctx, secrets.GetSecretRequest{
		ProjectID: project.LiveProjectID,
		SecretID:  createResp.CreatedSecret.SecretID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, createResp.CreatedSecret.SecretID, resp.Secret.SecretID)
	assert.Equal(t, createResp.CreatedSecret.LastFour, resp.Secret.LastFour)
}

func TestSecretsClient_Delete(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
		ProjectID: project.LiveProjectID,
	})
	require.NoError(t, err)

	_, err = client.Secrets.Delete(ctx, secrets.DeleteSecretRequest{
		ProjectID: project.LiveProjectID,
		SecretID:  createResp.CreatedSecret.SecretID,
	})

	// Assert
	assert.NoError(t, err)
}
