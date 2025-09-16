package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             testRedirectURL1,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
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
		assert.Equal(t, redirecturls.RedirectTypeLogin, resp.RedirectURL.ValidTypes[0].Type)
		assert.True(t, resp.RedirectURL.ValidTypes[0].IsDefault)
	})

	t.Run("create redirect URL with multiple types", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             testRedirectURL2,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
					IsDefault: true,
				},
				{
					Type:      redirecturls.RedirectTypeSignup,
					IsDefault: false,
				},
				{
					Type:      redirecturls.RedirectTypeInvite,
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
		typeMap := make(map[redirecturls.RedirectType]bool)
		for _, validType := range resp.RedirectURL.ValidTypes {
			typeMap[validType.Type] = validType.IsDefault
		}
		assert.Contains(t, typeMap, redirecturls.RedirectTypeLogin)
		assert.Contains(t, typeMap, redirecturls.RedirectTypeSignup)
		assert.Contains(t, typeMap, redirecturls.RedirectTypeInvite)
	})

	t.Run("create redirect URL with do not promote defaults", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             testRedirectURL3,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeResetPassword,
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
		assert.Equal(t, redirecturls.RedirectTypeResetPassword, resp.RedirectURL.ValidTypes[0].Type)
	})

	t.Run("create duplicate redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		duplicateURL := "https://duplicate.example.com/callback"

		// Create first redirect URL
		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             duplicateURL,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		// Act
		// Try to create the same URL again - should succeed but update the existing one
		_, err = client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             duplicateURL,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeSignup,
					IsDefault: true,
				},
			},
		})

		getresp, geterr := client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
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
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Create multiple redirect URLs
		url1 := "https://getall1.example.com/callback"
		url2 := "https://getall2.example.com/callback"

		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             url1,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		_, err = client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             url2,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeSignup,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		resp, err := client.RedirectURLs.GetAll(ctx, redirecturls.GetAllRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
		})

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.GreaterOrEqual(t, len(resp.RedirectURLs), 2)

		// Check that both URLs are present
		urlMap := make(map[string]bool)
		for _, redirectURL := range resp.RedirectURLs {
			urlMap[redirectURL.URL] = true
		}
		assert.Contains(t, urlMap, url1)
		assert.Contains(t, urlMap, url2)
	})

	t.Run("get all redirect URLs for empty project", func(t *testing.T) {
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		resp, err := client.RedirectURLs.GetAll(ctx, redirecturls.GetAllRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
		})

		require.NoError(t, err)
		assert.NotNil(t, resp)
		// Environments come with one default redirect URL
		assert.LessOrEqual(t, len(resp.RedirectURLs), 1)
	})
}

func TestRedirectURLsClient_Get(t *testing.T) {
	t.Run("get existing redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		getURL := "https://get.example.com/callback"

		// Create redirect URL first
		createResp, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             getURL,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
					IsDefault: true,
				},
				{
					Type:      redirecturls.RedirectTypeInvite,
					IsDefault: false,
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             getURL,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, createResp.RedirectURL.URL, resp.RedirectURL.URL)
		assert.Len(t, resp.RedirectURL.ValidTypes, 2)

		// Verify the types match
		typeMap := make(map[redirecturls.RedirectType]bool)
		for _, validType := range resp.RedirectURL.ValidTypes {
			typeMap[validType.Type] = validType.IsDefault
		}
		assert.Contains(t, typeMap, redirecturls.RedirectTypeLogin)
		assert.Contains(t, typeMap, redirecturls.RedirectTypeInvite)
		assert.True(t, typeMap[redirecturls.RedirectTypeLogin])
		assert.False(t, typeMap[redirecturls.RedirectTypeInvite])
	})
	t.Run("get existing redirect URL using query params", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		urlWithQueryParams := "https://localhost:3002/login?expires_at={}"
		// Create redirect URL first
		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			// Use one with query params to check that escaping is correct
			URL: urlWithQueryParams,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeInvite,
					IsDefault: false,
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
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
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
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
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		updateURL := "https://update.example.com/callback"

		// Create redirect URL first
		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             updateURL,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		// Act
		// Update with different types
		resp, err := client.RedirectURLs.Update(ctx, redirecturls.UpdateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             updateURL,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
					IsDefault: true,
				},
				{
					Type:      redirecturls.RedirectTypeSignup,
					IsDefault: true,
				},
				{
					Type:      redirecturls.RedirectTypeResetPassword,
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
		typeMap := make(map[redirecturls.RedirectType]bool)
		for _, validType := range resp.RedirectURL.ValidTypes {
			typeMap[validType.Type] = validType.IsDefault
		}
		assert.Contains(t, typeMap, redirecturls.RedirectTypeLogin)
		assert.Contains(t, typeMap, redirecturls.RedirectTypeSignup)
		assert.Contains(t, typeMap, redirecturls.RedirectTypeResetPassword)
	})

	t.Run("update non-existent redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Update(ctx, redirecturls.UpdateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             "https://nonexistent-update.example.com/callback",
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
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
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		deleteURL := "https://delete.example.com/callback"

		// Create redirect URL first
		_, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             deleteURL,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
					IsDefault: true,
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.RedirectURLs.Delete(ctx, redirecturls.DeleteRequest{
			ProjectSlug:          project.ProjectSlug,
			EnvironmentSlug:      TestEnvironment,
			URL:                  deleteURL,
			DoNotPromoteDefaults: false,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Verify redirect URL is deleted
		_, err = client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             deleteURL,
		})
		assert.Error(t, err)
		assert.ErrorContains(t, err, "404")
	})

	t.Run("delete non-existent redirect URL", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.RedirectURLs.Delete(ctx, redirecturls.DeleteRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			URL:             "https://nonexistent-delete.example.com/callback",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
