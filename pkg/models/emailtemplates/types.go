package emailtemplates

type TemplateType string

const (
	TemplateTypeLogin                    TemplateType = "LOGIN"
	TemplateTypeSignup                   TemplateType = "SIGNUP"
	TemplateTypeInvite                   TemplateType = "INVITE"
	TemplateTypeResetPassword            TemplateType = "RESET_PASSWORD"
	TemplateTypeOneTimePasscode          TemplateType = "ONE_TIME_PASSCODE"
	TemplateTypeOneTimePasscodeSignup    TemplateType = "ONE_TIME_PASSCODE_SIGNUP"
	TemplateTypeVerifyEmailPasswordReset TemplateType = "VERIFY_EMAIL_PASSWORD_RESET"
	TemplateTypeAll                      TemplateType = "ALL"
)

func TemplateTypes() []TemplateType {
	return []TemplateType{
		TemplateTypeLogin,
		TemplateTypeSignup,
		TemplateTypeInvite,
		TemplateTypeResetPassword,
		TemplateTypeOneTimePasscode,
		TemplateTypeOneTimePasscodeSignup,
		TemplateTypeAll,
	}
}

type TextAlignment string

const (
	TextAlignmentUnknown TextAlignment = "UNKNOWN_TEXT_ALIGNMENT"
	TextAlignmentLeft    TextAlignment = "LEFT"
	TextAlignmentCenter  TextAlignment = "CENTER"
)

func TextAlignments() []TextAlignment {
	return []TextAlignment{
		TextAlignmentUnknown,
		TextAlignmentLeft,
		TextAlignmentCenter,
	}
}

type FontFamily string

const (
	FontFamilyUnknown       FontFamily = "UNKNOWN_FONT_FAMILY"
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

func FontFamilies() []FontFamily {
	return []FontFamily{
		FontFamilyUnknown,
		FontFamilyArial,
		FontFamilyBrushScriptMT,
		FontFamilyCourierNew,
		FontFamilyGeorgia,
		FontFamilyHelvetica,
		FontFamilyTahoma,
		FontFamilyTimesNewRoman,
		FontFamilyTrebuchetMS,
		FontFamilyVerdana,
	}
}

// SenderInformation holds information about the sender of an email, such as the
// name, domain, and local part of the email address.
type SenderInformation struct {
	// FromLocalPart is the prefix of the sender’s email address, everything before the @ symbol (eg: first.last)
	FromLocalPart *string `json:"from_local_part,omitempty"`
	// FromDomain is the postfix of the sender’s email address, everything after the @ symbol (eg: stytch.com)
	FromDomain *string `json:"from_domain,omitempty"`
	// FromName is the sender of the email (eg: Login)
	FromName *string `json:"from_name,omitempty"`
	// ReplyToLocalPart is the prefix of the reply-to email address, everything before the @ symbol (eg: first.last)
	ReplyToLocalPart *string `json:"reply_to_local_part,omitempty"`
	// ReplyToName is the sender of the reply-to email address (eg: Support)
	ReplyToName *string `json:"reply_to_name,omitempty"`
}

// PrebuiltCustomization holds the customization options for prebuilt email templates.
type PrebuiltCustomization struct {
	// ButtonBorderRadius is the radius of the button border in the email body
	ButtonBorderRadius *float32 `json:"button_border_radius,omitempty"`
	// ButtonColor is the color of the button in the email body
	ButtonColor *string `json:"button_color,omitempty"`
	// ButtonTextColor is the color of the text in the button in the email body
	ButtonTextColor *string `json:"button_text_color,omitempty"`
	// FontFamily is the font type to be used in the email body
	FontFamily *FontFamily `json:"font_family,omitempty"`
	// TextAlignment is the alignment of the text in the email body
	TextAlignment *TextAlignment `json:"text_alignment,omitempty"`
}

// CustomHTMLCustomization holds the customization options for custom HTML email templates.
type CustomHTMLCustomization struct {
	// TemplateType is the type of email template this custom HTML customization is valid for
	TemplateType TemplateType `json:"template_type,omitempty"`
	// HTMLContent is the HTML content of the email body
	HTMLContent *string `json:"html_content,omitempty"`
	// PlaintextContent is the plaintext content of the email body
	PlaintextContent *string `json:"plaintext_content,omitempty"`
	// Subject is the subject line in the email template
	Subject *string `json:"subject,omitempty"`
}

// EmailTemplate represents an email template for use in Stytch's email products, such as Magic Links or OTPs.
type EmailTemplate struct {
	// TemplateID is a unique identifier to use for the template – this is how you'll refer to the template when sending
	// emails from your project or managing this template. It can never be changed after creation.
	TemplateID string `json:"template_id,omitempty"`
	// Name is a human-readable name of the template. This does not have to be unique.
	Name *string `json:"name,omitempty"`
	// SenderInformation is information about the email sender, such as the reply address or rendered name.
	// This is an optional field for PrebuiltCustomization, but required for CustomHTMLCustomization.
	SenderInformation *SenderInformation `json:"sender_information,omitempty"`

	// NOTE: Only *one of these fields* should be set.
	// PrebuiltCustomization is customization related to prebuilt fields (such as button color) for prebuilt email templates
	PrebuiltCustomization *PrebuiltCustomization `json:"prebuilt_customization,omitempty"`
	// CustomHTMLCustomization is customization defined for completely custom HTML email templates
	CustomHTMLCustomization *CustomHTMLCustomization `json:"custom_html_customization,omitempty"`
}

type CreateRequest struct {
	// ProjectID is the *live* project ID for which to create the project. This endpoint only works with a live project ID.
	// An email template will also be created for the respective test project.
	ProjectID string `json:"project_id,omitempty"`
	// EmailTemplate is the email template to be created
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type CreateResponse struct {
	//
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id,omitempty"`
	// EmailTemplate is the email template that was created
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type GetRequest struct {
	// ProjectID is the project ID that owns the template to retrieve
	ProjectID string `json:"project_id,omitempty"`
	// TemplateID is the unique template ID for the email template to retrieve
	TemplateID string `json:"template_id,omitempty"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id,omitempty"`
	// EmailTemplate is the email template that was retrieved
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type GetAllRequest struct {
	// ProjectID is the project ID for which to retrieve all email templates
	ProjectID string `json:"project_id,omitempty"`
}

type GetAllResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id,omitempty"`
	// EmailTemplates is a list of all active email templates for the project
	EmailTemplates []EmailTemplate `json:"email_templates,omitempty"`
}

type UpdateRequest struct {
	// ProjectID is the project ID that owns the template to be updated
	ProjectID string `json:"project_id,omitempty"`
	// EmailTemplate contains the updated email template. The template ID must match the template being updated.
	// NOTE: After creation, a Prebuilt template cannot be converted to a custom HTML template, and similarly, a
	// custom HTML template cannot be converted to a Prebuilt template.
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type UpdateResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id,omitempty"`
	// EmailTemplate is the newly updated email template
	EmailTemplate EmailTemplate `json:"email_template,omitempty"`
}

type DeleteRequest struct {
	// ProjectID is the project ID that owns the template to be deleted
	ProjectID string `json:"project_id,omitempty"`
	// TemplateID is the unique template ID for the email template to be deleted
	TemplateID string `json:"template_id,omitempty"`
}

type DeleteResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id,omitempty"`
}
