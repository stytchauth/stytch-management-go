package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stytchauth/stytch-management-go/v3/pkg/api/internal"
	cca "github.com/stytchauth/stytch-management-go/v3/pkg/models/countrycodeallowlist"
)

type CountryCodeAllowlistClient struct {
	client *internal.Client
}

func newCountryCodeAllowlistClient(c *internal.Client) *CountryCodeAllowlistClient {
	return &CountryCodeAllowlistClient{client: c}
}

func (c *CountryCodeAllowlistClient) GetAllowedSMSCountryCodes(
	ctx context.Context,
	body *cca.GetAllowedSMSCountryCodesRequest,
) (*cca.GetAllowedSMSCountryCodesResponse, error) {
	var resp cca.GetAllowedSMSCountryCodesResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/allowed_country_codes/sms", body.Project, body.Environment),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *CountryCodeAllowlistClient) GetAllowedWhatsAppCountryCodes(
	ctx context.Context,
	body *cca.GetAllowedWhatsAppCountryCodesRequest,
) (*cca.GetAllowedWhatsAppCountryCodesResponse, error) {
	var resp cca.GetAllowedWhatsAppCountryCodesResponse
	err := c.client.NewRequest(
		ctx,
		http.MethodGet,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/allowed_country_codes/whatsapp", body.Project, body.Environment),
		nil,
		nil,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *CountryCodeAllowlistClient) SetAllowedSMSCountryCodes(
	ctx context.Context,
	body *cca.SetAllowedSMSCountryCodesRequest,
) (*cca.SetAllowedSMSCountryCodesResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp cca.SetAllowedSMSCountryCodesResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/allowed_country_codes/sms", body.Project, body.Environment),
		nil,
		jsonBody,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}

func (c *CountryCodeAllowlistClient) SetAllowedWhatsAppCountryCodes(
	ctx context.Context,
	body *cca.SetAllowedWhatsAppCountryCodesRequest,
) (*cca.SetAllowedWhatsAppCountryCodesResponse, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	var resp cca.SetAllowedWhatsAppCountryCodesResponse
	err = c.client.NewRequest(
		ctx,
		http.MethodPost,
		fmt.Sprintf("/pwa/v3/projects/%s/environments/%s/allowed_country_codes/whatsapp", body.Project, body.Environment),
		nil,
		jsonBody,
		&resp)
	if err != nil {
		return nil, err
	}
	return &resp, err
}
