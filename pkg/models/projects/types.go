package projects

import "time"

type Vertical string

const (
	VerticalConsumer Vertical = "CONSUMER"
	VerticalB2B      Vertical = "B2B"
)

type Project struct {
	ID              string    `json:"id"`
	Name            string    `json:"name"`
	OAuthCallbackID string    `json:"oauth_callback_id"`
	Domain          string    `json:"domain"`
	Vertical        Vertical  `json:"vertical"`
	CreatedAt       time.Time `json:"created_at"`
}

type LiveAndTestProject struct {
	LiveProject Project `json:"live_project"`
	TestProject Project `json:"test_project"`
}

type CreateRequest struct {
	ProjectName string   `json:"project_name"`
	Vertical    Vertical `json:"vertical"`
}

type CreateResponse struct {
	StatusCode int                `json:"status_code"`
	RequestID  string             `json:"request_id"`
	Projects   LiveAndTestProject `json:"projects"`
}

type GetRequest struct {
	ProjectID string `json:"project_id"`
}

type GetResponse struct {
	StatusCode int     `json:"status_code"`
	RequestID  string  `json:"request_id"`
	Project    Project `json:"project"`
}

type GetAllRequest struct{}

type GetAllResponse struct {
	StatusCode int                  `json:"status_code"`
	RequestID  string               `json:"request_id"`
	Projects   []LiveAndTestProject `json:"projects"`
}

type DeleteRequest struct {
	ProjectID string `json:"project_id"`
}

type DeleteResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id"`
	ProjectID  string `json:"project_id"`
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
