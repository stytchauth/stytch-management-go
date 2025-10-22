package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/environments"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v3/pkg/models/rbacpolicy"
)

func getTestB2BPolicy() rbacpolicy.Policy {
	// Define some custom resources.
	resources := []rbacpolicy.Resource{
		{
			ResourceID:       "resource1",
			Description:      "Resource 1",
			AvailableActions: []string{"read", "write", "delete"},
		},
		{
			ResourceID:       "resource2",
			Description:      "Resource 2",
			AvailableActions: []string{"read", "write"},
		},
		{
			ResourceID:       "resource3",
			Description:      "Resource 3",
			AvailableActions: []string{"do_admin_things"},
		},
	}

	// Define some custom scopes.
	scopes := []rbacpolicy.Scope{
		{
			Scope:       "scope1",
			Description: "Scope 1",
			Permissions: []rbacpolicy.Permission{
				{
					ResourceID: "resource1",
					Actions:    []string{"read", "write"},
				},
			},
		},
		{
			Scope:       "scope2",
			Description: "Scope 2",
			Permissions: []rbacpolicy.Permission{
				{
					ResourceID: "resource2",
					Actions:    []string{"read"},
				},
			},
		},
	}

	// Define permissions for default roles
	adminPermissions := rbacpolicy.DefaultRole{
		Permissions: []rbacpolicy.Permission{
			{
				ResourceID: "resource1",
				Actions:    []string{"read", "write", "delete"},
			},
			{
				ResourceID: "resource2",
				Actions:    []string{"read", "write"},
			},
			{
				ResourceID: "resource3",
				Actions:    []string{"do_admin_things"},
			},
		},
	}
	memberPermissions := rbacpolicy.DefaultRole{
		Permissions: []rbacpolicy.Permission{
			{
				ResourceID: "resource1",
				Actions:    []string{"read"},
			},
			{
				ResourceID: "resource2",
				Actions:    []string{"read"},
			},
		},
	}

	writer := rbacpolicy.Role{
		RoleID: "writer_role",
		Permissions: []rbacpolicy.Permission{
			{
				ResourceID: "resource1",
				Actions:    []string{"read", "write"},
			},
			{
				ResourceID: "resource2",
				Actions:    []string{"read", "write"},
			},
		},
	}

	return rbacpolicy.Policy{
		StytchMember:    &memberPermissions,
		StytchAdmin:     &adminPermissions,
		CustomRoles:     []rbacpolicy.Role{writer},
		CustomResources: resources,
		CustomScopes:    scopes,
	}
}

func getTestB2CPolicy() rbacpolicy.Policy {
	// Define some custom resources.
	resources := []rbacpolicy.Resource{
		{
			ResourceID:       "resource1",
			Description:      "Resource 1",
			AvailableActions: []string{"read", "write", "delete"},
		},
		{
			ResourceID:       "resource2",
			Description:      "Resource 2",
			AvailableActions: []string{"read", "write"},
		},
		{
			ResourceID:       "resource3",
			Description:      "Resource 3",
			AvailableActions: []string{"do_admin_things"},
		},
	}

	// Define some custom scopes.
	scopes := []rbacpolicy.Scope{
		{
			Scope:       "scope1",
			Description: "Scope 1",
			Permissions: []rbacpolicy.Permission{
				{
					ResourceID: "resource1",
					Actions:    []string{"read", "write"},
				},
			},
		},
		{
			Scope:       "scope2",
			Description: "Scope 2",
			Permissions: []rbacpolicy.Permission{
				{
					ResourceID: "resource2",
					Actions:    []string{"read"},
				},
			},
		},
	}

	userPermissions := rbacpolicy.DefaultRole{
		Permissions: []rbacpolicy.Permission{
			{
				ResourceID: "resource1",
				Actions:    []string{"read", "write", "delete"},
			},
			{
				ResourceID: "resource2",
				Actions:    []string{"read", "write"},
			},
			{
				ResourceID: "resource3",
				Actions:    []string{"do_admin_things"},
			},
		},
	}

	writer := rbacpolicy.Role{
		RoleID: "writer_role",
		Permissions: []rbacpolicy.Permission{
			{
				ResourceID: "resource1",
				Actions:    []string{"read", "write"},
			},
			{
				ResourceID: "resource2",
				Actions:    []string{"read", "write"},
			},
		},
	}

	return rbacpolicy.Policy{
		StytchUser:      &userPermissions,
		CustomRoles:     []rbacpolicy.Role{writer},
		CustomResources: resources,
		CustomScopes:    scopes,
	}
}

func TestRBACPolicyClient_Get(t *testing.T) {
	t.Run("B2B policy", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		policy := getTestB2BPolicy()
		_, err := client.RBACPolicy.Set(context.Background(), rbacpolicy.SetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			StytchMember:    policy.StytchMember,
			StytchAdmin:     policy.StytchAdmin,
			CustomRoles:     policy.CustomRoles,
			CustomResources: policy.CustomResources,
			CustomScopes:    policy.CustomScopes,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.RBACPolicy.Get(context.Background(), rbacpolicy.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, policy.StytchMember, resp.Policy.StytchMember)
		assert.Equal(t, policy.StytchAdmin, resp.Policy.StytchAdmin)
		assert.Equal(t, policy.CustomRoles, resp.Policy.CustomRoles)
		assert.Equal(t, policy.CustomResources, resp.Policy.CustomResources)
		assert.Equal(t, policy.CustomScopes, resp.Policy.CustomScopes)
	})
	t.Run("B2C policy", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		policy := getTestB2CPolicy()
		_, err := client.RBACPolicy.Set(context.Background(), rbacpolicy.SetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			StytchUser:      policy.StytchUser,
			CustomRoles:     policy.CustomRoles,
			CustomResources: policy.CustomResources,
			CustomScopes:    policy.CustomScopes,
		})
		require.NoError(t, err)

		// Act
		resp, err := client.RBACPolicy.Get(context.Background(), rbacpolicy.GetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, policy.StytchUser, resp.Policy.StytchUser)
		assert.Equal(t, policy.CustomRoles, resp.Policy.CustomRoles)
		assert.Equal(t, policy.CustomResources, resp.Policy.CustomResources)
		assert.Equal(t, policy.CustomScopes, resp.Policy.CustomScopes)
	})
}

func TestRBACClient_SetPolicy(t *testing.T) {
	t.Run("B2B policy", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalB2B, environments.EnvironmentTypeTest)
		policy := getTestB2BPolicy()

		// Act
		resp, err := client.RBACPolicy.Set(context.Background(), rbacpolicy.SetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			StytchMember:    policy.StytchMember,
			StytchAdmin:     policy.StytchAdmin,
			CustomRoles:     policy.CustomRoles,
			CustomResources: policy.CustomResources,
			CustomScopes:    policy.CustomScopes,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, policy.StytchMember, resp.Policy.StytchMember)
		assert.Equal(t, policy.StytchAdmin, resp.Policy.StytchAdmin)
		assert.Equal(t, policy.CustomRoles, resp.Policy.CustomRoles)
		assert.Equal(t, policy.CustomResources, resp.Policy.CustomResources)
		assert.Equal(t, policy.CustomScopes, resp.Policy.CustomScopes)
	})
	t.Run("B2C policy", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		policy := getTestB2CPolicy()

		// Act
		resp, err := client.RBACPolicy.Set(context.Background(), rbacpolicy.SetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			StytchUser:      policy.StytchUser,
			CustomRoles:     policy.CustomRoles,
			CustomResources: policy.CustomResources,
			CustomScopes:    policy.CustomScopes,
		})

		// Assert
		assert.NoError(t, err)
		assert.Equal(t, policy.StytchUser, resp.Policy.StytchUser)
		assert.Equal(t, policy.CustomRoles, resp.Policy.CustomRoles)
		assert.Equal(t, policy.CustomResources, resp.Policy.CustomResources)
		assert.Equal(t, policy.CustomScopes, resp.Policy.CustomScopes)
	})
	t.Run("errors when fields irrelevant to vertical", func(t *testing.T) {
		// Arrange
		client := NewTestClient(t)
		env := client.DisposableEnvironment(projects.VerticalConsumer, environments.EnvironmentTypeTest)
		policy := getTestB2CPolicy()

		// Act
		resp, err := client.RBACPolicy.Set(context.Background(), rbacpolicy.SetRequest{
			ProjectSlug:     env.ProjectSlug,
			EnvironmentSlug: env.EnvironmentSlug,
			StytchUser:      policy.StytchUser,
			// Set StytchMember, which is irrelevant for Consumer projects.
			StytchMember:    policy.StytchUser,
			CustomRoles:     policy.CustomRoles,
			CustomResources: policy.CustomResources,
			CustomScopes:    policy.CustomScopes,
		})

		// Assert
		assert.ErrorContains(t, err, "invalid_role_for_vertical")
		assert.Nil(t, resp)
	})
}
