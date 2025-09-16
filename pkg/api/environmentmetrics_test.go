package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environmentmetrics"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

// We currently have no way of creating users / organizations / members *within* the project using
// the Management client, so this test is a bit simplistic.
func Test_EnvironmentMetricsGet(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
	ctx := context.Background()

	// Act
	resp, err := client.EnvironmentMetrics.Get(ctx, environmentmetrics.GetRequest{
		ProjectSlug:     env.ProjectSlug,
		EnvironmentSlug: env.EnvironmentSlug,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, environmentmetrics.Metrics{}, resp.Metrics)
}
