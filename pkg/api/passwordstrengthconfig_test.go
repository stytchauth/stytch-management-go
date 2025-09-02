package api_test

// import (
// 	"context"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"github.com/stytchauth/stytch-management-go/v3/pkg/models/passwordstrengthconfig"
// 	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
// )

// func Test_PasswordStrengthConfigGet(t *testing.T) {
// 	// Arrange
// 	client := NewTestClient(t)
// 	ctx := context.Background()
// 	project := client.DisposableProject(projects.VerticalB2B)
// 	_, err := client.PasswordStrengthConfig.Set(ctx, passwordstrengthconfig.SetRequest{
// 		ProjectID: project.LiveProjectID,
// 		PasswordStrengthConfig: passwordstrengthconfig.PasswordStrengthConfig{
// 			CheckBreachOnCreation:       true,
// 			CheckBreachOnAuthentication: true,
// 			ValidateOnAuthentication:    true,
// 			ValidationPolicy:            passwordstrengthconfig.ValidationPolicyLUDS,
// 			LudsMinPasswordLength:       12,
// 			LudsMinPasswordComplexity:   3,
// 		},
// 	})
// 	require.NoError(t, err)

// 	// Act
// 	resp, err := client.PasswordStrengthConfig.Get(ctx, passwordstrengthconfig.GetRequest{
// 		ProjectID: project.LiveProjectID,
// 	})

// 	// Assert
// 	assert.NoError(t, err)
// 	assert.True(t, resp.PasswordStrengthConfig.CheckBreachOnCreation)
// 	assert.Equal(t, 12, resp.PasswordStrengthConfig.LudsMinPasswordLength)
// }

// func Test_PasswordStrengthConfigSet(t *testing.T) {
// 	t.Run("luds", func(t *testing.T) {
// 		// Arrange
// 		client := NewTestClient(t)
// 		project := client.DisposableProject(projects.VerticalB2B)
// 		ctx := context.Background()

// 		// Act
// 		resp, err := client.PasswordStrengthConfig.Set(ctx, passwordstrengthconfig.SetRequest{
// 			ProjectID: project.LiveProjectID,
// 			PasswordStrengthConfig: passwordstrengthconfig.PasswordStrengthConfig{
// 				CheckBreachOnCreation:       true,
// 				CheckBreachOnAuthentication: true,
// 				ValidateOnAuthentication:    true,
// 				ValidationPolicy:            passwordstrengthconfig.ValidationPolicyLUDS,
// 				LudsMinPasswordLength:       12,
// 				LudsMinPasswordComplexity:   3,
// 			},
// 		})

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.True(t, resp.PasswordStrengthConfig.CheckBreachOnCreation)
// 		assert.Equal(t, 12, resp.PasswordStrengthConfig.LudsMinPasswordLength)
// 		assert.Equal(t, 3, resp.PasswordStrengthConfig.LudsMinPasswordComplexity)
// 	})
// 	t.Run("zxcvbn", func(t *testing.T) {
// 		// Arrange
// 		client := NewTestClient(t)
// 		project := client.DisposableProject(projects.VerticalB2B)
// 		ctx := context.Background()

// 		// Act
// 		resp, err := client.PasswordStrengthConfig.Set(ctx, passwordstrengthconfig.SetRequest{
// 			ProjectID: project.LiveProjectID,
// 			PasswordStrengthConfig: passwordstrengthconfig.PasswordStrengthConfig{
// 				CheckBreachOnCreation:       true,
// 				CheckBreachOnAuthentication: true,
// 				ValidateOnAuthentication:    true,
// 				ValidationPolicy:            passwordstrengthconfig.ValidationPolicyZXCVBN,
// 			},
// 		})

// 		// Assert
// 		assert.NoError(t, err)
// 		assert.True(t, resp.PasswordStrengthConfig.CheckBreachOnCreation)
// 		assert.Equal(t, passwordstrengthconfig.ValidationPolicyZXCVBN, resp.PasswordStrengthConfig.ValidationPolicy)
// 	})
// }
