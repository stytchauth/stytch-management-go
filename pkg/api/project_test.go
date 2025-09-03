package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

func Test_ProjectsCreate(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		resp, err := client.Project.Create(ctx, project.CreateRequest{
			Name:     "Test Project",
			Vertical: project.VerticalB2B,
		})
		t.Cleanup(func() {
			_, err := client.Project.Delete(ctx, project.DeleteRequest{
				Project: resp.Project.Project,
			})
			require.NoError(t, err)
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Test Project", resp.Project.Name)
		assert.Equal(t, project.VerticalB2B, resp.Project.Vertical)
	})
}

func Test_ProjectsGet(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.Project.Get(ctx, project.GetRequest{
			Project: project.Project,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, project.Name, resp.Project.Name)
		assert.Equal(t, project.VerticalB2B, resp.Project.Vertical)
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		resp, err := client.Project.Get(ctx, project.GetRequest{
			Project: "nonexistent-project",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func Test_ProjectsUpdate(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalB2B)
		ctx := context.Background()
		newProjectName := "Updated Project Name"

		// Act
		resp, err := client.Project.Update(ctx, project.UpdateRequest{
			Project: project.Project,
			Name:    &newProjectName,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, newProjectName, resp.Project.Name)
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()
		newProjectName := "Updated Project Name"

		// Act
		resp, err := client.Project.Update(ctx, project.UpdateRequest{
			Project: "nonexistent-project",
			Name:    &newProjectName,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func Test_ProjectsDelete(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()
		createResp, err := client.Project.Create(ctx, project.CreateRequest{
			Name:     "Test Project",
			Vertical: project.VerticalB2B,
		})
		require.NoError(t, err)

		// Act
		_, err = client.Project.Delete(ctx, project.DeleteRequest{
			Project: createResp.Project.Project,
		})

		// Assert
		assert.NoError(t, err)
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		_, err := client.Project.Delete(ctx, project.DeleteRequest{
			Project: "nonexistent-project",
		})

		// Assert
		assert.Error(t, err)
	})
}
