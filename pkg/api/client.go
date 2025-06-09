package api

import (
	"net/http"
	"strings"

	"github.com/stytchauth/stytch-management-go/v2/pkg/api/internal"
)

// This is the main entrypoint for interacting with the Stytch Management API
const defaultBaseURI = "https://management.stytch.com"

type API struct {
	client *internal.Client

	// These are the clients for all the different
	// resources available via the management API
	CountryCodeAllowlist   *CountryCodeAllowlistClient
	EmailTemplates         *EmailTemplatesClient
	JWTTemplates           *JWTTemplatesClient
	PasswordStrengthConfig *PasswordStrengthConfigClient
	Projects               *ProjectsClient
	ProjectMetrics         *ProjectMetricsClient
	PublicTokens           *PublicTokensClient
	RBACPolicy             *RBACPolicyClient
	RedirectURLs           *RedirectURLsClient
	SDK                    *SDKClient
	Secrets                *SecretsClient
}

type apiConfig struct {
	// The workspace key ID and secret required to authenticate
	// with the Stytch management API. These can be obtained
	// from the Stytch dashboard
	WorkspaceKeyID     string
	WorkspaceKeySecret string

	baseURI    string
	httpClient *http.Client
}

type APIOption func(*apiConfig)

func WithBaseURI(baseURI string) APIOption {
	return func(a *apiConfig) {
		a.baseURI = strings.TrimSuffix(baseURI, "/")
	}
}

func WithHTTPClient(client *http.Client) APIOption {
	return func(a *apiConfig) {
		a.httpClient = client
	}
}

// NewClient creates a new API client with the given workspace key ID and secret
func NewClient(workspaceKeyID string, workspaceKeySecret string, opts ...APIOption) *API {
	c := apiConfig{
		baseURI:    defaultBaseURI,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(&c)
	}

	client := internal.NewClient(internal.ClientConfig{
		WorkspaceKeyID:     workspaceKeyID,
		WorkspaceKeySecret: workspaceKeySecret,
		BaseURI:            c.baseURI,
		HTTPClient:         c.httpClient,
	})

	return &API{
		client:                 client,
		CountryCodeAllowlist:   newCountryCodeAllowlistClient(client),
		EmailTemplates:         newEmailTemplatesClient(client),
		JWTTemplates:           newJWTTemplatesClient(client),
		PasswordStrengthConfig: newPasswordStrengthConfigClient(client),
		Projects:               newProjectsClient(client),
		ProjectMetrics:         newProjectMetricsClient(client),
		PublicTokens:           newPublicTokensClient(client),
		RBACPolicy:             newRBACPolicyClient(client),
		RedirectURLs:           newRedirectURLsClient(client),
		SDK:                    newSDKClient(client),
		Secrets:                newSecretsClient(client),
	}
}

// NewAccessTokenClient creates a new API client with an access token.
func NewAccessTokenClient(accessToken string, opts ...APIOption) *API {
	c := apiConfig{
		baseURI:    defaultBaseURI,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(&c)
	}

	client := internal.NewClient(internal.ClientConfig{
		AccessToken: accessToken,
		BaseURI:     c.baseURI,
		HTTPClient:  c.httpClient,
	})

	return &API{
		client:                 client,
		CountryCodeAllowlist:   newCountryCodeAllowlistClient(client),
		EmailTemplates:         newEmailTemplatesClient(client),
		JWTTemplates:           newJWTTemplatesClient(client),
		PasswordStrengthConfig: newPasswordStrengthConfigClient(client),
		Projects:               newProjectsClient(client),
		ProjectMetrics:         newProjectMetricsClient(client),
		PublicTokens:           newPublicTokensClient(client),
		RBACPolicy:             newRBACPolicyClient(client),
		RedirectURLs:           newRedirectURLsClient(client),
		SDK:                    newSDKClient(client),
		Secrets:                newSecretsClient(client),
	}
}
