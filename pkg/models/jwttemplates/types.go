package jwttemplates

type TemplateType string

const (
	TemplateTypeSession TemplateType = "SESSION"
	TemplateTypeM2M     TemplateType = "M2M"
)

func TemplateTypes() []TemplateType {
	return []TemplateType{
		TemplateTypeSession,
		TemplateTypeM2M,
	}
}

// JWTTemplate represents a JWT template for a project. Templates are used to generate JSON objects by mapping
// custom metadata attributes to a specific format.
type JWTTemplate struct {
	// TemplateContent is the JWT template content
	TemplateContent string `json:"template_content"`
	// CustomAudience is an optional custom audience for the JWT template
	CustomAudience string `json:"custom_audience"`
	// TemplateType is the type of JWT template
	TemplateType TemplateType `json:"template_type"`
}

type GetRequest struct {
	// ProjectID is the unique ID of the project for which to retrieve the JWT template
	ProjectID string `json:"project_id"`
	// TemplateType is the type of JWT template to retrieve
	TemplateType TemplateType `json:"template_type"`
}

type GetResponse struct {
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// JWTTemplate is the JWT template for the requested TemplateType
	JWTTemplate JWTTemplate `json:"jwt_template"`
}

type SetRequest struct {
	// ProjectID is the unique ID of the project for which to set the JWT template
	ProjectID string `json:"project_id"`
	// JWTTemplate is the JWT template to set
	JWTTemplate JWTTemplate `json:"jwt_template"`
}

type SetResponse struct {
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// JWTTemplate is the JWT template that was set
	JWTTemplate JWTTemplate `json:"jwt_template"`
}
