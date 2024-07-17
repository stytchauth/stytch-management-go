package emailtemplates

type EmailTemplate struct {
}

type TestAndLiveEmailTemplate struct {
	LiveEmailTemplate EmailTemplate `json:"live_email_template"`
	TestEmailTemplate EmailTemplate `json:"test_email_template"`
}

type GetAllEmailTemplatesRequest struct {
	LiveProjectID string `json:"live_project_id"`
}

type GetEmailTemplatesResponse struct {
	StatusCode       int                        `json:"status_code"`
	RequestID        string                     `json:"request_id"`
	LiveProjectID    string                     `json:"project_id"`
	EmailTemplates   []TestAndLiveEmailTemplate `json:"email_templates"`
	DefaultTemplates []EmailTemplate            `json:"defaults"`
}

type CreateEmailTemplateRequest struct {
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
	VanityID  string `json:"vanity_id"`
	Method    string `json:"method"`
	Type      string `json:"type"`
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
	ButtonBorderRadius string `json:"button_border_radius"`
	FromLocalPart      string `json:"from_local_part"`
	FromDomain         string `json:"from_domain"`
	FromName           string `json:"from_name"`
	ReplyToLocalPart   string `json:"reply_to_local_part"`
	ReplyToName        string `json:"reply_to_name"`
	SecondarySubject   string `json:"secondary_subject"`
}

type DeleteEmailTemplateRequest struct {
	// XCXC Not sure if live or any
	LiveProjectID   string `json:"project_id"`
	EmailTemplateID string `json:"email_template_id"`
}
