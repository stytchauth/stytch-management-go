package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/redirecturls"
)

const (
	testRedirectURL1 = "https://localhost:3000/callback"
	testRedirectURL2 = "https://localhost:3001/auth/callback"
	testRedirectURL3 = "https://localhost:3002/login"
)

func TestRedirectURLsClient_Create(t *testing.T) {
	t.Run("create redirect URL with single type", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             testRedirectURL1,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeLogin,
					IsDefault: true,
				},
			},
			DoNotPromoteDefaults: false,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testRedirectURL1, resp.RedirectURL.URL)
		assert.Len(t, resp.RedirectURL.ValidTypes, 1)
		assert.Equal(t, redirecturls.RedirectURLTypeLogin, resp.RedirectURL.ValidTypes[0].Type)
		assert.True(t, resp.RedirectURL.ValidTypes[0].IsDefault)
	})

	t.Run("create redirect URL with multiple types", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             testRedirectURL2,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeLogin,
					IsDefault: true,
				},
				{
					Type:      redirecturls.RedirectURLTypeSignup,
					IsDefault: false,
				},
				{
					Type:      redirecturls.RedirectURLTypeInvite,
					IsDefault: true,
				},
			},
			DoNotPromoteDefaults: false,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testRedirectURL2, resp.RedirectURL.URL)
		assert.Len(t, resp.RedirectURL.ValidTypes, 3)

		// Check that all types are present
		typeMap := make(map[redirecturls.RedirectURLType]bool)
		for _, validType := range resp.RedirectURL.ValidTypes {
			typeMap[validType.Type] = validType.IsDefault
		}
		assert.Contains(t, typeMap, redirecturls.RedirectURLTypeLogin)
		assert.Contains(t, typeMap, redirecturls.RedirectURLTypeSignup)
		assert.Contains(t, typeMap, redirecturls.RedirectURLTypeInvite)
	})

	t.Run("create redirect URL with do not promote defaults", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             testRedirectURL3,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeResetPassword,
					IsDefault: false,
				},
			},
			DoNotPromoteDefaults: true,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, testRedirectURL3, resp.RedirectURL.URL)
		assert.Len(t, resp.RedirectURL.ValidTypes, 1)
		assert.Equal(t, redirecturls.RedirectURLTypeResetPassword, resp.RedirectURL.ValidTypes[0].Type)
	})

	t.Run("create duplicate redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		duplicateURL := "https://duplicate.example.com/callback"

		// Create first redirect URL
		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             duplicateURL,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeLogin,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		// Act
		// Try to create the same URL again - should succeed but update the existing one
		_, err = client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             duplicateURL,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeSignup,
					IsDefault: true,
				},
			},
		})

		getresp, geterr := client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             duplicateURL,
		})
		require.NoError(t, geterr)

		// Assert
		assert.NoError(t, err)
		assert.Len(t, getresp.RedirectURL.ValidTypes, 2)
	})
}

func TestRedirectURLsClient_GetAll(t *testing.T) {
	t.Run("get all redirect URLs", func(t *testing.T) {
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Create multiple redirect URLs
		url1 := "https://getall1.example.com/callback"
		url2 := "https://getall2.example.com/callback"

		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             url1,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeLogin,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		_, err = client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             url2,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeSignup,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		resp, err := client.RedirectURLs.GetAll(ctx, redirecturls.GetAllRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.GreaterOrEqual(t, len(resp.RedirectURLS), 2)

		// Check that both URLs are present
		urlMap := make(map[string]bool)
		for _, redirectURL := range resp.RedirectURLS {
			urlMap[redirectURL.URL] = true
		}
		assert.Contains(t, urlMap, url1)
		assert.Contains(t, urlMap, url2)
	})

	t.Run("get all redirect URLs for empty project", func(t *testing.T) {
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		resp, err := client.RedirectURLs.GetAll(ctx, redirecturls.GetAllRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		require.NoError(t, err)
		assert.NotNil(t, resp)
		// Environments come with one default redirect URL
		assert.LessOrEqual(t, len(resp.RedirectURLS), 1)
	})
}

func TestRedirectURLsClient_Get(t *testing.T) {
	t.Run("get existing redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		getURL := "https://get.example.com/callback"

		// Create redirect URL first
		createResp, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             getURL,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeLogin,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             getURL,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, createResp.RedirectURL.URL, resp.RedirectURL.URL)
		assert.Len(t, resp.RedirectURL.ValidTypes, 1)

		// Verify the types match
		typeMap := make(map[redirecturls.RedirectURLType]bool)
		for _, validType := range resp.RedirectURL.ValidTypes {
			typeMap[validType.Type] = validType.IsDefault
		}
		assert.Contains(t, typeMap, redirecturls.RedirectURLTypeLogin)
		assert.True(t, typeMap[redirecturls.RedirectURLTypeLogin])
	})
	t.Run("get existing redirect URL using query params", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		urlWithQueryParams := "https://localhost:3002/login?expires_at={}"
		// Create redirect URL first
		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			// Use one with query params to check that escaping is correct
			URL: urlWithQueryParams,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeInvite,
					IsDefault: false,
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             urlWithQueryParams,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, urlWithQueryParams, resp.RedirectURL.URL)
		assert.Len(t, resp.RedirectURL.ValidTypes, 1)
	})

	t.Run("get non-existent redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             "https://nonexistent.example.com/callback",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestRedirectURLsClient_Update(t *testing.T) {
	t.Run("update redirect URL valid types", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		updateURL := "https://update.example.com/callback"

		// Create redirect URL first
		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             updateURL,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeLogin,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		// Act
		// Update with different types
		resp, err := client.RedirectURLs.Update(ctx, redirecturls.UpdateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             updateURL,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeLogin,
					IsDefault: true,
				},
				{
					Type:      redirecturls.RedirectURLTypeSignup,
					IsDefault: true,
				},
				{
					Type:      redirecturls.RedirectURLTypeResetPassword,
					IsDefault: false,
				},
			},
			DoNotPromoteDefaults: false,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, updateURL, resp.RedirectURL.URL)
		assert.Len(t, resp.RedirectURL.ValidTypes, 3)

		// Verify all types are present
		typeMap := make(map[redirecturls.RedirectURLType]bool)
		for _, validType := range resp.RedirectURL.ValidTypes {
			typeMap[validType.Type] = validType.IsDefault
		}
		assert.Contains(t, typeMap, redirecturls.RedirectURLTypeLogin)
		assert.Contains(t, typeMap, redirecturls.RedirectURLTypeSignup)
		assert.Contains(t, typeMap, redirecturls.RedirectURLTypeResetPassword)
	})

	t.Run("update non-existent redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Update(ctx, redirecturls.UpdateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             "https://nonexistent-update.example.com/callback",
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeLogin,
					IsDefault: true,
				},
			},
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestRedirectURLsClient_Delete(t *testing.T) {
	t.Run("delete existing redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		deleteURL := "https://delete.example.com/callback"

		// Create redirect URL first
		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             deleteURL,
			ValidTypes: []*redirecturls.URLType{
				{
					Type:      redirecturls.RedirectURLTypeLogin,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.RedirectURLs.Delete(ctx, redirecturls.DeleteRequest{
			ProjectSlug:          env.ProjectSlug,
			EnvironmentSlug:      env.EnvironmentSlug,
			URL:                  deleteURL,
			DoNotPromoteDefaults: false,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Verify redirect URL is deleted
		_, err = client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             deleteURL,
		})
		assert.Error(t, err)
		assert.ErrorContains(t, err, "404")
	})

	t.Run("delete non-existent redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Delete(ctx, redirecturls.DeleteRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			URL:             "https://nonexistent-delete.example.com/callback",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
