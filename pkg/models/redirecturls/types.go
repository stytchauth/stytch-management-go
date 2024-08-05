package redirecturls

type RedirectType string

const (
	RedirectTypeLogin         RedirectType = "login"
	RedirectTypeSignup        RedirectType = "signup"
	RedirectTypeInvite        RedirectType = "invite"
	RedirectTypeResetPassword RedirectType = "reset_password"
	RedirectTypeDiscovery     RedirectType = "discovery"
)

type RedirectURL struct {
	MagicLinkURLID string         `json:"magic_link_url_id"`
	URL            string         `json:"url"`
	ValidTypes     []RedirectType `json:"valid_types"`
	DefaultTypes   []RedirectType `json:"default_types"`
}

type CreateRequest struct {
	ProjectID string       `json:"project_id"`
	URL       string       `json:"url"`
	Type      RedirectType `json:"type"`
	IsDefault bool         `json:"is_default"`
}

type CreateResponse struct {
	StatusCode     int    `json:"status_code"`
	RequestID      string `json:"request_id"`
	MagicLinkURLID string `json:"magic_link_url_id"`
}

type GetAllRequest struct {
	ProjectID string `json:"project_id"`
}

type GetAllResponse struct {
	StatusCode   int           `json:"status_code"`
	RequestID    string        `json:"request_id"`
	ProjectID    string        `json:"project_id"`
	RedirectURLs []RedirectURL `json:"redirect_urls"`
}

type RemoveValidTypeRequest struct {
	ProjectID string       `json:"project_id"`
	URL       string       `json:"url"`
	Type      RedirectType `json:"type"`
}

type RemoveValidTypeResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id"`
}

type DeleteRequest struct {
	ProjectID string `json:"project_id"`
	URL       string `json:"url"`
}

type DeleteResponse struct {
	StatusCode int    `json:"status_code"`
	RequestID  string `json:"request_id"`
}
