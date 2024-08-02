package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/redirecturls"
)

var defaultRedirectURL = "http://localhost:3000/authenticate"

func (c *testClient) createRedirectURL(projectID string, url string, redirectType redirecturls.RedirectType) {
	c.t.Helper()
	_, err := c.RedirectURLs.Create(context.Background(), redirecturls.CreateRequest{
		ProjectID: projectID,
		URL:       url,
		Type:      redirectType,
	})
	require.NoError(c.t, err)
}

func (c *testClient) getRedirectURL(projectID string, url string) redirecturls.RedirectURL {
	c.t.Helper()

	resp, err := c.RedirectURLs.GetAll(context.Background(), redirecturls.GetAllRequest{
		ProjectID: projectID,
	})
	require.NoError(c.t, err)

	for _, r := range resp.RedirectURLs {
		if r.URL == url {
			return r
		}
	}
	c.t.Fatalf("Redirect URL not found: %s", url)
	return redirecturls.RedirectURL{}
}

func TestRedirectURLsClient_Create(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	req := redirecturls.CreateRequest{
		ProjectID: project.TestProject.ID,
		URL:       "http://localhost:3000",
		Type:      redirecturls.RedirectTypeLogin,
		IsDefault: true,
	}

	// Act
	resp, err := client.RedirectURLs.Create(context.Background(), req)
	redirectURL := client.getRedirectURL(project.TestProject.ID, req.URL)

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, resp.MagicLinkURLID, redirectURL.MagicLinkURLID)
	assert.Equal(t, req.URL, redirectURL.URL)
	assert.Contains(t, redirectURL.ValidTypes, req.Type)
	assert.Contains(t, redirectURL.DefaultTypes, req.Type)
}

func TestRedirectURLsClient_GetAll(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	client.createRedirectURL(project.TestProject.ID, "http://localhost:3000", redirecturls.RedirectTypeLogin)
	client.createRedirectURL(project.TestProject.ID, "http://localhost:3000", redirecturls.RedirectTypeSignup)
	client.createRedirectURL(project.TestProject.ID, "http://localhost:3001", redirecturls.RedirectTypeInvite)
	expected := map[string][]redirecturls.RedirectType{
		"http://localhost:3000": {redirecturls.RedirectTypeLogin, redirecturls.RedirectTypeSignup},
		"http://localhost:3001": {redirecturls.RedirectTypeInvite},
	}

	// Act
	resp, err := client.RedirectURLs.GetAll(context.Background(), redirecturls.GetAllRequest{
		ProjectID: project.TestProject.ID,
	})
	actual := make(map[string][]redirecturls.RedirectType)
	for _, r := range resp.RedirectURLs {
		// Don't include the default redirect URL in the actual map
		if r.URL != defaultRedirectURL {
			actual[r.URL] = r.ValidTypes
		}
	}

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, expected, actual)
}

func TestRedirectURLsClient_RemoveValidType(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	client.createRedirectURL(project.TestProject.ID, "http://localhost:3000", redirecturls.RedirectTypeLogin)
	client.createRedirectURL(project.TestProject.ID, "http://localhost:3000", redirecturls.RedirectTypeSignup)

	// Act
	_, err := client.RedirectURLs.RemoveValidType(context.Background(), redirecturls.RemoveValidTypeRequest{
		ProjectID: project.TestProject.ID,
		URL:       "http://localhost:3000",
		Type:      redirecturls.RedirectTypeLogin,
	})
	redirectURL := client.getRedirectURL(project.TestProject.ID, "http://localhost:3000")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, []redirecturls.RedirectType{redirecturls.RedirectTypeSignup}, redirectURL.ValidTypes)
}
