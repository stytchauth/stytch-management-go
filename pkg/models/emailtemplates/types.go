package emailtemplates

type Method string

const (
	MethodCustomHTML            Method = "custom_html"
	MethodBuiltInCustomizations Method = "built_in_customizations"
)

type TemplateType string

const (
	TemplateTypeLogin                 TemplateType = "login"
	TemplateTypeSignup                TemplateType = "signup"
	TemplateTypeInvite                TemplateType = "invite"
	TemplateTypeResetPassword         TemplateType = "reset_password"
	TemplateTypeOneTimePasscode       TemplateType = "one_time_passcode"
	TemplateTypeOneTimePasscodeSignup TemplateType = "one_time_passcode_signup"
	TemplateTypeAll                   TemplateType = "all"
)

type EmailTemplate struct {
	ID                 string       `json:"id"`
	ProjectID          string       `json:"project_id"`
	Type               TemplateType `json:"type"`
	IsDefault          bool         `json:"is_default"`
	Method             Method       `json:"method"`
	Name               string       `json:"name"`
	VanityID           string       `json:"vanity_id"`
	Subject            string       `json:"subject"`
	PlaintextContent   string       `json:"plaintext_content"`
	HTMLContent        string       `json:"html_content"`
	ButtonColor        string       `json:"button_color"`
	ButtonTextColor    string       `json:"button_text_color"`
	FontFamily         string       `json:"font_family"`
	TextAlignment      string       `json:"text_alignment"`
	LogoSrc            string       `json:"logo_src"`
	ButtonBorderRadius int          `json:"button_border_radius"`
	FromLocalPart      string       `json:"from_local_part"`
	FromDomain         string       `json:"from_domain"`
	FromName           string       `json:"from_name"`
	ReplyToLocalPart   string       `json:"reply_to_local_part"`
	ReplyToName        string       `json:"reply_to_name"`
	SecondarySubject   string       `json:"secondary_subject"`
}

type EmailTemplateDefault struct {
	Type                      TemplateType `json:"type"`
	Locale                    string       `json:"locale"`
	HTMLContent               string       `json:"html_content"`
	PlaintextContent          string       `json:"plaintext_content"`
	Subject                   string       `json:"subject"`
	SecondarySubject          string       `json:"secondary_subject"`
	DiscoveryHTMLContent      string       `json:"discovery_html_content"`
	DiscoveryPlaintextContent string       `json:"discovery_plaintext_content"`
}

type TestAndLiveEmailTemplate struct {
	TestEmailTemplate EmailTemplate `json:"test_email_template"`
	LiveEmailTemplate EmailTemplate `json:"live_email_template"`
}

type GetAllEmailTemplatesRequest struct {
	LiveProjectID string `json:"project_id"`
}

type GetAllEmailTemplatesResponse struct {
	StatusCode       int                        `json:"status_code"`
	RequestID        string                     `json:"request_id"`
	LiveProjectID    string                     `json:"project_id"`
	EmailTemplates   []TestAndLiveEmailTemplate `json:"email_templates"`
	DefaultTemplates []EmailTemplateDefault     `json:"defaults"`
}

type CreateEmailTemplateRequest struct {
	LiveProjectID string       `json:"project_id"`
	Name          string       `json:"name"`
	VanityID      string       `json:"vanity_id"`
	Method        Method       `json:"method"`
	Type          TemplateType `json:"type"`
}

type CreateEmailTemplateResponse struct {
	StatusCode    int                      `json:"status_code"`
	RequestID     string                   `json:"request_id"`
	EmailTemplate TestAndLiveEmailTemplate `json:"email_template"`
}

type GetEmailTemplateRequest struct {
	ProjectID       string `json:"project_id"`
	EmailTemplateID string `json:"email_template_id"`
}

type GetEmailTemplateResponse struct {
	StatusCode         int    `json:"status_code"`
	RequestID          string `json:"request_id"`
	EmailTemplateID    string `json:"email_template_id"`
	Name               string `json:"name"`
	Subject            string `json:"subject"`
	PlaintextContent   string `json:"plaintext_content"`
	HTMLContent        string `json:"html_content"`
	ButtonColor        string `json:"button_color"`
	ButtonTextColor    string `json:"button_text_color"`
	FontFamily         string `json:"font_family"`
	TextAlignment      string `json:"text_alignment"`
	LogoSrc            string `json:"logo_src"`
	ButtonBorderRadius int    `json:"button_border_radius"`
	FromLocalPart      string `json:"from_local_part"`
	FromDomain         string `json:"from_domain"`
	FromName           string `json:"from_name"`
	ReplyToLocalPart   string `json:"reply_to_local_part"`
	ReplyToName        string `json:"reply_to_name"`
	SecondarySubject   string `json:"secondary_subject"`
}

type DeleteEmailTemplateRequest struct {
	LiveProjectID   string `json:"project_id"`
	EmailTemplateID string `json:"email_template_id"`
}

type DeleteEmailTemplateResponse struct {
	StatusCode          int    `json:"status_code"`
	RequestID           string `json:"request_id"`
	EmailTemplateID     string `json:"email_template_id"`
	TestEmailTemplateID string `json:"test_email_template_id"`
}

type UpdateEmailTemplateRequest struct {
	ProjectID          string `json:"project_id"`
	EmailTemplateID    string `json:"email_template_id"`
	ButtonColor        string `json:"button_color,omitempty"`
	ButtonTextColor    string `json:"button_text_color,omitempty"`
	FontFamily         string `json:"font_family,omitempty"`
	TextAlignment      string `json:"text_alignment,omitempty"`
	LogoSrc            string `json:"logo_src,omitempty"`
	ButtonBorderRadius int    `json:"button_border_radius,omitempty"`
	FromLocalPart      string `json:"from_local_part,omitempty"`
	FromDomain         string `json:"from_domain,omitempty"`
	FromName           string `json:"from_name,omitempty"`
	ReplyToLocalPart   string `json:"reply_to_local_part,omitempty"`
	ReplyToName        string `json:"reply_to_name,omitempty"`
	Name               string `json:"name,omitempty"`
	IsDefault          bool   `json:"is_default,omitempty"`
	HTMLContent        string `json:"html_content,omitempty"`
	PlaintextContent   string `json:"plaintext_content,omitempty"`
	Subject            string `json:"subject,omitempty"`
}

type UpdateEmailTemplateResponse struct {
	StatusCode      int    `json:"status_code"`
	RequestID       string `json:"request_id"`
	EmailTemplateID string `json:"email_template_id"`
}
