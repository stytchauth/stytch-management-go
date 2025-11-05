package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	migrationprojects "github.com/stytchauth/stytch-management-go/v3/pkg/models/migration/projects"
)

func TestMigrationClient_GetProjectsAndGetProject(t *testing.T) {
	client := NewTestClient(t)
	ctx := context.Background()

	projectsResp, err := client.V1ToV3MigrationClient.GetProjects(ctx, migrationprojects.GetProjectsRequest{})
	require.NoError(t, err)
	require.NotEmpty(t, projectsResp.Projects, "expect at least one project")

	legacyProject := projectsResp.Projects[0]
	assert.NotEmpty(t, legacyProject.LiveProjectID)
	assert.NotEmpty(t, legacyProject.LiveEnvironmentSlug)
	assert.NotEmpty(t, legacyProject.TestProjectID)
	assert.NotEmpty(t, legacyProject.TestEnvironmentSlug)
	assert.NotEmpty(t, legacyProject.ProjectSlug)

	projectResp, err := client.V1ToV3MigrationClient.GetProject(ctx, migrationprojects.GetProjectRequest{
		ProjectID: legacyProject.LiveProjectID,
	})
	require.NoError(t, err)

	assert.Equal(t, legacyProject.LiveProjectID, projectResp.Project.LiveProjectID)
	assert.Equal(t, legacyProject.LiveEnvironmentSlug, projectResp.Project.LiveEnvironmentSlug)
	assert.Equal(t, legacyProject.TestProjectID, projectResp.Project.TestProjectID)
	assert.Equal(t, legacyProject.TestEnvironmentSlug, projectResp.Project.TestEnvironmentSlug)
	assert.Equal(t, legacyProject.ProjectSlug, projectResp.Project.ProjectSlug)
}
