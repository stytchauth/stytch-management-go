package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/secrets"
)

func TestSecretsClient_CreateSecret(t *testing.T) {
	t.Run("create secret", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.Secrets.Create(ctx, secrets.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Secret.SecretID)
		assert.NotEmpty(t, resp.Secret.Secret)
		assert.True(t, len(resp.Secret.Secret) > 10)
		assert.False(t, resp.Secret.CreatedAt.IsZero())
	})
}

func TestSecretsClient_GetSecret(t *testing.T) {
	t.Run("get existing secret", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Create a secret first
		createResp, err := client.Secrets.Create(ctx, secrets.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.Secrets.Get(ctx, secrets.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			SecretID:        createResp.Secret.SecretID,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, createResp.Secret.SecretID, resp.Secret.SecretID)
		assert.NotEmpty(t, resp.Secret.LastFour)
		assert.False(t, resp.Secret.CreatedAt.IsZero())
		assert.Equal(t, createResp.Secret.Secret[len(createResp.Secret.Secret)-4:], resp.Secret.LastFour)
		assert.Equal(t, createResp.Secret.CreatedAt, resp.Secret.CreatedAt)
	})
	t.Run("secret does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.Secrets.Get(ctx, secrets.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			SecretID:        "secret-does-not-exist",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
	t.Run("missing secret ID", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.Secrets.Get(ctx, secrets.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			// SecretID is intentionally omitted.
		})

		// Assert
		assert.ErrorContains(t, err, "cannot be empty")
		assert.Nil(t, resp)
	})
}

func TestSecretsClient_GetAllSecrets(t *testing.T) {
	t.Run("get all secrets", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Create a few secrets first
		var createdSecrets []secrets.Secret
		for i := 0; i < 3; i++ {
			createResp, err := client.Secrets.Create(ctx, secrets.CreateRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
			})
			require.NoError(t, err)
			createdSecrets = append(createdSecrets, createResp.Secret)
		}

		// Act
		resp, err := client.Secrets.GetAll(ctx, secrets.GetAllRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
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
}

func TestSecretsClient_DeleteSecret(t *testing.T) {
	t.Run("delete existing secret", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Create a secret first
		createResp, err := client.Secrets.Create(ctx, secrets.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.Secrets.Delete(ctx, secrets.DeleteRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			SecretID:        createResp.Secret.SecretID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.RequestID)

		// Verify secret is deleted by trying to get it
		getResp, err := client.Secrets.Get(ctx, secrets.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			SecretID:        createResp.Secret.SecretID,
		})
		assert.Error(t, err)
		assert.Nil(t, getResp)
	})

	t.Run("secret does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		_, err := client.Secrets.Delete(ctx, secrets.DeleteRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			SecretID:        "secret-does-not-exist",
		})

		// Assert
		assert.Error(t, err)
	})
}
