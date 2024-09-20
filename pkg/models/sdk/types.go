package sdk

// DFPPASetting is a type specifying the mode for the Device Fingerprinting Protected Auth (DFPPA)
type DFPPASetting string

const (
	// DFPPASettingDisabled is the mode where DFPPA is disabled
	DFPPASettingDisabled DFPPASetting = "DISABLED"
	// DFPPASettingPassive is the mode where DFPPA verdicts are made, but no blocks are enforced
	DFPPASettingPassive DFPPASetting = "PASSIVE"
	// DFPPASettingEnabled is the mode where DFPPA verdicts are made and blocks are enforced
	DFPPASettingEnabled DFPPASetting = "ENABLED"
)

// DFPPAOnChallenge is a type specifying the action to take when a DFPPA "challenge" verdict is returned
type DFPPAOnChallengeAction string

const (
	// DFPPAOnChallengeActionAllow is the action to allow the request when a "challenge" verdict is returned
	DFPPAOnChallengeActionAllow DFPPAOnChallengeAction = "ALLOW"
	// DFPPAOnChallengeActionBlock is the action to block the request when a "challenge" verdict is returned
	DFPPAOnChallengeActionBlock DFPPAOnChallengeAction = "BLOCK"
	// DFPPAOnChallengeActionTriggerCaptcha is the action to trigger a CAPTCHA when a "challenge" verdict is returned
	DFPPAOnChallengeActionTriggerCaptcha DFPPAOnChallengeAction = "TRIGGER_CAPTCHA"
)

// AuthorizedB2BDomain is a type specifying the domain and slug pattern authorized for use in the SDK for a B2B project
type AuthorizedB2BDomain struct {
	// Domain is the domain name. Stytch uses the same-origin policy to determine matches.
	Domain string `json:"domain,omitempty"`
	// SlugPattern is the slug pattern which can be used to support authentication flows specific to each organization. An example
	// value here might be 'https://{{slug}}.example.com'. The value **must** include '{{slug}}' as a placeholder for the slug.
	SlugPattern string `json:"slug_pattern,omitempty"`
}

// SMSAutofillMetadata is a type specifying the metadata to use for autofill of SMS OTPs.
type SMSAutofillMetadata struct {
	// MetadataType is the type of metadata to use for autofill. This should be either "domain" or "hash".
	MetadataType string `json:"metadata_type"`
	// MetadataValue is the value of the metadata to use for autofill. This should be the associated domain name (for MetadataType "domain")
	// or application hash (for MetadataType "hash").
	MetadataValue string `json:"metadata_value"`
	// BundleID is the ID of the bundle to use for autofill. This should be the associated bundle ID.
	BundleID string `json:"bundle_id"`
}

type ConsumerBasicConfig struct {
	// Enabled is a boolean indicating whether the consumer project SDK is enabled. This allows the SDK to manage user and session data.
	Enabled bool `json:"enabled"`
	// CreateNewUsers is a boolean indicating whether new users can be created with the SDK.
	CreateNewUsers bool `json:"create_new_users"`
	// Domains is a list of domains authorized for use in the SDK.
	Domains []string `json:"domains"`
	// BundleIDs is a list of bundle IDs authorized for use in the SDK.
	BundleIDs []string `json:"bundle_ids"`
}

type ConsumerSessionsConfig struct {
	// MaxSessionDurationMinutes is the maximum session duration that can be created in minutes.
	MaxSessionDurationMinutes int32 `json:"max_session_duration_minutes"`
}

type ConsumerMagicLinksConfig struct {
	// LoginOrCreateEnabled is a boolean indicating whether login or create with magic links is enabled in the SDK.
	LoginOrCreateEnabled bool `json:"login_or_create_enabled"`
	// SendEnabled is a boolean indicating whether the magic links send endpoint is enabled in the SDK.
	SendEnabled bool `json:"send_enabled"`
	// PKCERequired is a boolean indicating whether PKCE is required for magic links. PKCE increases security by
	// introducing a one-time secret for each auth flow to ensure the user starts and completes each auth flow from
	// the same application on the device. This prevents a malicious app from intercepting a redirect and authenticating
	// with the users token. PKCE is enabled by default for mobile SDKs.
	PKCERequired bool `json:"pkce_required"`
}

type ConsumerOTPsConfig struct {
	// SMSLoginOrCreateEnabled is a boolean indicating whether the SMS OTP login or create endpoint is enabled in the SDK.
	SMSLoginOrCreateEnabled bool `json:"sms_login_or_create_enabled"`
	// WhatsAppLoginOrCreateEnabled is a boolean indicating whether the WhatsApp OTP login or create endpoint is enabled in the SDK.
	WhatsAppLoginOrCreateEnabled bool `json:"whatsapp_login_or_create_enabled"`
	// EmailLoginOrCreateEnabled is a boolean indicating whether the email OTP login or create endpoint is enabled in the SDK.
	EmailLoginOrCreateEnabled bool `json:"email_login_or_create_enabled"`
	// SMSSendEnabled is a boolean indicating whether the SMS OTP send endpoint is enabled in the SDK.
	SMSSendEnabled bool `json:"sms_send_enabled"`
	// WhatsAppSendEnabled is a boolean indicating whether the WhatsApp OTP send endpoint is enabled in the SDK.
	WhatsAppSendEnabled bool `json:"whatsapp_send_enabled"`
	// EmailSendEnabled is a boolean indicating whether the email OTP send endpoint is enabled in the SDK.
	EmailSendEnabled bool `json:"email_send_enabled"`

	// SMSAutofillMetadata is a list of metadata that can be used for autofill of SMS OTPs.
	SMSAutofillMetadata []SMSAutofillMetadata `json:"sms_autofill_metadata"`
}

type ConsumerOAuthConfig struct {
	// Enabled is a boolean indicating whether OAuth endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// PKCERequired is a boolean indicating whether PKCE is required for OAuth. PKCE increases security by
	// introducing a one-time secret for each auth flow to ensure the user starts and completes each auth flow from
	// the same application on the device. This prevents a malicious app from intercepting a redirect and authenticating
	// with the users token. PKCE is enabled by default for mobile SDKs.
	PKCERequired bool `json:"pkce_required"`
}

type ConsumerTOTPsConfig struct {
	// Enabled is a boolean indicating whether TOTP endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// CreateTOTPs is a boolean indicating whether TOTP creation is enabled in the SDK.
	CreateTOTPs bool `json:"create_totps"`
}

type ConsumerWebAuthnConfig struct {
	// Enabled is a boolean indicating whether WebAuthn endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// CreateWebAuthns is a boolean indicating whether WebAuthn creation is enabled in the SDK.
	CreateWebAuthns bool `json:"create_webauthns"`
}

type ConsumerCryptoWalletsConfig struct {
	// Enabled is a boolean indicating whether Crypto Wallets endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// SIWERequired is a boolean indicating whether Sign In With Ethereum is required for Crypto Wallets.
	SIWERequired bool `json:"siwe_required"`
}

type ConsumerDFPPAConfig struct {
	// Enabled is a boolean indicating whether Device Fingerprinting Protected Auth is enabled in the SDK.
	Enabled DFPPASetting `json:"enabled"`
	// OnChallenge is the action to take when a DFPPA "challenge" verdict is returned.
	OnChallenge DFPPAOnChallengeAction `json:"on_challenge"`
	// LookupTimeoutSeconds is how long to wait for a DFPPA lookup to complete before timing out.
	LookupTimeoutSeconds int32 `json:"lookup_timeout_seconds"`
}

type ConsumerBiometricsConfig struct {
	// Enabled is a boolean indicating whether biometrics endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// CreateBiometricsEnabled is a boolean indicating whether biometrics creation is enabled in the SDK.
	CreateBiometricsEnabled bool `json:"create_biometrics_enabled"`
}

type ConsumerPasswordsConfig struct {
	// Enabled is a boolean indicating whether password endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// PKCERequiredForPasswordResets is a boolean indicating whether PKCE is required for password resets. PKCE increases
	// security by introducing a one-time secret for each auth flow to ensure the user starts and completes each auth flow
	// from the same application on the device. This prevents a malicious app from intercepting a redirect and
	// authenticating with the users token. PKCE is enabled by default for mobile SDKs.
	PKCERequiredForPasswordResets bool `json:"pkce_required_for_password_resets"`
}

type ConsumerConfig struct {
	// Basic is the basic configuration for the consumer project SDK. This includes enabling the SDK.
	Basic *ConsumerBasicConfig `json:"basic,omitempty"`
	// Sessions is the session configuration for the consumer project SDK.
	Sessions *ConsumerSessionsConfig `json:"sessions,omitempty"`
	// MagicLinks is the magic links configuration for the consumer project SDK.
	MagicLinks *ConsumerMagicLinksConfig `json:"magic_links,omitempty"`
	// OTPs is the OTPs configuration for the consumer project SDK.
	OTPs *ConsumerOTPsConfig `json:"otps,omitempty"`
	// OAuth is the OAuth configuration for the consumer project SDK.
	OAuth *ConsumerOAuthConfig `json:"oauth,omitempty"`
	// TOTPs is the TOTPs configuration for the consumer project SDK.
	TOTPs *ConsumerTOTPsConfig `json:"totps,omitempty"`
	// WebAuthn is the WebAuthn configuration for the consumer project SDK.
	WebAuthn *ConsumerWebAuthnConfig `json:"webauthn,omitempty"`
	// CryptoWallets is the Crypto Wallets configuration for the consumer project SDK.
	CryptoWallets *ConsumerCryptoWalletsConfig `json:"crypto_wallets,omitempty"`
	// DFPPA is the Device Fingerprinting Protected Auth configuration for the consumer project SDK.
	DFPPA *ConsumerDFPPAConfig `json:"dfppa,omitempty"`
	// Biometrics is the biometrics configuration for the consumer project SDK.
	Biometrics *ConsumerBiometricsConfig `json:"biometrics,omitempty"`
	// Passwords is the passwords configuration for the consumer project SDK.
	Passwords *ConsumerPasswordsConfig `json:"passwords,omitempty"`
}

type B2BBasicConfig struct {
	// Enabled is a boolean indicating whether the B2B project SDK is enabled. This allows the SDK to manage user and session data.
	Enabled bool `json:"enabled"`
	// CreateNewMembers is a boolean indicating whether new members can be created with the SDK.
	CreateNewMembers bool `json:"create_new_members"`
	// AllowSelfOnboarding is a boolean indicating whether self-onboarding is allowed for members in the SDK.
	AllowSelfOnboarding bool `json:"allow_self_onboarding"`
	// EnableMemberPermissions is a boolean indicating whether member permissions RBAC are enabled in the SDK.
	EnableMemberPermissions bool `json:"enable_member_permissions"`
	// Domains is a list of domains authorized for use in the SDK.
	Domains []AuthorizedB2BDomain `json:"domains"`
	// BundleIDs is a list of bundle IDs authorized for use in the SDK.
	BundleIDs []string `json:"bundle_ids"`
}

type B2BSessionsConfig struct {
	// MaxSessionDurationMinutes is the maximum session duration that can be created in minutes.
	MaxSessionDurationMinutes int32 `json:"max_session_duration_minutes"`
}

type B2BMagicLinksConfig struct {
	// Enabled is a boolean indicating whether magic links endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// PKCERequired is a boolean indicating whether PKCE is required for magic links. PKCE increases security by
	// introducing a one-time secret for each auth flow to ensure the user starts and completes each auth flow from
	// the same application on the device. This prevents a malicious app from intercepting a redirect and authenticating
	// with the users token. PKCE is enabled by default for mobile SDKs.
	PKCERequired bool `json:"pkce_required"`
}

type B2BOAuthConfig struct {
	// Enabled is a boolean indicating whether OAuth endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// PKCERequired is a boolean indicating whether PKCE is required for OAuth. PKCE increases security by
	// introducing a one-time secret for each auth flow to ensure the user starts and completes each auth flow from
	// the same application on the device. This prevents a malicious app from intercepting a redirect and authenticating
	// with the users token. PKCE is enabled by default for mobile SDKs.
	PKCERequired bool `json:"pkce_required"`
}

type B2BTOTPsConfig struct {
	// Enabled is a boolean indicating whether TOTP endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// CreateTOTPs is a boolean indicating whether TOTP creation is enabled in the SDK.
	CreateTOTPs bool `json:"create_totps"`
}

type B2BSSOConfig struct {
	// Enabled is a boolean indicating whether SSO endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// PKCERequired is a boolean indicating whether PKCE is required for SSO. PKCE increases security by
	// introducing a one-time secret for each auth flow to ensure the user starts and completes each auth flow from
	// the same application on the device. This prevents a malicious app from intercepting a redirect and authenticating
	// with the users token. PKCE is enabled by default for mobile SDKs.
	PKCERequired bool `json:"pkce_required"`
}

type B2BOTPsConfig struct {
	// SMSEnabled is a boolean indicating whether the SMS OTP endpoints are enabled in the SDK.
	SMSEnabled bool `json:"sms_enabled"`
	// SMSAutofillMetadata is a list of metadata that can be used for autofill of SMS OTPs.
	SMSAutofillMetadata []SMSAutofillMetadata `json:"sms_autofill_metadata"`
}

type B2BDFPPAConfig struct {
	// Enabled is a boolean indicating whether Device Fingerprinting Protected Auth is enabled in the SDK.
	Enabled DFPPASetting `json:"enabled"`
	// OnChallenge is the action to take when a DFPPA "challenge" verdict is returned.
	OnChallenge DFPPAOnChallengeAction `json:"on_challenge"`
	// LookupTimeoutSeconds is how long to wait for a DFPPA lookup to complete before timing out.
	LookupTimeoutSeconds int32 `json:"lookup_timeout_seconds"`
}

type B2BPasswordsConfig struct {
	// Enabled is a boolean indicating whether password endpoints are enabled in the SDK.
	Enabled bool `json:"enabled"`
	// PKCERequiredForPasswordResets is a boolean indicating whether PKCE is required for password resets. PKCE increases
	// security by introducing a one-time secret for each auth flow to ensure the user starts and completes each auth flow
	// from the same application on the device. This prevents a malicious app from intercepting a redirect and
	// authenticating with the users token. PKCE is enabled by default for mobile SDKs.
	PKCERequiredForPasswordResets bool `json:"pkce_required_for_password_resets"`
}

type B2BConfig struct {
	// Basic is the basic configuration for the B2B project SDK. This includes enabling the SDK.
	Basic *B2BBasicConfig `json:"basic,omitempty"`
	// Sessions is the session configuration for the B2B project SDK.
	Sessions *B2BSessionsConfig `json:"sessions,omitempty"`
	// MagicLinks is the magic links configuration for the B2B project SDK.
	MagicLinks *B2BMagicLinksConfig `json:"magic_links,omitempty"`
	// OAuth is the OAuth configuration for the B2B project SDK.
	OAuth *B2BOAuthConfig `json:"oauth,omitempty"`
	// TOTPs is the TOTPs configuration for the B2B project SDK.
	TOTPs *B2BTOTPsConfig `json:"totps,omitempty"`
	// SSO is the SSO configuration for the B2B project SDK.
	SSO *B2BSSOConfig `json:"sso,omitempty"`
	// OTPs is the OTPs configuration for the B2B project SDK.
	OTPs *B2BOTPsConfig `json:"otps,omitempty"`
	// DFPPA is the Device Fingerprinting Protected Auth configuration for the B2B project SDK.
	DFPPA *B2BDFPPAConfig `json:"dfppa,omitempty"`
	// Passwords is the passwords configuration for the B2B project SDK.
	Passwords *B2BPasswordsConfig `json:"passwords,omitempty"`
}

type GetConsumerConfigRequest struct {
	// ProjectID is the ID of the consumer project.
	ProjectID string `json:"project_id"`
}

type GetConsumerConfigResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Config is the consumer project SDK configuration.
	Config ConsumerConfig `json:"config"`
}

type SetConsumerConfigRequest struct {
	// ProjectID is the ID of the consumer project.
	ProjectID string `json:"project_id"`
	// Config is the consumer project SDK configuration to set.
	Config ConsumerConfig `json:"config"`
}

type SetConsumerConfigResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Config is the consumer project SDK configuration that was set.
	Config ConsumerConfig `json:"config"`
}

type GetB2BConfigRequest struct {
	// ProjectID is the ID of the B2B project.
	ProjectID string `json:"project_id"`
}

type GetB2BConfigResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Config is the B2B project SDK configuration.
	Config B2BConfig `json:"config"`
}

type SetB2BConfigRequest struct {
	// ProjectID is the ID of the B2B project.
	ProjectID string `json:"project_id"`
	// Config is the B2B project SDK configuration to set.
	Config B2BConfig `json:"config"`
}

type SetB2BConfigResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Config is the B2B project SDK configuration that was set.
	Config B2BConfig `json:"config"`
}
