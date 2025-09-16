package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/sdk"
)

func makeTestConsumerConfig() sdk.ConsumerConfig {
	return sdk.ConsumerConfig{
		Basic: &sdk.ConsumerBasicConfig{
			Enabled:   true,
			Domains:   []string{"https://example.com"},
			BundleIDs: []string{"com.example.myapp"},
		},
		Sessions: &sdk.ConsumerSessionsConfig{
			MaxSessionDurationMinutes: 120,
		},
		MagicLinks: &sdk.ConsumerMagicLinksConfig{
			LoginOrCreateEnabled: true,
			SendEnabled:          true,
			PKCERequired:         false,
		},
		OTPs: &sdk.ConsumerOTPsConfig{
			SMSLoginOrCreateEnabled:      true,
			WhatsAppLoginOrCreateEnabled: false,
			EmailLoginOrCreateEnabled:    true,
			SMSSendEnabled:               true,
			WhatsAppSendEnabled:          false,
			EmailSendEnabled:             true,
			SMSAutofillMetadata: []sdk.SMSAutofillMetadata{
				{
					MetadataType:  sdk.SMSAutofillMetadataTypeDomain,
					MetadataValue: "myapp.com",
					BundleID:      "com.example.myapp",
				},
			},
		},
		OAuth: &sdk.ConsumerOAuthConfig{
			Enabled:      true,
			PKCERequired: false,
		},
		TOTPs: &sdk.ConsumerTOTPsConfig{
			Enabled:     true,
			CreateTOTPs: true,
		},
		WebAuthn: &sdk.ConsumerWebAuthnConfig{
			Enabled:         true,
			CreateWebAuthns: true,
		},
		CryptoWallets: &sdk.ConsumerCryptoWalletsConfig{
			Enabled:      true,
			SIWERequired: false,
		},
		// DFPPA sdk settings cannot be modified beyond defaults
		// unless the project has DFPPA enabled by Stytch.
		DFPPA: &sdk.ConsumerDFPPAConfig{
			Enabled:     sdk.DFPPASettingDisabled,
			OnChallenge: sdk.DFPPAOnChallengeActionAllow,
		},
		Biometrics: &sdk.ConsumerBiometricsConfig{
			Enabled:                 true,
			CreateBiometricsEnabled: true,
		},
		Passwords: &sdk.ConsumerPasswordsConfig{
			Enabled:                       true,
			PKCERequiredForPasswordResets: false,
		},
		Cookies: &sdk.ConsumerCookiesConfig{
			// Only disabled is supported, unless the project has
			// CNAMEs configured.
			HttpOnlyCookies: sdk.HttpOnlyCookiesSettingDisabled,
		},
	}
}

func makeTestB2BConfig() sdk.B2BConfig {
	return sdk.B2BConfig{
		Basic: &sdk.B2BBasicConfig{
			Enabled:                 true,
			AllowSelfOnboarding:     true,
			EnableMemberPermissions: true,
			Domains: []sdk.AuthorizedB2BDomain{
				{
					Domain:      "https://myb2bapp.com",
					SlugPattern: "https://{{slug}}.myb2bapp.com",
				},
			},
			BundleIDs: []string{"com.example.b2bapp"},
		},
		Sessions: &sdk.B2BSessionsConfig{
			MaxSessionDurationMinutes: 180,
		},
		MagicLinks: &sdk.B2BMagicLinksConfig{
			Enabled:      true,
			PKCERequired: true,
		},
		OAuth: &sdk.B2BOAuthConfig{
			Enabled:      true,
			PKCERequired: true,
		},
		TOTPs: &sdk.B2BTOTPsConfig{
			Enabled:     true,
			CreateTOTPs: true,
		},
		SSO: &sdk.B2BSSOConfig{
			Enabled:      true,
			PKCERequired: false,
		},
		OTPs: &sdk.B2BOTPsConfig{
			SMSEnabled:   true,
			EmailEnabled: true,
			SMSAutofillMetadata: []sdk.SMSAutofillMetadata{
				{
					MetadataType:  sdk.SMSAutofillMetadataTypeHash,
					MetadataValue: "abc123hash",
					BundleID:      "com.example.b2bapp",
				},
			},
		},
		// DFPPA sdk settings cannot be modified beyond defaults
		// unless the project has DFPPA enabled by Stytch.
		DFPPA: &sdk.B2BDFPPAConfig{
			Enabled:     sdk.DFPPASettingEnabled,
			OnChallenge: sdk.DFPPAOnChallengeActionBlock,
		},
		Passwords: &sdk.B2BPasswordsConfig{
			Enabled:                       true,
			PKCERequiredForPasswordResets: true,
		},
		Cookies: &sdk.B2BCookiesConfig{
			// Only disabled is supported, unless the project has
			// CNAMEs configured.
			HttpOnlyCookies: sdk.HttpOnlyCookiesSettingDisabled,
		},
	}
}

func TestSDKClient_GetConsumerConfig(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		expectedConfig := makeTestConsumerConfig()

		// First set the configuration
		_, err := client.SDK.SetConsumerConfig(ctx, sdk.SetConsumerConfigRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Config:          expectedConfig,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.SDK.GetConsumerConfig(ctx, sdk.GetConsumerConfigRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedConfig, resp.Config)
	})
}

func TestSDKClient_SetConsumerConfig(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		config := makeTestConsumerConfig()

		// Act
		resp, err := client.SDK.SetConsumerConfig(ctx, sdk.SetConsumerConfigRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Config:          config,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, config, resp.Config)
	})
	t.Run("invalid vertical config returns error", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalConsumer)
		ctx := context.Background()
		config := makeTestB2BConfig()
		config.Basic.Enabled = false

		// Act
		resp, err := client.SDK.SetB2BConfig(ctx, sdk.SetB2BConfigRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Config:          config,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestSDKClient_GetB2BConfig(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		expectedConfig := makeTestB2BConfig()

		// First set the configuration
		_, err := client.SDK.SetB2BConfig(ctx, sdk.SetB2BConfigRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Config:          expectedConfig,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.SDK.GetB2BConfig(ctx, sdk.GetB2BConfigRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, expectedConfig, resp.Config)
	})
}

func TestSDKClient_SetB2BConfig(t *testing.T) {
	t.Run("base case", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		config := makeTestB2BConfig()

		// Act
		resp, err := client.SDK.SetB2BConfig(ctx, sdk.SetB2BConfigRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Config:          config,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, config, resp.Config)
	})
	t.Run("invalid vertical config returns error", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()
		config := makeTestConsumerConfig()
		config.Basic.Enabled = false

		// Act
		resp, err := client.SDK.SetConsumerConfig(ctx, sdk.SetConsumerConfigRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Config:          config,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}
