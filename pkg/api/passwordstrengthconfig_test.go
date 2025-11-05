package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/passwordstrengthconfig"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
)

func TestPasswordStrengthConfigClient_Get(t *testing.T) {
	t.Run("get password strength config", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.PasswordStrengthConfig.Get(ctx, passwordstrengthconfig.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		// Assert
		assert.NoError(t, err)
		assert.NotEmpty(t, resp.RequestID)
		assert.NotEmpty(t, resp.PasswordStrengthConfig.ValidationPolicy)
	})
}

func TestPasswordStrengthConfigClient_Set(t *testing.T) {
	t.Run("set password strength config with LUDS policy", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.PasswordStrengthConfig.Set(ctx, passwordstrengthconfig.SetRequest{
			ProjectSlug:                 env.ProjectSlug,
			EnvironmentSlug:             env.EnvironmentSlug,
			CheckBreachOnCreation:       true,
			CheckBreachOnAuthentication: true,
			ValidateOnAuthentication:    true,
			ValidationPolicy:            passwordstrengthconfig.ValidationPolicyLUDS,
			LudsMinPasswordLength:       ptr(10),
			LudsMinPasswordComplexity:   ptr(3),
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, true, resp.PasswordStrengthConfig.CheckBreachOnCreation)
		assert.Equal(t, true, resp.PasswordStrengthConfig.CheckBreachOnAuthentication)
		assert.Equal(t, true, resp.PasswordStrengthConfig.ValidateOnAuthentication)
		assert.Equal(t, passwordstrengthconfig.ValidationPolicyLUDS, resp.PasswordStrengthConfig.ValidationPolicy)
		assert.Equal(t, 10, *resp.PasswordStrengthConfig.LudsMinPasswordLength)
		assert.Equal(t, 3, *resp.PasswordStrengthConfig.LudsMinPasswordComplexity)
	})

	t.Run("set password strength config with ZXCVBN policy", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		ctx := context.Background()

		// Act
		resp, err := client.PasswordStrengthConfig.Set(ctx, passwordstrengthconfig.SetRequest{
			ProjectSlug:                 env.ProjectSlug,
			EnvironmentSlug:             env.EnvironmentSlug,
			CheckBreachOnCreation:       false,
			CheckBreachOnAuthentication: false,
			ValidateOnAuthentication:    false,
			ValidationPolicy:            passwordstrengthconfig.ValidationPolicyZXCVBN,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, false, resp.PasswordStrengthConfig.CheckBreachOnCreation)
		assert.Equal(t, false, resp.PasswordStrengthConfig.CheckBreachOnAuthentication)
		assert.Equal(t, false, resp.PasswordStrengthConfig.ValidateOnAuthentication)
		assert.Equal(t, passwordstrengthconfig.ValidationPolicyZXCVBN, resp.PasswordStrengthConfig.ValidationPolicy)
		assert.Nil(t, resp.PasswordStrengthConfig.LudsMinPasswordLength)
		assert.Nil(t, resp.PasswordStrengthConfig.LudsMinPasswordComplexity)
	})
}
