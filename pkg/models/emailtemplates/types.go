package emailtemplates

type EmailTemplate struct {
	ID                 string `json:"id"`
	ProjectID          string `json:"project_id"`
	Type               string `json:"type"`
	IsDefault          bool   `json:"is_default"`
	Method             string `json:"method"`
	Name               string `json:"name"`
	VanityID           string `json:"vanity_id"`
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

type EmailTemplateDefault struct {
	Type                      string `json:"type"`
	Locale                    string `json:"locale"`
	HTMLContent               string `json:"html_content"`
	PlaintextContent          string `json:"plaintext_content"`
	Subject                   string `json:"subject"`
	SecondarySubject          string `json:"secondary_subject"`
	DiscoveryHTMLContent      string `json:"discovery_html_content"`
	DiscoveryPlaintextContent string `json:"discovery_plaintext_content"`
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
	LiveProjectID string `json:"project_id"`
	Name          string `json:"name"`
	VanityID      string `json:"vanity_id"`
	Method        string `json:"method"`
	Type          string `json:"type"`
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
	LiveProjectID   string `json:"project_id"`
	EmailTemplateID string `json:"email_template_id"`
}

type DeleteEmailTemplateResponse struct {
	StatusCode          int    `json:"status_code"`
	RequestID           string `json:"request_id"`
	EmailTemplateID     string `json:"email_template_id"`
	TestEmailTemplateID string `json:"test_email_template_id"`
}
