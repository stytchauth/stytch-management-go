package api

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/sdk"
)

type SDKClient struct {
	client *internal.Client
}

func newSDKClient(c *internal.Client) *SDKClient {
	return &SDKClient{client: c}
}

func (c *SDKClient) GetConfig(
	ctx context.Context,
	body sdk.GetConfigRequest,
) (*sdk.GetConfigResponse, error) {
	var res sdk.GetConfigResponse
	err := c.client.NewRequest(ctx, "GET", "/v1/projects/"+body.ProjectID+"/sdk", nil, nil, &res)
	return &res, err
}

func (c *SDKClient) SetConfig(
	ctx context.Context,
	body sdk.SetConfigRequest,
) (*sdk.SetConfigResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res sdk.SetConfigResponse
	err = c.client.NewRequest(ctx, "PUT", "/v1/projects/"+body.ProjectID+"/sdk", nil, jsonBody, &res)
	return &res, err
}
