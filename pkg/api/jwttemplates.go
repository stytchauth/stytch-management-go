package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/jwttemplates"
)

type JWTTemplatesClient struct {
	client *internal.Client
}

func newJWTTemplatesClient(c *internal.Client) *JWTTemplatesClient {
	return &JWTTemplatesClient{client: c}
}

// Get retrieves a JWT template for a project
func (c *JWTTemplatesClient) Get(
	ctx context.Context,
	body *jwttemplates.GetRequest,
) (*jwttemplates.GetResponse, error) {
	var resp jwttemplates.GetResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/jwt_templates/%s", body.ProjectSlug, body.EnvironmentSlug, body.JWTTemplateType),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Set updates a specific JWT template for a project
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
		http.MethodPut,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/jwt_templates/%s", body.ProjectSlug, body.EnvironmentSlug, body.JWTTemplateType),
		nil,
		jsonBody,
		&res)
	if err != nil {
		return nil, err
	}
	return &res, nil
}
