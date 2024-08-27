package passwordstrengthconfig

type ValidationPolicy string

const (
	ValidationPolicyLUDS   ValidationPolicy = "LUDS"
	ValidationPolicyZXCVBN ValidationPolicy = "ZXCVBN"
)

type PasswordStrengthConfig struct {
	CheckBreachOnCreation       bool             `json:"check_breach_on_creation"`
	CheckBreachOnAuthentication bool             `json:"check_breach_on_authentication"`
	ValidateOnAuthentication    bool             `json:"validate_on_authentication"`
	ValidationPolicy            ValidationPolicy `json:"validation_policy"`
	LudsMinPasswordLength       int              `json:"luds_min_password_length"`
	LudsMinPasswordComplexity   int              `json:"luds_min_password_complexity"`
}

type GetRequest struct {
	ProjectID string `json:"project_id"`
}

type GetResponse struct {
	StatusCode             int                    `json:"status_code"`
	RequestID              string                 `json:"request_id"`
	PasswordStrengthConfig PasswordStrengthConfig `json:"password_strength_config"`
}

type SetRequest struct {
	ProjectID              string                 `json:"project_id"`
	PasswordStrengthConfig PasswordStrengthConfig `json:"password_strength_config"`
}

type SetResponse struct {
	StatusCode             int                    `json:"status_code"`
	RequestID              string                 `json:"request_id"`
	PasswordStrengthConfig PasswordStrengthConfig `json:"password_strength_config"`
}
