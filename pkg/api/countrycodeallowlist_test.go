package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	cca "github.com/stytchauth/stytch-management-go/v3/pkg/models/countrycodeallowlist"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

func TestCountryCodeAllowlistClient_GetAllowedSMSCountryCodes(t *testing.T) {
	t.Run("default country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalB2B)
		ctx := context.Background()
		expected := []string{"CA", "US"}

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				Project:     project.Project,
				Environment: "test",
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("get country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}
		_, err := client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx,
			&cca.SetAllowedSMSCountryCodesRequest{
				Project:      project.Project,
				Environment:  "test",
				CountryCodes: expected,
			})
		require.NoError(t, err)

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				Project:     project.Project,
				Environment: "test",
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				Project:     "project-does-not-exist",
				Environment: "test",
			})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestCountryCodeAllowlistClient_GetAllowedWhatsAppCountryCodes(t *testing.T) {
	t.Run("default country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "US"}

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				Project:     project.Project,
				Environment: "test",
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("get country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}
		_, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				Project:      project.Project,
				Environment:  "test",
				CountryCodes: expected,
			})
		require.NoError(t, err)

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				Project:     project.Project,
				Environment: "test",
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("B2B WhatsApp not supported", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalB2B)
		ctx := context.Background()

		// Act
		_, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				Project:     project.Project,
				Environment: "test",
			})

		// Assert
		assert.ErrorContains(t, err, "country_code_allowlist_b2b_whatsapp_not_supported")
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				Project:     "project-does-not-exist",
				Environment: "test",
			})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestCountryCodeAllowlistClient_SetAllowedSMSCountryCodes(t *testing.T) {
	t.Run("set country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}

		// Act
		setResp, err := client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx,
			&cca.SetAllowedSMSCountryCodesRequest{
				Project:      project.Project,
				Environment:  "test",
				CountryCodes: expected,
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, setResp.CountryCodes)

		getResp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				Project:     project.Project,
				Environment: "test",
			})
		require.NoError(t, err)
		assert.Equal(t, expected, getResp.CountryCodes)
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		_, err := client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx,
			&cca.SetAllowedSMSCountryCodesRequest{
				Project:      "project-does-not-exist",
				Environment:  "test",
				CountryCodes: []string{"CA", "MX", "US"},
			})

		assert.Error(t, err)
	})
}

func TestCountryCodeAllowlistClient_SetAllowedWhatsAppCountryCodes(t *testing.T) {
	t.Run("set country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}

		// Act
		setResp, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				Project:      project.Project,
				Environment:  "test",
				CountryCodes: expected,
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, setResp.CountryCodes)

		getResp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				Project:     project.Project,
				Environment: "test",
			})
		require.NoError(t, err)
		assert.Equal(t, expected, getResp.CountryCodes)
	})
	t.Run("B2B WhatsApp not supported", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(project.VerticalB2B)
		ctx := context.Background()

		// Act
		_, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				Project:      project.Project,
				Environment:  "test",
				CountryCodes: []string{"CA", "MX", "US"},
			})

		// Assert
		assert.ErrorContains(t, err, "country_code_allowlist_b2b_whatsapp_not_supported")
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		_, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				Project:      "project-does-not-exist",
				Environment:  "test",
				CountryCodes: []string{"CA", "MX", "US"},
			})

		assert.Error(t, err)
	})
}
