package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environmentmetrics"
)

type EnvironmentMetricsClient struct {
	client *internal.Client
}

func newEnvironmentMetricsClient(c *internal.Client) *EnvironmentMetricsClient {
	return &EnvironmentMetricsClient{
		client: c,
	}
}

// Get retrieves metrics for an environment.
func (c *EnvironmentMetricsClient) Get(
	ctx context.Context,
	body environmentmetrics.GetRequest,
) (*environmentmetrics.GetResponse, error) {
	var resp environmentmetrics.GetResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/metrics", body.Project, body.Environment),
		nil,
		nil,
		&resp)
	return &resp, err
}
