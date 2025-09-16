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

// RedirectTypes returns a list of all the possible RedirectType values.
func RedirectTypes() []RedirectType {
	return []RedirectType{
		RedirectTypeLogin,
		RedirectTypeSignup,
		RedirectTypeInvite,
		RedirectTypeResetPassword,
		RedirectTypeDiscovery,
	}
}

// URLRedirectType holds information for a specific kind of redirect.
type URLRedirectType struct {
	// Type is one of the RedirectType values.
	Type RedirectType `json:"type"`
	// IsDefault is true if this is the default redirect type, false otherwise.
	IsDefault bool `json:"is_default"`
}

// RedirectURL holds information for a specific redirect URL and all its redirect types.
type RedirectURL struct {
	// URL is the URL to which to redirect.
	URL string `json:"url"`
	// ValidTypes is a list of all the URLRedirectType available for this URL.
	ValidTypes []URLRedirectType `json:"valid_types"`
}

type CreateRequest struct {
	// ProjectSlug is the slug of the project for which to create the redirect URL.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to create the redirect URL.
	EnvironmentSlug string `json:"-"`
	// URL is the URL to which to redirect.
	URL string `json:"url"`
	// ValidTypes is a list of all the URLRedirectType available for this URL.
	ValidTypes []URLRedirectType `json:"valid_types"`
	// DoNotPromoteDefaults is used to suppress the automatic "promotion" of a RedirectURL to the
	// default if no other RedirectURL exists for the given type. This is primarily intended for use
	// with stytchauth/terraform-provider-stytch to allow Terraform provisioning to be idempotent. For
	// a Create request, the default behavior is to promote the RedirectURL to the default if no other
	// RedirectURL exists for the given type, even if `IsDefault` was set to false.
	//
	// WARNING: If this is set to true, it is possible to have valid RedirectURLs for a given type but
	// *no* default RedirectURL for that type. If no default exists for a given type, using an API
	// endpoint that uses redirect URLs (such as sending a magic link), you will need to explicitly
	// specify which redirect URL should be used.
	DoNotPromoteDefaults bool `json:"do_not_promote_defaults"`
}

type CreateResponse struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// RedirectURL is the object that was created.
	RedirectURL RedirectURL `json:"redirect_url"`
}

type GetRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve the redirect URL.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to retrieve the redirect URL.
	EnvironmentSlug string `json:"-"`
	// URL is the redirect URL to get.
	URL string `json:"url"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// RedirectURL is the object that was retrieved.
	RedirectURL RedirectURL `json:"redirect_url"`
}

type GetAllRequest struct {
	// ProjectSlug is the slug of the project for which to retrieve all redirect URLs.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to retrieve all redirect URLs.
	EnvironmentSlug string `json:"-"`
}

type GetAllResponse struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// RedirectURLs is a list of all the redirect URLs for the project.
	RedirectURLs []RedirectURL `json:"redirect_urls"`
}

type DeleteRequest struct {
	// ProjectSlug is the slug of the project for which to delete the redirect URL.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to delete the redirect URL.
	EnvironmentSlug string `json:"-"`
	// URL is the redirect URL to delete.
	URL string `json:"url"`
	// DoNotPromoteDefaults is used to suppress the automatic "promotion" of a RedirectURL to the
	// default if no other RedirectURL exists for the given type. This is primarily intended for use
	// with stytchauth/terraform-provider-stytch to allow Terraform provisioning to be idempotent. For
	// a Delete request, the default behavior is to promote some other valid redirect URL for the
	// given type to the new default if *this* URL was the current default.
	//
	// WARNING: If this is set to true, it is possible to have valid RedirectURLs for a given type but
	// *no* default RedirectURL for that type. If no default exists for a given type, using an API
	// endpoint that uses redirect URLs (such as sending a magic link), you will need to explicitly
	// specify which redirect URL should be used.
	DoNotPromoteDefaults bool `json:"do_not_promote_defaults"`
}

type DeleteResponse struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
}

type UpdateRequest struct {
	// ProjectSlug is the slug of the project for which to update the redirect URL.
	ProjectSlug string `json:"-"`
	// EnvironmentSlug is the slug of the environment for which to update the redirect URL.
	EnvironmentSlug string `json:"-"`
	// URL is the URL to which to redirect.
	URL string `json:"url"`
	// ValidTypes is a list of all the URLRedirectType available for this URL.
	ValidTypes []URLRedirectType `json:"valid_types"`
	// DoNotPromoteDefaults is used to suppress the automatic "promotion" of a RedirectURL to the
	// default if no other RedirectURL exists for the given type. This is primarily intended for use
	// with stytchauth/terraform-provider-stytch to allow Terraform provisioning to be idempotent. For
	// an Update request, the default behavior is a combination of what happens for Create and Delete
	// requests:
	// - If the RedirectURL is having a new valid type added, it may be promoted to the default if no
	//   other RedirectURL exists for that type.
	// - If the RedirectURL is having a valid type removed and is *currently* the default for that
	//   type, a new default will be promoted from the remaining URLs for that type.
	//
	// WARNING: If this is set to true, it is possible to have valid RedirectURLs for a given type but
	// *no* default RedirectURL for that type. If no default exists for a given type, using an API
	// endpoint that uses redirect URLs (such as sending a magic link), you will need to explicitly
	// specify which redirect URL should be used.
	DoNotPromoteDefaults bool `json:"do_not_promote_defaults"`
}

type UpdateResponse struct {
	// StatusCode is the HTTP status code of the response.
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request.
	RequestID string `json:"request_id"`
	// RedirectURL is the object that was updated.
	RedirectURL RedirectURL `json:"redirect_url"`
}
