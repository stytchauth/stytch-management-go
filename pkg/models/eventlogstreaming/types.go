package eventlogstreaming

type DestinationType string

const (
	DestinationTypeDatadog     DestinationType = "DATADOG"
	DestinationTypeGrafanaLoki DestinationType = "GRAFANA_LOKI"
)

// DestinationTypes returns a list of all supported destination types.
func DestinationTypes() []DestinationType {
	return []DestinationType{
		DestinationTypeDatadog,
		DestinationTypeGrafanaLoki,
	}
}

// DestinationConfig represents the possible destination configurations.
type DestinationConfig struct {
	Datadog     *DatadogConfig     `json:"datadog"`
	GrafanaLoki *GrafanaLokiConfig `json:"grafana_loki"`
}

// DestinationConfigMasked represents the possible destination configurations, with relevant
// credentials masked.
type DestinationConfigMasked struct {
	Datadog     *DatadogConfigMasked     `json:"datadog"`
	GrafanaLoki *GrafanaLokiConfigMasked `json:"grafana_loki"`
}

// DatadogSite is the site of the Datadog account to which to send events.
type DatadogSite string

const (
	DatadogSiteUS  DatadogSite = "US"
	DatadogSiteUS3 DatadogSite = "US3"
	DatadogSiteUS5 DatadogSite = "US5"
	DatadogSiteEU  DatadogSite = "EU"
	DatadogSiteAP1 DatadogSite = "AP1"
)

// DatadogSites returns a list of all supported Datadog sites.
func DatadogSites() []DatadogSite {
	return []DatadogSite{
		DatadogSiteUS,
		DatadogSiteUS3,
		DatadogSiteUS5,
		DatadogSiteEU,
		DatadogSiteAP1,
	}
}

// DatadogConfig is the configuration for sending events to a Datadog account.
// All values must be provided in any Create or Update request.
type DatadogConfig struct {
	// Site is one of the supported DatadogSite constants.
	Site DatadogSite `json:"site"`
	// APIKey is the API key for submitting logs to a Datadog account.
	APIKey string `json:"api_key"`
}

type DatadogConfigMasked struct {
	// Site is one of the supported DatadogSite constants.
	Site DatadogSite `json:"site"`
	// APIKeyLastFour is the last four characters of the API key in use.
	APIKeyLastFour string `json:"api_key_last_four"`
}

type StreamingStatus string

const (
	StreamingStatusActive   StreamingStatus = "ACTIVE"
	StreamingStatusDisabled StreamingStatus = "DISABLED"
)

// StreamingStatuses returns a list of all supported streaming statuses.
func StreamingStatuses() []StreamingStatus {
	return []StreamingStatus{
		StreamingStatusActive,
		StreamingStatusDisabled,
	}
}

// EventLogStreamingConfig is the configuration for sending events to a destination.
type EventLogStreamingConfig struct {
	DestinationType   DestinationType   `json:"destination_type"`
	DestinationConfig DestinationConfig `json:"destination_config"`
	StreamingStatus   StreamingStatus   `json:"streaming_status"`
}

// EventLogStreamingConfigMasked is the configuration for sending events to a destination, with
// relevant credentials masked.
type EventLogStreamingConfigMasked struct {
	DestinationType   DestinationType         `json:"destination_type"`
	DestinationConfig DestinationConfigMasked `json:"destination_config"`
	StreamingStatus   StreamingStatus         `json:"streaming_status"`
}

// GrafanaLokiConfig is the configuration for sending events to a Grafana Loki instance.
// All values must be provided in any Create or Update request
type GrafanaLokiConfig struct {
	// Hostname is the hostname of the Grafana Loki instance to which to send events.
	Hostname string `json:"hostname"`
	// Username is the username for authenticating the request to a Grafana Loki instance.
	Username string `json:"username"`
	// Password is the password for authenticating the request to a Grafana Loki instance.
	Password string `json:"password"`
}

type GrafanaLokiConfigMasked struct {
	// Hostname is the hostname of the Grafana Loki instance to which to send events.
	Hostname string `json:"hostname"`
	// Username is the username for authenticating the request to a Grafana Loki instance.
	Username string `json:"username"`
	// PasswordLastFour is the last four characters of the password in use.
	PasswordLastFour string `json:"password_last_four"`
}

type CreateEventLogStreamingRequest struct {
	// ProjectSlug is the slug of the project for which to create the event log streaming config.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to create the event log streaming
	// config.
	EnvironmentSlug string `json:"-"`
	// DestinationType is the type of destination to which to send events.
	DestinationType DestinationType `json:"destination_type"`
	// DestinationConfig is the configuration for the destination to which to send events.
	DestinationConfig DestinationConfig `json:"destination_config"`
}

type CreateEventLogStreamingResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// EventLogStreamingConfig is the configuration that was created.
	EventLogStreamingConfig EventLogStreamingConfig `json:"event_log_streaming_config"`
}

type GetEventLogStreamingRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve the event log streaming config.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to retrieve the event log streaming
	// config.
	EnvironmentSlug string `json:"-"`
	// DestinationType is the destination type of the event log streaming config to retrieve.
	DestinationType DestinationType `json:"destination_type"`
}

type GetEventLogStreamingResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// EventLogStreamingConfig is the configuration retrieved, with masked credentials.
	EventLogStreamingConfig EventLogStreamingConfigMasked `json:"event_log_streaming_config"`
}

type UpdateEventLogStreamingRequest struct {
	// ProjectSlug is the slug of the project for which to update the event log streaming config.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to update the event log streaming
	// config.
	EnvironmentSlug string `json:"-"`
	// DestinationType is the type of destination for which to update the event log streaming config.
	DestinationType DestinationType `json:"destination_type"`
	// DestinationType is the destination type of the event log streaming config to update.
	DestinationConfig DestinationConfig `json:"destination_config"`
}

type UpdateEventLogStreamingResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// EventLogStreamingConfig is the configuration that was updated.
	EventLogStreamingConfig EventLogStreamingConfig `json:"event_log_streaming_config"`
}

type DeleteEventLogStreamingRequest struct {
	// ProjectSlug is the slug of the project for which to delete the event log streaming config.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to delete the event log streaming
	// config.
	EnvironmentSlug string `json:"-"`
	// DestinationType is the destination type of the event log streaming config to delete.
	DestinationType DestinationType `json:"destination_type"`
}

type DeleteEventLogStreamingResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
}

type EnableEventLogStreamingRequest struct {
	// ProjectSlug is the slug of the project for which to enable the event log streaming config.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to enable the event log streaming
	// config.
	EnvironmentSlug string `json:"-"`
	// DestinationType is the destination type of the event log streaming config to enable.
	DestinationType DestinationType `json:"destination_type"`
}

type EnableEventLogStreamingResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
}

type DisableEventLogStreamingRequest struct {
	// ProjectSlug is the slug of the project for which to disable the event log streaming config.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to disable the event log streaming
	// config.
	EnvironmentSlug string `json:"-"`
	// DestinationType is the destination type of the event log streaming config to disable.
	DestinationType DestinationType `json:"destination_type"`
}

type DisableEventLogStreamingResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
}
