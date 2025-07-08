package trustedtokenprofiles

type TrustedTokenProfile struct {
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
	// Name is the name of the trusted token profile
	Name string `json:"name"`
	// Audience is the audience for the trusted token profile
	Audience string `json:"audience"`
	// Issuer is the issuer for the trusted token profile
	Issuer string `json:"issuer"`
	// JwksUrl is the JWKS URL for the trusted token profile (optional)
	JwksUrl *string `json:"jwks_url,omitempty"`
	// AttributeMapping is the attribute mapping for the trusted token profile (optional)
	AttributeMapping map[string]interface{} `json:"attribute_mapping,omitempty"`
	// PemFiles is a list of PEM file identifiers
	PemFiles []PEMFile `json:"pem_files"`
	// PublicKeyType is the type of public key
	PublicKeyType string `json:"public_key_type"`
}

type PEMFile struct {
	// KeyID is the unique identifier for the PEM file
	PemFileID string `json:"pem_file_id"`
	// PublicKey is the public key content
	PublicKey string `json:"public_key"`
}

type CreateTrustedTokenProfileRequest struct {
	// ProjectID is the project to create the trusted token profile for
	ProjectID string `json:"-"`
	// Name is the name of the trusted token profile
	Name string `json:"name"`
	// Audience is the audience for the trusted token profile
	Audience string `json:"audience"`
	// Issuer is the issuer for the trusted token profile
	Issuer string `json:"issuer"`
	// JwksUrl is the JWKS URL for the trusted token profile (optional)
	JwksUrl *string `json:"jwks_url,omitempty"`
	// AttributeMapping is the attribute mapping for the trusted token profile (optional)
	AttributeMapping map[string]interface{} `json:"attribute_mapping,omitempty"`
	// PemFiles is a list of PEM file identifiers
	PemFiles []string `json:"pem_files"`
	// PublicKeyType is the type of public key
	PublicKeyType string `json:"public_key_type"`
}

type CreateTrustedTokenProfileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// ProfileID is the id of the trusted token profile that was created
	ProfileID string `json:"profile_id"`
}

type GetTrustedTokenProfileRequest struct {
	// ProjectID is the project to retrieve the trusted token profile for
	ProjectID string `json:"-"`
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
	// ProjectID is the project to list the trusted token profiles for
	ProjectID string `json:"project_id"`
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
	// ProjectID is the project to update the trusted token profile for
	ProjectID string `json:"-"`
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
	// Name is the name of the trusted token profile (optional)
	Name string `json:"name,omitempty"`
	// Audience is the audience for the trusted token profile (optional)
	Audience string `json:"audience,omitempty"`
	// Issuer is the issuer for the trusted token profile (optional)
	Issuer string `json:"issuer,omitempty"`
	// JwksUrl is the JWKS URL for the trusted token profile (optional)
	JwksUrl *string `json:"jwks_url,omitempty"`
	// AttributeMapping is the attribute mapping for the trusted token profile (optional)
	AttributeMapping map[string]interface{} `json:"attribute_mapping,omitempty"`
}

type UpdateTrustedTokenProfileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
	// ProfileID is the id of the trusted token profile that was updated
	ProfileID string `json:"profile_id"`
}

type DeleteTrustedTokenProfileRequest struct {
	// ProjectID is the project to delete the trusted token profile for
	ProjectID string `json:"-"`
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
}

type DeleteTrustedTokenProfileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
}

type CreatePEMFileRequest struct {
	// ProjectID is the project to create the PEM file for
	ProjectID string `json:"-"`
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
	// PemFileID is the id of the PEM file that was created
	PEMFileID string `json:"pem_file_id"`
}

type DeletePEMFileRequest struct {
	// ProjectID is the project to delete the PEM file for
	ProjectID string `json:"-"`
	// ProfileID is the unique identifier for the trusted token profile
	ProfileID string `json:"profile_id"`
	// PemFileId is the unique identifier for the PEM file
	PemFileID string `json:"pem_file_id"`
}

type DeletePEMFileResponse struct {
	// StatusCode is the HTTP status code for the response
	StatusCode int `json:"status_code"`
	// RequestID is a unique identifier to help with debugging the request
	RequestID string `json:"request_id"`
}
