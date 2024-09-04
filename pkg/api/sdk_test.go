package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/sdk"
)

func makeTestConsumerConfig(t *testing.T) sdk.ConsumerConfig {
	t.Helper()
	return sdk.ConsumerConfig{
		Basic: &sdk.Consumer_BasicConfig{
			Enabled:        true,
			CreateNewUsers: true,
			Domains:        []string{"https://example.com"},
			BundleIDs:      []string{"com.example.app"},
		},
		Sessions: &sdk.Consumer_SessionsConfig{
			Enabled:                   true,
			MaxSessionDurationMinutes: 60,
		},
		MagicLinks: &sdk.Consumer_MagicLinksConfig{
			LoginOrCreateEnabled: true,
			SendEnabled:          false,
			PKCERequired:         false,
		},
		OTPs: &sdk.Consumer_OTPsConfig{
			SMSLoginOrCreateEnabled:      true,
			WhatsAppLoginOrCreateEnabled: false,
			EmailLoginOrCreateEnabled:    false,
			SMSSendEnabled:               false,
			WhatsAppSendEnabled:          false,
			EmailSendEnabled:             false,
			SMSAutofillMetadata:          []sdk.SMSAutofillMetadata{},
		},
		OAuth: &sdk.Consumer_OAuthConfig{
			Enabled:      true,
			PKCERequired: true,
		},
		TOTPs: &sdk.Consumer_TOTPsConfig{
			CreateTOTPs: true,
			Enabled:     true,
		},
		WebAuthn: &sdk.Consumer_WebAuthnConfig{
			CreateWebAuthns: false,
			Enabled:         false,
		},
		CryptoWallets: &sdk.Consumer_CryptoWalletsConfig{
			Enabled:      false,
			SIWERequired: false,
		},
		DFPPA: &sdk.Consumer_DFPPAConfig{
			Enabled:              sdk.DFPPASettingEnabled,
			OnChallenge:          sdk.DFPPAOnChallengeActionTriggerCaptcha,
			LookupTimeoutSeconds: 10,
		},
		Biometrics: &sdk.Consumer_BiometricsConfig{
			CreateBiometricsEnabled: true,
			Enabled:                 true,
		},
		Passwords: &sdk.Consumer_PasswordsConfig{
			Enabled:                       true,
			PKCERequiredForPasswordResets: true,
		},
	}
}

func makeTestB2BConfig(t *testing.T) sdk.B2BConfig {
	t.Helper()
	return sdk.B2BConfig{
		Basic: &sdk.B2B_BasicConfig{
			Enabled:                 true,
			CreateNewMembers:        true,
			AllowSelfOnboarding:     true,
			EnableMemberPermissions: true,
			Domains: []sdk.AuthorizedB2BDomain{
				{
					Domain:      "https://example.com",
					SlugPattern: "https://{{slug}}.example.com",
				},
			},
			BundleIDs: []string{"com.example.app"},
		},
		Sessions: &sdk.B2B_SessionsConfig{
			Enabled:                   true,
			MaxSessionDurationMinutes: 60,
		},
		MagicLinks: &sdk.B2B_MagicLinksConfig{
			Enabled:      true,
			PKCERequired: true,
		},
		OAuth: &sdk.B2B_OAuthConfig{
			Enabled:      true,
			PKCERequired: true,
		},
		TOTPs: &sdk.B2B_TOTPsConfig{
			CreateTOTPs: true,
			Enabled:     true,
		},
		SSO: &sdk.B2B_SSOConfig{
			Enabled:      true,
			PKCERequired: false,
		},
		OTPs: &sdk.B2B_OTPsConfig{
			SMSEnabled:          false,
			SMSAutofillMetadata: []sdk.SMSAutofillMetadata{},
		},
		DFPPA: &sdk.B2B_DFPPAConfig{
			Enabled:              sdk.DFPPASettingEnabled,
			OnChallenge:          sdk.DFPPAOnChallengeActionTriggerCaptcha,
			LookupTimeoutSeconds: 10,
		},
		Passwords: &sdk.B2B_PasswordsConfig{
			Enabled:                       false,
			PKCERequiredForPasswordResets: false,
		},
	}
}

func TestSDKClient_GetConsumerConfig(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalConsumer)
	config := makeTestConsumerConfig(t)
	_, err := client.SDK.SetConsumerConfig(context.Background(), sdk.SetConsumerConfigRequest{
		ProjectID: project.LiveProjectID,
		Config:    config,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.SDK.GetConsumerConfig(context.Background(), sdk.GetConsumerConfigRequest{
		ProjectID: project.LiveProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, config, resp.Config)
}

func TestSDKClient_SetConsumerConfig(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalConsumer)
	config := makeTestConsumerConfig(t)

	// Act
	resp, err := client.SDK.SetConsumerConfig(context.Background(), sdk.SetConsumerConfigRequest{
		ProjectID: project.LiveProjectID,
		Config:    config,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, config, resp.Config)
}

func TestSDKClient_GetB2BConfig(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	config := makeTestB2BConfig(t)
	_, err := client.SDK.SetB2BConfig(context.Background(), sdk.SetB2BConfigRequest{
		ProjectID: project.LiveProjectID,
		Config:    config,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.SDK.GetB2BConfig(context.Background(), sdk.GetB2BConfigRequest{
		ProjectID: project.LiveProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, config, resp.Config)
}

func TestSDKClient_SetB2BConfig(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	config := makeTestB2BConfig(t)

	// Act
	resp, err := client.SDK.SetB2BConfig(context.Background(), sdk.SetB2BConfigRequest{
		ProjectID: project.LiveProjectID,
		Config:    config,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, config, resp.Config)
}
