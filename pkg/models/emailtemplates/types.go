package emailtemplates

type TemplateType string

const (
	TemplateTypeLogin                 TemplateType = "LOGIN"
	TemplateTypeSignup                TemplateType = "SIGNUP"
	TemplateTypeInvite                TemplateType = "INVITE"
	TemplateTypeResetPassword         TemplateType = "RESET_PASSWORD"
	TemplateTypeOneTimePasscode       TemplateType = "ONE_TIME_PASSCODE"
	TemplateTypeOneTimePasscodeSignup TemplateType = "ONE_TIME_PASSCODE_SIGNUP"
	TemplateTypeAll                   TemplateType = "ALL"
)

type TextAlignment string

const (
	TextAlignmentUnknown TextAlignment = "UNKNOWN_TEXT_ALIGNMENT"
	TextAlignmentLeft    TextAlignment = "LEFT"
	TextAlignmentCenter  TextAlignment = "CENTER"
)

type FontFamily string

const (
	FontFamilyUnknown                  = "UNKNOWN_FONT_FAMILY"
	FontFamilyArial         FontFamily = "ARIAL"
	FontFamilyBrushScriptMT FontFamily = "BRUSH_SCRIPT_MT"
	FontFamilyCourierNew    FontFamily = "COURIER_NEW"
	FontFamilyGeorgia       FontFamily = "GEORGIA"
	FontFamilyHelvetica     FontFamily = "HELVETICA"
	FontFamilyTahoma        FontFamily = "TAHOMA"
	FontFamilyTimesNewRoman FontFamily = "TIMES_NEW_ROMAN"
	FontFamilyTrebuchetMS   FontFamily = "TREBUCHET_MS"
	FontFamilyVerdana       FontFamily = "VERDANA"
)

type SenderInformation struct {
	FromLocalPart    *string `json:"from_local_part,omitempty"`
	FromDomain       *string `json:"from_domain,omitempty"`
	FromName         *string `json:"from_name,omitempty"`
	ReplyToLocalPart *string `json:"reply_to_local_part,omitempty"`
	ReplyToName      *string `json:"reply_to_name,omitempty"`
}

type PrebuiltCustomization struct {
	ButtonBorderRadius *float32       `json:"button_border_radius,omitempty"`
	ButtonColor        *string        `json:"button_color,omitempty"`
	ButtonTextColor    *string        `json:"button_text_color,omitempty"`
	FontFamily         *FontFamily    `json:"font_family,omitempty"`
	LogoSrc            *string        `json:"logo_src,omitempty"`
	TextAlignment      *TextAlignment `json:"text_alignment,omitempty"`
}

type CustomHTMLCustomization struct {
	TemplateType     TemplateType `json:"template_type,omitempty"`
	HTMLContent      *string      `json:"html_content,omitempty"`
	PlaintextContent *string      `json:"plaintext_content,omitempty"`
	Subject          *string      `json:"subject,omitempty"`
}

type EmailTemplate struct {
	TemplateID        string             `json:"template_id,omitempty"`
	Name              *string            `json:"name,omitempty"`
	SenderInformation *SenderInformation `json:"sender_information,omitempty"`

	// NOTE: Only *one of these fields* should be set.
	PrebuiltCustomization   *PrebuiltCustomization   `json:"prebuilt_customization,omitempty"`
	CustomHTMLCustomization *CustomHTMLCustomization `json:"custom_html_customization,omitempty"`
}

type CreateRequest struct {
	ProjectID     string        `json:"project_id,omitempty"`
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type CreateResponse struct {
	StatusCode    int           `json:"status_code"`
	RequestID     string        `json:"request_id,omitempty"`
	ProjectID     string        `json:"project_id,omitempty"`
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type GetRequest struct {
	ProjectID  string `json:"project_id,omitempty"`
	TemplateID string `json:"template_id,omitempty"`
}

type GetResponse struct {
	StatusCode    int           `json:"status_code"`
	RequestID     string        `json:"request_id,omitempty"`
	ProjectID     string        `json:"project_id,omitempty"`
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type GetAllRequest struct {
	ProjectID string `json:"project_id,omitempty"`
}

type GetAllResponse struct {
	StatusCode     int             `json:"status_code"`
	RequestID      string          `json:"request_id,omitempty"`
	ProjectID      string          `json:"project_id,omitempty"`
	EmailTemplates []EmailTemplate `json:"email_templates,omitempty"`
}

type UpdateRequest struct {
	ProjectID     string        `json:"project_id,omitempty"`
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type UpdateResponse struct {
	StatusCode    int           `json:"status_code"`
	RequestID     string        `json:"request_id,omitempty"`
	ProjectID     string        `json:"project_id,omitempty"`
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type DeleteRequest struct {
	ProjectID  string `json:"project_id,omitempty"`
	TemplateID string `json:"template_id,omitempty"`
}

type DeleteResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id,omitempty"`
	ProjectID  string `json:"project_id,omitempty"`
	TemplateID string `json:"template_id,omitempty"`
}
