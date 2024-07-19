package api_test

import (
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/emailtemplates"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
	"testing"
)

func createNewEmailTemplate(
	t *testing.T,
	ctx context.Context,
	client *testClient,
	project projects.Project,
) (emailtemplates.CreateEmailTemplateResponse, error) {
	requestBody := emailtemplates.CreateEmailTemplateRequest{
		LiveProjectID: project.ProjectID,
		Name:          "Test email template",
		VanityID:      "test_email_template",
		Method:        emailtemplates.MethodBuiltInCustomizations,
		Type:          emailtemplates.TemplateTypeAll,
	}
	resp, err := client.EmailTemplates.Create(ctx, requestBody)
	require.NoError(t, err)
	return *resp, nil
}

func TestEmailTemplatesClient_Create(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	requestBody := emailtemplates.CreateEmailTemplateRequest{
		LiveProjectID: project.ProjectID,
		Name:          "Test email template",
		VanityID:      "test_email_template",
		Method:        emailtemplates.MethodBuiltInCustomizations,
		Type:          emailtemplates.TemplateTypeAll,
	}

	// Act
	resp, err := client.EmailTemplates.Create(ctx, requestBody)

	// Assert
	assert.NoError(t, err)
	// both Live and Test templates should have the same values
	assert.Equal(t, requestBody.Name, resp.EmailTemplate.LiveEmailTemplate.Name)
	assert.Equal(t, requestBody.VanityID, resp.EmailTemplate.TestEmailTemplate.VanityID)
	assert.Equal(t, requestBody.Method, resp.EmailTemplate.LiveEmailTemplate.Method)
	assert.Equal(t, requestBody.Type, resp.EmailTemplate.TestEmailTemplate.Type)
}

func TestEmailTemplatesClient_Get(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp, err := createNewEmailTemplate(t, ctx, client, project)

	// Act
	resp, err := client.EmailTemplates.Get(ctx, emailtemplates.GetEmailTemplateRequest{
		ProjectID:       project.ProjectID,
		EmailTemplateID: createResp.EmailTemplate.LiveEmailTemplate.ID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, createResp.EmailTemplate.LiveEmailTemplate.ID, resp.EmailTemplateID)
	assert.Equal(t, createResp.EmailTemplate.LiveEmailTemplate.Name, resp.Name)
}

func TestEmailTemplatesClient_GetAll(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp, err := createNewEmailTemplate(t, ctx, client, project)

	// Act
	resp, err := client.EmailTemplates.GetAll(ctx, emailtemplates.GetAllEmailTemplatesRequest{
		LiveProjectID: project.ProjectID,
	})
	var liveTemplatesVanityIDs []string
	for _, template := range resp.EmailTemplates {
		liveTemplatesVanityIDs = append(liveTemplatesVanityIDs, template.LiveEmailTemplate.VanityID)
	}

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, liveTemplatesVanityIDs, createResp.EmailTemplate.LiveEmailTemplate.VanityID)
}

func TestEmailTemplatesClient_Update(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp, err := createNewEmailTemplate(t, ctx, client, project)

	// Act
	_, err = client.EmailTemplates.Update(ctx, emailtemplates.UpdateEmailTemplateRequest{
		ProjectID:       project.ProjectID,
		EmailTemplateID: createResp.EmailTemplate.LiveEmailTemplate.ID,
		Name:            "Updated email template",
		FromName:        "Myself",
	})

	// Assert
	assert.NoError(t, err)
	getResp, err := client.EmailTemplates.Get(ctx, emailtemplates.GetEmailTemplateRequest{
		ProjectID:       project.ProjectID,
		EmailTemplateID: createResp.EmailTemplate.LiveEmailTemplate.ID,
	})
	assert.NoError(t, err)
	assert.Equal(t, "Updated email template", getResp.Name)
	assert.Equal(t, "Myself", getResp.FromName)
}

func TestEmailTemplatesClient_Delete(t *testing.T) {
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp, err := createNewEmailTemplate(t, ctx, client, project)

	// Act
	resp, err := client.EmailTemplates.Delete(ctx, emailtemplates.DeleteEmailTemplateRequest{
		LiveProjectID:   project.ProjectID,
		EmailTemplateID: createResp.EmailTemplate.LiveEmailTemplate.ID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, createResp.EmailTemplate.LiveEmailTemplate.ID, resp.EmailTemplateID)
	assert.Equal(t, createResp.EmailTemplate.TestEmailTemplate.ID, resp.TestEmailTemplateID)
}
