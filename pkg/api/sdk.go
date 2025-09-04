package api

import (
	"context"
	"encoding/json"
	"fmt"

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
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/sdk/consumer", body.Project, body.Environment),
		nil,
		nil,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
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
	err = c.client.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/sdk/consumer", body.Project, body.Environment),
		nil,
		jsonBody,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}

// GetB2BConfig retrieves the SDK configuration for a B2B project environment
func (c *SDKClient) GetB2BConfig(
	ctx context.Context,
	body sdk.GetB2BConfigRequest,
) (*sdk.GetB2BConfigResponse, error) {
	var res sdk.GetB2BConfigResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/sdk/b2b", body.Project, body.Environment),
		nil,
		nil,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
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
	err = c.client.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/sdk/b2b", body.Project, body.Environment),
		nil,
		jsonBody,
		&res,
	)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
