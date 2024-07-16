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

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Test project", resp.Project.ProjectSettings.ProjectName)
	assert.Equal(t, projects.VerticalB2B, resp.Project.Vertical)
}

func Test_ProjectsGet(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	ctx := context.Background()
	createResp, err := client.Projects.Create(ctx, projects.CreateRequest{
		ProjectName: "Get project test",
		Vertical:    projects.VerticalB2B,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.Projects.Get(ctx, projects.GetRequest{
		ProjectID: createResp.Project.ProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Get project test", resp.Project.ProjectSettings.ProjectName)
	assert.Equal(t, projects.VerticalB2B, resp.Project.Vertical)
}
