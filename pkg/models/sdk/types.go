package sdk

type AuthSetting string

const (
	AuthSettingDisabled      AuthSetting = "disabled"
	AuthSettingSecondaryOnly AuthSetting = "secondary_only"
	AuthSettingPrimaryOnly   AuthSetting = "primary_only"
	AuthSettingAlways        AuthSetting = "always"
)

type DFPProtectedAuthSetting string

const (
	DFPProtectedAuthDisabled DFPProtectedAuthSetting = "DISABLED"
	DFPProtectedAuthPassive  DFPProtectedAuthSetting = "PASSIVE"
	DFPProtectedAuthEnabled  DFPProtectedAuthSetting = "ENABLED"
)

type DFPProtectedAuthChallengeSetting string

const (
	DFPProtectedAuthChallengeSettingAllow   DFPProtectedAuthChallengeSetting = "ALLOW"
	DFPProtectedAuthChallengeSettingBlock   DFPProtectedAuthChallengeSetting = "BLOCK"
	DFPProtectedAuthChallengeSettingCaptcha DFPProtectedAuthChallengeSetting = "CAPTCHA"
)

type AuthorizedB2BDomain struct {
	Domain      string `json:"domain,omitempty"`
	SlugPattern string `json:"slug_pattern,omitempty"`
}

type SMSAutofillMetadata struct {
	AppDomain string `json:"app_domain,omitempty"`
	AppHash   string `json:"app_hash,omitempty"`
}

type Config struct {
	ManageUserData                          bool                             `json:"manage_user_data,omitempty"`
	ManageSessionData                       bool                             `json:"manage_session_data,omitempty"`
	EmailMagicLinks                         AuthSetting                      `json:"email_magic_links,omitempty"`
	SMSOTPs                                 AuthSetting                      `json:"sms_otps,omitempty"`
	WhatsappOTPs                            AuthSetting                      `json:"whatsapp_otps,omitempty"`
	EmailOTPs                               AuthSetting                      `json:"email_otps,omitempty"`
	OAuth                                   AuthSetting                      `json:"oauth,omitempty"`
	CreateTOTPEnabled                       bool                             `json:"create_totp_enabled,omitempty"`
	TOTPs                                   AuthSetting                      `json:"totps,omitempty"`
	CreateWebauthnEnabled                   bool                             `json:"create_webauthn_enabled,omitempty"`
	Webauthns                               AuthSetting                      `json:"webauthns,omitempty"`
	CreateNewUsers                          bool                             `json:"create_new_users,omitempty"`
	CryptoWallets                           AuthSetting                      `json:"crypto_wallets,omitempty"`
	MaxSessionDurationMinutes               int                              `json:"max_session_duration_minutes,omitempty"`
	PKCERequiredForEmailMagicLinks          bool                             `json:"pkce_required_for_email_magic_links,omitempty"`
	PKCERequiredForOAuth                    bool                             `json:"pkce_required_for_oauth,omitempty"`
	PKCERequiredForPasswordResets           bool                             `json:"pkce_required_for_password_resets,omitempty"`
	Passwords                               AuthSetting                      `json:"passwords,omitempty"`
	CreateBiometricsEnabled                 bool                             `json:"create_biometrics_enabled,omitempty"`
	Biometrics                              AuthSetting                      `json:"biometrics,omitempty"`
	EmailMagicLinksSend                     AuthSetting                      `json:"email_magic_links_send,omitempty"`
	SMSOTPsSend                             AuthSetting                      `json:"sms_otps_send,omitempty"`
	WhatsappOTPsSend                        AuthSetting                      `json:"whatsapp_otps_send,omitempty"`
	EmailOTPsSend                           AuthSetting                      `json:"email_otps_send,omitempty"`
	EnableGenericB2BLoginForEmailMagicLinks bool                             `json:"enable_generic_b2b_login_for_email_magic_links,omitempty"`
	EnableGenericB2BLoginForEmailOTPs       bool                             `json:"enable_generic_b2b_login_for_email_otps,omitempty"`
	EnableGenericB2BLoginForOAuth           bool                             `json:"enable_generic_b2b_login_for_oauth,omitempty"`
	EnableOrganizationAuthSettings          bool                             `json:"enable_organization_auth_settings,omitempty"`
	AllowSelfOnboarding                     bool                             `json:"allow_self_onboarding,omitempty"`
	SSO                                     AuthSetting                      `json:"sso,omitempty"`
	PKCERequiredForSSO                      bool                             `json:"pkce_required_for_sso,omitempty"`
	DiscoveryEnabled                        bool                             `json:"discovery_enabled,omitempty"`
	EnableGenericB2BLoginForPasswords       bool                             `json:"enable_generic_b2b_login_for_passwords,omitempty"`
	EnableB2BUseMemberPermissions           bool                             `json:"enable_b2b_use_member_permissions,omitempty"`
	EnableSCIMConnections                   bool                             `json:"enable_scim_connections,omitempty"`
	Domains                                 []string                         `json:"domains,omitempty"`
	BundleIDs                               []string                         `json:"bundle_ids,omitempty"`
	B2BDomains                              []AuthorizedB2BDomain            `json:"b2b_domains,omitempty"`
	DFPProtectedAuthEnabled                 DFPProtectedAuthSetting          `json:"dfp_protected_auth_enabled,omitempty"`
	DFPProtectedAuthOnChallenge             DFPProtectedAuthChallengeSetting `json:"dfp_protected_auth_on_challenge,omitempty"`
	DFPProtectedAuthLookupTimeoutSeconds    int                              `json:"dfp_protected_auth_lookup_timeout_seconds,omitempty"`
	SIWERequiredForCryptoWallets            bool                             `json:"siwe_required_for_crypto_wallets,omitempty"`
	SMSAutofillMetadata                     SMSAutofillMetadata              `json:"sms_autofill_metadata,omitempty"`
}

type GetConfigRequest struct {
	ProjectID string `json:"project_id"`
}

type GetConfigResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id"`
	Config     Config `json:"config"`
}

type SetConfigRequest struct {
	ProjectID string `json:"project_id"`
	Config    Config `json:"config"`
}

type SetConfigResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id"`
}
