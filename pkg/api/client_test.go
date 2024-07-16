package api_test

import (
	"os"
	"testing"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api"
)

// NewTestClient is a test helper function that returns a new API client.
// It relies on the environment variables STYTCH_WORKSPACE_KEY_ID and STYTCH_WORKSPACE_KEY_SECRET being set.
func NewTestClient(t *testing.T) *api.API {
	t.Helper()

	keyID := os.Getenv("STYTCH_WORKSPACE_KEY_ID")
	keySecret := os.Getenv("STYTCH_WORKSPACE_KEY_SECRET")
	if keyID == "" || keySecret == "" {
		t.Skip("STYTCH_WORKSPACE_KEY_ID and STYTCH_WORKSPACE_KEY_SECRET environment variables are required for this test")
	}

	return api.NewClient(keyID, keySecret)
}

// General testing utilities

func GetProjectID(t *testing.T) string {
	t.Helper()
	projectID := os.Getenv("STYTCH_PROJECT_ID")
	if projectID == "" {
		t.Skip("STYTCH_PROJECT_ID environment variable is required for this test")
	}
	return projectID
}
