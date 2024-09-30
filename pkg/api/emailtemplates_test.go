package api_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/pkg/models/emailtemplates"
	"github.com/stytchauth/stytch-management-go/pkg/models/projects"
)

func randomID(t *testing.T) string {
	t.Helper()
	return fmt.Sprintf("test_template_%d", rand.Int())
}

func createNewEmailTemplate(
	t *testing.T,
	ctx context.Context,
	client *testClient,
	project projects.Project,
) emailtemplates.CreateResponse {
	t.Helper()
	requestBody := emailtemplates.CreateRequest{
		ProjectID: project.LiveProjectID,
		EmailTemplate: emailtemplates.EmailTemplate{
			TemplateID: randomID(t),
			Name:       ptr("Test email template"),
			PrebuiltCustomization: &emailtemplates.PrebuiltCustomization{
				ButtonColor:     ptr("blue"),
				ButtonTextColor: ptr("white"),
			},
		},
	}
	resp, err := client.EmailTemplates.Create(ctx, requestBody)
	require.NoError(t, err)
	return *resp
}

func TestEmailTemplatesClient_Create(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	requestBody := emailtemplates.CreateRequest{
		ProjectID: project.LiveProjectID,
		EmailTemplate: emailtemplates.EmailTemplate{
			TemplateID: randomID(t),
			Name:       ptr("Test email template"),
			PrebuiltCustomization: &emailtemplates.PrebuiltCustomization{
				ButtonColor:     ptr("blue"),
				ButtonTextColor: ptr("white"),
			},
		},
	}

	// Act
	resp, err := client.EmailTemplates.Create(ctx, requestBody)

	// Assert
	assert.NoError(t, err)
	// both Live and Test templates should have the same values
	assert.Equal(t, requestBody.EmailTemplate.Name, resp.EmailTemplate.Name)
	assert.Equal(t, requestBody.EmailTemplate.TemplateID, resp.EmailTemplate.TemplateID)
	assert.NotNil(t, resp.EmailTemplate.PrebuiltCustomization)
}

func TestEmailTemplatesClient_Get(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp := createNewEmailTemplate(t, ctx, client, project)

	// Act
	resp, err := client.EmailTemplates.Get(ctx, emailtemplates.GetRequest{
		ProjectID:  project.LiveProjectID,
		TemplateID: createResp.EmailTemplate.TemplateID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, createResp.EmailTemplate.TemplateID, resp.EmailTemplate.TemplateID)
	assert.Equal(t, createResp.EmailTemplate.Name, resp.EmailTemplate.Name)
}

func TestEmailTemplatesClient_GetAll(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp := createNewEmailTemplate(t, ctx, client, project)
	createResp2 := createNewEmailTemplate(t, ctx, client, project)

	// Act
	resp, err := client.EmailTemplates.GetAll(ctx, emailtemplates.GetAllRequest{
		ProjectID: project.LiveProjectID,
	})
	var liveTemplatesIDs []string
	for _, template := range resp.EmailTemplates {
		liveTemplatesIDs = append(liveTemplatesIDs, template.TemplateID)
	}

	// Assert
	assert.NoError(t, err)
	assert.Contains(t, liveTemplatesIDs, createResp.EmailTemplate.TemplateID)
	assert.Contains(t, liveTemplatesIDs, createResp2.EmailTemplate.TemplateID)
}

func TestEmailTemplatesClient_Update(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp := createNewEmailTemplate(t, ctx, client, project)

	// Act
	_, err := client.EmailTemplates.Update(ctx, emailtemplates.UpdateRequest{
		ProjectID: project.LiveProjectID,
		EmailTemplate: emailtemplates.EmailTemplate{
			TemplateID: createResp.EmailTemplate.TemplateID,
			Name:       ptr("Updated email template"),
			SenderInformation: &emailtemplates.SenderInformation{
				FromName: ptr("Myself"),
			},
		},
	})
	getResp, getErr := client.EmailTemplates.Get(ctx, emailtemplates.GetRequest{
		ProjectID:  project.LiveProjectID,
		TemplateID: createResp.EmailTemplate.TemplateID,
	})

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, getErr)
	assert.Equal(t, "Updated email template", *getResp.EmailTemplate.Name)
	assert.Equal(t, "Myself", *getResp.EmailTemplate.SenderInformation.FromName)
}

func TestEmailTemplatesClient_Delete(t *testing.T) {
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()
	createResp := createNewEmailTemplate(t, ctx, client, project)

	// Act
	_, err := client.EmailTemplates.Delete(ctx, emailtemplates.DeleteRequest{
		ProjectID:  project.LiveProjectID,
		TemplateID: createResp.EmailTemplate.TemplateID,
	})

	// Assert
	assert.NoError(t, err)
}
