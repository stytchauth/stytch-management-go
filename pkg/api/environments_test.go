package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

func Test_EnvironmentsCreate(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.Environments.Create(ctx, environments.CreateRequest{
			Project:                  project.Project,
			Name:                     "Test Environment",
			Type:                     environments.EnvironmentTypeTest,
			CrossOrgPasswordsEnabled: ptr(true),
			UserImpersonationEnabled: ptr(true),
		})
		t.Cleanup(func() {
			_, err := client.Environments.Delete(ctx, environments.DeleteRequest{
				Project:     project.Project,
				Environment: resp.Environment.Environment,
			})
			require.NoError(t, err)
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Test Environment", resp.Environment.Name)
		assert.Equal(t, environments.EnvironmentTypeTest, resp.Environment.Type)
		assert.Equal(t, true, resp.Environment.CrossOrgPasswordsEnabled)
		assert.Equal(t, true, resp.Environment.UserImpersonationEnabled)
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
			Project:     env.Project,
			Environment: env.Environment,
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
			Project:     project.Project,
			Environment: "nonexistent-environment",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
	t.Run("missing environment", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.Environments.Get(ctx, environments.GetRequest{
			Project: project.Project,
			// Environment is intentionally omitted.
		})

		// Assert
		assert.ErrorContains(t, err, "environment")
		assert.Nil(t, resp)
	})
}

func hasEnvironment(environments []environments.Environment, target environments.Environment) bool {
	for _, e := range environments {
		if e.Environment == target.Environment {
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
			Project: project.Project,
			Name:    "Another Test Environment",
			Type:    environments.EnvironmentTypeTest,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.Environments.GetAll(ctx, environments.GetAllRequest{
			Project: project.Project,
		})

		// Assert
		assert.NoError(t, err)
		// The project is created with both a live and test environment, and we additionally created a
		// test environment, so we expect at least 3 environments to be returned.
		assert.Equal(t, 3, len(resp.Environments))
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
		resp, err := client.Environments.Update(ctx, environments.UpdateRequest{
			Project:                  env.Project,
			Environment:              env.Environment,
			Name:                     &newEnvironmentName,
			CrossOrgPasswordsEnabled: ptr(true),
			UserImpersonationEnabled: ptr(true),
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, newEnvironmentName, resp.Environment.Name)
		assert.Equal(t, true, resp.Environment.CrossOrgPasswordsEnabled)
		assert.Equal(t, true, resp.Environment.UserImpersonationEnabled)
	})
	t.Run("does not overwrite existing values", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()
		newEnvironmentName := "Updated Environment Name"

		// Act
		resp, err := client.Environments.Update(ctx, environments.UpdateRequest{
			Project:                  env.Project,
			Environment:              env.Environment,
			CrossOrgPasswordsEnabled: ptr(true),
			UserImpersonationEnabled: ptr(true),
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, env.Name, resp.Environment.Name)
		assert.Equal(t, true, resp.Environment.CrossOrgPasswordsEnabled)
		assert.Equal(t, true, resp.Environment.UserImpersonationEnabled)

		// Act
		// Update again, but specify only the name.
		resp, err = client.Environments.Update(ctx, environments.UpdateRequest{
			Project:     env.Project,
			Environment: env.Environment,
			Name:        &newEnvironmentName,
		})

		// Assert
		assert.NoError(t, err)
		// The other values should remain unchanged.
		assert.Equal(t, newEnvironmentName, resp.Environment.Name)
		assert.Equal(t, true, resp.Environment.CrossOrgPasswordsEnabled)
		assert.Equal(t, true, resp.Environment.UserImpersonationEnabled)
	})
}

func Test_EnvironmentsDelete(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		createResp, err := client.Environments.Create(ctx, environments.CreateRequest{
			Project: project.Project,
			Name:    "Test Environment",
			Type:    environments.EnvironmentTypeTest,
		})
		require.NoError(t, err)

		// Act
		_, err = client.Environments.Delete(ctx, environments.DeleteRequest{
			Project:     project.Project,
			Environment: createResp.Environment.Environment,
		})

		// Assert
		assert.NoError(t, err)
		// Verify the environment is actually deleted.
		_, err = client.Environments.Get(ctx, environments.GetRequest{
			Project:     project.Project,
			Environment: createResp.Environment.Environment,
		})
		assert.Error(t, err)
	})
}
