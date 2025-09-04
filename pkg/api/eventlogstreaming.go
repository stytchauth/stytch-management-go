package api

import (
	"context"
	"encoding/json"

	"github.com/stytchauth/stytch-management-go/v2/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/eventlogstreaming"
)

type EventLogStreamingClient struct {
	client *internal.Client
}

func newEventLogStreamingClient(c *internal.Client) *EventLogStreamingClient {
	return &EventLogStreamingClient{client: c}
}

// Get retrieves the event log streaming config for a project
func (c *EventLogStreamingClient) Get(
	ctx context.Context,
	body eventlogstreaming.GetEventLogStreamingRequest,
) (*eventlogstreaming.GetEventLogStreamingResponse, error) {
	var resp eventlogstreaming.GetEventLogStreamingResponse
	err := c.client.NewRequest(
		ctx,
		"GET",
		"/v1/projects/"+body.ProjectID+"/event_log_streaming/"+string(body.DestinationType),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Create creates the event log streaming config for a project
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
		"POST",
		"/v1/projects/"+body.ProjectID+"/event_log_streaming",
		nil,
		jsonBody,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Update updates the event log streaming config for a project
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
		"PUT",
		"/v1/projects/"+body.ProjectID+"/event_log_streaming/"+string(body.DestinationType),
		nil,
		jsonBody,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Delete deletes the event log streaming config for a project
func (c *EventLogStreamingClient) Delete(
	ctx context.Context,
	body eventlogstreaming.DeleteEventLogStreamingRequest,
) (*eventlogstreaming.DeleteEventLogStreamingResponse, error) {
	var resp eventlogstreaming.DeleteEventLogStreamingResponse
	err := c.client.NewRequest(
		ctx,
		"DELETE",
		"/v1/projects/"+body.ProjectID+"/event_log_streaming/"+string(body.DestinationType),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *EventLogStreamingClient) Enable(
	ctx context.Context,
	body eventlogstreaming.EnableEventLogStreamingRequest,
) (*eventlogstreaming.EnableEventLogStreamingResponse, error) {
	var resp eventlogstreaming.EnableEventLogStreamingResponse
	err := c.client.NewRequest(
		ctx,
		"POST",
		"/v1/projects/"+body.ProjectID+"/event_log_streaming/"+string(body.DestinationType)+"/enable",
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *EventLogStreamingClient) Disable(
	ctx context.Context,
	body eventlogstreaming.DisableEventLogStreamingRequest,
) (*eventlogstreaming.DisableEventLogStreamingResponse, error) {
	var resp eventlogstreaming.DisableEventLogStreamingResponse
	err := c.client.NewRequest(
		ctx,
		"POST",
		"/v1/projects/"+body.ProjectID+"/event_log_streaming/"+string(body.DestinationType)+"/disable",
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
