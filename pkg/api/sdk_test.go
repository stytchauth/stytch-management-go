package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/sdk"
)

func makeTestConfig(t *testing.T) sdk.Config {
	t.Helper()
	return sdk.Config{
		AllowSelfOnboarding:                  true,
		B2BDomains:                           []sdk.AuthorizedB2BDomain{},
		Biometrics:                           sdk.AuthSettingSecondaryOnly,
		BundleIDs:                            []string{},
		CreateBiometricsEnabled:              true,
		CreateNewUsers:                       true,
		CreateTOTPEnabled:                    true,
		CreateWebauthnEnabled:                true,
		CryptoWallets:                        sdk.AuthSettingAlways,
		DFPProtectedAuthEnabled:              sdk.DFPProtectedAuthDisabled,
		DFPProtectedAuthLookupTimeoutSeconds: 10,
		DFPProtectedAuthOnChallenge:          sdk.DFPProtectedAuthChallengeSettingAllow,
		Domains:                              []string{"http://localhost:3000", "http://localhost:3001"},
		EmailMagicLinks:                      sdk.AuthSettingAlways,
		EmailMagicLinksSend:                  sdk.AuthSettingAlways,
		EmailOTPs:                            sdk.AuthSettingAlways,
		EmailOTPsSend:                        sdk.AuthSettingAlways,
		EnableB2BUseMemberPermissions:        true,
		ManageSessionData:                    true,
		ManageUserData:                       true,
		MaxSessionDurationMinutes:            60,
		OAuth:                                sdk.AuthSettingAlways,
		Passwords:                            sdk.AuthSettingAlways,
		SMSAutofillMetadata:                  []sdk.SMSAutofillMetadata{},
		SMSOTPs:                              sdk.AuthSettingAlways,
		SMSOTPsSend:                          sdk.AuthSettingAlways,
		SSO:                                  sdk.AuthSettingAlways,
		TOTPs:                                sdk.AuthSettingSecondaryOnly,
		Webauthns:                            sdk.AuthSettingSecondaryOnly,
		WhatsappOTPs:                         sdk.AuthSettingAlways,
		WhatsappOTPsSend:                     sdk.AuthSettingAlways,
	}
}

func TestSDKClient_GetConfig(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	config := makeTestConfig(t)
	_, err := client.SDK.SetConfig(context.Background(), sdk.SetConfigRequest{
		ProjectID: project.LiveProject.ID,
		Config:    config,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.SDK.GetConfig(context.Background(), sdk.GetConfigRequest{
		ProjectID: project.LiveProject.ID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, config, resp.Config)
}

func TestSDKClient_SetConfig(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	config := makeTestConfig(t)

	// Act
	_, err := client.SDK.SetConfig(context.Background(), sdk.SetConfigRequest{
		ProjectID: project.LiveProject.ID,
		Config:    config,
	})
	getResp, getErr := client.SDK.GetConfig(context.Background(), sdk.GetConfigRequest{
		ProjectID: project.LiveProject.ID,
	})

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, getErr)
	assert.Equal(t, config, getResp.Config)
}
