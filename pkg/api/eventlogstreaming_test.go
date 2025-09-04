package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/eventlogstreaming"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

func TestEventLogStreamingClient_Create(t *testing.T) {
	t.Run("create datadog config", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteUS,
					APIKey: "1234567890abcdef1234567890abcdef",
				},
			},
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, eventlogstreaming.DestinationTypeDatadog, resp.EventLogStreamingConfig.DestinationType)
		assert.Equal(t, eventlogstreaming.DatadogSiteUS, resp.EventLogStreamingConfig.DestinationConfig.Datadog.Site)
		assert.Equal(t, "1234567890abcdef1234567890abcdef", resp.EventLogStreamingConfig.DestinationConfig.Datadog.APIKey)
		assert.Equal(t, eventlogstreaming.StreamingStatusDisabled, resp.EventLogStreamingConfig.StreamingStatus)
	})

	t.Run("create grafana loki config", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeGrafanaLoki,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				GrafanaLoki: &eventlogstreaming.GrafanaLokiConfig{
					Hostname: "logs-prod-us-central1.grafana.net",
					Username: "test-user",
					Password: "test-password-12345678",
				},
			},
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, eventlogstreaming.DestinationTypeGrafanaLoki, resp.EventLogStreamingConfig.DestinationType)
		assert.Equal(t, "logs-prod-us-central1.grafana.net", resp.EventLogStreamingConfig.DestinationConfig.GrafanaLoki.Hostname)
		assert.Equal(t, "test-user", resp.EventLogStreamingConfig.DestinationConfig.GrafanaLoki.Username)
		assert.Equal(t, "test-password-12345678", resp.EventLogStreamingConfig.DestinationConfig.GrafanaLoki.Password)
		assert.Equal(t, eventlogstreaming.StreamingStatusDisabled, resp.EventLogStreamingConfig.StreamingStatus)
	})
}

func TestEventLogStreamingClient_Get(t *testing.T) {
	t.Run("get datadog config", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create config first
		_, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteUS,
					APIKey: "1234567890abcdef1234567890abcdef",
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.EventLogStreaming.Get(ctx, eventlogstreaming.GetEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, eventlogstreaming.DestinationTypeDatadog, resp.EventLogStreamingConfig.DestinationType)
		assert.Equal(t, eventlogstreaming.DatadogSiteUS, resp.EventLogStreamingConfig.DestinationConfig.Datadog.Site)
		assert.Equal(t, "cdef", resp.EventLogStreamingConfig.DestinationConfig.Datadog.APIKeyLastFour)
		assert.Equal(t, eventlogstreaming.StreamingStatusDisabled, resp.EventLogStreamingConfig.StreamingStatus)
	})

	t.Run("config does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EventLogStreaming.Get(ctx, eventlogstreaming.GetEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestEventLogStreamingClient_Update(t *testing.T) {
	t.Run("update datadog config", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create config first
		_, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteUS,
					APIKey: "1234567890abcdef1234567890abcdef",
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.EventLogStreaming.Update(ctx, eventlogstreaming.UpdateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteEU,
					APIKey: "abcdefabcdefabcdefabcdefabcdefab",
				},
			},
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, eventlogstreaming.DestinationTypeDatadog, resp.EventLogStreamingConfig.DestinationType)
		assert.Equal(t, eventlogstreaming.DatadogSiteEU, resp.EventLogStreamingConfig.DestinationConfig.Datadog.Site)
		assert.Equal(t, "abcdefabcdefabcdefabcdefabcdefab", resp.EventLogStreamingConfig.DestinationConfig.Datadog.APIKey)
	})

	t.Run("update grafana loki config", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create Loki config first
		_, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeGrafanaLoki,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				GrafanaLoki: &eventlogstreaming.GrafanaLokiConfig{
					Hostname: "logs-prod-us-central1.grafana.net",
					Username: "initial-user",
					Password: "initial-password-12345678",
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.EventLogStreaming.Update(ctx, eventlogstreaming.UpdateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeGrafanaLoki,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				GrafanaLoki: &eventlogstreaming.GrafanaLokiConfig{
					Hostname: "logs-prod-eu-west-0.grafana.net",
					Username: "updated-user",
					Password: "updated-password-87654321",
				},
			},
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, eventlogstreaming.DestinationTypeGrafanaLoki, resp.EventLogStreamingConfig.DestinationType)
		assert.Equal(t, "logs-prod-eu-west-0.grafana.net", resp.EventLogStreamingConfig.DestinationConfig.GrafanaLoki.Hostname)
		assert.Equal(t, "updated-user", resp.EventLogStreamingConfig.DestinationConfig.GrafanaLoki.Username)
		assert.Equal(t, "updated-password-87654321", resp.EventLogStreamingConfig.DestinationConfig.GrafanaLoki.Password)
	})

	t.Run("update loki config when datadog config exists", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create Datadog config first
		_, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteUS,
					APIKey: "1234567890abcdef1234567890abcdef",
				},
			},
		})
		require.NoError(t, err)

		// Act - Try to update with Loki config when Datadog exists
		resp, err := client.EventLogStreaming.Update(ctx, eventlogstreaming.UpdateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeGrafanaLoki,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				GrafanaLoki: &eventlogstreaming.GrafanaLokiConfig{
					Hostname: "logs-prod-eu-west-0.grafana.net",
					Username: "test-user",
					Password: "test-password-12345678",
				},
			},
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("config does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EventLogStreaming.Update(ctx, eventlogstreaming.UpdateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteUS,
					APIKey: "1234567890abcdef1234567890abcdef",
				},
			},
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestEventLogStreamingClient_Enable(t *testing.T) {
	t.Run("enable config", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create config first
		_, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteUS,
					APIKey: "1234567890abcdef1234567890abcdef",
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.EventLogStreaming.Enable(ctx, eventlogstreaming.EnableEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.RequestID)
	})

	t.Run("config does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EventLogStreaming.Enable(ctx, eventlogstreaming.EnableEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestEventLogStreamingClient_Disable(t *testing.T) {
	t.Run("disable config", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create and enable config first
		_, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteUS,
					APIKey: "1234567890abcdef1234567890abcdef",
				},
			},
		})
		require.NoError(t, err)

		_, err = client.EventLogStreaming.Enable(ctx, eventlogstreaming.EnableEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.EventLogStreaming.Disable(ctx, eventlogstreaming.DisableEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.RequestID)
	})

	t.Run("config does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EventLogStreaming.Disable(ctx, eventlogstreaming.DisableEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestEventLogStreamingClient_Delete(t *testing.T) {
	t.Run("delete config", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Create config first
		_, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteUS,
					APIKey: "1234567890abcdef1234567890abcdef",
				},
			},
		})
		require.NoError(t, err)

		// Act
		resp, err := client.EventLogStreaming.Delete(ctx, eventlogstreaming.DeleteEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.RequestID)

		// Verify config is deleted by trying to get it
		getResp, err := client.EventLogStreaming.Get(ctx, eventlogstreaming.GetEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})
		assert.Error(t, err)
		assert.Nil(t, getResp)
	})

	t.Run("config does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()

		// Act
		resp, err := client.EventLogStreaming.Delete(ctx, eventlogstreaming.DeleteEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
