package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v1/pkg/models/rbac"
)

func getTestPolicy(t *testing.T) rbac.Policy {
	t.Helper()

	resources := []rbac.PolicyResource{
		{
			ResourceID:  "resource1",
			Description: "Resource 1",
			Actions:     []string{"read", "write", "delete"},
		},
		{
			ResourceID:  "resource2",
			Description: "Resource 2",
			Actions:     []string{"read", "write"},
		},
		{
			ResourceID:  "resource3",
			Description: "Resource 3",
			Actions:     []string{"do_admin_things"},
		},
	}

	admin := rbac.PolicyRole{
		RoleID:      "admin_role",
		Description: "Admin role",
		Permissions: []rbac.PolicyRolePermission{
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
	writer := rbac.PolicyRole{
		RoleID:      "writer_role",
		Description: "Writer role",
		Permissions: []rbac.PolicyRolePermission{
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
	viewer := rbac.PolicyRole{
		RoleID:      "viewer_role",
		Description: "Viewer role",
		Permissions: []rbac.PolicyRolePermission{
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

	return rbac.Policy{
		DefaultRole:           viewer,
		OrganizationAdminRole: admin,
		CustomRoles:           []rbac.PolicyRole{writer},
		CustomResources:       resources,
	}
}

func TestRBACClient_GetPolicy(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	policy := getTestPolicy(t)
	_, err := client.RBAC.SetPolicy(context.Background(), rbac.SetPolicyRequest{
		ProjectID: project.LiveProjectID,
		Policy:    policy,
	})
	require.NoError(t, err)

	// Act
	resp, err := client.RBAC.GetPolicy(context.Background(), rbac.GetPolicyRequest{
		ProjectID: project.LiveProjectID,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, policy.DefaultRole, resp.Policy.DefaultRole)
	assert.Equal(t, policy.OrganizationAdminRole, resp.Policy.OrganizationAdminRole)
	assert.Equal(t, policy.CustomRoles, resp.Policy.CustomRoles)
	assert.Equal(t, policy.CustomResources, resp.Policy.CustomResources)
}

func TestRBACClient_SetPolicy(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	policy := getTestPolicy(t)

	// Act
	resp, err := client.RBAC.SetPolicy(context.Background(), rbac.SetPolicyRequest{
		ProjectID: project.LiveProjectID,
		Policy:    policy,
	})

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, policy.DefaultRole, resp.Policy.DefaultRole)
	assert.Equal(t, policy.OrganizationAdminRole, resp.Policy.OrganizationAdminRole)
	assert.Equal(t, policy.CustomRoles, resp.Policy.CustomRoles)
	assert.Equal(t, policy.CustomResources, resp.Policy.CustomResources)
}
