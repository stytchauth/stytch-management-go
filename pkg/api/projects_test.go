package api_test

import (
	"context"
	"fmt"
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
			ProjectID: resp.Project.ProjectID,
		})
		require.NoError(t, err)
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "Test project", resp.Project.ProjectSettings.ProjectName)
	assert.Equal(t, projects.VerticalB2B, resp.Project.Vertical)
}

func Test_ProjectsGet(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()

	// Act
	resp, err := client.Projects.Get(ctx, projects.GetRequest{
		ProjectID: project.ProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, project.ProjectSettings.ProjectName, resp.Project.ProjectSettings.ProjectName)
	assert.Equal(t, projects.VerticalB2B, resp.Project.Vertical)
}

func Test_ProjectsGetPasswordStrengthPolicy(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	ctx := context.Background()
	project := client.DisposableProject(projects.VerticalB2B)
	_, err := client.Projects.SetPasswordStrengthPolicy(ctx, projects.SetPasswordStrengthPolicyRequest{
		ProjectID: project.ProjectID,
		PasswordConfig: projects.PasswordStrengthConfig{
			CheckBreachOnCreate:         true,
			CheckBreachOnAuthentication: true,
			ValidateOnAuthentication:    true,
			ValidationPolicy:            projects.PasswordValidationPolicyLuds,
			LudsMinCount:                ptr(12),
			LudsComplexity:              ptr(3),
		},
	})
	require.NoError(t, err)

	// Act
	resp, err := client.Projects.GetPasswordStrengthPolicy(ctx, projects.GetPasswordStrengthPolicyRequest{
		ProjectID: project.ProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.True(t, resp.PasswordConfig.CheckBreachOnCreate)
	fmt.Printf("Resp is %+v\n\n\n", resp)
	assert.Equal(t, 12, *resp.PasswordConfig.LudsMinCount)
}

func Test_ProjectsSetPasswordStrengthPolicy(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()

	// Act
	resp, err := client.Projects.SetPasswordStrengthPolicy(ctx, projects.SetPasswordStrengthPolicyRequest{
		ProjectID: project.ProjectID,
		PasswordConfig: projects.PasswordStrengthConfig{
			CheckBreachOnCreate:         true,
			CheckBreachOnAuthentication: true,
			ValidateOnAuthentication:    true,
			ValidationPolicy:            projects.PasswordValidationPolicyLuds,
			LudsMinCount:                ptr(12),
			LudsComplexity:              ptr(3),
		},
	})

	// Assert
	assert.NoError(t, err)
	assert.True(t, resp.PasswordConfig.CheckBreachOnCreate)
	assert.Equal(t, 12, *resp.PasswordConfig.LudsMinCount)
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
	resp, err := client.Projects.Delete(ctx, projects.DeleteRequest{
		ProjectID: createResp.Project.ProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, createResp.Project.ProjectID, resp.ProjectID)
}
