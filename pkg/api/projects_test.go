package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/pkg/models/projects"
)

func Test_ProjectsCreate(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	ctx := context.Background()

	// Act
	resp, err := client.Projects.Create(ctx, projects.CreateRequest{
		ProjectName:                  "Test project",
		Vertical:                     projects.VerticalB2B,
		TestUserImpersonationEnabled: true,
		LiveUserImpersonationEnabled: false,
		TestCrossOrgPasswordsEnabled: true,
		LiveCrossOrgPasswordsEnabled: true,
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
	assert.True(t, resp.Project.TestUserImpersonationEnabled)
	assert.False(t, resp.Project.LiveUserImpersonationEnabled)
	assert.True(t, resp.Project.TestCrossOrgPasswordsEnabled)
	assert.True(t, resp.Project.LiveCrossOrgPasswordsEnabled)
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
	assert.False(t, project.TestUserImpersonationEnabled)
	assert.False(t, project.LiveUserImpersonationEnabled)
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
	assert.False(t, project.LiveCrossOrgPasswordsEnabled)
	ctx := context.Background()
	newProjectName := "The new project v2"

	// Act
	resp, err := client.Projects.Update(ctx, projects.UpdateRequest{
		ProjectID:                    project.LiveProjectID,
		Name:                         newProjectName,
		LiveUserImpersonationEnabled: ptr(true),
		LiveUseCrossOrgPasswords:     ptr(true),
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, newProjectName, resp.Project.Name)
	assert.True(t, resp.Project.LiveUserImpersonationEnabled)
	assert.True(t, resp.Project.LiveCrossOrgPasswordsEnabled)
}

func Test_ProjectsUpdateDoesNotOverwriteValues(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	assert.False(t, project.LiveCrossOrgPasswordsEnabled)
	ctx := context.Background()
	newProjectName := "The new project v2"

	// Act
	resp, err := client.Projects.Update(ctx, projects.UpdateRequest{
		ProjectID:                    project.LiveProjectID,
		Name:                         newProjectName,
		LiveUserImpersonationEnabled: ptr(true),
		TestUserImpersonationEnabled: ptr(true),
		LiveUseCrossOrgPasswords:     ptr(true),
		TestUseCrossOrgPasswords:     ptr(true),
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, newProjectName, resp.Project.Name)
	assert.True(t, resp.Project.LiveUserImpersonationEnabled)
	assert.True(t, resp.Project.TestUserImpersonationEnabled)
	assert.True(t, resp.Project.LiveCrossOrgPasswordsEnabled)
	assert.True(t, resp.Project.TestCrossOrgPasswordsEnabled)

	// Act again to check if the values are not overwritten
	newProjectName = "The new project v2.1"
	resp, err = client.Projects.Update(ctx, projects.UpdateRequest{
		ProjectID: project.LiveProjectID,
		Name:      newProjectName,
	})
	// Assert
	assert.NoError(t, err)
	assert.Equal(t, newProjectName, resp.Project.Name)
	assert.True(t, resp.Project.LiveUserImpersonationEnabled)
	assert.True(t, resp.Project.LiveCrossOrgPasswordsEnabled)
}
