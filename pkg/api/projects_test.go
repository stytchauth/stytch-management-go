package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
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
