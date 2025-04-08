package api_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/projects"
	"github.com/stytchauth/stytch-management-go/v2/pkg/models/rbacpolicy"
)

func getTestPolicy(t *testing.T, client *testClient, projectID string) rbacpolicy.Policy {
	t.Helper()
	// The StytchMember and StytchAdmin RoleIDs and Descriptions cannot be modified, so we first
	// look up their current value when constructing our new RBAC policy.
	resp, err := client.RBACPolicy.Get(context.Background(), rbacpolicy.GetRequest{
		ProjectID: projectID,
	})
	require.NoError(t, err)

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
	viewer := rbacpolicy.Role{
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
		StytchMember:    viewer,
		StytchAdmin:     admin,
		CustomRoles:     []rbacpolicy.Role{writer},
		CustomResources: resources,
	}
}

func TestRBACPolicyClient_Get(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	policy := getTestPolicy(t, client, project.LiveProjectID)
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
}

func TestRBACClient_SetPolicy(t *testing.T) {
	// Arrange
	client := NewTestClient(t)
	project := client.DisposableProject(projects.VerticalB2B)
	policy := getTestPolicy(t, client, project.LiveProjectID)

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
}
