package trustedtokenprofiles

type TrustedTokenProfile struct {
	// ID is the unique identifier for the trusted token profile
	ID string `json:"profile_id"`
	// Name is the name of the trusted token profile
	Name string `json:"name"`
	// Audience is the audience for the trusted token profile
	Audience string `json:"audience"`
	// Issuer is the issuer for the trusted token profile
	Issuer string `json:"issuer"`
	// JwksURL is the JWKS URL for the trusted token profile
	JwksURL string `json:"jwks_url"`
	// AttributeMapping is the attribute mapping for the trusted token profile
	AttributeMapping map[string]interface{} `json:"attribute_mapping"`
	// PEMFiles is a list of PEM files
	PEMFiles []PEMFile `json:"pem_files"`
	// PublicKeyType is the type of public key
	PublicKeyType string `json:"public_key_type"`
}

type PEMFile struct {
	// ID is the unique identifier for the PEM file
	ID string `json:"pem_file_id"`
	// PublicKey is the public key content
	PublicKey string `json:"public_key"`
}

type CreateTrustedTokenProfileRequest struct {
	// Project is the project to create the trusted token profile for
	Project string `json:"-"`
	// Environment is the environment (e.g., "test", "live") to create the trusted token profile for
	Environment string `json:"-"`
	// Name is the name of the trusted token profile
	Name string `json:"name"`
	// Audience is the audience for the trusted token profile
	Audience string `json:"audience"`
	// Issuer is the issuer for the trusted token profile
	Issuer string `json:"issuer"`
	// JwksURL is the JWKS URL for the trusted token profile (optional)
	JwksURL *string `json:"jwks_url,omitempty"`
	// AttributeMapping is the attribute mapping for the trusted token profile (optional)
	AttributeMapping map[string]interface{} `json:"attribute_mapping,omitempty"`
	// PublicKeyType is the type of public key
	PublicKeyType string `json:"public_key_type"`
	// PEMFiles is a list of PEM files
	PEMFiles []string `json:"pem_files"`
	// CanJITProvision is a boolean indicating whether the trusted token profile can be provisioned JIT
	CanJITProvision bool `json:"can_jit_provision"`
}

type CreateTrustedTokenProfileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// TrustedTokenProfile is the trusted token profile that was created
	TrustedTokenProfile TrustedTokenProfile `json:"profile"`
}

type GetTrustedTokenProfileRequest struct {
	// Project is the project to retrieve the trusted token profile for
	Project string `json:"-"`
	// Environment is the environment (e.g., "test", "live") to retrieve the trusted token profile for
	Environment string `json:"-"`
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
}

type GetTrustedTokenProfileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// TrustedTokenProfile is the trusted token profile that was retrieved
	TrustedTokenProfile TrustedTokenProfile `json:"trusted_token_profile"`
}

type ListTrustedTokenProfilesRequest struct {
	// Project is the project to list the trusted token profiles for
	Project string `json:"-"`
	// Environment is the environment (e.g., "test", "live") to list the trusted token profiles for
	Environment string `json:"-"`
}

type ListTrustedTokenProfilesResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// TrustedTokenProfiles is a list of all trusted token profiles for the project
	TrustedTokenProfiles []TrustedTokenProfile `json:"trusted_token_profiles"`
}

type UpdateTrustedTokenProfileRequest struct {
	// Project is the project to update the trusted token profile for
	Project string `json:"-"`
	// Environment is the environment (e.g., "test", "live") to update the trusted token profile for
	Environment string `json:"-"`
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
	// Name is the name of the trusted token profile (optional)
	Name *string `json:"name,omitempty"`
	// Audience is the audience for the trusted token profile (optional)
	Audience *string `json:"audience,omitempty"`
	// Issuer is the issuer for the trusted token profile (optional)
	Issuer *string `json:"issuer,omitempty"`
	// JwksURL is the JWKS URL for the trusted token profile (optional)
	JwksURL *string `json:"jwks_url"`
	// AttributeMapping is the attribute mapping for the trusted token profile (optional)
	AttributeMapping map[string]interface{} `json:"attribute_mapping"`
	// CanJITProvision is a boolean indicating whether the trusted token profile can be provisioned JIT
	CanJITProvision *bool `json:"can_jit_provision,omitempty"`
}

type UpdateTrustedTokenProfileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// TrustedTokenProfile is the trusted token profile that was updated
	TrustedTokenProfile TrustedTokenProfile `json:"profile"`
}

type DeleteTrustedTokenProfileRequest struct {
	// Project is the project to delete the trusted token profile for
	Project string `json:"-"`
	// Environment is the environment (e.g., "test", "live") to delete the trusted token profile for
	Environment string `json:"-"`
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
}

type DeleteTrustedTokenProfileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
}

// PEM File management

type CreatePEMFileRequest struct {
	// Project is the project to create the PEM file for
	Project string `json:"-"`
	// Environment is the environment (e.g., "test", "live") to create the PEM file for
	Environment string `json:"-"`
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
	// PublicKey is the public key to create
	PublicKey string `json:"public_key"`
}

type CreatePEMFileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// PEMFile is the PEM file that was created
	PEMFile PEMFile `json:"pem_file"`
}

type GetPEMFileRequest struct {
	// Project is the project to get the PEM file for
	Project string `json:"-"`
	// Environment is the environment (e.g., "test", "live") to get the PEM file for
	Environment string `json:"-"`
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
	// PEMFileID is the unique identifier for the PEM file
	PEMFileID string `json:"pem_file_id"`
}

type GetPEMFileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// PEMFile is the PEM file that was retrieved
	PEMFile PEMFile `json:"pem_file"`
}

type DeletePEMFileRequest struct {
	// Project is the project to delete the PEM file for
	Project string `json:"-"`
	// Environment is the environment (e.g., "test", "live") to delete the PEM file for
	Environment string `json:"-"`
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
	// PEMFileID is the unique identifier for the PEM file
	PEMFileID string `json:"pem_file_id"`
}

type DeletePEMFileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
}
