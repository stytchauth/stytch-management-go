package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/jwttemplates"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

func TestJWTTemplatesClient_GetJWTTemplate(t *testing.T) {
	t.Run("get session template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		templateContent := `{"custom_user_id": "user-123", "custom_email": "test@example.com"}`
		customAudience := "my-custom-audience"

		// First set a template
		_, err := client.JWTTemplates.Set(ctx, &jwttemplates.SetRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			JWTTemplateType: jwttemplates.TemplateTypeSession,
			TemplateContent: templateContent,
			CustomAudience:  customAudience,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.JWTTemplates.Get(ctx, &jwttemplates.GetRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			JWTTemplateType: jwttemplates.TemplateTypeSession,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, templateContent, resp.JWTTemplate.TemplateContent)
		assert.Equal(t, customAudience, resp.JWTTemplate.CustomAudience)
		assert.Equal(t, jwttemplates.TemplateTypeSession, resp.JWTTemplate.JWTTemplateType)
	})

	t.Run("get m2m template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		templateContent := `{"custom_org_id": "org-456", "custom_name": "Test Organization"}`
		customAudience := "m2m-audience"

		// First set a template
		_, err := client.JWTTemplates.Set(ctx, &jwttemplates.SetRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			JWTTemplateType: jwttemplates.TemplateTypeM2M,
			TemplateContent: templateContent,
			CustomAudience:  customAudience,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.JWTTemplates.Get(ctx, &jwttemplates.GetRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			JWTTemplateType: jwttemplates.TemplateTypeM2M,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, templateContent, resp.JWTTemplate.TemplateContent)
		assert.Equal(t, customAudience, resp.JWTTemplate.CustomAudience)
		assert.Equal(t, jwttemplates.TemplateTypeM2M, resp.JWTTemplate.JWTTemplateType)
	})
}

func TestJWTTemplatesClient_SetJWTTemplate(t *testing.T) {
	t.Run("set session template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		templateContent := `{"custom_user_id": "user-123", "custom_email": "test@example.com"}`
		customAudience := "my-custom-audience"

		// Act
		setResp, err := client.JWTTemplates.Set(ctx, &jwttemplates.SetRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			JWTTemplateType: jwttemplates.TemplateTypeSession,
			TemplateContent: templateContent,
			CustomAudience:  customAudience,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, templateContent, setResp.JWTTemplate.TemplateContent)
		assert.Equal(t, customAudience, setResp.JWTTemplate.CustomAudience)
		assert.Equal(t, jwttemplates.TemplateTypeSession, setResp.JWTTemplate.JWTTemplateType)
	})

	t.Run("set m2m template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		templateContent := `{"custom_org_id": "org-456", "custom_name": "Test Organization"}`
		customAudience := "m2m-audience"

		// Act
		setResp, err := client.JWTTemplates.Set(ctx, &jwttemplates.SetRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			JWTTemplateType: jwttemplates.TemplateTypeM2M,
			TemplateContent: templateContent,
			CustomAudience:  customAudience,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, templateContent, setResp.JWTTemplate.TemplateContent)
		assert.Equal(t, customAudience, setResp.JWTTemplate.CustomAudience)
		assert.Equal(t, jwttemplates.TemplateTypeM2M, setResp.JWTTemplate.JWTTemplateType)
	})
}
