package api_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
// 	"github.com/stytchauth/stytch-management-go/v3/pkg/models/sdk"
// )

// func makeTestConsumerConfig(t *testing.T) sdk.ConsumerConfig {
// 	t.Helper()
// 	return sdk.ConsumerConfig{
// 		Basic: &sdk.ConsumerBasicConfig{
// 			Enabled:   true,
// 			Domains:   []string{"https://example.com"},
// 			BundleIDs: []string{"com.example.app"},
// 		},
// 		Sessions: &sdk.ConsumerSessionsConfig{
// 			MaxSessionDurationMinutes: 60,
// 		},
// 		MagicLinks: &sdk.ConsumerMagicLinksConfig{
// 			LoginOrCreateEnabled: true,
// 			SendEnabled:          false,
// 			PKCERequired:         false,
// 		},
// 		OTPs: &sdk.ConsumerOTPsConfig{
// 			SMSLoginOrCreateEnabled:      true,
// 			WhatsAppLoginOrCreateEnabled: false,
// 			EmailLoginOrCreateEnabled:    false,
// 			SMSSendEnabled:               false,
// 			WhatsAppSendEnabled:          false,
// 			EmailSendEnabled:             false,
// 			SMSAutofillMetadata:          []sdk.SMSAutofillMetadata{},
// 		},
// 		OAuth: &sdk.ConsumerOAuthConfig{
// 			Enabled:      true,
// 			PKCERequired: true,
// 		},
// 		TOTPs: &sdk.ConsumerTOTPsConfig{
// 			CreateTOTPs: true,
// 			Enabled:     true,
// 		},
// 		WebAuthn: &sdk.ConsumerWebAuthnConfig{
// 			CreateWebAuthns: false,
// 			Enabled:         false,
// 		},
// 		CryptoWallets: &sdk.ConsumerCryptoWalletsConfig{
// 			Enabled:      false,
// 			SIWERequired: false,
// 		},
// 		// This cannot be modified beyond defaults
// 		// unless the project uses DFPPA
// 		DFPPA: &sdk.ConsumerDFPPAConfig{
// 			Enabled:     sdk.DFPPASettingDisabled,
// 			OnChallenge: sdk.DFPPAOnChallengeActionAllow,
// 		},
// 		Biometrics: &sdk.ConsumerBiometricsConfig{
// 			CreateBiometricsEnabled: true,
// 			Enabled:                 true,
// 		},
// 		Passwords: &sdk.ConsumerPasswordsConfig{
// 			Enabled:                       true,
// 			PKCERequiredForPasswordResets: true,
// 		},
// 		Cookies: &sdk.ConsumerCookiesConfig{
// 			// This can only be Disabled unless the project has
// 			// CNAMEs configured
// 			HttpOnlyCookies: sdk.HttpOnlyCookiesSettingDisabled,
// 		},
// 	}
// }

// func makeTestB2BConfig(t *testing.T) sdk.B2BConfig {
// 	t.Helper()
// 	return sdk.B2BConfig{
// 		Basic: &sdk.B2BBasicConfig{
// 			Enabled:                 true,
// 			AllowSelfOnboarding:     true,
// 			EnableMemberPermissions: true,
// 			Domains: []sdk.AuthorizedB2BDomain{
// 				{
// 					Domain:      "https://example.com",
// 					SlugPattern: "https://{{slug}}.example.com",
// 				},
// 			},
// 			BundleIDs: []string{"com.example.app"},
// 		},
// 		Sessions: &sdk.B2BSessionsConfig{
// 			MaxSessionDurationMinutes: 60,
// 		},
// 		MagicLinks: &sdk.B2BMagicLinksConfig{
// 			Enabled:      true,
// 			PKCERequired: true,
// 		},
// 		OAuth: &sdk.B2BOAuthConfig{
// 			Enabled:      true,
// 			PKCERequired: true,
// 		},
// 		TOTPs: &sdk.B2BTOTPsConfig{
// 			CreateTOTPs: true,
// 			Enabled:     true,
// 		},
// 		SSO: &sdk.B2BSSOConfig{
// 			Enabled:      true,
// 			PKCERequired: false,
// 		},
// 		OTPs: &sdk.B2BOTPsConfig{
// 			SMSEnabled:          false,
// 			SMSAutofillMetadata: []sdk.SMSAutofillMetadata{},
// 			EmailEnabled:        false,
// 		},
// 		// These cannot be modified beyond defaults
// 		// unless the project is using DFPPA
// 		DFPPA: &sdk.B2BDFPPAConfig{
// 			Enabled:     sdk.DFPPASettingDisabled,
// 			OnChallenge: sdk.DFPPAOnChallengeActionAllow,
// 		},
// 		Passwords: &sdk.B2BPasswordsConfig{
// 			Enabled:                       false,
// 			PKCERequiredForPasswordResets: false,
// 		},
// 		Cookies: &sdk.B2BCookiesConfig{
// 			// This can only be Disabled unless the project has
// 			// CNAMEs configured
// 			HttpOnlyCookies: sdk.HttpOnlyCookiesSettingDisabled,
// 		},
// 	}
// }

// func TestSDKClient_GetConsumerConfig(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	project := client.DisposableProject(projects.VerticalConsumer)
// 	config := makeTestConsumerConfig(t)
// 	_, err := client.SDK.SetConsumerConfig(context.Background(), sdk.SetConsumerConfigRequest{
// 		ProjectID: project.LiveProjectID,
// 		Config:    config,
// 	})
// 	require.NoError(t, err)

// 	// Act
// 	resp, err := client.SDK.GetConsumerConfig(context.Background(), sdk.GetConsumerConfigRequest{
// 		ProjectID: project.LiveProjectID,
// 	})

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.Equal(t, config, resp.Config)
// }

// func TestSDKClient_SetConsumerConfig(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	project := client.DisposableProject(projects.VerticalConsumer)
// 	config := makeTestConsumerConfig(t)

// 	// Act
// 	resp, err := client.SDK.SetConsumerConfig(context.Background(), sdk.SetConsumerConfigRequest{
// 		ProjectID: project.LiveProjectID,
// 		Config:    config,
// 	})

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.Equal(t, config, resp.Config)
// }

// func TestSDKClient_GetB2BConfig(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	project := client.DisposableProject(projects.VerticalB2B)
// 	config := makeTestB2BConfig(t)
// 	_, err := client.SDK.SetB2BConfig(context.Background(), sdk.SetB2BConfigRequest{
// 		ProjectID: project.LiveProjectID,
// 		Config:    config,
// 	})
// 	require.NoError(t, err)

// 	// Act
// 	resp, err := client.SDK.GetB2BConfig(context.Background(), sdk.GetB2BConfigRequest{
// 		ProjectID: project.LiveProjectID,
// 	})

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.Equal(t, config, resp.Config)
// }

// func TestSDKClient_SetB2BConfig(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	project := client.DisposableProject(projects.VerticalB2B)
// 	config := makeTestB2BConfig(t)

// 	// Act
// 	resp, err := client.SDK.SetB2BConfig(context.Background(), sdk.SetB2BConfigRequest{
// 		ProjectID: project.LiveProjectID,
// 		Config:    config,
// 	})

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.Equal(t, config, resp.Config)
// }
