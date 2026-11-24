package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projecttwilioapikeys"
)

// These tests require the Twilio API keys feature flag to be enabled for the workspace that the
// test credentials belong to (in CI: workspace-staging-b949968e-11d2-4410-9faa-a64bb87b960d).

const (
	testTwilioAPIKeySid    = "SK00000000000000000000000000000000"
	testTwilioAPIKeySecret = "test_twilio_api_key_secret"
)

func TestProjectTwilioAPIKeysClient_CreateTwilioAPIKey(t *testing.T) {
	t.Run("create twilio api key", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.ProjectTwilioAPIKeys.CreateTwilioAPIKey(ctx, projecttwilioapikeys.CreateTwilioAPIKeyRequest{
			ProjectSlug:        env.ProjectSlug,
			EnvironmentSlug:    env.EnvironmentSlug,
			TwilioAPIKeySid:    testTwilioAPIKeySid,
			TwilioAPIKeySecret: testTwilioAPIKeySecret,
		})

		// Assert
		require.NoError(t, err)
		assert.Equal(t, testTwilioAPIKeySid, resp.TwilioAPIKeySid)
	})
	t.Run("empty project slug returns error", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		resp, err := client.ProjectTwilioAPIKeys.CreateTwilioAPIKey(ctx, projecttwilioapikeys.CreateTwilioAPIKeyRequest{
			EnvironmentSlug:    "environment",
			TwilioAPIKeySid:    testTwilioAPIKeySid,
			TwilioAPIKeySecret: testTwilioAPIKeySecret,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
	t.Run("empty environment slug returns error", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		resp, err := client.ProjectTwilioAPIKeys.CreateTwilioAPIKey(ctx, projecttwilioapikeys.CreateTwilioAPIKeyRequest{
			ProjectSlug:        "project",
			TwilioAPIKeySid:    testTwilioAPIKeySid,
			TwilioAPIKeySecret: testTwilioAPIKeySecret,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestProjectTwilioAPIKeysClient_GetTwilioAPIKey(t *testing.T) {
	t.Run("get existing twilio api key", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()
		_, err := client.ProjectTwilioAPIKeys.CreateTwilioAPIKey(ctx, projecttwilioapikeys.CreateTwilioAPIKeyRequest{
			ProjectSlug:        env.ProjectSlug,
			EnvironmentSlug:    env.EnvironmentSlug,
			TwilioAPIKeySid:    testTwilioAPIKeySid,
			TwilioAPIKeySecret: testTwilioAPIKeySecret,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.ProjectTwilioAPIKeys.GetTwilioAPIKey(ctx, projecttwilioapikeys.GetTwilioAPIKeyRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		// Assert
		require.NoError(t, err)
		assert.Equal(t, testTwilioAPIKeySid, resp.TwilioAPIKeySid)
	})
	t.Run("get twilio api key when none exists returns error", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.ProjectTwilioAPIKeys.GetTwilioAPIKey(ctx, projecttwilioapikeys.GetTwilioAPIKeyRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
	t.Run("empty slugs return error", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		resp, err := client.ProjectTwilioAPIKeys.GetTwilioAPIKey(ctx, projecttwilioapikeys.GetTwilioAPIKeyRequest{})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
