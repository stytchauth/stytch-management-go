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
	body emailtemplates.GetAllEmailTemplatesRequest,
) (*emailtemplates.GetAllEmailTemplatesResponse, error) {
	var resp emailtemplates.GetAllEmailTemplatesResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/projects/%s/email_templates", body.LiveProjectID),
		nil,
		nil,
		&resp)

	return &resp, err
}

func (c *EmailTemplatesClient) Get(
	ctx context.Context,
	body emailtemplates.GetEmailTemplateRequest,
) (*emailtemplates.GetEmailTemplateResponse, error) {
	var resp emailtemplates.GetEmailTemplateResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		fmt.Sprintf("/v1/projects/%s/email_templates/%s", body.ProjectID, body.EmailTemplateID),
		nil,
		nil,
		&resp)

	return &resp, err
}

func (c *EmailTemplatesClient) Create(
	ctx context.Context,
	body emailtemplates.CreateEmailTemplateRequest,
) (*emailtemplates.CreateEmailTemplateResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var res emailtemplates.CreateEmailTemplateResponse
	err = c.client.NewRequest(
		ctx,
		"POST",
		fmt.Sprintf("/v1/projects/%s/email_templates", body.LiveProjectID),
		nil,
		jsonBody,
		&res)
	return &res, err
}

func (c *EmailTemplatesClient) Delete(
	ctx context.Context,
	body emailtemplates.DeleteEmailTemplateRequest,
) (*emailtemplates.DeleteEmailTemplateResponse, error) {
	var res emailtemplates.DeleteEmailTemplateResponse
	err := c.client.NewRequest(
		ctx,
		"DELETE",
		fmt.Sprintf("/v1/projects/%s/email_templates/%s", body.LiveProjectID, body.EmailTemplateID),
		nil,
		nil,
		&res)
	return &res, err
}
