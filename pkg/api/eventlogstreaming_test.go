package api_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stytchauth/stytch-management-go/v3/pkg/models/eventlogstreaming"
// 	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
// )

// func (c *testClient) createEventLogStreamingConfig(
// 	projectID string,
// 	destinationType eventlogstreaming.DestinationType,
// 	destinationConfig eventlogstreaming.DestinationConfig,
// ) {
// 	c.t.Helper()
// 	_, err := c.EventLogStreaming.Create(context.Background(), eventlogstreaming.CreateEventLogStreamingRequest{
// 		ProjectID:         projectID,
// 		DestinationType:   destinationType,
// 		DestinationConfig: destinationConfig,
// 	})
// 	require.NoError(c.t, err)
// }

// func (c *testClient) cleanupEventLogStreamingConfig(projectID string, destinationType eventlogstreaming.DestinationType) {
// 	c.t.Helper()
// 	c.t.Cleanup(func() {
// 		_, err := c.EventLogStreaming.Delete(context.Background(), eventlogstreaming.DeleteEventLogStreamingRequest{
// 			ProjectID:       projectID,
// 			DestinationType: destinationType,
// 		})
// 		require.NoError(c.t, err)
// 	})
// }

// func TestEventLogStreamingClient_Create(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	project := client.DisposableProject(projects.VerticalB2B)
// 	destinationConfig := eventlogstreaming.DestinationConfig{
// 		Datadog: &eventlogstreaming.DatadogConfig{
// 			APIKey: "1234567890abcdef1234567890abcdef",
// 			Site:   eventlogstreaming.DatadogSiteUS,
// 		},
// 	}

// 	t.Run("happy path", func(t *testing.T) {
// 		// Act
// 		resp, err := client.EventLogStreaming.Create(context.Background(), eventlogstreaming.CreateEventLogStreamingRequest{
// 			ProjectID:         project.TestProjectID,
// 			DestinationType:   eventlogstreaming.DestinationTypeDatadog,
// 			DestinationConfig: destinationConfig,
// 		})
// 		client.cleanupEventLogStreamingConfig(project.TestProjectID, eventlogstreaming.DestinationTypeDatadog)

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.Equal(t, eventlogstreaming.DestinationTypeDatadog, resp.EventLogStreamingConfig.DestinationType)
// 		assert.Equal(t, destinationConfig, resp.EventLogStreamingConfig.DestinationConfig)
// 		assert.Equal(t, eventlogstreaming.StreamingStatusDisabled, resp.EventLogStreamingConfig.StreamingStatus)
// 	})

// 	t.Run("incorrect destination type for config returns error", func(t *testing.T) {
// 		// Act
// 		_, err := client.EventLogStreaming.Create(context.Background(), eventlogstreaming.CreateEventLogStreamingRequest{
// 			ProjectID:         project.TestProjectID,
// 			DestinationType:   eventlogstreaming.DestinationTypeGrafanaLoki,
// 			DestinationConfig: destinationConfig,
// 		})

// 		// Assert
// 		assert.Error(t, err)
// 	})
// }

// func TestEventLogStreamingClient_Get(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	project := client.DisposableProject(projects.VerticalB2B)
// 	destinationConfig := eventlogstreaming.DestinationConfig{
// 		Datadog: &eventlogstreaming.DatadogConfig{
// 			APIKey: "1234567890abcdef1234567890abcdef",
// 			Site:   eventlogstreaming.DatadogSiteUS,
// 		},
// 	}
// 	client.createEventLogStreamingConfig(project.TestProjectID, eventlogstreaming.DestinationTypeDatadog, destinationConfig)
// 	client.cleanupEventLogStreamingConfig(project.TestProjectID, eventlogstreaming.DestinationTypeDatadog)

// 	// Act
// 	resp, err := client.EventLogStreaming.Get(context.Background(), eventlogstreaming.GetEventLogStreamingRequest{
// 		ProjectID:       project.TestProjectID,
// 		DestinationType: eventlogstreaming.DestinationTypeDatadog,
// 	})

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.Equal(t, eventlogstreaming.DestinationTypeDatadog, resp.EventLogStreamingConfig.DestinationType)
// 	// The destination is masked, so we need to check one at a time
// 	assert.Equal(t, eventlogstreaming.StreamingStatusDisabled, resp.EventLogStreamingConfig.StreamingStatus)
// 	assert.Equal(t, eventlogstreaming.DatadogSiteUS, resp.EventLogStreamingConfig.DestinationConfig.Datadog.Site)
// 	assert.Equal(t, "cdef", resp.EventLogStreamingConfig.DestinationConfig.Datadog.APIKeyLastFour)
// }

// func TestEventLogStreamingClient_Update(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	project := client.DisposableProject(projects.VerticalB2B)
// 	destinationConfig := eventlogstreaming.DestinationConfig{
// 		Datadog: &eventlogstreaming.DatadogConfig{
// 			APIKey: "1234567890abcdef1234567890abcdef",
// 			Site:   eventlogstreaming.DatadogSiteUS5,
// 		},
// 	}
// 	client.createEventLogStreamingConfig(project.TestProjectID, eventlogstreaming.DestinationTypeDatadog, destinationConfig)
// 	client.cleanupEventLogStreamingConfig(project.TestProjectID, eventlogstreaming.DestinationTypeDatadog)

// 	// Act
// 	resp, err := client.EventLogStreaming.Update(context.Background(), eventlogstreaming.UpdateEventLogStreamingRequest{
// 		ProjectID:       project.TestProjectID,
// 		DestinationType: eventlogstreaming.DestinationTypeDatadog,
// 		DestinationConfig: eventlogstreaming.DestinationConfig{
// 			Datadog: &eventlogstreaming.DatadogConfig{
// 				APIKey: "00000000000000000000000000000000",
// 				Site:   eventlogstreaming.DatadogSiteUS,
// 			},
// 		},
// 	})

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.Equal(t, resp.EventLogStreamingConfig.DestinationType, eventlogstreaming.DestinationTypeDatadog)
// 	assert.Equal(t, resp.EventLogStreamingConfig.DestinationConfig.Datadog.Site, eventlogstreaming.DatadogSiteUS)
// 	assert.Equal(t, resp.EventLogStreamingConfig.DestinationConfig.Datadog.APIKey, "00000000000000000000000000000000")
// 	assert.Equal(t, resp.EventLogStreamingConfig.StreamingStatus, eventlogstreaming.StreamingStatusDisabled)
// }

// func TestEventLogStreamingClient_Enable(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	project := client.DisposableProject(projects.VerticalB2B)
// 	destinationConfig := eventlogstreaming.DestinationConfig{
// 		Datadog: &eventlogstreaming.DatadogConfig{
// 			APIKey: "1234567890abcdef1234567890abcdef",
// 			Site:   eventlogstreaming.DatadogSiteUS5,
// 		},
// 	}
// 	client.createEventLogStreamingConfig(project.TestProjectID, eventlogstreaming.DestinationTypeDatadog, destinationConfig)
// 	client.cleanupEventLogStreamingConfig(project.TestProjectID, eventlogstreaming.DestinationTypeDatadog)

// 	// Act
// 	_, err := client.EventLogStreaming.Enable(context.Background(), eventlogstreaming.EnableEventLogStreamingRequest{
// 		ProjectID:       project.TestProjectID,
// 		DestinationType: eventlogstreaming.DestinationTypeDatadog,
// 	})

// 	// Assert
// 	assert.NoError(t, err)
// }

// func TestEventLogStreamingClient_Disable(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	project := client.DisposableProject(projects.VerticalB2B)
// 	destinationConfig := eventlogstreaming.DestinationConfig{
// 		Datadog: &eventlogstreaming.DatadogConfig{
// 			APIKey: "1234567890abcdef1234567890abcdef",
// 			Site:   eventlogstreaming.DatadogSiteUS5,
// 		},
// 	}
// 	client.createEventLogStreamingConfig(project.TestProjectID, eventlogstreaming.DestinationTypeDatadog, destinationConfig)
// 	client.cleanupEventLogStreamingConfig(project.TestProjectID, eventlogstreaming.DestinationTypeDatadog)
// 	// Must enable before disabling a destination
// 	_, err := client.EventLogStreaming.Enable(context.Background(), eventlogstreaming.EnableEventLogStreamingRequest{
// 		ProjectID:       project.TestProjectID,
// 		DestinationType: eventlogstreaming.DestinationTypeDatadog,
// 	})
// 	require.NoError(t, err)

// 	// Act
// 	_, err = client.EventLogStreaming.Disable(context.Background(), eventlogstreaming.DisableEventLogStreamingRequest{
// 		ProjectID:       project.TestProjectID,
// 		DestinationType: eventlogstreaming.DestinationTypeDatadog,
// 	})

// 	// Assert
// 	assert.NoError(t, err)
// }
