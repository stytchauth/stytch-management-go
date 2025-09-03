package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/secrets"
)

func TestSecretsClient_CreateSecret(t *testing.T) {
	t.Run("create secret", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
			Project:     project.Project,
			Environment: "test",
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.CreatedSecret.SecretID)
		assert.NotEmpty(t, resp.CreatedSecret.Secret)
		assert.True(t, len(resp.CreatedSecret.Secret) > 10)
		assert.False(t, resp.CreatedSecret.CreatedAt.IsZero())
	})
}

func TestSecretsClient_GetSecret(t *testing.T) {
	t.Run("get existing secret", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create a secret first
		createResp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
			Project:     project.Project,
			Environment: "test",
		})
		require.NoError(t, err)

		// Act
		resp, err := client.Secrets.Get(ctx, secrets.GetSecretRequest{
			Project:     project.Project,
			Environment: "test",
			SecretID:    createResp.CreatedSecret.SecretID,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, createResp.CreatedSecret.SecretID, resp.Secret.SecretID)
		assert.NotEmpty(t, resp.Secret.LastFour)
		assert.False(t, resp.Secret.CreatedAt.IsZero())
		assert.Equal(t, createResp.CreatedSecret.Secret[len(createResp.CreatedSecret.Secret)-4:], resp.Secret.LastFour)
		assert.Equal(t, createResp.CreatedSecret.CreatedAt, resp.Secret.CreatedAt)
	})

	t.Run("secret does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.Secrets.Get(ctx, secrets.GetSecretRequest{
			Project:     project.Project,
			Environment: "test",
			SecretID:    "secret-does-not-exist",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestSecretsClient_GetAllSecrets(t *testing.T) {
	t.Run("get all secrets", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create a few secrets first
		var createdSecrets []secrets.CreatedSecret
		for i := 0; i < 3; i++ {
			createResp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
				Project:     project.Project,
				Environment: "test",
			})
			require.NoError(t, err)
			createdSecrets = append(createdSecrets, createResp.CreatedSecret)
		}

		// Act
		resp, err := client.Secrets.GetAll(ctx, secrets.GetAllSecretsRequest{
			Project:     project.Project,
			Environment: "test",
		})

		// Assert
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.Secrets), 3)

		// Verify all created secrets are returned
		secretIDs := make(map[string]bool)
		for _, secret := range resp.Secrets {
			secretIDs[secret.SecretID] = true
			assert.NotEmpty(t, secret.LastFour)
			assert.False(t, secret.CreatedAt.IsZero())
		}

		for _, createdSecret := range createdSecrets {
			assert.True(t, secretIDs[createdSecret.SecretID], "Created secret %s not found in response", createdSecret.SecretID)
		}
	})

	t.Run("only the default secret exists", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.Secrets.GetAll(ctx, secrets.GetAllSecretsRequest{
			Project:     project.Project,
			Environment: "test",
		})

		// Assert
		assert.NoError(t, err)
		assert.Len(t, resp.Secrets, 1)
	})
}

func TestSecretsClient_DeleteSecret(t *testing.T) {
	t.Run("delete existing secret", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create a secret first
		createResp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
			Project:     project.Project,
			Environment: "test",
		})
		require.NoError(t, err)

		// Act
		resp, err := client.Secrets.Delete(ctx, secrets.DeleteSecretRequest{
			Project:     project.Project,
			Environment: "test",
			SecretID:    createResp.CreatedSecret.SecretID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.RequestID)

		// Verify secret is deleted by trying to get it
		getResp, err := client.Secrets.Get(ctx, secrets.GetSecretRequest{
			Project:     project.Project,
			Environment: "test",
			SecretID:    createResp.CreatedSecret.SecretID,
		})
		assert.Error(t, err)
		assert.Nil(t, getResp)
	})

	t.Run("secret does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		_, err := client.Secrets.Delete(ctx, secrets.DeleteSecretRequest{
			Project:     project.Project,
			Environment: "test",
			SecretID:    "secret-does-not-exist",
		})

		// Assert
		assert.Error(t, err)
	})
}
