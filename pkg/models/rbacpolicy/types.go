package rbacpolicy

type Policy struct {
	// StytchMember is the default role given to members within the project
	StytchMember Role `json:"stytch_member"`
	// StytchAdmin is the role assigned to admins within an organization
	StytchAdmin Role `json:"stytch_admin"`
	// StytchResources consists of resources created by Stytch that always exist.
	// This field will be returned in relevant Policy objects but can never be overridden or deleted.
	StytchResources []Resource `json:"stytch_resources"`
	// CustomRoles are additional roles that exist within the project beyond the stytch_member or stytch_admin roles
	CustomRoles []Role `json:"custom_roles"`
	// CustomResources are resources that exist within the project beyond those defined within the stytch_resources
	CustomResources []Resource `json:"custom_resources"`
	// CustomScopes are scopes that within the project that are used in Connected Apps
	CustomScopes []Scope `json:"custom_scopes"`
}

type Role struct {
	// RoleID is a human-readable name that is unique within the project
	RoleID string `json:"role_id"`
	// Description is a description for the role
	Description string `json:"description"`
	// Permissions are the permissions granted to this role for resources within the project
	Permissions []Permission `json:"permissions"`
}

type Resource struct {
	// ResourceID is a human-readable name that is unique within the project
	ResourceID string `json:"resource_id"`
	// Description is a description for the resource
	Description string `json:"description"`
	// AvailableActions are the actions that can be granted for this resource
	AvailableActions []string `json:"available_actions"`
}

type Scope struct {
	// Scope is a human-readable name that is unique within the project
	Scope string `json:"scope"`
	// Description is a description for the scope
	Description string `json:"description"`
	// Permissions are the permissions required to be granted this scope for resources within the project
	Permissions []Permission `json:"permissions"`
}

type Permission struct {
	// ResourceID is the ID of the resource that the role can perform actions on
	ResourceID string `json:"resource_id"`
	// Actions is an array of actions that the role can perform on the given resource
	Actions []string `json:"actions"`
}

type GetRequest struct {
	// ProjectID is the ID of the project to get the RBAC policy for
	ProjectID string `json:"project_id"`
}

type GetResponse struct {
	// StatusCode is the HTTP status code for the request
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Policy is the RBAC policy for the project
	Policy Policy `json:"policy"`
}

type SetRequest struct {
	// ProjectID is the ID of the project to set the RBAC policy for
	ProjectID string `json:"project_id"`
	// Policy is the RBAC policy to set for the project
	Policy Policy `json:"policy"`
}

type SetResponse struct {
	// StatusCode is the HTTP status code for the request
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// Policy is the RBAC policy for the project
	Policy Policy `json:"policy"`
}
