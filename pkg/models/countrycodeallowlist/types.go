package countrycodeallowlist

type DeliveryMethod string

const (
	DeliveryMethodSMS      DeliveryMethod = "sms"
	DeliveryMethodWhatsApp DeliveryMethod = "whatsapp"
)

// DeliveryMethods returns a list of all supported delivery methods.
func DeliveryMethods() []DeliveryMethod {
	return []DeliveryMethod{
		DeliveryMethodSMS,
		DeliveryMethodWhatsApp,
	}
}

var DefaultCountryCodes = []string{"CA", "US"}

type GetAllowedSMSCountryCodesRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve allowed SMS country codes.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to retrieve allowed SMS country codes.
	EnvironmentSlug string `json:"-"`
}

type GetAllowedSMSCountryCodesResponse struct {
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// CountryCodes is a list of country codes that are allowed for SMS.
	CountryCodes []string `json:"country_codes"`
}

type GetAllowedWhatsAppCountryCodesRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve allowed WhatsApp country codes.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to retrieve allowed WhatsApp country
	// codes.
	EnvironmentSlug string `json:"-"`
}

type GetAllowedWhatsAppCountryCodesResponse struct {
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// CountryCodes is a list of country codes that are allowed for WhatsApp.
	CountryCodes []string `json:"country_codes"`
}

type SetAllowedSMSCountryCodesRequest struct {
	// ProjectSlug is the slug of the project for which to set allowed SMS country codes.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to set allowed SMS country codes.
	EnvironmentSlug string `json:"-"`
	// CountryCodes is a list of country codes to set as allowed for SMS.
	CountryCodes []string `json:"country_codes"`
}

type SetAllowedSMSCountryCodesResponse struct {
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// CountryCodes is a list of country codes that are allowed for SMS.
	CountryCodes []string `json:"country_codes"`
}

type SetAllowedWhatsAppCountryCodesRequest struct {
	// ProjectSlug is the slug of the project for which to set allowed WhatsApp country codes.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to set allowed WhatsApp country codes.
	EnvironmentSlug string `json:"-"`
	// CountryCodes is a list of country codes to set as allowed for WhatsApp.
	CountryCodes []string `json:"country_codes"`
}

type SetAllowedWhatsAppCountryCodesResponse struct {
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// CountryCodes is a list of country codes that are allowed for WhatsApp.
	CountryCodes []string `json:"country_codes"`
}
