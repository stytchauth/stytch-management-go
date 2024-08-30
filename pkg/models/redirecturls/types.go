package redirecturls

// RedirectType are the different types of redirect available.
type RedirectType string

const (
	RedirectTypeLogin         RedirectType = "LOGIN"
	RedirectTypeSignup        RedirectType = "SIGNUP"
	RedirectTypeInvite        RedirectType = "INVITE"
	RedirectTypeResetPassword RedirectType = "RESET_PASSWORD"
	// RedirectTypeDiscovery is used for the discovery endpoint exclusively in B2B projects.
	RedirectTypeDiscovery RedirectType = "DISCOVERY"
)

// URLRedirectType holds information for a specific kind of redirect.
type URLRedirectType struct {
	// Type is one of the RedirectType values.
	Type RedirectType `json:"type"`
	// IsDefault is true if this is the default redirect type, false otherwise.
	IsDefault bool `json:"is_default"`
}

// RedirectURL holds information for a specific redirect URL and all its redirect types
type RedirectURL struct {
	// URL is the URL to redirect to.
	URL string `json:"url"`
	// ValidTypes is a list of all the URLRedirectType available for this object
	ValidTypes []URLRedirectType `json:"valid_types"`
}

type CreateRequest struct {
	// ProjectID is the ID of the project to create the redirect URL for
	ProjectID string `json:"project_id"`
	// RedirectURL is the object that will be created
	RedirectURL RedirectURL `json:"redirect_url"`
}

type CreateResponse struct {
	// StatusCode is the HTTP status code of the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// RedirectURL is the object that was created
	RedirectURL RedirectURL `json:"redirect_url"`
}

type GetRequest struct {
	// ProjectID is the ID of the project to get the redirect URL from
	ProjectID string `json:"project_id"`
	// URL is the redirect URL to get
	URL string `json:"url"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code of the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// RedirectURL is the object that was retrieved
	RedirectURL RedirectURL `json:"redirect_url"`
}

type GetAllRequest struct {
	// ProjectID is the ID of the project to get all the redirect URLs from
	ProjectID string `json:"project_id"`
}

type GetAllResponse struct {
	// StatusCode is the HTTP status code of the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// ProjectID is the ID of the project for the redirect URLs
	ProjectID string `json:"project_id"`
	// RedirectURLs is a list of all the redirect URLs for the project
	RedirectURLs []RedirectURL `json:"redirect_urls"`
}

type DeleteRequest struct {
	// ProjectID is the ID of the project to delete the redirect URL from
	ProjectID string `json:"project_id"`
	// URL is the redirect URL to delete
	URL string `json:"url"`
}

type DeleteResponse struct {
	// StatusCode is the HTTP status code of the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
}

type UpdateRequest struct {
	// ProjectID is the ID of the project to update the redirect URL in
	ProjectID string `json:"project_id"`
	// RedirectURL is the object that will be updated
	RedirectURL RedirectURL `json:"redirect_url"`
}

type UpdateResponse struct {
	// StatusCode is the HTTP status code of the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// RedirectURL is the object that was updated
	RedirectURL RedirectURL `json:"redirect_url"`
}
