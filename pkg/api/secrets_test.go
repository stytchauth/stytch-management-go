package api_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/secrets"
	"testing"
)

func TestSecretsClient_Create(t *testing.T) {
	// Arrange
	projectID := GetProjectID(t)
	client := NewTestClient(t)
	ctx := context.Background()

	// Act
	_, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
		ProjectID: projectID,
	})

	// Assert
	assert.NoError(t, err)
}

func TestSecretsClient_GetAll(t *testing.T) {
	// Arrange
	projectID := GetProjectID(t)
	client := NewTestClient(t)
	ctx := context.Background()
	createResp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
		ProjectID: projectID,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.Secrets.GetAll(ctx, secrets.GetAllSecretsRequest{
		ProjectID: projectID,
	})
	var secretIDs []string
	for _, secret := range resp.Secrets {
		secretIDs = append(secretIDs, secret.SecretID)
	}

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, secretIDs, createResp.SecretID)

}

func TestSecretsClient_Delete(t *testing.T) {
	// Arrange
	projectID := GetProjectID(t)
	client := NewTestClient(t)
	ctx := context.Background()
	createResp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
		ProjectID: projectID,
	})
	require.NoError(t, err)

	resp, err := client.Secrets.Delete(ctx, secrets.DeleteSecretRequest{
		ProjectID: projectID,
		SecretID:  createResp.SecretID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, createResp.SecretID, resp.SecretID)
}
