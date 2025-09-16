package jwttemplates

type JWTTemplateType string

const (
	JWTTemplateTypeSession JWTTemplateType = "SESSION"
	JWTTemplateTypeM2M     JWTTemplateType = "M2M"
)

// JWTTemplateTypes returns a list of all supported JWT template types.
func JWTTemplateTypes() []JWTTemplateType {
	return []JWTTemplateType{
		JWTTemplateTypeSession,
		JWTTemplateTypeM2M,
	}
}

// JWTTemplate represents a JWT template for a project. Templates are used to generate JSON objects
// by mapping custom metadata attributes to a specific format.
type JWTTemplate struct {
	// TemplateContent is the JWT template content.
	TemplateContent string `json:"template_content"`
	// CustomAudience is an optional custom audience for the JWT template.
	CustomAudience string `json:"custom_audience"`
	// JWTTemplateType is the type of JWT template.
	JWTTemplateType JWTTemplateType `json:"jwt_template_type"`
}

type GetRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve the JWT template.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to retrieve the JWT template.
	EnvironmentSlug string `json:"environment"`
	// JWTTemplateType is the type of JWT template to retrieve.
	JWTTemplateType JWTTemplateType `json:"jwt_template_type"`
}

type GetResponse struct {
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// JWTTemplate is the JWT template for the requested JWTTemplateType.
	JWTTemplate JWTTemplate `json:"jwt_template"`
}

type SetRequest struct {
	// ProjectSlug is the slug of the project for which to set the JWT template.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to set the JWT template.
	EnvironmentSlug string `json:"-"`
	// JWTTemplateType is the type of JWT template.
	JWTTemplateType JWTTemplateType `json:"jwt_template_type"`
	// TemplateContent is the JWT template content.
	TemplateContent string `json:"template_content"`
	// CustomAudience is an optional custom audience for the JWT template.
	CustomAudience string `json:"custom_audience"`
}

type SetResponse struct {
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// StatusCode is the HTTP status code for the response.
	StatusCode int `json:"status_code"`
	// JWTTemplate is the JWT template that was set.
	JWTTemplate JWTTemplate `json:"jwt_template"`
}
