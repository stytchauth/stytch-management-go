package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/trustedtokenprofiles"
)

type TrustedTokenProfilesClient struct {
	client *internal.Client
}

func newTrustedTokenProfilesClient(c *internal.Client) *TrustedTokenProfilesClient {
	return &TrustedTokenProfilesClient{client: c}
}

// Get retrieves the trusted token profile for an environment.
func (c *TrustedTokenProfilesClient) Get(
	ctx context.Context,
	body *trustedtokenprofiles.GetTrustedTokenProfileRequest,
) (*trustedtokenprofiles.GetTrustedTokenProfileResponse, error) {
	var resp trustedtokenprofiles.GetTrustedTokenProfileResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/trusted_token_profiles/%s", body.Project, body.Environment, body.ProfileID),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// List retrieves all the trusted token profiles for an environment.
func (c *TrustedTokenProfilesClient) List(
	ctx context.Context,
	body *trustedtokenprofiles.ListTrustedTokenProfilesRequest,
) (*trustedtokenprofiles.ListTrustedTokenProfilesResponse, error) {
	var resp trustedtokenprofiles.ListTrustedTokenProfilesResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/trusted_token_profiles", body.Project, body.Environment),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Create creates a trusted token profile for an environment.
func (c *TrustedTokenProfilesClient) Create(
	ctx context.Context,
	body *trustedtokenprofiles.CreateTrustedTokenProfileRequest,
) (*trustedtokenprofiles.CreateTrustedTokenProfileResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp trustedtokenprofiles.CreateTrustedTokenProfileResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/trusted_token_profiles", body.Project, body.Environment),
		nil,
		jsonBody,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Update updates a trusted token profile for an environment.
func (c *TrustedTokenProfilesClient) Update(
	ctx context.Context,
	body *trustedtokenprofiles.UpdateTrustedTokenProfileRequest,
) (*trustedtokenprofiles.UpdateTrustedTokenProfileResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp trustedtokenprofiles.UpdateTrustedTokenProfileResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPatch,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/trusted_token_profiles/%s", body.Project, body.Environment, body.ProfileID),
		nil,
		jsonBody,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// Delete deletes a trusted token profile for an environment.
func (c *TrustedTokenProfilesClient) Delete(
	ctx context.Context,
	body *trustedtokenprofiles.DeleteTrustedTokenProfileRequest,
) (*trustedtokenprofiles.DeleteTrustedTokenProfileResponse, error) {
	var resp trustedtokenprofiles.DeleteTrustedTokenProfileResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/trusted_token_profiles/%s", body.Project, body.Environment, body.ProfileID),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// CreatePEM creates a PEM file for a trusted token profile for an environment.
func (c *TrustedTokenProfilesClient) CreatePEM(
	ctx context.Context,
	body *trustedtokenprofiles.CreatePEMFileRequest,
) (*trustedtokenprofiles.CreatePEMFileResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp trustedtokenprofiles.CreatePEMFileResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/trusted_token_profiles/%s/keys", body.Project, body.Environment, body.ProfileID),
		nil,
		jsonBody,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// GetPEM retrieves a PEM file for a trusted token profile for an environment.
func (c *TrustedTokenProfilesClient) GetPEM(
	ctx context.Context,
	body *trustedtokenprofiles.GetPEMFileRequest,
) (*trustedtokenprofiles.GetPEMFileResponse, error) {
	var resp trustedtokenprofiles.GetPEMFileResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/trusted_token_profiles/%s/keys/%s", body.Project, body.Environment, body.ProfileID, body.PEMFileID),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

// DeletePEM deletes a PEM file for a trusted token profile for an environment.
func (c *TrustedTokenProfilesClient) DeletePEM(
	ctx context.Context,
	body *trustedtokenprofiles.DeletePEMFileRequest,
) (*trustedtokenprofiles.DeletePEMFileResponse, error) {
	var resp trustedtokenprofiles.DeletePEMFileResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodDelete,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/trusted_token_profiles/%s/keys/%s", body.Project, body.Environment, body.ProfileID, body.PEMFileID),
		nil,
		nil,
		&resp,
	)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
