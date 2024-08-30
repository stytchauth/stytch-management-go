package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/redirecturls"
)

func (c *testClient) createRedirectURL(projectID string, url string, redirectType redirecturls.RedirectType) {
	c.t.Helper()
	_, err := c.RedirectURLs.Create(context.Background(), redirecturls.CreateRequest{
		ProjectID: projectID,
		RedirectURL: redirecturls.RedirectURL{
			URL: url,
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirectType,
					IsDefault: true,
				},
			},
		},
	})
	require.NoError(c.t, err)
}

func TestRedirectURLsClient_Create(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	newRedirectUrl := redirecturls.RedirectURL{
		URL: "http://localhost:3000",
		ValidTypes: []redirecturls.URLRedirectType{
			{
				Type:      redirecturls.RedirectTypeLogin,
				IsDefault: true,
			},
		},
	}

	// Act
	resp, err := client.RedirectURLs.Create(
		context.Background(),
		redirecturls.CreateRequest{
			ProjectID:   project.TestProjectID,
			RedirectURL: newRedirectUrl,
		},
	)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, resp.RedirectURL, newRedirectUrl)
}

func TestRedirectURLsClient_GetAll(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	client.createRedirectURL(project.TestProjectID, "http://localhost:3000", redirecturls.RedirectTypeLogin)
	client.createRedirectURL(project.TestProjectID, "http://localhost:3000", redirecturls.RedirectTypeSignup)
	client.createRedirectURL(project.TestProjectID, "http://localhost:3001", redirecturls.RedirectTypeInvite)

	expected := []redirecturls.RedirectURL{
		{
			URL: "http://localhost:3000",
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeLogin,
					IsDefault: true,
				},
				{
					Type:      redirecturls.RedirectTypeSignup,
					IsDefault: true,
				},
			},
		},
		{
			URL: "http://localhost:3001",
			ValidTypes: []redirecturls.URLRedirectType{
				{
					Type:      redirecturls.RedirectTypeInvite,
					IsDefault: true,
				},
			},
		},
	}

	// Act
	resp, err := client.RedirectURLs.GetAll(context.Background(), redirecturls.GetAllRequest{
		ProjectID: project.TestProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.ElementsMatch(t, expected, resp.RedirectURLs)
}

func TestRedirectURLsClient_Delete(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	client.createRedirectURL(project.TestProjectID, "http://localhost:3000", redirecturls.RedirectTypeLogin)

	// Act
	_, err := client.RedirectURLs.Delete(context.Background(), redirecturls.DeleteRequest{
		ProjectID: project.TestProjectID,
		URL:       "http://localhost:3000",
	})

	// Assert
	assert.NoError(t, err)
}

func TestRedirectURLsClient_Get(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	newRedirectUrl := redirecturls.RedirectURL{
		URL: "http://localhost:3000",
		ValidTypes: []redirecturls.URLRedirectType{
			{
				Type:      redirecturls.RedirectTypeLogin,
				IsDefault: true,
			},
		},
	}
	client.createRedirectURL(project.TestProjectID, newRedirectUrl.URL, redirecturls.RedirectTypeLogin)

	// Act
	resp, err := client.RedirectURLs.Get(
		context.Background(),
		redirecturls.GetRequest{
			ProjectID: project.TestProjectID,
			URL:       newRedirectUrl.URL,
		},
	)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, newRedirectUrl, resp.RequestID)
}

func TestRedirectURLsClient_Update(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	updatedRedirectUrl := redirecturls.RedirectURL{
		URL: "http://localhost:3000",
		ValidTypes: []redirecturls.URLRedirectType{
			{
				Type:      redirecturls.RedirectTypeInvite,
				IsDefault: true,
			},
			{
				Type:      redirecturls.RedirectTypeSignup,
				IsDefault: true,
			},
		},
	}
	client.createRedirectURL(project.TestProjectID, "http://localhost:3000", redirecturls.RedirectTypeLogin)

	// Act
	resp, err := client.RedirectURLs.Update(
		context.Background(),
		redirecturls.UpdateRequest{
			ProjectID:   project.TestProjectID,
			RedirectURL: updatedRedirectUrl,
		},
	)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, updatedRedirectUrl, resp.RequestID)
}
