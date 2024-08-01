package rbac

type PolicyRolePermission struct {
	ResourceID string   `json:"resource_id"`
	Actions    []string `json:"actions"`
}

type PolicyRole struct {
	RoleID      string                 `json:"role_id"`
	Description string                 `json:"description"`
	Permissions []PolicyRolePermission `json:"permissions"`
}

type PolicyResource struct {
	ResourceID  string   `json:"resource_id"`
	Description string   `json:"description"`
	Actions     []string `json:"actions"`
}

type Policy struct {
	DefaultRole           PolicyRole       `json:"default_role"`
	OrganizationAdminRole PolicyRole       `json:"organization_admin_role"`
	StytchResources       []PolicyResource `json:"stytch_resources"`
	CustomRoles           []PolicyRole     `json:"custom_roles"`
	CustomResources       []PolicyResource `json:"custom_resources"`
}

type GetPolicyRequest struct {
	ProjectID string `json:"project_id"`
}

type GetPolicyResponse struct {
	RequestID string `json:"request_id"`
	Policy    Policy `json:"policy"`
}

type SetPolicyRequest struct {
	ProjectID string `json:"project_id"`
	Policy    Policy `json:"policy"`
}

type SetPolicyResponse struct {
	RequestID string `json:"request_id"`
	Policy    Policy `json:"policy"`
}
