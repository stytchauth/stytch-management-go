package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
)

func Test_ProjectsCreate(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	ctx := context.Background()

	// Act
	resp, err := client.Projects.Create(ctx, projects.CreateRequest{
		ProjectName: "Test project",
		Vertical:    projects.VerticalB2B,
	})
	t.Cleanup(func() {
		_, err := client.Projects.Delete(ctx, projects.DeleteRequest{
			ProjectID: resp.Project.LiveProjectID,
		})
		require.NoError(t, err)
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Test project", resp.Project.Name)
	assert.Equal(t, projects.VerticalB2B, resp.Project.Vertical)
}

func Test_ProjectsGet(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()

	// Act
	resp, err := client.Projects.Get(ctx, projects.GetRequest{
		ProjectID: project.LiveProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, project.Name, resp.Project.Name)
	assert.Equal(t, projects.VerticalB2B, resp.Project.Vertical)
}

func Test_ProjectsDelete(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	ctx := context.Background()
	createResp, err := client.Projects.Create(ctx, projects.CreateRequest{
		ProjectName: "Delete project test",
		Vertical:    projects.VerticalB2B,
	})
	require.NoError(t, err)

	// Act
	_, err = client.Projects.Delete(ctx, projects.DeleteRequest{
		ProjectID: createResp.Project.LiveProjectID,
	})

	// Assert
	assert.NoError(t, err)
}

func Test_ProjectsUpdate(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	newProjectName := "The new project v2"

	// Act
	resp, err := client.Projects.Update(ctx, projects.UpdateRequest{
		ProjectID: project.LiveProjectID,
		Name:      newProjectName,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, newProjectName, resp.Project.Name)
}
