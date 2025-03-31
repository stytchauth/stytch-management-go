package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/pkg/models/rbacpolicy"
)

func getTestPolicy(t *testing.T) rbacpolicy.Policy {
	t.Helper()

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

	admin := rbacpolicy.Role{
		RoleID:      "admin_role",
		Description: "Admin role",
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
		RoleID:      "writer_role",
		Description: "Writer role",
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
	viewer := rbacpolicy.Role{
		RoleID:      "viewer_role",
		Description: "Viewer role",
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

	writerScope := rbacpolicy.Scope{
		Scope:       "write:documents",
		Description: "Write documents",
		Permissions: []rbacpolicy.Permission{
			{
				ResourceID: "resource1",
				Actions:    []string{"write"},
			},
			{
				ResourceID: "resource2",
				Actions:    []string{"write"},
			},
		},
	}

	return rbacpolicy.Policy{
		StytchMember:    viewer,
		StytchAdmin:     admin,
		CustomRoles:     []rbacpolicy.Role{writer},
		CustomResources: resources,
		CustomScopes:    []rbacpolicy.Scope{writerScope},
	}
}

func TestRBACPolicyClient_Get(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	policy := getTestPolicy(t)
	_, err := client.RBACPolicy.Set(context.Background(), rbacpolicy.SetRequest{
		ProjectID: project.LiveProjectID,
		Policy:    policy,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.RBACPolicy.Get(context.Background(), rbacpolicy.GetRequest{
		ProjectID: project.LiveProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, policy.StytchMember, resp.Policy.StytchMember)
	assert.Equal(t, policy.StytchAdmin, resp.Policy.StytchAdmin)
	assert.Equal(t, policy.CustomRoles, resp.Policy.CustomRoles)
	assert.Equal(t, policy.CustomResources, resp.Policy.CustomResources)
	assert.Equal(t, policy.CustomScopes, resp.Policy.CustomScopes)
}

func TestRBACClient_SetPolicy(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	policy := getTestPolicy(t)

	// Act
	resp, err := client.RBACPolicy.Set(context.Background(), rbacpolicy.SetRequest{
		ProjectID: project.LiveProjectID,
		Policy:    policy,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, policy.StytchMember, resp.Policy.StytchMember)
	assert.Equal(t, policy.StytchAdmin, resp.Policy.StytchAdmin)
	assert.Equal(t, policy.CustomRoles, resp.Policy.CustomRoles)
	assert.Equal(t, policy.CustomResources, resp.Policy.CustomResources)
	assert.Equal(t, policy.CustomScopes, resp.Policy.CustomScopes)
}
