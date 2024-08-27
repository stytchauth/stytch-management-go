package jwttemplates

type TemplateType string

const (
	TemplateTypeSession TemplateType = "SESSION"
	TemplateTypeM2M     TemplateType = "M2M"
)

type JWTTemplate struct {
	TemplateContent string       `json:"template_content"`
	CustomAudience  string       `json:"custom_audience"`
	TemplateType    TemplateType `json:"template_type"`
}

type GetRequest struct {
	ProjectID    string       `json:"project_id"`
	TemplateType TemplateType `json:"template_type"`
}

type GetResponse struct {
	RequestID   string      `json:"request_id"`
	StatusCode  int         `json:"status_code"`
	JWTTemplate JWTTemplate `json:"jwt_template"`
}

type SetRequest struct {
	ProjectID   string      `json:"project_id"`
	JWTTemplate JWTTemplate `json:"jwt_template"`
}

type SetResponse struct {
	RequestID   string      `json:"request_id"`
	StatusCode  int         `json:"status_code"`
	JWTTemplate JWTTemplate `json:"jwt_template"`
}
