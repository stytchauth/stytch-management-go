package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/trustedtokenprofiles"
)

const (
	// Sample JWK for testing
	testJWKSURL = "https://example.com/.well-known/jwks.json"

	// Sample PEM public key for testing
	testPEMKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890abcdef
-----END PUBLIC KEY-----`
)

func TestTrustedTokenProfilesClient_Create(t *testing.T) {
	t.Run("create with JWK", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		jwksURL := testJWKSURL

		// Act
		resp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "Test JWK Profile",
			Audience:        "test-audience",
			Issuer:          "test-issuer",
			JwksURL:         &jwksURL,
			PublicKeyType:   "jwk",
			PEMFiles:        []string{},
			CanJITProvision: true,
			AttributeMapping: map[string]interface{}{
				"user_id": "sub",
				"email":   "email",
			},
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "Test JWK Profile", resp.TrustedTokenProfile.Name)
		assert.Equal(t, "test-audience", resp.TrustedTokenProfile.Audience)
		assert.Equal(t, "test-issuer", resp.TrustedTokenProfile.Issuer)
		assert.Equal(t, testJWKSURL, resp.TrustedTokenProfile.JwksURL)
		assert.Equal(t, "jwk", resp.TrustedTokenProfile.PublicKeyType)
		assert.NotEmpty(t, resp.TrustedTokenProfile.ID)
		assert.True(t, resp.TrustedTokenProfile.CanJITProvision)
	})

	t.Run("create with PEM files", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "Test PEM Profile",
			Audience:        "test-audience-pem",
			Issuer:          "test-issuer-pem",
			PublicKeyType:   "pem",
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: false,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "Test PEM Profile", resp.TrustedTokenProfile.Name)
		assert.Equal(t, "test-audience-pem", resp.TrustedTokenProfile.Audience)
		assert.Equal(t, "test-issuer-pem", resp.TrustedTokenProfile.Issuer)
		assert.Equal(t, "pem", resp.TrustedTokenProfile.PublicKeyType)
		assert.NotEmpty(t, resp.TrustedTokenProfile.ID)
		assert.False(t, resp.TrustedTokenProfile.CanJITProvision)
		assert.Len(t, resp.TrustedTokenProfile.PEMFiles, 1)
	})
}

func TestTrustedTokenProfilesClient_Get(t *testing.T) {
	t.Run("get existing profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		jwksURL := testJWKSURL
		createResp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "Test Get Profile",
			Audience:        "get-test-audience",
			Issuer:          "get-test-issuer.com",
			JwksURL:         &jwksURL,
			PublicKeyType:   "jwk",
			PEMFiles:        []string{},
			CanJITProvision: false,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.Get(ctx, &trustedtokenprofiles.GetTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       createResp.TrustedTokenProfile.ID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, createResp.TrustedTokenProfile.ID, resp.TrustedTokenProfile.ID)
		assert.Equal(t, "Test Get Profile", resp.TrustedTokenProfile.Name)
		assert.Equal(t, "get-test-audience", resp.TrustedTokenProfile.Audience)
		assert.Equal(t, "get-test-issuer.com", resp.TrustedTokenProfile.Issuer)
	})

	t.Run("get non-existent profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.TrustedTokenProfiles.Get(ctx, &trustedtokenprofiles.GetTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       "non-existent-profile-id",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("missing profile ID", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.TrustedTokenProfiles.Get(ctx, &trustedtokenprofiles.GetTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			// ProfileID is intentionally omitted.
		})

		// Assert
		assert.ErrorContains(t, err, "profile ID")
		assert.Nil(t, resp)
	})
}

func TestTrustedTokenProfilesClient_GetAll(t *testing.T) {
	t.Run("get all profiles", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		jwksURL := testJWKSURL

		profile1, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "List Test Profile 1",
			Audience:        "list-test-audience-1",
			Issuer:          "list-test-issuer-1",
			JwksURL:         &jwksURL,
			PublicKeyType:   "jwk",
			PEMFiles:        []string{},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		profile2, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "List Test Profile 2",
			Audience:        "list-test-audience-2",
			Issuer:          "list-test-issuer-2",
			PublicKeyType:   "pem",
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: false,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.GetAll(ctx, &trustedtokenprofiles.GetAllTrustedTokenProfilesRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.TrustedTokenProfiles, 2)

		var foundProfile1, foundProfile2 bool
		for _, profile := range resp.TrustedTokenProfiles {
			if profile.ID == profile1.TrustedTokenProfile.ID {
				foundProfile1 = true
				assert.Equal(t, "List Test Profile 1", profile.Name)
			}
			if profile.ID == profile2.TrustedTokenProfile.ID {
				foundProfile2 = true
				assert.Equal(t, "List Test Profile 2", profile.Name)
			}
		}
		assert.True(t, foundProfile1)
		assert.True(t, foundProfile2)
	})
}

func TestTrustedTokenProfilesClient_Update(t *testing.T) {
	t.Run("update profile name and audience", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		jwksURL := testJWKSURL
		createResp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "Update Test Profile",
			Audience:        "update-test-audience",
			Issuer:          "update-test-issuer",
			JwksURL:         &jwksURL,
			PublicKeyType:   "jwk",
			PEMFiles:        []string{},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		updatedName := "Updated Profile Name"
		updatedAudience := "updated-audience"
		updatedJITProvision := false

		// Act
		resp, err := client.TrustedTokenProfiles.Update(ctx, &trustedtokenprofiles.UpdateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       createResp.TrustedTokenProfile.ID,
			Name:            &updatedName,
			Audience:        &updatedAudience,
			CanJITProvision: &updatedJITProvision,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, createResp.TrustedTokenProfile.ID, resp.TrustedTokenProfile.ID)
		assert.Equal(t, "Updated Profile Name", resp.TrustedTokenProfile.Name)
		assert.Equal(t, "updated-audience", resp.TrustedTokenProfile.Audience)
		assert.Equal(t, "update-test-issuer", resp.TrustedTokenProfile.Issuer) // Should remain unchanged
	})
}

func TestTrustedTokenProfilesClient_Delete(t *testing.T) {
	t.Run("delete existing profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		jwksURL := testJWKSURL
		createResp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "Delete Test Profile",
			Audience:        "delete-test-audience",
			Issuer:          "delete-test-issuer",
			JwksURL:         &jwksURL,
			PublicKeyType:   "jwk",
			PEMFiles:        []string{},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.Delete(ctx, &trustedtokenprofiles.DeleteTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       createResp.TrustedTokenProfile.ID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Verify profile is deleted
		getResp, err := client.TrustedTokenProfiles.Get(ctx, &trustedtokenprofiles.GetTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       createResp.TrustedTokenProfile.ID,
		})
		assert.Error(t, err)
		assert.Nil(t, getResp)
	})
}

func TestTrustedTokenProfilesClient_CreatePEM(t *testing.T) {
	t.Run("create PEM file for existing PEM profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Create a profile first with initial PEM file
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "PEM Test Profile",
			Audience:        "pem-test-audience",
			Issuer:          "pem-test-issuer",
			PublicKeyType:   "pem",
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Act
		// Add another PEM file to the profile
		resp, err := client.TrustedTokenProfiles.CreatePEM(ctx, &trustedtokenprofiles.CreatePEMFileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       profileResp.TrustedTokenProfile.ID,
			PublicKey:       testPEMKey,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.PEMFile.ID)
		assert.Equal(t, testPEMKey, resp.PEMFile.PublicKey)
	})

	t.Run("create PEM file for JWK profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		jwksURL := testJWKSURL
		// Create a JWK profile first
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "JWK to PEM Test Profile",
			Audience:        "jwk-pem-test-audience",
			Issuer:          "jwk-pem-test-issuer",
			JwksURL:         &jwksURL,
			PublicKeyType:   "jwk",
			PEMFiles:        []string{},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Act
		// Try to add a PEM file to the JWK profile
		resp, err := client.TrustedTokenProfiles.CreatePEM(ctx, &trustedtokenprofiles.CreatePEMFileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       profileResp.TrustedTokenProfile.ID,
			PublicKey:       testPEMKey,
		})

		// Arrange
		// This should fail since the trusted token profile key is not compatible
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("create PEM file for non-existent profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Act
		resp, err := client.TrustedTokenProfiles.CreatePEM(ctx, &trustedtokenprofiles.CreatePEMFileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       "non-existent-profile-id",
			PublicKey:       testPEMKey,
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestTrustedTokenProfilesClient_GetPEM(t *testing.T) {
	t.Run("get existing PEM file", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Create a profile first with initial PEM file
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "PEM Get Test Profile",
			Audience:        "pem-get-test-audience",
			Issuer:          "pem-get-test-issuer",
			PublicKeyType:   "pem",
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Use the PEM file that was created with the profile
		pemFileID := profileResp.TrustedTokenProfile.PEMFiles[0].ID

		// Act
		resp, err := client.TrustedTokenProfiles.GetPEM(ctx, &trustedtokenprofiles.GetPEMFileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       profileResp.TrustedTokenProfile.ID,
			PEMFileID:       pemFileID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, pemFileID, resp.PEMFile.ID)
		assert.Equal(t, testPEMKey, resp.PEMFile.PublicKey)
	})

	t.Run("get non-existent PEM file", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Create a profile first with initial PEM file
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "PEM Get Test Profile",
			Audience:        "pem-get-test-audience",
			Issuer:          "pem-get-test-issuer",
			PublicKeyType:   "pem",
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.GetPEM(ctx, &trustedtokenprofiles.GetPEMFileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       profileResp.TrustedTokenProfile.ID,
			PEMFileID:       "non-existent-pem-file-id",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})
}

func TestTrustedTokenProfilesClient_DeletePEM(t *testing.T) {
	t.Run("delete existing PEM file", func(t *testing.T) {
		client := NewTestClient(t)
		project := client.DisposableProject(projects.VerticalB2B)
		ctx := context.Background()

		// Create a profile first with two PEM files so we can delete one
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, &trustedtokenprofiles.CreateTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			Name:            "PEM Delete Test Profile",
			Audience:        "pem-delete-test-audience",
			Issuer:          "pem-delete-test-issuer",
			PublicKeyType:   "pem",
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Add another PEM file so we can delete one
		createPEMResp, err := client.TrustedTokenProfiles.CreatePEM(ctx, &trustedtokenprofiles.CreatePEMFileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       profileResp.TrustedTokenProfile.ID,
			PublicKey:       testPEMKey,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.DeletePEM(ctx, &trustedtokenprofiles.DeletePEMFileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       profileResp.TrustedTokenProfile.ID,
			PEMFileID:       createPEMResp.PEMFile.ID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Assert
		// Verify PEM file is deleted
		getProfileResp, err := client.TrustedTokenProfiles.Get(ctx, &trustedtokenprofiles.GetTrustedTokenProfileRequest{
			ProjectSlug:     project.ProjectSlug,
			EnvironmentSlug: TestEnvironment,
			ProfileID:       profileResp.TrustedTokenProfile.ID,
		})
		require.NoError(t, err)

		assert.Len(t, getProfileResp.TrustedTokenProfile.PEMFiles, 1)
		assert.NotEqual(t, createPEMResp.PEMFile.ID, getProfileResp.TrustedTokenProfile.PEMFiles[0].ID)
	})
}
