package api_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stytchauth/stytch-management-go/pkg/models/projects"
	"testing"
)

func Test_BugBash(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)

	// Act
	ctx := context.Background()
	resp, err := client.Projects.Update(ctx, projects.UpdateRequest{
		ProjectID:            project.LiveProjectID,
		UseCrossOrgPasswords: false,
		Name:                 "Edited",
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.False(t, resp.Project.UseCrossOrgPasswords)

	ctx = context.Background()
	resp, err = client.Projects.Update(ctx, projects.UpdateRequest{
		ProjectID:            project.LiveProjectID,
		UseCrossOrgPasswords: true,
		Name:                 "Edited 2",
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, 200, resp.StatusCode)
	assert.True(t, resp.Project.UseCrossOrgPasswords)
}
