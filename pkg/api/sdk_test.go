package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/sdk"
)

func getTestConfig(t *testing.T) sdk.Config {
	t.Helper()
	return sdk.Config{
		ManageUserData:                 true,
		ManageSessionData:              true,
		EmailMagicLinks:                sdk.AuthSettingAlways,
		SMSOTPs:                        sdk.AuthSettingSecondaryOnly,
		WhatsappOTPs:                   sdk.AuthSettingDisabled,
		EmailOTPs:                      sdk.AuthSettingSecondaryOnly,
		OAuth:                          sdk.AuthSettingPrimaryOnly,
		CreateTOTPEnabled:              true,
		TOTPs:                          sdk.AuthSettingDisabled,
		CreateWebauthnEnabled:          true,
		Webauthns:                      sdk.AuthSettingDisabled,
		CreateNewUsers:                 true,
		CryptoWallets:                  sdk.AuthSettingAlways,
		MaxSessionDurationMinutes:      60,
		PKCERequiredForEmailMagicLinks: true,
		PKCERequiredForOAuth:           true,
		PKCERequiredForPasswordResets:  true,
		Passwords:                      sdk.AuthSettingAlways,
		CreateBiometricsEnabled:        true,
		Biometrics:                     sdk.AuthSettingDisabled,
		EmailMagicLinksSend:            sdk.AuthSettingAlways,
		SMSOTPsSend:                    sdk.AuthSettingAlways,
		WhatsappOTPsSend:               sdk.AuthSettingDisabled,
		EmailOTPsSend:                  sdk.AuthSettingAlways,
		SSO:                            sdk.AuthSettingDisabled,
		Domains:                        []string{},
		BundleIDs:                      []string{},
		B2BDomains:                     []sdk.AuthorizedB2BDomain{},
		DFPProtectedAuthEnabled:        sdk.DFPProtectedAuthDisabled,
		DFPProtectedAuthOnChallenge:    sdk.DFPProtectedAuthChallengeSettingAllow,
	}
}

func TestSDKClient_GetConfig(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	config := getTestConfig(t)
	_, err := client.SDK.SetConfig(context.Background(), sdk.SetConfigRequest{
		ProjectID: project.ProjectID,
		Config:    config,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.SDK.GetConfig(context.Background(), sdk.GetConfigRequest{
		ProjectID: project.ProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, config, resp.Config)
}

func TestSDKClient_SetConfig(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	config := getTestConfig(t)

	// Act
	_, err := client.SDK.SetConfig(context.Background(), sdk.SetConfigRequest{
		ProjectID: project.ProjectID,
		Config:    config,
	})
	getResp, getErr := client.SDK.GetConfig(context.Background(), sdk.GetConfigRequest{
		ProjectID: project.ProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.NoError(t, getErr)
	assert.Equal(t, config, getResp.Config)
}
