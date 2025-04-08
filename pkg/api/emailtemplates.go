package api

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stytchauth/stytch-management-go/v2/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/emailtemplates"
)

type EmailTemplatesClient struct {
	client *internal.Client
}

func newEmailTemplatesClient(c *internal.Client) *EmailTemplatesClient {
	return &EmailTemplatesClient{client: c}
}

// GetAll retrieves all email templates for a project
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

// Get retrieves an email template for a project
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

// Create creates an email template for both a live and a test project
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

// Delete deletes an email template for a project
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

// Update updates an email template for a project
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
