package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
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
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		jwksURL := testJWKSURL

		// Act
		resp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "Test JWK Profile",
			Audience:        "test-audience",
			Issuer:          "test-issuer",
			JWKSURL:         &jwksURL,
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypeJwk,
			PEMFiles:        []string{},
			CanJITProvision: true,
			AttributeMapping: &map[string]any{
				"user_id": "sub",
				"email":   "email",
			},
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "Test JWK Profile", resp.Profile.Name)
		assert.Equal(t, "test-audience", resp.Profile.Audience)
		assert.Equal(t, "test-issuer", resp.Profile.Issuer)
		assert.Equal(t, testJWKSURL, *resp.Profile.JWKSURL)
		assert.Equal(t, trustedtokenprofiles.PublicKeyTypeJwk, resp.Profile.PublicKeyType)
		assert.NotEmpty(t, resp.Profile.ProfileID)
		assert.True(t, resp.Profile.CanJITProvision)
	})

	t.Run("create with PEM files", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "Test PEM Profile",
			Audience:        "test-audience-pem",
			Issuer:          "test-issuer-pem",
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypePem,
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: false,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, "Test PEM Profile", resp.Profile.Name)
		assert.Equal(t, "test-audience-pem", resp.Profile.Audience)
		assert.Equal(t, "test-issuer-pem", resp.Profile.Issuer)
		assert.Equal(t, trustedtokenprofiles.PublicKeyTypePem, resp.Profile.PublicKeyType)
		assert.NotEmpty(t, resp.Profile.ProfileID)
		assert.False(t, resp.Profile.CanJITProvision)
		assert.Len(t, resp.Profile.PEMFiles, 1)
	})
}

func TestTrustedTokenProfilesClient_Get(t *testing.T) {
	t.Run("get existing profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		jwksURL := testJWKSURL
		createResp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "Test Get Profile",
			Audience:        "get-test-audience",
			Issuer:          "get-test-issuer.com",
			JWKSURL:         &jwksURL,
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypeJwk,
			PEMFiles:        []string{},
			CanJITProvision: false,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.Get(ctx, trustedtokenprofiles.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       createResp.Profile.ProfileID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, createResp.Profile.ProfileID, resp.Profile.ProfileID)
		assert.Equal(t, "Test Get Profile", resp.Profile.Name)
		assert.Equal(t, "get-test-audience", resp.Profile.Audience)
		assert.Equal(t, "get-test-issuer.com", resp.Profile.Issuer)
	})

	t.Run("get non-existent profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.TrustedTokenProfiles.Get(ctx, trustedtokenprofiles.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       "non-existent-profile-id",
		})

		// Assert
		assert.Error(t, err)
		assert.Nil(t, resp)
	})

	t.Run("missing profile ID", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.TrustedTokenProfiles.Get(ctx, trustedtokenprofiles.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			// ProfileID is intentionally omitted.
		})

		// Assert
		assert.ErrorContains(t, err, "cannot be empty")
		assert.Nil(t, resp)
	})
}

func TestTrustedTokenProfilesClient_GetAll(t *testing.T) {
	t.Run("get all profiles", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		jwksURL := testJWKSURL

		profile1, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "List Test Profile 1",
			Audience:        "list-test-audience-1",
			Issuer:          "list-test-issuer-1",
			JWKSURL:         &jwksURL,
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypeJwk,
			PEMFiles:        []string{},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		profile2, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "List Test Profile 2",
			Audience:        "list-test-audience-2",
			Issuer:          "list-test-issuer-2",
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypePem,
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: false,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.GetAll(ctx, trustedtokenprofiles.GetAllRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Profiles, 2)

		var foundProfile1, foundProfile2 bool
		for _, profile := range resp.Profiles {
			if profile.ProfileID == profile1.Profile.ProfileID {
				foundProfile1 = true
				assert.Equal(t, "List Test Profile 1", profile.Name)
			}
			if profile.ProfileID == profile2.Profile.ProfileID {
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
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		jwksURL := testJWKSURL
		createResp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "Update Test Profile",
			Audience:        "update-test-audience",
			Issuer:          "update-test-issuer",
			JWKSURL:         &jwksURL,
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypeJwk,
			PEMFiles:        []string{},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		updatedName := "Updated Profile Name"
		updatedAudience := "updated-audience"
		updatedJITProvision := false

		// Act
		resp, err := client.TrustedTokenProfiles.Update(ctx, trustedtokenprofiles.UpdateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       createResp.Profile.ProfileID,
			Name:            &updatedName,
			Audience:        &updatedAudience,
			CanJITProvision: &updatedJITProvision,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, createResp.Profile.ProfileID, resp.Profile.ProfileID)
		assert.Equal(t, "Updated Profile Name", resp.Profile.Name)
		assert.Equal(t, "updated-audience", resp.Profile.Audience)
		assert.Equal(t, "update-test-issuer", resp.Profile.Issuer) // Should remain unchanged
	})
}

func TestTrustedTokenProfilesClient_Delete(t *testing.T) {
	t.Run("delete existing profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		jwksURL := testJWKSURL
		createResp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "Delete Test Profile",
			Audience:        "delete-test-audience",
			Issuer:          "delete-test-issuer",
			JWKSURL:         &jwksURL,
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypeJwk,
			PEMFiles:        []string{},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.Delete(ctx, trustedtokenprofiles.DeleteRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       createResp.Profile.ProfileID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Verify profile is deleted
		getResp, err := client.TrustedTokenProfiles.Get(ctx, trustedtokenprofiles.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       createResp.Profile.ProfileID,
		})
		assert.Error(t, err)
		assert.Nil(t, getResp)
	})
}

func TestTrustedTokenProfilesClient_CreatePEM(t *testing.T) {
	t.Run("create PEM file for existing PEM profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Create a profile first with initial PEM file
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "PEM Test Profile",
			Audience:        "pem-test-audience",
			Issuer:          "pem-test-issuer",
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypePem,
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Act
		// Add another PEM file to the profile
		resp, err := client.TrustedTokenProfiles.CreatePEMFile(ctx, trustedtokenprofiles.CreatePEMFileRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       profileResp.Profile.ProfileID,
			PublicKey:       testPEMKey,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.NotEmpty(t, resp.PEMFile.PEMFileID)
		assert.Equal(t, testPEMKey, resp.PEMFile.PublicKey)
	})

	t.Run("create PEM file for JWK profile", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		jwksURL := testJWKSURL
		// Create a JWK profile first
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "JWK to PEM Test Profile",
			Audience:        "jwk-pem-test-audience",
			Issuer:          "jwk-pem-test-issuer",
			JWKSURL:         &jwksURL,
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypeJwk,
			PEMFiles:        []string{},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Act
		// Try to add a PEM file to the JWK profile
		resp, err := client.TrustedTokenProfiles.CreatePEMFile(ctx, trustedtokenprofiles.CreatePEMFileRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       profileResp.Profile.ProfileID,
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
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.TrustedTokenProfiles.CreatePEMFile(ctx, trustedtokenprofiles.CreatePEMFileRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
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
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Create a profile first with initial PEM file
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "PEM Get Test Profile",
			Audience:        "pem-get-test-audience",
			Issuer:          "pem-get-test-issuer",
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypePem,
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Use the PEM file that was created with the profile
		pemFileID := profileResp.Profile.PEMFiles[0].PEMFileID

		// Act
		resp, err := client.TrustedTokenProfiles.GetPEMFile(ctx, trustedtokenprofiles.GetPEMFileRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       profileResp.Profile.ProfileID,
			PEMFileID:       pemFileID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Equal(t, pemFileID, resp.PEMFile.PEMFileID)
		assert.Equal(t, testPEMKey, resp.PEMFile.PublicKey)
	})

	t.Run("get non-existent PEM file", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Create a profile first with initial PEM file
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "PEM Get Test Profile",
			Audience:        "pem-get-test-audience",
			Issuer:          "pem-get-test-issuer",
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypePem,
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.GetPEMFile(ctx, trustedtokenprofiles.GetPEMFileRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       profileResp.Profile.ProfileID,
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
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Create a profile first with two PEM files so we can delete one
		profileResp, err := client.TrustedTokenProfiles.Create(ctx, trustedtokenprofiles.CreateRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			Name:            "PEM Delete Test Profile",
			Audience:        "pem-delete-test-audience",
			Issuer:          "pem-delete-test-issuer",
			PublicKeyType:   trustedtokenprofiles.PublicKeyTypePem,
			PEMFiles:        []string{testPEMKey},
			CanJITProvision: true,
		})
		require.NoError(t, err)

		// Add another PEM file so we can delete one
		createPEMResp, err := client.TrustedTokenProfiles.CreatePEMFile(ctx, trustedtokenprofiles.CreatePEMFileRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       profileResp.Profile.ProfileID,
			PublicKey:       testPEMKey,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.TrustedTokenProfiles.DeletePEMFile(ctx, trustedtokenprofiles.DeletePEMFileRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       profileResp.Profile.ProfileID,
			PEMFileID:       createPEMResp.PEMFile.PEMFileID,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotNil(t, resp)

		// Assert
		// Verify PEM file is deleted
		getProfileResp, err := client.TrustedTokenProfiles.Get(ctx, trustedtokenprofiles.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			ProfileID:       profileResp.Profile.ProfileID,
		})
		require.NoError(t, err)

		assert.Len(t, getProfileResp.Profile.PEMFiles, 1)
		assert.NotEqual(t, createPEMResp.PEMFile.PEMFileID, getProfileResp.Profile.PEMFiles[0].PEMFileID)
	})
}
