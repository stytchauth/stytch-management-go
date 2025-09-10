package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/publictokens"
)

func TestPublicTokensClient_Create(t *testing.T) {
	t.Run("create public token", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.PublicTokens.Create(ctx, publictokens.CreateRequest{
			Project:     project.Project,
			Environment: TestEnvironment,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.PublicToken.PublicToken)
		assert.True(t, len(resp.PublicToken.PublicToken) > 10)
		assert.False(t, resp.PublicToken.CreatedAt.IsZero())
	})
}

func TestPublicTokensClient_Get(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create a public token.
		createResp, err := client.PublicTokens.Create(ctx, publictokens.CreateRequest{
			Project:     project.Project,
			Environment: TestEnvironment,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.PublicTokens.Get(ctx, publictokens.GetRequest{
			Project:     project.Project,
			Environment: TestEnvironment,
			PublicToken: createResp.PublicToken.PublicToken,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, createResp.PublicToken, resp.PublicToken)
	})
	t.Run("missing public token", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.PublicTokens.Get(ctx, publictokens.GetRequest{
			Project:     project.Project,
			Environment: TestEnvironment,
			// PublicToken field is intentionally omitted.
		})

		// Assert
		assert.ErrorContains(t, err, "public token")
		assert.Nil(t, resp)
	})
}

func TestPublicTokensClient_GetAll(t *testing.T) {
	t.Run("get all public tokens", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create a few public tokens first
		var createdTokens []publictokens.PublicToken
		for i := 0; i < 3; i++ {
			createResp, err := client.PublicTokens.Create(ctx, publictokens.CreateRequest{
				Project:     project.Project,
				Environment: TestEnvironment,
			})
			require.NoError(t, err)
			createdTokens = append(createdTokens, createResp.PublicToken)
		}

		// Act
		resp, err := client.PublicTokens.GetAll(ctx, publictokens.GetAllRequest{
			Project:     project.Project,
			Environment: TestEnvironment,
		})

		// Assert
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.PublicTokens), 3)

		// Verify all created tokens are returned
		tokenValues := make(map[string]bool)
		for _, token := range resp.PublicTokens {
			tokenValues[token.PublicToken] = true
			assert.NotEmpty(t, token.PublicToken)
			assert.False(t, token.CreatedAt.IsZero())
		}

		for _, createdToken := range createdTokens {
			assert.True(t, tokenValues[createdToken.PublicToken], "Created token %s not found in response", createdToken.PublicToken)
		}
	})
}

func TestPublicTokensClient_Delete(t *testing.T) {
	t.Run("delete existing public token", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create a public token first
		createResp, err := client.PublicTokens.Create(ctx, publictokens.CreateRequest{
			Project:     project.Project,
			Environment: TestEnvironment,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.PublicTokens.Delete(ctx, publictokens.DeleteRequest{
			Project:     project.Project,
			Environment: TestEnvironment,
			PublicToken: createResp.PublicToken.PublicToken,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.RequestID)

		// Verify token is deleted by checking GetAll doesn't include it
		getAllResp, err := client.PublicTokens.GetAll(ctx, publictokens.GetAllRequest{
			Project:     project.Project,
			Environment: TestEnvironment,
		})
		require.NoError(t, err)

		for _, token := range getAllResp.PublicTokens {
			assert.NotEqual(t, createResp.PublicToken.PublicToken, token.PublicToken, "Deleted token should not appear in GetAll response")
		}
	})

	t.Run("public token does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.PublicTokens.Delete(ctx, publictokens.DeleteRequest{
			Project:     project.Project,
			Environment: TestEnvironment,
			PublicToken: "public-token-does-not-exist",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
