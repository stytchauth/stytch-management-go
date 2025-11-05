package api_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/emailtemplates"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

func randomID() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return fmt.Sprintf("test-template-%d", r.Intn(1000000))
}

func makeTestPrebuiltTemplate(templateID string) emailtemplates.PrebuiltCustomization {
	return emailtemplates.PrebuiltCustomization{
		ButtonBorderRadius: ptr(float32(5.0)),
		ButtonColor:        ptr("#007BFF"),
		ButtonTextColor:    ptr("#FFFFFF"),
		FontFamily:         emailtemplates.FontFamilyArial,
		TextAlignment:      emailtemplates.TextAlignmentCenter,
	}
}

func senderInformation() *emailtemplates.SenderInformation {
	return &emailtemplates.SenderInformation{
		FromLocalPart:    ptr("noreply"),
		FromName:         ptr("No Reply"),
		ReplyToLocalPart: ptr("support"),
		ReplyToName:      ptr("Support Team"),
	}
}

// NOTE: Custom HTML Templates can only be managed in
// projects with verified domains, so they are excluded
// from these tests.

func TestEmailTemplatesClient_Create(t *testing.T) {
	t.Run("create prebuilt template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		templateID := randomID()
		template := makeTestPrebuiltTemplate(templateID)

		// Act
		resp, err := client.EmailTemplates.Create(ctx, emailtemplates.CreateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            templateID,
			Name:                  ptr("Test Prebuilt Template"),
			SenderInformation:     senderInformation(),
			PrebuiltCustomization: &template,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, templateID, resp.EmailTemplate.TemplateID)
		assert.Equal(t, "Test Prebuilt Template", *resp.EmailTemplate.Name)
		assert.NotNil(t, resp.EmailTemplate.SenderInformation)
		assert.NotNil(t, resp.EmailTemplate.PrebuiltCustomization)
		assert.Nil(t, resp.EmailTemplate.CustomHTMLCustomization)
	})
}

func TestEmailTemplatesClient_Get(t *testing.T) {
	t.Run("get existing prebuilt template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		templateID := randomID()
		template := makeTestPrebuiltTemplate(templateID)

		// Create template first
		createResp, err := client.EmailTemplates.Create(ctx, emailtemplates.CreateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            templateID,
			Name:                  ptr("Test Prebuilt Template"),
			SenderInformation:     senderInformation(),
			PrebuiltCustomization: &template,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.EmailTemplates.Get(ctx, emailtemplates.GetRequest{
			ProjectSlug: project.ProjectSlug,
			TemplateID:  templateID,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, createResp.EmailTemplate.TemplateID, resp.EmailTemplate.TemplateID)
		assert.Equal(t, createResp.EmailTemplate.Name, resp.EmailTemplate.Name)
		assert.NotNil(t, resp.EmailTemplate.PrebuiltCustomization)
		assert.Nil(t, resp.EmailTemplate.CustomHTMLCustomization)
	})
	t.Run("get non-existent template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EmailTemplates.Get(ctx, emailtemplates.GetRequest{
			ProjectSlug: project.ProjectSlug,
			TemplateID:  "non-existent-template",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
	t.Run("missing template ID", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EmailTemplates.Get(ctx, emailtemplates.GetRequest{
			ProjectSlug: project.ProjectSlug,
			// TemplateID is intentionally omitted.
		})

		// Assert
		assert.ErrorContains(t, err, "cannot be empty")
		assert.Nil(t, resp)
	})
}

func TestEmailTemplatesClient_GetAll(t *testing.T) {
	t.Run("get all templates", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create multiple templates
		template1ID := randomID()
		template2ID := randomID()

		template1 := makeTestPrebuiltTemplate(template1ID)
		template2 := makeTestPrebuiltTemplate(template2ID)

		_, err := client.EmailTemplates.Create(ctx, emailtemplates.CreateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            template1ID,
			Name:                  ptr("Test Prebuilt Template"),
			SenderInformation:     senderInformation(),
			PrebuiltCustomization: &template1,
		})
		require.NoError(t, err)

		_, err = client.EmailTemplates.Create(ctx, emailtemplates.CreateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            template2ID,
			Name:                  ptr("Test Prebuilt Template 2"),
			SenderInformation:     senderInformation(),
			PrebuiltCustomization: &template2,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.EmailTemplates.GetAll(ctx, emailtemplates.GetAllRequest{
			ProjectSlug: project.ProjectSlug,
		})

		// Assert
		assert.NoError(t, err)
		assert.GreaterOrEqual(t, len(resp.EmailTemplates), 2)

		// Check that both templates are present
		templateIDs := make(map[string]bool)
		for _, template := range resp.EmailTemplates {
			templateIDs[template.TemplateID] = true
		}
		assert.Contains(t, templateIDs, template1ID)
		assert.Contains(t, templateIDs, template2ID)
	})
}

func TestEmailTemplatesClient_Update(t *testing.T) {
	t.Run("update prebuilt template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		templateID := randomID()
		template := makeTestPrebuiltTemplate(templateID)

		// Create template first
		_, err := client.EmailTemplates.Create(ctx, emailtemplates.CreateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            templateID,
			Name:                  ptr("Test Prebuilt Template"),
			SenderInformation:     senderInformation(),
			PrebuiltCustomization: &template,
		})
		require.NoError(t, err)

		// Act - update the template
		newName := "Updated Prebuilt Template"
		updatedTemplate := emailtemplates.PrebuiltCustomization{
			ButtonBorderRadius: ptr(float32(10.0)),
			ButtonColor:        ptr("#FF0000"),
			ButtonTextColor:    ptr("#000000"),
			FontFamily:         emailtemplates.FontFamilyHelvetica,
			TextAlignment:      emailtemplates.TextAlignmentLeft,
		}

		resp, err := client.EmailTemplates.Update(ctx, emailtemplates.UpdateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            templateID,
			Name:                  &newName,
			PrebuiltCustomization: &updatedTemplate,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, templateID, resp.EmailTemplate.TemplateID)
		assert.Equal(t, newName, *resp.EmailTemplate.Name)
		assert.NotNil(t, resp.EmailTemplate.PrebuiltCustomization)
		assert.Equal(t, updatedTemplate, *resp.EmailTemplate.PrebuiltCustomization)
		assert.Nil(t, resp.EmailTemplate.CustomHTMLCustomization)
	})

	t.Run("update non-existent template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		newName := "Non-existent Template"
		resp, err := client.EmailTemplates.Update(ctx, emailtemplates.UpdateRequest{
			ProjectSlug: project.ProjectSlug,
			TemplateID:  "non-existent-template",
			Name:        &newName,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestEmailTemplatesClient_Delete(t *testing.T) {
	t.Run("delete existing template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		templateID := randomID()
		template := makeTestPrebuiltTemplate(templateID)

		// Create template first
		_, err := client.EmailTemplates.Create(ctx, emailtemplates.CreateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            templateID,
			Name:                  ptr("Test Prebuilt Template"),
			SenderInformation:     senderInformation(),
			PrebuiltCustomization: &template,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.EmailTemplates.Delete(ctx, emailtemplates.DeleteRequest{
			ProjectSlug: project.ProjectSlug,
			TemplateID:  templateID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Verify template is deleted by trying to get it
		_, err = client.EmailTemplates.Get(ctx, emailtemplates.GetRequest{
			ProjectSlug: project.ProjectSlug,
			TemplateID:  templateID,
		})
		assert.Error(t, err)
	})
}

func TestEmailTemplatesClient_SetDefault(t *testing.T) {
	t.Run("set default template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		templateID := randomID()
		template := makeTestPrebuiltTemplate(templateID)

		// Create template first
		_, err := client.EmailTemplates.Create(ctx, emailtemplates.CreateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            templateID,
			Name:                  ptr("Test Default Template"),
			SenderInformation:     senderInformation(),
			PrebuiltCustomization: &template,
		})
		require.NoError(t, err)

		// Act - set as default
		resp, err := client.EmailTemplates.SetDefault(ctx, emailtemplates.SetDefaultRequest{
			ProjectSlug:       project.ProjectSlug,
			EmailTemplateType: emailtemplates.TemplateTypePrebuilt,
			TemplateID:        templateID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})

	t.Run("set default with non-existent template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EmailTemplates.SetDefault(ctx, emailtemplates.SetDefaultRequest{
			ProjectSlug:       project.ProjectSlug,
			EmailTemplateType: emailtemplates.TemplateTypeLogin,
			TemplateID:        "non-existent-template",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestEmailTemplatesClient_GetDefault(t *testing.T) {
	t.Run("get default template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		templateID := randomID()
		template := makeTestPrebuiltTemplate(templateID)

		// Create template first
		_, err := client.EmailTemplates.Create(ctx, emailtemplates.CreateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            templateID,
			Name:                  ptr("Test Default Template"),
			SenderInformation:     senderInformation(),
			PrebuiltCustomization: &template,
		})
		require.NoError(t, err)

		// Set as default
		_, err = client.EmailTemplates.SetDefault(ctx, emailtemplates.SetDefaultRequest{
			ProjectSlug:       project.ProjectSlug,
			EmailTemplateType: emailtemplates.TemplateTypePrebuilt,
			TemplateID:        templateID,
		})
		require.NoError(t, err)

		// Act - get default
		resp, err := client.EmailTemplates.GetDefault(ctx, emailtemplates.GetDefaultRequest{
			ProjectSlug:       project.ProjectSlug,
			EmailTemplateType: emailtemplates.TemplateTypePrebuilt,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, templateID, resp.TemplateID)
	})

	t.Run("get default for type with no default set", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EmailTemplates.GetDefault(ctx, emailtemplates.GetDefaultRequest{
			ProjectSlug:       project.ProjectSlug,
			EmailTemplateType: emailtemplates.TemplateTypeSignup,
		})

		// Assert - this should return an error when no default is set
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestEmailTemplatesClient_UnsetDefault(t *testing.T) {
	t.Run("unsetting prebuilt template returns error", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		templateID := randomID()
		template := makeTestPrebuiltTemplate(templateID)

		// Create template first
		_, err := client.EmailTemplates.Create(ctx, emailtemplates.CreateRequest{
			ProjectSlug:           project.ProjectSlug,
			TemplateID:            templateID,
			Name:                  ptr("Test Default Template"),
			SenderInformation:     senderInformation(),
			PrebuiltCustomization: &template,
		})
		require.NoError(t, err)

		// Act - unset default
		resp, err := client.EmailTemplates.UnsetDefault(ctx, emailtemplates.UnsetDefaultRequest{
			ProjectSlug:       project.ProjectSlug,
			EmailTemplateType: emailtemplates.TemplateTypePrebuilt,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("unset default for type with no default set", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EmailTemplates.UnsetDefault(ctx, emailtemplates.UnsetDefaultRequest{
			ProjectSlug:       project.ProjectSlug,
			EmailTemplateType: emailtemplates.TemplateTypeSignup,
		})

		// Assert - this succeeds even if no default was set
		assert.NoError(t, err)
		assert.NotNil(t, resp)
	})
}
