package passwordstrengthconfig

type ValidationPolicy string

const (
	ValidationPolicyLUDS   ValidationPolicy = "LUDS"
	ValidationPolicyZXCVBN ValidationPolicy = "ZXCVBN"
)

func ValidationPolicies() []ValidationPolicy {
	return []ValidationPolicy{
		ValidationPolicyLUDS,
		ValidationPolicyZXCVBN,
	}
}

// PasswordStrengthConfig is the configuration for password strength requirements used in password-based authentication
type PasswordStrengthConfig struct {
	// CheckBreachOnCreation is a flag to check whether to use the HaveIBeenPwned database to detect password breaches
	// when a user first creates their password.
	CheckBreachOnCreation bool `json:"check_breach_on_creation"`
	// CheckBreachOnAuthentication denotes whether to use the HaveIBeenPwned database to detect password breaches when
	// a user authenticates.
	CheckBreachOnAuthentication bool `json:"check_breach_on_authentication"`
	// ValidateOnAuthentication notes whether to require a password reset on authentication if a user's current password
	// no longer meets the project's current policy requirements.
	ValidateOnAuthentication bool `json:"validate_on_authentication"`
	// ValidationPolicy is the policy to use for password validation
	ValidationPolicy ValidationPolicy `json:"validation_policy"`
	// LudsMinPasswordLength is the minimum number of characters in a password if using a LUDS validation_policy. This field
	// is ignored when using the ZXCVBN validation_policy. If present, this value must be a number in the range [8, 32].
	LudsMinPasswordLength int `json:"luds_min_password_length"`
	// LudsMinPasswordComplexity is the minimum number of "character types" in a password (Lowercase, Uppercase, Digits,
	// Symbols) when using a LUDS validation_policy. If present, this must be a number in the range [1, 4].
	LudsMinPasswordComplexity int `json:"luds_min_password_complexity"`
}

type GetRequest struct {
	// ProjectID is the unique ID of the project for which to retrieve the password strength config
	ProjectID string `json:"project_id"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// PasswordStrengthConfig is the password strength configuration for the project
	PasswordStrengthConfig PasswordStrengthConfig `json:"password_strength_config"`
}

type SetRequest struct {
	// ProjectID is the unique ID of the project for which to set the password strength config
	ProjectID string `json:"project_id"`
	// PasswordStrengthConfig is the new password strength configuration to be set for the project
	PasswordStrengthConfig PasswordStrengthConfig `json:"password_strength_config"`
}

type SetResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// PasswordStrengthConfig is the new password strength configuration for the project
	PasswordStrengthConfig PasswordStrengthConfig `json:"password_strength_config"`
}
