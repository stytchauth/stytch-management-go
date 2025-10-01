package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/api"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

func ptr[T any](v T) *T {
	return &v
}

type testClient struct {
	t *testing.T
	*api.API
}

// This is the name of the first live environment created
// when a new project is created from the dashboard. We choose the same name in tests.
const LiveEnvironment string = "production"

// NewTestClient is a test helper function that returns a new API client.
// It relies on the environment variables STYTCH_WORKSPACE_KEY_ID and STYTCH_WORKSPACE_KEY_SECRET being set.
func NewTestClient(t *testing.T) *testClient {
	t.Helper()

	keyID := os.Getenv("STYTCH_WORKSPACE_KEY_ID")
	keySecret := os.Getenv("STYTCH_WORKSPACE_KEY_SECRET")
	if keyID == "" || keySecret == "" {
		t.Skip("STYTCH_WORKSPACE_KEY_ID and STYTCH_WORKSPACE_KEY_SECRET environment variables are required for this test")
	}

	var opts []api.APIOption
	if baseURI := os.Getenv("STYTCH_WORKSPACE_BASE_URI"); baseURI != "" {
		opts = append(opts, api.WithBaseURI(baseURI))
	}

	return &testClient{
		t:   t,
		API: api.NewClient(keyID, keySecret, opts...),
	}
}

func (c *testClient) DisposableProject(vertical projects.Vertical) projects.Project {
	c.t.Helper()
	ctx := context.Background()
	resp, err := c.Projects.Create(ctx, projects.CreateRequest{
		Name:     "Disposable Project",
		Vertical: vertical,
	})
	require.NoError(c.t, err)

	// Create a live environment since otherwise we cannot create disposable test environments.
	_, err = c.Environments.Create(ctx, environments.CreateRequest{
		ProjectSlug:     resp.Project.ProjectSlug,
		Name:            LiveEnvironment,
		Type:            environments.EnvironmentTypeLive,
		EnvironmentSlug: ptr(LiveEnvironment),
	})
	require.NoError(c.t, err)

	c.t.Cleanup(func() {
		_, err := c.Projects.Delete(ctx, projects.DeleteRequest{
			ProjectSlug: resp.Project.ProjectSlug,
		})
		require.NoError(c.t, err)
	})

	return resp.Project
}

func (c *testClient) DisposableEnvironment(
	vertical projects.Vertical, environmentType environments.EnvironmentType,
) environments.Environment {
	c.t.Helper()
	project := c.DisposableProject(vertical)
	ctx := context.Background()

	envResp, err := c.Environments.GetAll(ctx, environments.GetAllRequest{
		ProjectSlug: project.ProjectSlug,
	})
	require.NoError(c.t, err)

	// Projects are created with both a live and test environment, so return the one that matches the
	// requested type (if it exists).
	for _, env := range envResp.Environments {
		if env.Type == environmentType {
			return env
		}
	}

	// Otherwise, we need to create a new one because one of that type does not exist
	createResp, err := c.Environments.Create(ctx, environments.CreateRequest{
		ProjectSlug: project.ProjectSlug,
		Name:        "Disposable Environment",
		Type:        environmentType,
	})
	require.NoError(c.t, err)
	return createResp.Environment
}
