package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	cca "github.com/stytchauth/stytch-management-go/v2/pkg/models/countrycodeallowlist"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"
)

func TestCountryCodeAllowlistClient_GetAllowedSMSCountryCodes(t *testing.T) {
	t.Run("default country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		expected := []string{"CA", "US"}

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				ProjectID: project.LiveProjectID,
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("get country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}
		_, err := client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx,
			&cca.SetAllowedSMSCountryCodesRequest{
				ProjectID:    project.LiveProjectID,
				CountryCodes: expected,
			})
		require.NoError(t, err)

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				ProjectID: project.LiveProjectID,
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
				ProjectID: "project-does-not-exist",
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
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "US"}

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				ProjectID: project.LiveProjectID,
			})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("get country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}
		_, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				ProjectID:    project.LiveProjectID,
				CountryCodes: expected,
			})
		require.NoError(t, err)

		// Act
		resp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				ProjectID: project.LiveProjectID,
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
		resp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				ProjectID: "project-does-not-exist",
			})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestCountryCodeAllowlistClient_SetAllowedSMSCountryCodes(t *testing.T) {
	t.Run("get country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}

		// Act
		_, err := client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx,
			&cca.SetAllowedSMSCountryCodesRequest{
				ProjectID:    project.LiveProjectID,
				CountryCodes: expected,
			})

		// Assert
		assert.NoError(t, err)
		resp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx,
			&cca.GetAllowedSMSCountryCodesRequest{
				ProjectID: project.LiveProjectID,
			})
		require.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		_, err := client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx,
			&cca.SetAllowedSMSCountryCodesRequest{
				ProjectID:    "project-does-not-exist",
				CountryCodes: []string{"CA", "MX", "US"},
			})

		assert.Error(t, err)
	})
}

func TestCountryCodeAllowlistClient_SetAllowedWhatsAppCountryCodes(t *testing.T) {
	t.Run("get country codes", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		expected := []string{"CA", "MX", "US"}

		// Act
		_, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				ProjectID:    project.LiveProjectID,
				CountryCodes: expected,
			})

		// Assert
		assert.NoError(t, err)
		resp, err := client.CountryCodeAllowlist.GetAllowedWhatsAppCountryCodes(ctx,
			&cca.GetAllowedWhatsAppCountryCodesRequest{
				ProjectID: project.LiveProjectID,
			})
		require.NoError(t, err)
		assert.Equal(t, expected, resp.CountryCodes)
	})
	t.Run("project does not exist", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		ctx := context.Background()

		// Act
		_, err := client.CountryCodeAllowlist.SetAllowedWhatsAppCountryCodes(ctx,
			&cca.SetAllowedWhatsAppCountryCodesRequest{
				ProjectID:    "project-does-not-exist",
				CountryCodes: []string{"CA", "MX", "US"},
			})

		assert.Error(t, err)
	})
}
