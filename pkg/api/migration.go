package api

import (
	"context"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	migrationprojects "github.com/stytchauth/stytch-management-go/v3/pkg/models/migration/projects"
)

// V1ToV3MigrationClient exposes legacy endpoints that are only used to support provider migrations.
type V1ToV3MigrationClient struct {
	client *internal.Client
}

func newMigrationClient(c *internal.Client) *V1ToV3MigrationClient {
	return &V1ToV3MigrationClient{
		client: c,
	}
}

// GetProjects retrieves all projects' identifiers from the PWA v1 endpoint.
// In order to get a map between PWA v1 and PWA v3 identifiers.
func (c *V1ToV3MigrationClient) GetProjects(
	ctx context.Context,
	_ migrationprojects.GetProjectsRequest,
) (*migrationprojects.GetProjectsResponse, error) {
	var res migrationprojects.GetProjectsResponse
	err := c.client.NewRequest(ctx, http.MethodGet, "/web/v1/projects", nil, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetProject retrieves the project details with both PWA V1 and PWA V3 identifiers for the provided PWA V1 project ID.
func (c *V1ToV3MigrationClient) GetProject(
	ctx context.Context,
	body migrationprojects.GetProjectRequest,
) (*migrationprojects.GetProjectResponse, error) {
	var res migrationprojects.GetProjectResponse
	err := c.client.NewRequest(ctx, http.MethodGet, "/web/v1/projects/"+body.ProjectID, nil, nil, &res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
