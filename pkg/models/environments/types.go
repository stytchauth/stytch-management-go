package environments

import (
	"time"
)

type EnvironmentType string

const (
	EnvironmentTypeLive EnvironmentType = "LIVE"
	EnvironmentTypeTest EnvironmentType = "TEST"
)

func EnvironmentTypes() []EnvironmentType {
	return []EnvironmentType{
		EnvironmentTypeLive,
		EnvironmentTypeTest,
	}
}

// Environment represents an Environment within a Stytch Project.
type Environment struct {
	// EnvironmentSlug is the immutable unique identifier (slug) of the environment.
	EnvironmentSlug string `json:"environment"`
	// ProjectSlug is the immutable unique identifier (slug) of the project.
	ProjectSlug string `json:"project"`
	// Name is the environment's name.
	Name string `json:"name"`
	// Type is the environment's type. See EnvironmentTypes() for possible values.
	Type EnvironmentType `json:"type"`
	// CreatedAt is the ISO-8601 timestamp for when the environment was created.
	CreatedAt time.Time `json:"created_at"`

	// Configuration fields for the Environment are below.

	// OAuthCallbackID is the callback ID used in OAuth requests for the environment.
	OauthCallbackID string `json:"oauth_callback_id"`
	// CrossOrgPasswordsEnabled indicates whether the environment uses cross-org passwords.
	CrossOrgPasswordsEnabled bool `json:"cross_org_passwords_enabled"`
	// UserImpersonationEnabled indicates whether user impersonation is enabled for the environment.
	UserImpersonationEnabled bool `json:"user_impersonation_enabled"`
	// ZeroDowntimeSessionMigrationURL is the OIDC-compliant UserInfo endpoint for session migration.
	ZeroDowntimeSessionMigrationURL string `json:"zero_downtime_session_migration_url"`

	// User locking fields.
	// UserLockSelfServeEnabled indicates whether users in the environment who get locked out should
	// automatically get an unlock email magic link.
	UserLockSelfServeEnabled bool `json:"user_lock_self_serve_enabled"`
	// UserLockThreshold represents the number of failed authenticate attempts that will cause a user
	// in the environment to be locked. Defaults to 10.
	UserLockThreshold int32 `json:"user_lock_threshold"`
	// UserLockTTL represents the time in seconds that the user in the environment remains locked once
	// the lock is set. Defaults to 1 hour (3600 seconds).
	UserLockTTL int32 `json:"user_lock_ttl"`

	// IDP fields.
	// IDPAuthorizationURL is the OpenID Configuration endpoint for Connected Apps for the
	// environment.
	IDPAuthorizationURL string `json:"idp_authorization_url"`
	// IDPDynamicClientRegistrationEnabled indicates whether the project has opted in to Dynamic
	// Client Registration (DCR) for Connected Apps.
	IDPDynamicClientRegistrationEnabled bool `json:"idp_dynamic_client_registration_enabled"`
	// IDPDynamicClientRegistrationAccessTokenTemplateContent is the access token template to use for
	// clients created through Dynamic Client Registration (DCR).
	IDPDynamicClientRegistrationAccessTokenTemplateContent string `json:"idp_dynamic_client_registration_access_token_template_content"`
}

type CreateRequest struct {
	// ProjectSlug is the slug of the project for which to create the environment.
	ProjectSlug string `json:"-"`
	// Name is the name of the environment.
	Name string `json:"name"`
	// Type is the environment's type.
	Type EnvironmentType `json:"type"`
	// CrossOrgPasswordsEnabled indicates whether the environment should use cross-org passwords.
	CrossOrgPasswordsEnabled *bool `json:"cross_org_passwords_enabled,omitempty"`
	// UserImpersonationEnabled indicates whether user impersonation should be enabled for the
	// environment.
	UserImpersonationEnabled *bool `json:"user_impersonation_enabled,omitempty"`
	// ZeroDowntimeSessionMigrationURL is the OIDC-compliant UserInfo endpoint for session migration.
	ZeroDowntimeSessionMigrationURL *string `json:"zero_downtime_session_migration_url,omitempty"`

	// User locking fields.
	// UserLockSelfServeEnabled indicates whether users in the environment who get locked out should
	// automatically get an unlock email magic link.
	UserLockSelfServeEnabled *bool `json:"user_lock_self_serve_enabled,omitempty"`
	// UserLockThreshold represents the number of failed authenticate attempts that will cause a user
	// in the environment to be locked. Defaults to 10.
	UserLockThreshold *int32 `json:"user_lock_threshold,omitempty"`
	// UserLockTTL represents the time in seconds that the user in the environment remains locked once
	// the lock is set. Defaults to 1 hour (3600 seconds).
	UserLockTTL *int32 `json:"user_lock_ttl,omitempty"`

	// IDP fields.
	// IDPAuthorizationURL is the OpenID Configuration endpoint for Connected Apps for the
	// environment.
	IDPAuthorizationURL *string `json:"idp_authorization_url,omitempty"`
	// IDPDynamicClientRegistrationEnabled indicates whether the project has opted in to Dynamic
	// Client Registration (DCR) for Connected Apps.
	IDPDynamicClientRegistrationEnabled *bool `json:"idp_dynamic_client_registration_enabled,omitempty"`
	// IDPDynamicClientRegistrationAccessTokenTemplateContent is the access token template to use for
	// clients created through Dynamic Client Registration (DCR).
	IDPDynamicClientRegistrationAccessTokenTemplateContent *string `json:"idp_dynamic_client_registration_access_token_template_content,omitempty"`
}

type CreateResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// Environment contains the details for the newly created environment.
	Environment Environment `json:"environment"`
}

type GetRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve the environment.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment to retrieve.
	EnvironmentSlug string `json:"-"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// Environment contains the details for the requested environment.
	Environment Environment `json:"environment"`
}

type GetAllRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve environments.
	ProjectSlug string `json:"-"`
}

type GetAllResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// Environments is the list of environments that belong to the specified project.
	Environments []Environment `json:"environments"`
}

type UpdateRequest struct {
	// ProjectSlug is the slug of the project for which to update the environment.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment to update.
	EnvironmentSlug string `json:"-"`
	// Name is the name of the environment.
	Name *string `json:"name"`
	// CrossOrgPasswordsEnabled indicates whether the environment should use cross-org passwords.
	CrossOrgPasswordsEnabled *bool `json:"cross_org_passwords_enabled,omitempty"`
	// UserImpersonationEnabled indicates whether user impersonation should be enabled for the
	// environment.
	UserImpersonationEnabled *bool `json:"user_impersonation_enabled,omitempty"`
	// ZeroDowntimeSessionMigrationURL is the OIDC-compliant UserInfo endpoint for session migration.
	ZeroDowntimeSessionMigrationURL *string `json:"zero_downtime_session_migration_url,omitempty"`

	// User locking fields.
	// UserLockSelfServeEnabled indicates whether users in the environment who get locked out should
	// automatically get an unlock email magic link.
	UserLockSelfServeEnabled *bool `json:"user_lock_self_serve_enabled,omitempty"`
	// UserLockThreshold represents the number of failed authenticate attempts that will cause a user
	// in the environment to be locked. Defaults to 10.
	UserLockThreshold *int32 `json:"user_lock_threshold,omitempty"`
	// UserLockTTL represents the time in seconds that the user in the environment remains locked once
	// the lock is set. Defaults to 1 hour (3600 seconds).
	UserLockTTL *int32 `json:"user_lock_ttl,omitempty"`

	// IDP fields.
	// IDPAuthorizationURL is the OpenID Configuration endpoint for Connected Apps for the
	// environment.
	IDPAuthorizationURL *string `json:"idp_authorization_url,omitempty"`
	// IDPDynamicClientRegistrationEnabled indicates whether the project has opted in to Dynamic
	// Client Registration (DCR) for Connected Apps.
	IDPDynamicClientRegistrationEnabled *bool `json:"idp_dynamic_client_registration_enabled,omitempty"`
	// IDPDynamicClientRegistrationAccessTokenTemplateContent is the access token template to use for
	// clients created through Dynamic Client Registration (DCR).
	IDPDynamicClientRegistrationAccessTokenTemplateContent *string `json:"idp_dynamic_client_registration_access_token_template_content,omitempty"`
}

type UpdateResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// Environment contains the details for the updated environment.
	Environment Environment `json:"environment"`
}

type DeleteRequest struct {
	// ProjectSlug is the slug of the project for which to delete the environment.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment to delete.
	EnvironmentSlug string `json:"-"`
}

type DeleteResponse struct {
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
}
