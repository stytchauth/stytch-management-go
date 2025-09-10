package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/emailtemplates"
)

type EmailTemplatesClient struct {
	client *internal.Client
}

func newEmailTemplatesClient(c *internal.Client) *EmailTemplatesClient {
	return &EmailTemplatesClient{client: c}
}

// Create creates an email template for a project.
func (c *EmailTemplatesClient) Create(
	ctx context.Context,
	body emailtemplates.CreateRequest,
) (*emailtemplates.CreateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp emailtemplates.CreateResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/email_templates", body.Project),
		nil,
		jsonBody,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Get retrieves an email template for a project.
func (c *EmailTemplatesClient) Get(
	ctx context.Context,
	body emailtemplates.GetRequest,
) (*emailtemplates.GetResponse, error) {
	if body.TemplateID == "" {
		return nil, fmt.Errorf("missing template ID")
	}
	var resp emailtemplates.GetResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/email_templates/%s", body.Project, body.TemplateID),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAll retrieves all email templates for a project.
func (c *EmailTemplatesClient) GetAll(
	ctx context.Context,
	body emailtemplates.GetAllRequest,
) (*emailtemplates.GetAllResponse, error) {
	var resp emailtemplates.GetAllResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/email_templates", body.Project),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Update updates an email template for a project.
func (c *EmailTemplatesClient) Update(
	ctx context.Context,
	body emailtemplates.UpdateRequest,
) (*emailtemplates.UpdateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp emailtemplates.UpdateResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("/pwa/v3/projects/%s/email_templates/%s", body.Project, body.TemplateID),
		nil,
		jsonBody,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Delete deletes an email template for a project.
func (c *EmailTemplatesClient) Delete(
	ctx context.Context,
	body emailtemplates.DeleteRequest,
) (*emailtemplates.DeleteResponse, error) {
	var resp emailtemplates.DeleteResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("/pwa/v3/projects/%s/email_templates/%s", body.Project, body.TemplateID),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetDefault sets the default email template for a specific template type in a project.
func (c *EmailTemplatesClient) SetDefault(
	ctx context.Context,
	body emailtemplates.SetDefaultRequest,
) (*emailtemplates.SetDefaultResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp emailtemplates.SetDefaultResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/default_email_templates/%s", body.Project, body.EmailTemplateType),
		nil,
		jsonBody,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDefault retrieves the default email template for a specific template type in a project.
func (c *EmailTemplatesClient) GetDefault(
	ctx context.Context,
	body emailtemplates.GetDefaultRequest,
) (*emailtemplates.GetDefaultResponse, error) {
	var resp emailtemplates.GetDefaultResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/default_email_templates/%s", body.Project, body.EmailTemplateType),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsetDefault removes the default email template for a specific template type in a project.
func (c *EmailTemplatesClient) UnsetDefault(
	ctx context.Context,
	body emailtemplates.UnsetDefaultRequest,
) (*emailtemplates.UnsetDefaultResponse, error) {
	var resp emailtemplates.UnsetDefaultResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("/pwa/v3/projects/%s/default_email_templates/%s", body.Project, body.EmailTemplateType),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
