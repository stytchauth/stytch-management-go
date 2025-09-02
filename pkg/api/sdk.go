package api

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/sdk"
)

type SDKClient struct {
	client *internal.Client
}

func newSDKClient(c *internal.Client) *SDKClient {
	return &SDKClient{client: c}
}

// GetConsumerConfig retrieves the SDK configuration for a B2C project environment
func (c *SDKClient) GetConsumerConfig(
	ctx context.Context,
	body sdk.GetConsumerConfigRequest,
) (*sdk.GetConsumerConfigResponse, error) {
	var res sdk.GetConsumerConfigResponse
	err := c.client.NewRequest(ctx, "GET", "/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/sdk/consumer", nil, nil, &res)
	return &res, err
}

// SetConsumerConfig updates the SDK configuration for a B2C project environment
func (c *SDKClient) SetConsumerConfig(
	ctx context.Context,
	body sdk.SetConsumerConfigRequest,
) (*sdk.SetConsumerConfigResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res sdk.SetConsumerConfigResponse
	err = c.client.NewRequest(ctx, "PUT", "/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/sdk/consumer", nil, jsonBody, &res)
	return &res, err
}

// GetB2BConfig retrieves the SDK configuration for a B2B project environment
func (c *SDKClient) GetB2BConfig(
	ctx context.Context,
	body sdk.GetB2BConfigRequest,
) (*sdk.GetB2BConfigResponse, error) {
	var res sdk.GetB2BConfigResponse
	err := c.client.NewRequest(ctx, "GET", "/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/sdk/b2b", nil, nil, &res)
	return &res, err
}

// SetB2BConfig updates the SDK configuration for a B2B project environment
func (c *SDKClient) SetB2BConfig(
	ctx context.Context,
	body sdk.SetB2BConfigRequest,
) (*sdk.SetB2BConfigResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res sdk.SetB2BConfigResponse
	err = c.client.NewRequest(ctx, "PUT", "/pwa/v3/projects/"+body.Project+"/environments/"+body.Environment+"/sdk/b2b", nil, jsonBody, &res)
	return &res, err
}
