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
	if err != nil {
		return nil, err
	}

	var res projects.CreateResponse
	err = c.client.NewRequest(ctx, "POST", "/v1/projects", nil, jsonBody, &res)
	return &res, err
}

func (c *ProjectsClient) Get(
	ctx context.Context,
	body projects.GetRequest,
) (*projects.GetResponse, error) {
	var res projects.GetResponse
	err := c.client.NewRequest(ctx, "GET", "/v1/projects/"+body.ProjectID, nil, nil, &res)
	return &res, err
}

func (c *ProjectsClient) GetAll(
	ctx context.Context,
	body projects.GetAllRequest,
) (*projects.GetAllResponse, error) {
	var res projects.GetAllResponse
	err := c.client.NewRequest(ctx, "GET", "/v1/projects/", nil, nil, &res)
	return &res, err
}

func (c *ProjectsClient) Delete(
	ctx context.Context,
	body projects.DeleteRequest,
) (*projects.DeleteResponse, error) {
	var res projects.DeleteResponse
	err := c.client.NewRequest(ctx, "DELETE", "/v1/projects/"+body.ProjectID, nil, nil, &res)
	return &res, err
}
