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

func getTestB2BPolicy(
	t *testing.T, client *testClient, project string, env string,
) rbacpolicy.Policy {
	t.Helper()
	// The StytchMember and StytchAdmin RoleIDs and Descriptions cannot be modified, so we first
	// look up their current values when constructing our new RBAC policy.
	resp, err := client.RBACPolicy.Get(context.Background(), rbacpolicy.GetRequest{
		ProjectSlug:     project,
		EnvironmentSlug: env,
	})
	require.NoError(t, err)

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

	// Add custom resources to default and custom roles.
	admin := rbacpolicy.Role{
		RoleID:      resp.Policy.StytchAdmin.RoleID,
		Description: resp.Policy.StytchAdmin.Description,
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
	reader := rbacpolicy.Role{
		RoleID:      resp.Policy.StytchMember.RoleID,
		Description: resp.Policy.StytchMember.Description,
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

	return rbacpolicy.Policy{
		StytchMember:    &reader,
		StytchAdmin:     &admin,
		CustomRoles:     []rbacpolicy.Role{writer},
		CustomResources: resources,
		CustomScopes:    scopes,
	}
}

func getTestB2CPolicy(
	t *testing.T, client *testClient, project string, env string,
) rbacpolicy.Policy {
	t.Helper()
	// The StytchUser RoleID and Description cannot be modified, so we first look up its current
	// values when constructing our new RBAC policy.
	resp, err := client.RBACPolicy.Get(context.Background(), rbacpolicy.GetRequest{
		ProjectSlug:     project,
		EnvironmentSlug: env,
	})
	require.NoError(t, err)

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

	// Add custom resources to default and custom roles.
	user := rbacpolicy.Role{
		RoleID:      resp.Policy.StytchUser.RoleID,
		Description: resp.Policy.StytchUser.Description,
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
		StytchUser:      &user,
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
		policy := getTestB2BPolicy(t, client, env.ProjectSlug, env.EnvironmentSlug)
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
		policy := getTestB2CPolicy(t, client, env.ProjectSlug, env.EnvironmentSlug)
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
		policy := getTestB2BPolicy(t, client, env.ProjectSlug, env.EnvironmentSlug)

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
		policy := getTestB2CPolicy(t, client, env.ProjectSlug, env.EnvironmentSlug)

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
		policy := getTestB2CPolicy(t, client, env.ProjectSlug, env.EnvironmentSlug)

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
