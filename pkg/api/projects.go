package api

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
)

type ProjectsClient struct {
	client *internal.Client
}

func newProjectsClient(c *internal.Client) *ProjectsClient {
	return &ProjectsClient{
		client: c,
	}
}

func (c *ProjectsClient) Create(
	ctx context.Context,
	body projects.CreateRequest,
) (*projects.CreateResponse, error) {
	jsonBody, err := json.Marshal(body)

	var res projects.CreateResponse
	err = c.client.NewRequest(ctx, "POST", "/v1/projects", nil, jsonBody, &res)
	return &res, err
}
