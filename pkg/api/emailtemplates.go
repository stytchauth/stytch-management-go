package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/emailtemplates"
)

type EmailTemplatesClient struct {
	client *internal.Client
}

func newEmailTemplatesClient(c *internal.Client) *EmailTemplatesClient {
	return &EmailTemplatesClient{client: c}
}

func (c *EmailTemplatesClient) GetAll(
	ctx context.Context,
	body emailtemplates.GetAllRequest,
) (*emailtemplates.GetAllResponse, error) {
	var resp emailtemplates.GetAllResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/projects/%s/email_templates", body.ProjectID),
		nil,
		nil,
		&resp)

	return &resp, err
}

func (c *EmailTemplatesClient) Get(
	ctx context.Context,
	body emailtemplates.GetRequest,
) (*emailtemplates.GetResponse, error) {
	var resp emailtemplates.GetResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/projects/%s/email_templates/%s", body.ProjectID, body.TemplateID),
		nil,
		nil,
		&resp)

	return &resp, err
}

func (c *EmailTemplatesClient) Create(
	ctx context.Context,
	body emailtemplates.CreateRequest,
) (*emailtemplates.CreateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res emailtemplates.CreateResponse
	err = c.client.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/projects/%s/email_templates", body.ProjectID),
		nil,
		jsonBody,
		&res)
	return &res, err
}

func (c *EmailTemplatesClient) Delete(
	ctx context.Context,
	body emailtemplates.DeleteRequest,
) (*emailtemplates.DeleteResponse, error) {
	var res emailtemplates.DeleteResponse
	err := c.client.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/projects/%s/email_templates/%s", body.ProjectID, body.TemplateID),
		nil,
		nil,
		&res)
	return &res, err
}

func (c *EmailTemplatesClient) Update(
	ctx context.Context,
	body emailtemplates.UpdateRequest,
) (*emailtemplates.UpdateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res emailtemplates.UpdateResponse
	err = c.client.NewRequest(
		ctx,
		"PUT",
		fmt.Sprintf("/v1/projects/%s/email_templates/%s", body.ProjectID, body.EmailTemplate.TemplateID),
		nil,
		jsonBody,
		&res)
	return &res, err
}
