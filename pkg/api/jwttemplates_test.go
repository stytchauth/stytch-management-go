package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/jwttemplates"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"
)

func TestJWTTemplatesClient_Set(t *testing.T) {
	t.Run("sessions template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		expected := jwttemplates.JWTTemplate{
			TemplateContent: "{\"bid\": {{member.trusted_metadata.billing_id}}}",
			CustomAudience:  "audience",
			TemplateType:    jwttemplates.TemplateTypeSession,
		}

		// Act
		resp, err := client.JWTTemplates.Set(ctx, &jwttemplates.SetRequest{
			ProjectID:   project.LiveProjectID,
			JWTTemplate: expected,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, 200)
		assert.Equal(t, expected, resp.JWTTemplate)
	})
	t.Run("m2m template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		expected := jwttemplates.JWTTemplate{
			TemplateContent: "{\"tier\": {{ client.trusted_metadata.subscription_tier }}}",
			CustomAudience:  "audience",
			TemplateType:    jwttemplates.TemplateTypeM2M,
		}

		// Act
		resp, err := client.JWTTemplates.Set(ctx, &jwttemplates.SetRequest{
			ProjectID:   project.LiveProjectID,
			JWTTemplate: expected,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, 200)
		assert.Equal(t, expected, resp.JWTTemplate)
	})
}

func TestJWTTemplatesClient_Get(t *testing.T) {
	t.Run("sessions template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		expected := jwttemplates.JWTTemplate{
			TemplateContent: "{\"bid\": {{member.trusted_metadata.billing_id}}}",
			CustomAudience:  "audience",
			TemplateType:    jwttemplates.TemplateTypeSession,
		}
		_, err := client.JWTTemplates.Set(ctx, &jwttemplates.SetRequest{
			ProjectID:   project.LiveProjectID,
			JWTTemplate: expected,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.JWTTemplates.Get(ctx, &jwttemplates.GetRequest{
			ProjectID:    project.LiveProjectID,
			TemplateType: jwttemplates.TemplateTypeSession,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, 200)
		assert.Equal(t, expected, resp.JWTTemplate)
	})
	t.Run("m2m template", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		expected := jwttemplates.JWTTemplate{
			TemplateContent: "{\"tier\": {{ client.trusted_metadata.subscription_tier }}}",
			CustomAudience:  "audience",
			TemplateType:    jwttemplates.TemplateTypeM2M,
		}
		_, err := client.JWTTemplates.Set(ctx, &jwttemplates.SetRequest{
			ProjectID:   project.LiveProjectID,
			JWTTemplate: expected,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.JWTTemplates.Get(ctx, &jwttemplates.GetRequest{
			ProjectID:    project.LiveProjectID,
			TemplateType: jwttemplates.TemplateTypeM2M,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, resp.StatusCode, 200)
		assert.Equal(t, expected, resp.JWTTemplate)
	})
}
