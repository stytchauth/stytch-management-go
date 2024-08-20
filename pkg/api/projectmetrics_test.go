package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projectmetrics"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
)

// We currently have no way of creating users/orgs/members *within* the project
// with the Management client, so this test is a bit simplistic.
func Test_ProjectMetricsGet(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	ctx := context.Background()

	// Act
	resp, err := client.ProjectMetrics.Get(ctx, projectmetrics.GetRequest{
		ProjectID: project.LiveProject.ID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, projectmetrics.Metrics{}, resp.Metrics)
}
