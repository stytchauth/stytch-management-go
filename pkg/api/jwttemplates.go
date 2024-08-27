package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/jwttemplates"
)

type JWTTemplatesClient struct {
	client *internal.Client
}

func newJWTTemplatesClient(c *internal.Client) *JWTTemplatesClient {
	return &JWTTemplatesClient{client: c}
}

func (c *JWTTemplatesClient) Get(
	ctx context.Context,
	body *jwttemplates.GetRequest,
) (*jwttemplates.GetResponse, error) {
	var resp jwttemplates.GetResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/projects/%s/jwt_templates/%s", body.ProjectID, body.TemplateType),
		nil,
		nil,
		&resp)

	return &resp, err
}

func (c *JWTTemplatesClient) Set(
	ctx context.Context,
	body *jwttemplates.SetRequest,
) (*jwttemplates.SetResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res jwttemplates.SetResponse
	err = c.client.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/projects/%s/jwt_templates/%s", body.ProjectID, body.TemplateType),
		nil,
		jsonBody,
		&res)
	return &res, err
}
