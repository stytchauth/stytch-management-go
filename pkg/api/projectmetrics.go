package api

import (
	"context"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projectmetrics"
)

type ProjectMetricsClient struct {
	client *internal.Client
}

func newProjectMetricsClient(c *internal.Client) *ProjectMetricsClient {
	return &ProjectMetricsClient{
		client: c,
	}
}

// Get retrieves metrics for a project
func (c *ProjectMetricsClient) Get(
	ctx context.Context,
	body projectmetrics.GetRequest,
) (*projectmetrics.GetResponse, error) {
	var res projectmetrics.GetResponse
	err := c.client.NewRequest(ctx, "GET", "/v1/projects/"+body.ProjectID+"/project_metrics", nil, nil, &res)
	return &res, err
}
