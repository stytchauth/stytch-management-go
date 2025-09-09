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
		resp, err := client.Projects.Create(ctx, projects.CreateRequest{
			Name:     "Test Project",
			Vertical: projects.VerticalB2B,
		})
		t.Cleanup(func() {
			_, err := client.Projects.Delete(ctx, projects.DeleteRequest{
				Project: resp.Project.Project,
			})
			require.NoError(t, err)
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, "Test Project", resp.Project.Name)
		assert.Equal(t, projects.VerticalB2B, resp.Project.Vertical)
	})
}

func Test_ProjectsGet(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.Projects.Get(ctx, projects.GetRequest{
			Project: project.Project,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, project.Name, resp.Project.Name)
		assert.Equal(t, projects.VerticalB2B, resp.Project.Vertical)
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		resp, err := client.Projects.Get(ctx, projects.GetRequest{
			Project: "nonexistent-project",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
	t.Run("missing project", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		resp, err := client.Projects.Get(ctx, projects.GetRequest{
			// Project is intentionally omitted.
		})

		// Assert
		assert.ErrorContains(t, err, "project")
		assert.Nil(t, resp)
	})
}

func hasProject(projects []projects.Project, target projects.Project) bool {
	for _, p := range projects {
		if p.Project == target.Project {
			return p.Name == target.Name && p.Vertical == target.Vertical
		}
	}
	return false
}

func Test_ProjectsGetAll(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project1 := client.DisposableProject(projects.VerticalB2B)
		project2 := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.Projects.GetAll(ctx, projects.GetAllRequest{})

		// Assert
		assert.NoError(t, err)
		// The test workspace may have other projects in it, so we just check that there are at least
		// 2 projects.
		assert.GreaterOrEqual(t, len(resp.Projects), 2)
		// Similarly, we check that the two projects we created are in the returned list.
		assert.True(t, hasProject(resp.Projects, project1))
		assert.True(t, hasProject(resp.Projects, project2))
	})
}

func Test_ProjectsUpdate(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		newProjectName := "Updated Project Name"

		// Act
		resp, err := client.Projects.Update(ctx, projects.UpdateRequest{
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
		resp, err := client.Projects.Update(ctx, projects.UpdateRequest{
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
		createResp, err := client.Projects.Create(ctx, projects.CreateRequest{
			Name:     "Test Project",
			Vertical: projects.VerticalB2B,
		})
		require.NoError(t, err)

		// Act
		_, err = client.Projects.Delete(ctx, projects.DeleteRequest{
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
		_, err := client.Projects.Delete(ctx, projects.DeleteRequest{
			Project: "nonexistent-project",
		})

		// Assert
		assert.Error(t, err)
	})
}
