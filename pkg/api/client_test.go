package api_test

import (
	"context"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/pkg/api"
	"github.com/stytchauth/stytch-management-go/pkg/models/projects"
)

func ptr[T any](v T) *T {
	return &v
}

type testClient struct {
	t *testing.T
	*api.API
}

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
		ProjectName: "Disposable project",
		Vertical:    vertical,
	})
	require.NoError(c.t, err)

	c.t.Cleanup(func() {
		_, err := c.Projects.Delete(ctx, projects.DeleteRequest{
			ProjectID: resp.Project.LiveProjectID,
		})
		require.NoError(c.t, err)
	})

	return resp.Project
}
