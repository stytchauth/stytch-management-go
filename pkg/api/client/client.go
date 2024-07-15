package client

import (
	"net/http"

	"github.com/stytchauth/stytch-management-go/v1/pkg/api/internal"
)

const defaultBaseURI = "https://management.stytch.com"

type API struct {
	client *internal.Client
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

func New(workspaceKeyID string, workspaceKeySecret string, opts ...APIOption) *API {
	c := apiConfig{
		baseURI:    defaultBaseURI,
		httpClient: &http.Client{},
	}

	for _, opt := range opts {
		opt(&c)
	}

	return &API{
		client: internal.NewClient(internal.ClientConfig{
			WorkspaceKeyID:     workspaceKeyID,
			WorkspaceKeySecret: workspaceKeySecret,
			BaseURI:            c.baseURI,
			HTTPClient:         c.httpClient,
		}),
	}
}
