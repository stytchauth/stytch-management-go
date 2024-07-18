package projects

import "time"

type Vertical string

const (
	VerticalConsumer Vertical = "CONSUMER"
	VerticalB2B      Vertical = "B2B"
)

type CreateRequest struct {
	ProjectName string   `json:"project_name"`
	Vertical    Vertical `json:"vertical"`
}

type ProjectSettings struct {
	ProjectName string    `json:"project_name"`
	CreatedAt   time.Time `json:"created_at"`
}

type Project struct {
	ProjectID                           string          `json:"project_id"`
	TestProjectID                       string          `json:"test_project_id"`
	ProjectSettings                     ProjectSettings `json:"project_settings"`
	Domain                              string          `json:"domain"`
	OAuthCallbackID                     string          `json:"oauth_callback_id"`
	TestOAuthCallbackID                 string          `json:"test_oauth_callback_id"`
	EnableCustomLogo                    bool            `json:"enable_custom_logo"`
	DisableEmailWatermark               bool            `json:"disable_email_watermark"`
	Vertical                            Vertical        `json:"vertical"`
	ZeroDownTimeSessionMigrationURL     string          `json:"zero_downtime_session_migration_url"`
	TestZeroDownTimeSessionMigrationURL string          `json:"test_zero_downtime_session_migration_url"`
}

type CreateResponse struct {
	StatusCode      int     `json:"status_code"`
	RequestID       string  `json:"request_id"`
	ProjectUserID   string  `json:"project_user_id"`
	Project         Project `json:"project"`
	AccountVerified bool    `json:"account_verified"`
	RedirectEnabled bool    `json:"redirect_enabled"`
}

type GetRequest struct {
	ProjectID string `json:"project_id"`
}

type GetResponse struct {
	StatusCode            int     `json:"status_code"`
	RequestID             string  `json:"request_id"`
	Project               Project `json:"project"`
	DisableEmailWatermark bool    `json:"disable_email_watermark"`
	DisableSDKWatermark   bool    `json:"disable_sdk_watermark"`
	AccountVerified       bool    `json:"account_verified"`
	RedirectEnabled       bool    `json:"redirect_enabled"`
	EnableCustomLogo      bool    `json:"enable_custom_logo"`
}

type PasswordValidationPolicy string

const (
	PasswordValidationPolicyZxcvbn PasswordValidationPolicy = "zxcvbn"
	PasswordValidationPolicyLuds   PasswordValidationPolicy = "luds"
)

type PasswordStrengthConfig struct {
	CheckBreachOnCreate         bool                     `json:"check_breach_on_create"`
	CheckBreachOnAuthentication bool                     `json:"check_breach_on_authentication"`
	ValidateOnAuthentication    bool                     `json:"validate_on_authentication"`
	ValidationPolicy            PasswordValidationPolicy `json:"validation_policy"`
	LudsMinCount                *int                     `json:"luds_min_count"`
	LudsComplexity              *int                     `json:"luds_complexity"`
}

type GetPasswordStrengthPolicyRequest struct {
	ProjectID string `json:"project_id"`
}

type GetPasswordStrengthPolicyResponse struct {
	StatusCode     int                    `json:"status_code"`
	RequestID      string                 `json:"request_id"`
	ProjectID      string                 `json:"project_id"`
	PasswordConfig PasswordStrengthConfig `json:"password_config"`
}

type SetPasswordStrengthPolicyRequest struct {
	ProjectID      string                 `json:"project_id"`
	PasswordConfig PasswordStrengthConfig `json:"password_config"`
}

type SetPasswordStrengthPolicyResponse struct {
	StatusCode     int                    `json:"status_code"`
	RequestID      string                 `json:"request_id"`
	PasswordConfig PasswordStrengthConfig `json:"password_config"`
}
