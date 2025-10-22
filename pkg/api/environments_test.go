package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/stytcherror"
)

func Test_EnvironmentsCreate(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		slug := "custom-slug"
		zeroDowntimeSessionMigrationURL := "https://example.com/migration"
		resp, err := client.Environments.Create(ctx, environments.CreateRequest{
			ProjectSlug:                     project.ProjectSlug,
			Name:                            "Test Environment",
			Type:                            environments.EnvironmentTypeTest,
			EnvironmentSlug:                 &slug,
			CrossOrgPasswordsEnabled:        ptr(true),
			UserImpersonationEnabled:        ptr(true),
			ZeroDowntimeSessionMigrationURL: ptr(zeroDowntimeSessionMigrationURL),
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Test Environment", resp.Environment.Name)
		assert.Equal(t, environments.EnvironmentTypeTest, resp.Environment.Type)
		assert.Equal(t, slug, resp.Environment.EnvironmentSlug)
		assert.True(t, resp.Environment.CrossOrgPasswordsEnabled)
		assert.True(t, resp.Environment.UserImpersonationEnabled)
		assert.Equal(t, zeroDowntimeSessionMigrationURL, resp.Environment.ZeroDowntimeSessionMigrationURL)
	})
	t.Run("user locking fields", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		userLockThreshold := int32(5)
		userLockTTL := int32(600)
		resp, err := client.Environments.Create(ctx, environments.CreateRequest{
			ProjectSlug:              project.ProjectSlug,
			Name:                     "Test Environment",
			Type:                     environments.EnvironmentTypeTest,
			UserLockSelfServeEnabled: ptr(true),
			UserLockThreshold:        ptr(userLockThreshold),
			UserLockTTL:              ptr(userLockTTL),
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Test Environment", resp.Environment.Name)
		assert.Equal(t, environments.EnvironmentTypeTest, resp.Environment.Type)
		assert.True(t, resp.Environment.UserLockSelfServeEnabled)
		assert.Equal(t, userLockThreshold, resp.Environment.UserLockThreshold)
		assert.Equal(t, userLockTTL, resp.Environment.UserLockTTL)
	})
	t.Run("IDP fields", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		idpAuthorizationURL := "https://example.com/idp"
		idpTemplateContent := "{\"field\": {{ user.user_id }} }"
		resp, err := client.Environments.Create(ctx, environments.CreateRequest{
			ProjectSlug:                         project.ProjectSlug,
			Name:                                "Test Environment",
			Type:                                environments.EnvironmentTypeTest,
			IDPAuthorizationURL:                 &idpAuthorizationURL,
			IDPDynamicClientRegistrationEnabled: ptr(true),
			IDPDynamicClientRegistrationAccessTokenTemplateContent: ptr(idpTemplateContent),
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Test Environment", resp.Environment.Name)
		assert.Equal(t, environments.EnvironmentTypeTest, resp.Environment.Type)
		assert.Equal(t, idpAuthorizationURL, resp.Environment.IDPAuthorizationURL)
		assert.True(t, resp.Environment.IDPDynamicClientRegistrationEnabled)
		assert.Equal(t, idpTemplateContent, resp.Environment.IDPDynamicClientRegistrationAccessTokenTemplateContent)
	})
}

func Test_EnvironmentsGet(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.Environments.Get(ctx, environments.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, env.Name, resp.Environment.Name)
	})
	t.Run("environment does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.Environments.Get(ctx, environments.GetRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: "nonexistent-environment",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)

		// Check that the error is well-formatted.
		var stytchErr stytcherror.Error
		assert.ErrorAs(t, err, &stytchErr)
		assert.NotEmpty(t, stytchErr.RequestID)
		assert.Equal(t, 404, stytchErr.StatusCode)
		assert.Contains(t, stytchErr.ErrorType, "not_found")
	})
	t.Run("missing environment", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.Environments.Get(ctx, environments.GetRequest{
			ProjectSlug: project.ProjectSlug,
			// Environment is intentionally omitted.
		})

		// Assert
		assert.ErrorContains(t, err, "environment")
		assert.Nil(t, resp)
	})
}

func hasEnvironment(environments []environments.Environment, target environments.Environment) bool {
	for _, e := range environments {
		if e.EnvironmentSlug == target.EnvironmentSlug {
			return e.Name == target.Name && e.Type == target.Type
		}
	}
	return false
}

func Test_EnvironmentsGetAll(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		createEnvResp, err := client.Environments.Create(ctx, environments.CreateRequest{
			ProjectSlug: project.ProjectSlug,
			Name:        "Another Test Environment",
			Type:        environments.EnvironmentTypeTest,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.Environments.GetAll(ctx, environments.GetAllRequest{
			ProjectSlug: project.ProjectSlug,
		})

		// Assert
		assert.NoError(t, err)
		// The disposable project is created with only a live environment, and we additionally created a
		// test environment, so we expect 2 environments to be returned.
		assert.Equal(t, 2, len(resp.Environments))
		// Check that the created environment is in the returned list.
		assert.True(t, hasEnvironment(resp.Environments, createEnvResp.Environment))
	})
}

func Test_EnvironmentsUpdate(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()
		newEnvironmentName := "Updated Environment Name"

		// Act
		zeroDowntimeSessionMigrationURL := "https://example.com/migration"
		resp, err := client.Environments.Update(ctx, environments.UpdateRequest{
			ProjectSlug:                     env.ProjectSlug,
			EnvironmentSlug:                 env.EnvironmentSlug,
			Name:                            &newEnvironmentName,
			CrossOrgPasswordsEnabled:        ptr(true),
			UserImpersonationEnabled:        ptr(true),
			ZeroDowntimeSessionMigrationURL: &zeroDowntimeSessionMigrationURL,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, newEnvironmentName, resp.Environment.Name)
		assert.True(t, resp.Environment.CrossOrgPasswordsEnabled)
		assert.True(t, resp.Environment.UserImpersonationEnabled)
		assert.Equal(t, zeroDowntimeSessionMigrationURL, resp.Environment.ZeroDowntimeSessionMigrationURL)
	})
	t.Run("user locking fields", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		userLockThreshold := int32(5)
		userLockTTL := int32(600)
		resp, err := client.Environments.Update(ctx, environments.UpdateRequest{
			ProjectSlug:              env.ProjectSlug,
			EnvironmentSlug:          env.EnvironmentSlug,
			UserLockSelfServeEnabled: ptr(true),
			UserLockThreshold:        ptr(userLockThreshold),
			UserLockTTL:              ptr(userLockTTL),
		})

		// Assert
		assert.NoError(t, err)
		assert.True(t, resp.Environment.UserLockSelfServeEnabled)
		assert.Equal(t, userLockThreshold, resp.Environment.UserLockThreshold)
		assert.Equal(t, userLockTTL, resp.Environment.UserLockTTL)
	})
	t.Run("IDP fields", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		idpAuthorizationURL := "https://example.com/idp"
		idpTemplateContent := "{\"field\": {{ user.user_id }} }"
		resp, err := client.Environments.Update(ctx, environments.UpdateRequest{
			ProjectSlug:                         env.ProjectSlug,
			EnvironmentSlug:                     env.EnvironmentSlug,
			IDPAuthorizationURL:                 &idpAuthorizationURL,
			IDPDynamicClientRegistrationEnabled: ptr(true),
			IDPDynamicClientRegistrationAccessTokenTemplateContent: ptr(idpTemplateContent),
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, idpAuthorizationURL, resp.Environment.IDPAuthorizationURL)
		assert.True(t, resp.Environment.IDPDynamicClientRegistrationEnabled)
		assert.Equal(t, idpTemplateContent, resp.Environment.IDPDynamicClientRegistrationAccessTokenTemplateContent)
	})
	t.Run("does not overwrite existing values", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()
		newEnvironmentName := "Updated Environment Name"

		// Act
		resp, err := client.Environments.Update(ctx, environments.UpdateRequest{
			ProjectSlug:              env.ProjectSlug,
			EnvironmentSlug:          env.EnvironmentSlug,
			CrossOrgPasswordsEnabled: ptr(true),
			UserImpersonationEnabled: ptr(true),
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, env.Name, resp.Environment.Name)
		assert.True(t, resp.Environment.CrossOrgPasswordsEnabled)
		assert.True(t, resp.Environment.UserImpersonationEnabled)

		// Act
		// Update again, but specify only the name.
		resp, err = client.Environments.Update(ctx, environments.UpdateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            &newEnvironmentName,
		})

		// Assert
		assert.NoError(t, err)
		// The other values should remain unchanged.
		assert.Equal(t, newEnvironmentName, resp.Environment.Name)
		assert.True(t, resp.Environment.CrossOrgPasswordsEnabled)
		assert.True(t, resp.Environment.UserImpersonationEnabled)
	})
}

func Test_EnvironmentsDelete(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		createResp, err := client.Environments.Create(ctx, environments.CreateRequest{
			ProjectSlug: project.ProjectSlug,
			Name:        "Test Environment",
			Type:        environments.EnvironmentTypeTest,
		})
		require.NoError(t, err)

		// Act
		_, err = client.Environments.Delete(ctx, environments.DeleteRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: createResp.Environment.EnvironmentSlug,
		})

		// Assert
		assert.NoError(t, err)
		// Verify the environment is actually deleted.
		_, err = client.Environments.Get(ctx, environments.GetRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: createResp.Environment.EnvironmentSlug,
		})
		assert.Error(t, err)
	})
}
