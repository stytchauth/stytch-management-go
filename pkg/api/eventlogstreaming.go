package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/eventlogstreaming"
)

type EventLogStreamingClient struct {
	client *internal.Client
}

func newEventLogStreamingClient(c *internal.Client) *EventLogStreamingClient {
	return &EventLogStreamingClient{client: c}
}

// Get retrieves an event log streaming config for an environment.
func (c *EventLogStreamingClient) Get(
	ctx context.Context,
	body eventlogstreaming.GetEventLogStreamingRequest,
) (*eventlogstreaming.GetEventLogStreamingResponse, error) {
	var resp eventlogstreaming.GetEventLogStreamingResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/event_log_streaming/%s", body.Project, body.Environment, string(body.DestinationType)),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Create creates an event log streaming config for an environment.
func (c *EventLogStreamingClient) Create(
	ctx context.Context,
	body eventlogstreaming.CreateEventLogStreamingRequest,
) (*eventlogstreaming.CreateEventLogStreamingResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp eventlogstreaming.CreateEventLogStreamingResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/event_log_streaming", body.Project, body.Environment),
		nil,
		jsonBody,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Update updates an event log streaming config for an environment.
func (c *EventLogStreamingClient) Update(
	ctx context.Context,
	body eventlogstreaming.UpdateEventLogStreamingRequest,
) (*eventlogstreaming.UpdateEventLogStreamingResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp eventlogstreaming.UpdateEventLogStreamingResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPut,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/event_log_streaming/%s", body.Project, body.Environment, string(body.DestinationType)),
		nil,
		jsonBody,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Delete deletes an event log streaming config for an environment.
func (c *EventLogStreamingClient) Delete(
	ctx context.Context,
	body eventlogstreaming.DeleteEventLogStreamingRequest,
) (*eventlogstreaming.DeleteEventLogStreamingResponse, error) {
	var resp eventlogstreaming.DeleteEventLogStreamingResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/event_log_streaming/%s", body.Project, body.Environment, string(body.DestinationType)),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Enable starts streaming event logs for an environment to a destination.
func (c *EventLogStreamingClient) Enable(
	ctx context.Context,
	body eventlogstreaming.EnableEventLogStreamingRequest,
) (*eventlogstreaming.EnableEventLogStreamingResponse, error) {
	var resp eventlogstreaming.EnableEventLogStreamingResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/event_log_streaming/%s/enable", body.Project, body.Environment, string(body.DestinationType)),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Disable stops streaming event logs for an environment to a destination.
func (c *EventLogStreamingClient) Disable(
	ctx context.Context,
	body eventlogstreaming.DisableEventLogStreamingRequest,
) (*eventlogstreaming.DisableEventLogStreamingResponse, error) {
	var resp eventlogstreaming.DisableEventLogStreamingResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/event_log_streaming/%s/disable", body.Project, body.Environment, string(body.DestinationType)),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
