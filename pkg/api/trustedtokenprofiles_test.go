package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/trustedtokenprofiles"
)

func (c *testClient) createTrustedTokenProfileJWK(
	projectID string,
	name string,
	audience string,
	issuer string,
	jwksUrl string,
) string {
	c.t.Helper()
	resp, err := c.TrustedTokenProfiles.Create(context.Background(), &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
		ProjectID:     projectID,
		Name:          name,
		Audience:      audience,
		Issuer:        issuer,
		PublicKeyType: "jwk",
		JwksURL:       ptr(jwksUrl),
	})
	require.NoError(c.t, err)
	return resp.TrustedTokenProfile.ID
}

func (c *testClient) createTrustedTokenProfilePEM(
	projectID string,
	name string,
	audience string,
	issuer string,
	publicKey string,
) string {
	c.t.Helper()
	resp, err := c.TrustedTokenProfiles.Create(context.Background(), &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
		ProjectID:     projectID,
		Name:          name,
		Audience:      audience,
		Issuer:        issuer,
		PublicKeyType: "pem",
		PEMFiles:      []string{publicKey},
	})
	require.NoError(c.t, err)
	return resp.TrustedTokenProfile.ID
}

func (c *testClient) cleanupTrustedTokenProfile(projectID string, profileID string) {
	c.t.Helper()
	c.t.Cleanup(func() {
		_, err := c.TrustedTokenProfiles.Delete(context.Background(), &trustedtokenprofiles.DeleteTrustedTokenProfileRequest{
			ProjectID: projectID,
			ProfileID: profileID,
		})
		require.NoError(c.t, err)
	})
}

func TestTrustedTokenProfilesClient_Create(t *testing.T) {
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)

	t.Run("happy path - with JWK", func(t *testing.T) {
		// Act
		resp, err := client.TrustedTokenProfiles.Create(context.Background(), &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectID:     project.TestProjectID,
			Name:          "Test Profile",
			Audience:      "test-audience",
			Issuer:        "https://test-issuer.com",
			PublicKeyType: "jwk",
			JwksURL:       ptr("https://test-issuer.com/.well-known/jwks.json"),
			AttributeMapping: map[string]interface{}{
				"email": "myemail@example.com",
				"name":  "myname",
			},
		})

		// Assert
		require.NoError(t, err)
		assert.NotEmpty(t, resp.TrustedTokenProfile.ID)
		assert.Equal(t, "Test Profile", resp.TrustedTokenProfile.Name)
		assert.Equal(t, "test-audience", resp.TrustedTokenProfile.Audience)
		assert.Equal(t, "https://test-issuer.com", resp.TrustedTokenProfile.Issuer)
		assert.Equal(t, "jwk", resp.TrustedTokenProfile.PublicKeyType)
		assert.Equal(t, "https://test-issuer.com/.well-known/jwks.json", *resp.TrustedTokenProfile.JwksURL)
		assert.Equal(t, map[string]interface{}{
			"email": "myemail@example.com",
			"name":  "myname",
		}, resp.TrustedTokenProfile.AttributeMapping)
		client.cleanupTrustedTokenProfile(project.TestProjectID, resp.TrustedTokenProfile.ID)
	})

	t.Run("happy path - with PEM", func(t *testing.T) {
		// Act
		resp, err := client.TrustedTokenProfiles.Create(context.Background(), &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectID:     project.TestProjectID,
			Name:          "Test Profile with PEM",
			Audience:      "pem-test-audience",
			Issuer:        "https://pem-test-issuer.com",
			PublicKeyType: "pem",
			PEMFiles: []string{
				"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4f5wg5l2hKsTeNem/V41\nfGnJm6gOdrj8ym3rFkEjWT2btYK36hY+c2QKfPU5O7w=\n-----END PUBLIC KEY-----",
			},
		})

		// Assert
		require.NoError(t, err)
		assert.NotEmpty(t, resp.TrustedTokenProfile.ID)
		assert.Equal(t, "Test Profile with PEM", resp.TrustedTokenProfile.Name)
		assert.Equal(t, "pem-test-audience", resp.TrustedTokenProfile.Audience)
		assert.Equal(t, "https://pem-test-issuer.com", resp.TrustedTokenProfile.Issuer)
		assert.Equal(t, "pem", resp.TrustedTokenProfile.PublicKeyType)
		assert.Equal(t, []string{
			"-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA4f5wg5l2hKsTeNem/V41\nfGnJm6gOdrj8ym3rFkEjWT2btYK36hY+c2QKfPU5O7w=\n-----END PUBLIC KEY-----",
		}, resp.TrustedTokenProfile.PEMFiles)
		client.cleanupTrustedTokenProfile(project.TestProjectID, resp.TrustedTokenProfile.ID)
	})
}

func TestTrustedTokenProfilesClient_Get(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	profileID := client.createTrustedTokenProfileJWK(project.TestProjectID, "Test Profile Get", "test-profile-get-audience", "https://test-profile-get-issuer.com", "https://test-profile-get-issuer.com/.well-known/jwks.json")
	client.cleanupTrustedTokenProfile(project.TestProjectID, profileID)

	// Act
	resp, err := client.TrustedTokenProfiles.Get(context.Background(), &trustedtokenprofiles.GetTrustedTokenProfileRequest{
		ProjectID: project.TestProjectID,
		ProfileID: profileID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, profileID, resp.TrustedTokenProfile.ID)
	assert.Equal(t, "Test Profile Get", resp.TrustedTokenProfile.Name)
	assert.Equal(t, "test-profile-get-audience", resp.TrustedTokenProfile.Audience)
	assert.Equal(t, "https://test-profile-get-issuer.com", resp.TrustedTokenProfile.Issuer)
	assert.Equal(t, "jwk", resp.TrustedTokenProfile.PublicKeyType)
	assert.Equal(t, "https://test-profile-get-issuer.com/.well-known/jwks.json", *resp.TrustedTokenProfile.JwksURL)
}

func TestTrustedTokenProfilesClient_List(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	profileID1 := client.createTrustedTokenProfileJWK(project.TestProjectID, "Test Profile 1", "test-profile-1-audience", "https://test-profile-1-issuer.com", "https://test-profile-1-issuer.com/.well-known/jwks.json")
	profileID2 := client.createTrustedTokenProfileJWK(project.TestProjectID, "Test Profile 2", "test-profile-2-audience", "https://test-profile-2-issuer.com", "https://test-profile-2-issuer.com/.well-known/jwks.json")
	client.cleanupTrustedTokenProfile(project.TestProjectID, profileID1)
	client.cleanupTrustedTokenProfile(project.TestProjectID, profileID2)

	// Act
	resp, err := client.TrustedTokenProfiles.List(context.Background(), &trustedtokenprofiles.ListTrustedTokenProfilesRequest{
		ProjectID: project.TestProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Len(t, resp.TrustedTokenProfiles, 2)

	var profile1, profile2 *trustedtokenprofiles.TrustedTokenProfile
	if resp.TrustedTokenProfiles[0].ID == profileID1 {
		profile1 = &resp.TrustedTokenProfiles[0]
		profile2 = &resp.TrustedTokenProfiles[1]
	} else {
		profile1 = &resp.TrustedTokenProfiles[1]
		profile2 = &resp.TrustedTokenProfiles[0]
	}

	assert.Equal(t, profileID1, profile1.ID)
	assert.Equal(t, "Test Profile 1", profile1.Name)
	assert.Equal(t, "test-profile-1-audience", profile1.Audience)
	assert.Equal(t, "https://test-profile-1-issuer.com", profile1.Issuer)
	assert.Equal(t, "jwk", profile1.PublicKeyType)
	assert.Equal(t, "https://test-profile-1-issuer.com/.well-known/jwks.json", *profile1.JwksURL)

	assert.Equal(t, profileID2, profile2.ID)
	assert.Equal(t, "Test Profile 2", profile2.Name)
	assert.Equal(t, "test-profile-2-audience", profile2.Audience)
	assert.Equal(t, "https://test-profile-2-issuer.com", profile2.Issuer)
	assert.Equal(t, "jwk", profile2.PublicKeyType)
	assert.Equal(t, "https://test-profile-2-issuer.com/.well-known/jwks.json", *profile2.JwksURL)
}

func TestTrustedTokenProfilesClient_Update(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		profileID := client.createTrustedTokenProfileJWK(project.TestProjectID, "Test Profile Update", "test-profile-update-audience", "https://test-profile-update-issuer.com", "https://test-profile-update-issuer.com/.well-known/jwks.json")
		client.cleanupTrustedTokenProfile(project.TestProjectID, profileID)

		// Act
		resp, err := client.TrustedTokenProfiles.Update(context.Background(), &trustedtokenprofiles.UpdateTrustedTokenProfileRequest{
			ProjectID: project.TestProjectID,
			ProfileID: profileID,
			Name:      "Updated Profile",
			Audience:  "updated-profile-audience",
			Issuer:    "https://updated-profile-issuer.com",
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, profileID, resp.TrustedTokenProfile.ID)
		assert.Equal(t, "Updated Profile", resp.TrustedTokenProfile.Name)
		assert.Equal(t, "updated-profile-audience", resp.TrustedTokenProfile.Audience)
		assert.Equal(t, "https://updated-profile-issuer.com", resp.TrustedTokenProfile.Issuer)
		assert.Equal(t, "jwk", resp.TrustedTokenProfile.PublicKeyType)
		assert.Equal(t, "https://test-profile-update-issuer.com/.well-known/jwks.json", *resp.TrustedTokenProfile.JwksURL)
	})
	t.Run("happy path - only some fields updated", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		profileID := client.createTrustedTokenProfileJWK(project.TestProjectID, "Test Profile Update 2", "test-profile-update-2-audience", "https://test-profile-update-2-issuer.com", "https://test-profile-update-2-issuer.com/.well-known/jwks.json")
		client.cleanupTrustedTokenProfile(project.TestProjectID, profileID)

		// Act
		resp, err := client.TrustedTokenProfiles.Update(context.Background(), &trustedtokenprofiles.UpdateTrustedTokenProfileRequest{
			ProjectID: project.TestProjectID,
			ProfileID: profileID,
			Name:      "Updated Profile 3",
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, profileID, resp.TrustedTokenProfile.ID)
		assert.Equal(t, "Updated Profile 3", resp.TrustedTokenProfile.Name)
		assert.Equal(t, "test-profile-update-2-audience", resp.TrustedTokenProfile.Audience)
		assert.Equal(t, "https://test-profile-update-2-issuer.com", resp.TrustedTokenProfile.Issuer)
		assert.Equal(t, "jwk", resp.TrustedTokenProfile.PublicKeyType)
		assert.Equal(t, "https://test-profile-update-2-issuer.com/.well-known/jwks.json", *resp.TrustedTokenProfile.JwksURL)
	})
}

func TestTrustedTokenProfilesClient_Delete(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	profileID := client.createTrustedTokenProfileJWK(project.TestProjectID, "Test Profile Delete", "test-profile-delete-audience", "https://test-profile-delete-issuer.com", "https://test-profile-delete-issuer.com/.well-known/jwks.json")

	// Act
	_, err := client.TrustedTokenProfiles.Delete(context.Background(), &trustedtokenprofiles.DeleteTrustedTokenProfileRequest{
		ProjectID: project.TestProjectID,
		ProfileID: profileID,
	})

	// Assert
	assert.NoError(t, err)
}

func TestTrustedTokenProfilesClient_CreatePEM(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	initialPublicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAKEFBBOCAQ8AAIICCgKCAKAA1234567890abcdef\n-----END PUBLIC KEY-----"
	profileID := client.createTrustedTokenProfilePEM(project.TestProjectID, "Test Profile Create PEM", "test-profile-create-pem-audience", "https://test-profile-create-pem-issuer.com", initialPublicKey)
	client.cleanupTrustedTokenProfile(project.TestProjectID, profileID)

	samplePublicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AAIICCgKCAQEA1234567890abcdef\n-----END PUBLIC KEY-----"

	// Act
	resp, err := client.TrustedTokenProfiles.CreatePEM(context.Background(), &trustedtokenprofiles.CreatePEMFileRequest{
		ProjectID: project.TestProjectID,
		ProfileID: profileID,
		PublicKey: samplePublicKey,
	})

	// Assert
	assert.NoError(t, err)
	assert.NotEmpty(t, resp.PEMFile.ID)
	assert.Equal(t, samplePublicKey, resp.PEMFile.PublicKey)
}

func TestTrustedTokenProfilesClient_DeletePEM(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	initialPublicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAKEFBBOCAQ8AAIICCgKCAKAA1234567890abcdef\n-----END PUBLIC KEY-----"
	profileID := client.createTrustedTokenProfilePEM(project.TestProjectID, "Test Profile Delete PEM", "test-profile-delete-pem-audience", "https://test-profile-delete-pem-issuer.com", initialPublicKey)
	client.cleanupTrustedTokenProfile(project.TestProjectID, profileID)

	samplePublicKey := "-----BEGIN PUBLIC KEY-----\nMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AAIICCgKCAKAA1234567890abcdef\n-----END PUBLIC KEY-----"

	createResp, err := client.TrustedTokenProfiles.CreatePEM(context.Background(), &trustedtokenprofiles.CreatePEMFileRequest{
		ProjectID: project.TestProjectID,
		ProfileID: profileID,
		PublicKey: samplePublicKey,
	})
	require.NoError(t, err)

	// Act
	_, err = client.TrustedTokenProfiles.DeletePEM(context.Background(), &trustedtokenprofiles.DeletePEMFileRequest{
		ProjectID: project.TestProjectID,
		ProfileID: profileID,
		PEMFileID: createResp.PEMFile.ID,
	})

	// Assert
	assert.NoError(t, err)
}
