package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/countrycodeallowlist"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/emailtemplates"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environmentmetrics"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/eventlogstreaming"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/jwttemplates"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/passwordstrengthconfig"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/publictokens"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/rbacpolicy"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/redirecturls"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/sdk"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/secrets"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/trustedtokenprofiles"
)

// Test_Integration is a comprehensive integration test that touches as many methods as possible from the Stytch
// Management API v3. This test covers the full lifecycle of workspace management.
func Test_Integration(t *testing.T) {
	client := NewTestClient(t)
	ctx := context.Background()
	project := client.DisposableProject(projects.VerticalB2B)

	// Get the project
	getResp, err := client.Projects.Get(ctx, projects.GetRequest{Project: project.Project})
	require.NoError(t, err)
	assert.Equal(t, project.Project, getResp.Project.Project)

	// Update the project
	updateResp, err := client.Projects.Update(ctx, projects.UpdateRequest{
		Project: project.Project,
		Name:    ptr("Updated Project"),
	})
	require.NoError(t, err)
	assert.Equal(t, "Updated Project", updateResp.Project.Name)

	// Get all projects
	getAllResp, err := client.Projects.GetAll(ctx, projects.GetAllRequest{})
	require.NoError(t, err)
	assert.GreaterOrEqual(t, len(getAllResp.Projects), 1)

	// Get all environments (should have live and test by default)
	envGetAllResp, err := client.Environments.GetAll(ctx, environments.GetAllRequest{
		Project: project.Project,
	})
	require.NoError(t, err)
	assert.Len(t, envGetAllResp.Environments, 2) // live and test

	var liveEnv environments.Environment
	for _, env := range envGetAllResp.Environments {
		if env.Type == environments.EnvironmentTypeLive {
			liveEnv = env
			break
		}
	}

	// Get individual environments
	liveGetResp, err := client.Environments.Get(ctx, environments.GetRequest{
		Project:     project.Project,
		Environment: liveEnv.Environment,
	})
	require.NoError(t, err)
	assert.Equal(t, liveEnv.Environment, liveGetResp.Environment.Environment)

	// Create a custom environment
	customEnvResp, err := client.Environments.Create(ctx, environments.CreateRequest{
		Project: project.Project,
		Name:    "Custom Env",
		Type:    environments.EnvironmentTypeTest,
	})
	require.NoError(t, err)
	customEnv := customEnvResp.Environment

	// Update the custom environment
	_, err = client.Environments.Update(ctx, environments.UpdateRequest{
		Project:     project.Project,
		Environment: customEnv.Environment,
		Name:        ptr("Updated Custom Env"),
	})
	require.NoError(t, err)

	t.Run("Secrets", func(t *testing.T) {
		// Create a secret in the custom environment
		secretResp, err := client.Secrets.Create(ctx, secrets.CreateSecretRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.NotEmpty(t, secretResp.CreatedSecret.SecretID)

		// Get the secret
		getSecretResp, err := client.Secrets.Get(ctx, secrets.GetSecretRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
			SecretID:    secretResp.CreatedSecret.SecretID,
		})
		require.NoError(t, err)
		assert.Equal(t, secretResp.CreatedSecret.SecretID, getSecretResp.Secret.SecretID)

		// Get all secrets
		getAllSecretsResp, err := client.Secrets.GetAll(ctx, secrets.GetAllSecretsRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(getAllSecretsResp.Secrets), 1)

		// Delete the secret
		_, err = client.Secrets.Delete(ctx, secrets.DeleteSecretRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
			SecretID:    secretResp.CreatedSecret.SecretID,
		})
		require.NoError(t, err)
	})

	t.Run("PublicTokens", func(t *testing.T) {
		// Get all public tokens
		pubTokensResp, err := client.PublicTokens.GetAll(ctx, publictokens.GetAllRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(pubTokensResp.PublicTokens), 0)

		// If there are public tokens, get one of them
		if len(pubTokensResp.PublicTokens) > 0 {
			firstToken := pubTokensResp.PublicTokens[0]
			getTokenResp, err := client.PublicTokens.Get(ctx, publictokens.GetRequest{
				Project:     project.Project,
				Environment: customEnv.Environment,
				PublicToken: firstToken.PublicToken,
			})
			require.NoError(t, err)
			assert.Equal(t, firstToken.PublicToken, getTokenResp.PublicToken.PublicToken)
		}
	})

	t.Run("RedirectURLs", func(t *testing.T) {
		testURL := "https://customer.example.com/callback"

		// Create a redirect URL
		createURLResp, err := client.RedirectURLs.Create(ctx, redirecturls.CreateRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
			URL:         testURL,
			ValidTypes: []redirecturls.URLRedirectType{
				{Type: redirecturls.RedirectTypeLogin, IsDefault: true},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, testURL, createURLResp.RedirectURL.URL)

		// Get all redirect URLs
		getAllURLsResp, err := client.RedirectURLs.GetAll(ctx, redirecturls.GetAllRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(getAllURLsResp.RedirectURLs), 1)

		// Get the specific redirect URL
		getURLResp, err := client.RedirectURLs.Get(ctx, redirecturls.GetRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
			URL:         testURL,
		})
		require.NoError(t, err)
		assert.Equal(t, testURL, getURLResp.RedirectURL.URL)

		// Update the redirect URL
		updateURLResp, err := client.RedirectURLs.Update(ctx, redirecturls.UpdateRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
			URL:         testURL,
			ValidTypes: []redirecturls.URLRedirectType{
				{Type: redirecturls.RedirectTypeLogin, IsDefault: true},
				{Type: redirecturls.RedirectTypeSignup, IsDefault: false},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, testURL, updateURLResp.RedirectURL.URL)
		assert.Len(t, updateURLResp.RedirectURL.ValidTypes, 2)

		// Delete the redirect URL
		_, err = client.RedirectURLs.Delete(ctx, redirecturls.DeleteRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
			URL:         testURL,
		})
		require.NoError(t, err)
	})

	t.Run("CountryCodeAllowlist", func(t *testing.T) {
		// Get current SMS allowlist
		getSMSResp, err := client.CountryCodeAllowlist.GetAllowedSMSCountryCodes(ctx, &countrycodeallowlist.GetAllowedSMSCountryCodesRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.NotNil(t, getSMSResp.CountryCodes)

		// Set SMS allowlist
		setSMSResp, err := client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx, &countrycodeallowlist.SetAllowedSMSCountryCodesRequest{
			Project:      project.Project,
			Environment:  customEnv.Environment,
			CountryCodes: []string{"US", "CA", "GB", "AU"},
		})
		require.NoError(t, err)
		assert.Contains(t, setSMSResp.CountryCodes, "US")
		assert.Contains(t, setSMSResp.CountryCodes, "CA")

		// Reset SMS allowlist
		_, err = client.CountryCodeAllowlist.SetAllowedSMSCountryCodes(ctx, &countrycodeallowlist.SetAllowedSMSCountryCodesRequest{
			Project:      project.Project,
			Environment:  customEnv.Environment,
			CountryCodes: countrycodeallowlist.DefaultCountryCodes,
		})
		require.NoError(t, err)
	})

	t.Run("EmailTemplates", func(t *testing.T) {
		// Get all email templates
		getTemplatesResp, err := client.EmailTemplates.GetAll(ctx, emailtemplates.GetAllRequest{
			Project: project.Project,
		})
		require.NoError(t, err)

		// If there are templates, update one
		if len(getTemplatesResp.EmailTemplates) > 0 {
			template := getTemplatesResp.EmailTemplates[0]

			// Update the template
			_, err = client.EmailTemplates.Update(ctx, emailtemplates.UpdateRequest{
				Project:    project.Project,
				TemplateID: template.TemplateID,
			})
			require.NoError(t, err)

			// Get the updated template
			getTemplateResp, err := client.EmailTemplates.Get(ctx, emailtemplates.GetRequest{
				Project:    project.Project,
				TemplateID: template.TemplateID,
			})
			require.NoError(t, err)
			assert.Equal(t, template.TemplateID, getTemplateResp.EmailTemplate.TemplateID)
		}
	})

	t.Run("JWTTemplates", func(t *testing.T) {
		// Set/Update JWT template
		jwtContent := `{"custom_user_id": "user-123", "custom_email": "test@example.com"}`
		_, err = client.JWTTemplates.Set(ctx, &jwttemplates.SetRequest{
			Project:         project.Project,
			Environment:     customEnv.Environment,
			TemplateType:    jwttemplates.TemplateTypeSession,
			TemplateContent: jwtContent,
		})
		require.NoError(t, err)

		// Get JWT template for session type
		getJWTTemplateResp, err := client.JWTTemplates.Get(ctx, &jwttemplates.GetRequest{
			Project:      project.Project,
			Environment:  customEnv.Environment,
			TemplateType: jwttemplates.TemplateTypeSession,
		})
		require.NoError(t, err)
		assert.Equal(t, jwttemplates.TemplateTypeSession, getJWTTemplateResp.JWTTemplate.TemplateType)
		assert.Equal(t, jwtContent, getJWTTemplateResp.JWTTemplate.TemplateContent)
	})

	t.Run("PasswordStrengthConfig", func(t *testing.T) {
		// Get current password strength config
		getPSCResp, err := client.PasswordStrengthConfig.Get(ctx, passwordstrengthconfig.GetRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.NotNil(t, getPSCResp.PasswordStrengthConfig)

		// Set/Update password strength config (just ensure it can be set)
		_, err = client.PasswordStrengthConfig.Set(ctx, passwordstrengthconfig.SetRequest{
			Project:                   project.Project,
			Environment:               customEnv.Environment,
			ValidationPolicy:          passwordstrengthconfig.ValidationPolicyLUDS,
			LudsMinPasswordLength:     ptr(12),
			LudsMinPasswordComplexity: ptr(4),
		})
		require.NoError(t, err)
	})

	t.Run("RBACPolicy", func(t *testing.T) {
		// Get RBAC policy
		getRBACResp, err := client.RBACPolicy.Get(ctx, rbacpolicy.GetRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.NotNil(t, getRBACResp.Policy)
	})

	t.Run("SDK", func(t *testing.T) {
		// Get B2B SDK config since we're testing a B2B project
		getB2BSDKResp, err := client.SDK.GetB2BConfig(ctx, sdk.GetB2BConfigRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.NotNil(t, getB2BSDKResp)

		// Set B2B SDK config
		_, err = client.SDK.SetB2BConfig(ctx, sdk.SetB2BConfigRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
	})

	t.Run("TrustedTokenProfiles", func(t *testing.T) {
		// Get all trusted token profiles
		getTTPResp, err := client.TrustedTokenProfiles.GetAll(ctx, &trustedtokenprofiles.GetAllTrustedTokenProfilesRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.GreaterOrEqual(t, len(getTTPResp.TrustedTokenProfiles), 0)

		// Create a trusted token profile with a sample RSA public key
		sampleRSAPublicKey := `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4f5wg5l2hKsTeNem/V41
fGnJm6gOdrj8ym3rFkEjWT2btYkUWzTdjqFMAYXbWnDxYi5qeJm9qNJKN/Wc5jNi
v6mHwV0KxD5IG7P3UXL9aK3D4lJ3CaYlQx2Y/YEfnVZ9qW8NfWv8dC4vTxuKcKy4
FI/EH9Z6L8jkq5XG6PL0yVG5T7eAW8f8F4l2LNm8TQhD8X/dFjJg/vhT8CvGhbDk
z4TXJK5u9K8FsKv7QV9P5wGhUK4e3XRPY2qNxV7PKxJ9pYgNxJ6bKh0nfQcQYa7K
bE5Uj+MuUDFQ4g3W+5v9S0s2SG8GpN2p8qUhzJGpKzGkPE4BZ6Q4GwsKkfTl8Oce
QIDAQAB
-----END PUBLIC KEY-----`

		createTTPResp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			Project:         project.Project,
			Environment:     customEnv.Environment,
			Name:            "Test Profile",
			Audience:        "audience-test",
			Issuer:          "test-issuer",
			PublicKeyType:   "RSA",
			PEMFiles:        []string{sampleRSAPublicKey},
			CanJITProvision: false,
		})
		require.NoError(t, err)
		assert.Equal(t, "Test Profile", createTTPResp.TrustedTokenProfile.Name)

		profileID := createTTPResp.TrustedTokenProfile.ID

		// Get the trusted token profile
		getTTPByIDResp, err := client.TrustedTokenProfiles.Get(ctx, &trustedtokenprofiles.GetTrustedTokenProfileRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
			ProfileID:   profileID,
		})
		require.NoError(t, err)
		assert.Equal(t, profileID, getTTPByIDResp.TrustedTokenProfile.ID)

		// Update the trusted token profile
		updateTTPResp, err := client.TrustedTokenProfiles.Update(ctx, &trustedtokenprofiles.UpdateTrustedTokenProfileRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
			ProfileID:   profileID,
			Name:        ptr("Updated Profile"),
		})
		require.NoError(t, err)
		assert.Equal(t, "Updated Profile", updateTTPResp.TrustedTokenProfile.Name)

		// Delete the trusted token profile
		_, err = client.TrustedTokenProfiles.Delete(ctx, &trustedtokenprofiles.DeleteTrustedTokenProfileRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
			ProfileID:   profileID,
		})
		require.NoError(t, err)
	})

	t.Run("EnvironmentMetrics", func(t *testing.T) {
		// Get environment metrics
		getMetricsResp, err := client.EnvironmentMetrics.Get(ctx, environmentmetrics.GetRequest{
			Project:     project.Project,
			Environment: customEnv.Environment,
		})
		require.NoError(t, err)
		assert.NotNil(t, getMetricsResp.Metrics)
	})

	t.Run("EventLogStreaming", func(t *testing.T) {
		// Create a Datadog event log streaming configuration
		createELSResp, err := client.EventLogStreaming.Create(ctx, eventlogstreaming.CreateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteUS,
					APIKey: "1234567890abcdef1234567890abcdef",
				},
			},
		})
		require.NoError(t, err)
		assert.Equal(t, eventlogstreaming.DestinationTypeDatadog, createELSResp.EventLogStreamingConfig.DestinationType)

		// Get the event log streaming configuration
		getELSResp, err := client.EventLogStreaming.Get(ctx, eventlogstreaming.GetEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})
		require.NoError(t, err)
		assert.Equal(t, eventlogstreaming.DestinationTypeDatadog, getELSResp.EventLogStreamingConfig.DestinationType)

		// Update the event log streaming configuration
		_, err = client.EventLogStreaming.Update(ctx, eventlogstreaming.UpdateEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
			DestinationConfig: eventlogstreaming.DestinationConfig{
				Datadog: &eventlogstreaming.DatadogConfig{
					Site:   eventlogstreaming.DatadogSiteEU,
					APIKey: "abcdef1234567890abcdef1234567890",
				},
			},
		})
		require.NoError(t, err)

		// Delete the event log streaming configuration
		_, err = client.EventLogStreaming.Delete(ctx, eventlogstreaming.DeleteEventLogStreamingRequest{
			Project:         project.Project,
			Environment:     TestEnvironment,
			DestinationType: eventlogstreaming.DestinationTypeDatadog,
		})
		require.NoError(t, err)
	})

	// Clean up the custom environment
	_, err = client.Environments.Delete(ctx, environments.DeleteRequest{
		Project:     project.Project,
		Environment: customEnv.Environment,
	})
	require.NoError(t, err)
}
