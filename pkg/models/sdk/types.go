package sdk

type DFPPASetting string

const (
	DFPPASettingDisabled DFPPASetting = "DISABLED"
	DFPPASettingPassive  DFPPASetting = "PASSIVE"
	DFPPASettingEnabled  DFPPASetting = "ENABLED"
)

type DFPPAOnChallengeAction string

const (
	DFPPAOnChallengeActionAllow          DFPPAOnChallengeAction = "ALLOW"
	DFPPAOnChallengeActionBlock          DFPPAOnChallengeAction = "BLOCK"
	DFPPAOnChallengeActionTriggerCaptcha DFPPAOnChallengeAction = "TRIGGER_CAPTCHA"
)

type AuthorizedB2BDomain struct {
	Domain      string `json:"domain,omitempty"`
	SlugPattern string `json:"slug_pattern,omitempty"`
}

type SMSAutofillMetadata struct {
	MetadataType  string `json:"metadata_type"`
	MetadataValue string `json:"metadata_value"`
	BundleID      string `json:"bundle_id"`
	ID            string `json:"id"`
}

type Consumer_BasicConfig struct {
	Enabled        bool     `json:"enabled"`
	CreateNewUsers bool     `json:"create_new_users"`
	Domains        []string `json:"domains"`
	BundleIDs      []string `json:"bundle_ids"`
}

type Consumer_SessionsConfig struct {
	Enabled                   bool  `json:"enabled"`
	MaxSessionDurationMinutes int32 `json:"max_session_duration_minutes"`
}

type Consumer_MagicLinksConfig struct {
	LoginOrCreateEnabled bool `json:"login_or_create_enabled"`
	SendEnabled          bool `json:"send_enabled"`
	PKCERequired         bool `json:"pkce_required"`
}

type Consumer_OTPsConfig struct {
	SMSLoginOrCreateEnabled      bool `json:"sms_login_or_create_enabled"`
	WhatsAppLoginOrCreateEnabled bool `json:"whatsapp_login_or_create_enabled"`
	EmailLoginOrCreateEnabled    bool `json:"email_login_or_create_enabled"`
	SMSSendEnabled               bool `json:"sms_send_enabled"`
	WhatsAppSendEnabled          bool `json:"whatsapp_send_enabled"`
	EmailSendEnabled             bool `json:"email_send_enabled"`

	SMSAutofillMetadata []SMSAutofillMetadata `json:"sms_autofill_metadata"`
}

type Consumer_OAuthConfig struct {
	Enabled      bool `json:"enabled"`
	PKCERequired bool `json:"pkce_required"`
}

type Consumer_TOTPsConfig struct {
	CreateTOTPs bool `json:"create_totps"`
	Enabled     bool `json:"enabled"`
}

type Consumer_WebAuthnConfig struct {
	CreateWebAuthns bool `json:"create_webauthns"`
	Enabled         bool `json:"enabled"`
}

type Consumer_CryptoWalletsConfig struct {
	Enabled      bool `json:"enabled"`
	SIWERequired bool `json:"siwe_required"`
}

type Consumer_DFPPAConfig struct {
	Enabled              DFPPASetting           `json:"enabled"`
	OnChallenge          DFPPAOnChallengeAction `json:"on_challenge"`
	LookupTimeoutSeconds int32                  `json:"lookup_timeout_seconds"`
}

type Consumer_BiometricsConfig struct {
	CreateBiometricsEnabled bool `json:"create_biometrics_enabled"`
	Enabled                 bool `json:"enabled"`
}

type Consumer_PasswordsConfig struct {
	Enabled                       bool `json:"enabled"`
	PKCERequiredForPasswordResets bool `json:"pkce_required_for_password_resets"`
}

type ConsumerConfig struct {
	Basic         *Consumer_BasicConfig         `json:"basic,omitempty"`
	Sessions      *Consumer_SessionsConfig      `json:"sessions,omitempty"`
	MagicLinks    *Consumer_MagicLinksConfig    `json:"magic_links,omitempty"`
	OTPs          *Consumer_OTPsConfig          `json:"otps,omitempty"`
	OAuth         *Consumer_OAuthConfig         `json:"oauth,omitempty"`
	TOTPs         *Consumer_TOTPsConfig         `json:"totps,omitempty"`
	WebAuthn      *Consumer_WebAuthnConfig      `json:"webauthn,omitempty"`
	CryptoWallets *Consumer_CryptoWalletsConfig `json:"crypto_wallets,omitempty"`
	DFPPA         *Consumer_DFPPAConfig         `json:"dfppa,omitempty"`
	Biometrics    *Consumer_BiometricsConfig    `json:"biometrics,omitempty"`
	Passwords     *Consumer_PasswordsConfig     `json:"passwords,omitempty"`
}

type B2B_BasicConfig struct {
	Enabled                 bool                  `json:"enabled"`
	CreateNewMembers        bool                  `json:"create_new_members"`
	AllowSelfOnboarding     bool                  `json:"allow_self_onboarding"`
	EnableMemberPermissions bool                  `json:"enable_member_permissions"`
	Domains                 []AuthorizedB2BDomain `json:"domains"`
	BundleIDs               []string              `json:"bundle_ids"`
}

type B2B_SessionsConfig struct {
	Enabled                   bool  `json:"enabled"`
	MaxSessionDurationMinutes int32 `json:"max_session_duration_minutes"`
}

type B2B_MagicLinksConfig struct {
	Enabled      bool `json:"enabled"`
	PKCERequired bool `json:"pkce_required"`
}

type B2B_OAuthConfig struct {
	Enabled      bool `json:"enabled"`
	PKCERequired bool `json:"pkce_required"`
}

type B2B_TOTPsConfig struct {
	CreateTOTPs bool `json:"create_totps"`
	Enabled     bool `json:"enabled"`
}

type B2B_SSOConfig struct {
	Enabled      bool `json:"enabled"`
	PKCERequired bool `json:"pkce_required"`
}

type B2B_OTPsConfig struct {
	SMSEnabled          bool                  `json:"sms_enabled"`
	SMSAutofillMetadata []SMSAutofillMetadata `json:"sms_autofill_metadata"`
}

type B2B_DFPPAConfig struct {
	Enabled              DFPPASetting           `json:"enabled"`
	OnChallenge          DFPPAOnChallengeAction `json:"on_challenge"`
	LookupTimeoutSeconds int32                  `json:"lookup_timeout_seconds"`
}

type B2B_PasswordsConfig struct {
	Enabled                       bool `json:"enabled"`
	PKCERequiredForPasswordResets bool `json:"pkce_required_for_password_resets"`
}

type B2BConfig struct {
	Basic      *B2B_BasicConfig      `json:"basic,omitempty"`
	Sessions   *B2B_SessionsConfig   `json:"sessions,omitempty"`
	MagicLinks *B2B_MagicLinksConfig `json:"magic_links,omitempty"`
	OAuth      *B2B_OAuthConfig      `json:"oauth,omitempty"`
	TOTPs      *B2B_TOTPsConfig      `json:"totps,omitempty"`
	SSO        *B2B_SSOConfig        `json:"sso,omitempty"`
	OTPs       *B2B_OTPsConfig       `json:"otps,omitempty"`
	DFPPA      *B2B_DFPPAConfig      `json:"dfppa,omitempty"`
	Passwords  *B2B_PasswordsConfig  `json:"passwords,omitempty"`
}

type GetConsumerConfigRequest struct {
	ProjectID string `json:"project_id"`
}

type GetConsumerConfigResponse struct {
	StatusCode int            `json:"status_code"`
	RequestID  string         `json:"request_id"`
	Config     ConsumerConfig `json:"config"`
}

type SetConsumerConfigRequest struct {
	ProjectID string         `json:"project_id"`
	Config    ConsumerConfig `json:"config"`
}

type SetConsumerConfigResponse struct {
	StatusCode int            `json:"status_code"`
	RequestID  string         `json:"request_id"`
	Config     ConsumerConfig `json:"config"`
}

type GetB2BConfigRequest struct {
	ProjectID string `json:"project_id"`
}

type GetB2BConfigResponse struct {
	StatusCode int       `json:"status_code"`
	RequestID  string    `json:"request_id"`
	Config     B2BConfig `json:"config"`
}

type SetB2BConfigRequest struct {
	ProjectID string    `json:"project_id"`
	Config    B2BConfig `json:"config"`
}

type SetB2BConfigResponse struct {
	StatusCode int       `json:"status_code"`
	RequestID  string    `json:"request_id"`
	Config     B2BConfig `json:"config"`
}
