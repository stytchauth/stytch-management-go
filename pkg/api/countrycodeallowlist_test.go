package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	cca "github.com/stytchauth/stytch-management-go/v3/pkg/models/countrycodeallowlist"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

func TestCountryCodeAllowlistClient_GetAllowedSMSCountryCodes(t *testing.T) {
	t.Run("default country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()
		expected := []string{"CA", "US"}

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("get country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}
		_, err := client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx,
			&cca.SetAllowedSMSCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
				CountryCodes:    expected,
			})
		require.NoError(t, err)

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
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
				ProjectSlug:     "project-does-not-exist",
				EnvironmentSlug: "test",
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
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()
		expected := []string{"CA", "US"}

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("get country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}
		_, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
				CountryCodes:    expected,
			})
		require.NoError(t, err)

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("B2B WhatsApp not supported", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		_, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
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
				ProjectSlug:     "project-does-not-exist",
				EnvironmentSlug: "test",
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
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}

		// Act
		setResp, err := client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx,
			&cca.SetAllowedSMSCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
				CountryCodes:    expected,
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, setResp.CountryCodes)

		getResp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
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
				ProjectSlug:     "project-does-not-exist",
				EnvironmentSlug: "test",
				CountryCodes:    []string{"CA", "MX", "US"},
			})

		assert.Error(t, err)
	})
}

func TestCountryCodeAllowlistClient_SetAllowedWhatsAppCountryCodes(t *testing.T) {
	t.Run("set country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}

		// Act
		setResp, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
				CountryCodes:    expected,
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, setResp.CountryCodes)

		getResp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
			})
		require.NoError(t, err)
		assert.Equal(t, expected, getResp.CountryCodes)
	})
	t.Run("B2B WhatsApp not supported", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		_, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				ProjectSlug:     env.ProjectSlug,
				EnvironmentSlug: env.EnvironmentSlug,
				CountryCodes:    []string{"CA", "MX", "US"},
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
				ProjectSlug:     "project-does-not-exist",
				EnvironmentSlug: "test",
				CountryCodes:    []string{"CA", "MX", "US"},
			})

		assert.Error(t, err)
	})
}
