package api

import (
	"context"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environmentmetrics"
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
	body environmentmetrics.GetRequest,
) (*environmentmetrics.GetResponse, error) {
	var res environmentmetrics.GetResponse
	err := c.client.NewRequest(ctx, "GET", "/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/metrics", nil, nil, &res)
	return &res, err
}
