package api

import (
	"net/http"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
)

const defaultBaseURI = "https://management.stytch.com"

type API struct {
	client *internal.Client

	EmailTemplates *EmailTemplatesClient
	Projects       *ProjectsClient
	ProjectMetrics *ProjectMetricsClient
	PublicTokens   *PublicTokensClient
	RBAC           *RBACClient
	RedirectURLs   *RedirectURLsClient
	SDK            *SDKClient
	Secrets        *SecretsClient
}

type apiConfig struct {
	WorkspaceKeyID     string
	WorkspaceKeySecret string

	baseURI    string
	httpClient *http.Client
}

type APIOption func(*apiConfig)

func WithBaseURI(baseURI string) APIOption {
	return func(a *apiConfig) {
		a.baseURI = baseURI
	}
}

func WithHTTPClient(client *http.Client) APIOption {
	return func(a *apiConfig) {
		a.httpClient = client
	}
}

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
		client:         client,
		EmailTemplates: newEmailTemplatesClient(client),
		Projects:       newProjectsClient(client),
		ProjectMetrics: newProjectMetricsClient(client),
		PublicTokens:   newPublicTokensClient(client),
		RBAC:           newRBACClient(client),
		RedirectURLs:   newRedirectURLsClient(client),
		SDK:            newSDKClient(client),
		Secrets:        newSecretsClient(client),
	}
}
