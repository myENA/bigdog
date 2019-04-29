package certificate

// API Version: v8_0

type Certificate struct {
	CertificasSigningRequest *common.GenericRef `json:"certificasSigningRequest,omitempty"`
	Data                     *string            `json:"data,omitempty"`
	Description              *string            `json:"description,omitempty"`
	ID                       *string            `json:"id,omitempty"`
	Information              *string            `json:"information,omitempty"`
	IntermediateData         []string           `json:"intermediateData,omitempty"`
	Name                     *string            `json:"name,omitempty"`
	Passphrase               *string            `json:"passphrase,omitempty"`
	PrivateKeyData           *string            `json:"privateKeyData,omitempty"`
	PublicKey                *string            `json:"publicKey,omitempty"`
	RootData                 *string            `json:"rootData,omitempty"`
}

type CertificateList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type CertificateListType struct {
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type CertificatesSigningRequest struct {
	City             *string `json:"city,omitempty"`
	CommonName       *string `json:"commonName,omitempty"`
	CountryCode      *string `json:"countryCode,omitempty"`
	Description      *string `json:"description,omitempty"`
	Email            *string `json:"email,omitempty"`
	ID               *string `json:"id,omitempty"`
	Name             *string `json:"name,omitempty"`
	Organization     *string `json:"organization,omitempty"`
	OrganizationUnit *string `json:"organizationUnit,omitempty"`
	State            *string `json:"state,omitempty"`
}

type CertSetting struct {
	ServiceCertificates []*ServiceCertificates `json:"serviceCertificates,omitempty"`
}

type CreateCert struct {
	CertificasSigningRequest *common.GenericRef `json:"certificasSigningRequest,omitempty"`
	Data                     *string            `json:"data,omitempty"`
	Description              *string            `json:"description,omitempty"`
	IntermediateData         []string           `json:"intermediateData,omitempty"`
	Name                     *string            `json:"name,omitempty"`
	Passphrase               *string            `json:"passphrase,omitempty"`
	PrivateKeyData           *string            `json:"privateKeyData,omitempty"`
	RootData                 *string            `json:"rootData,omitempty"`
}

type CreateCSR struct {
	City             *string `json:"city,omitempty"`
	CommonName       *string `json:"commonName,omitempty"`
	CountryCode      *string `json:"countryCode,omitempty"`
	Description      *string `json:"description,omitempty"`
	Email            *string `json:"email,omitempty"`
	Name             *string `json:"name,omitempty"`
	Organization     *string `json:"organization,omitempty"`
	OrganizationUnit *string `json:"organizationUnit,omitempty"`
	State            *string `json:"state,omitempty"`
}

type CreateTrustedCAChain struct {
	Description   *string  `json:"description,omitempty"`
	InterCertData []string `json:"interCertData,omitempty"`
	Name          *string  `json:"name,omitempty"`
	RootCertData  *string  `json:"rootCertData,omitempty"`
}

type CsrList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type CsrListType struct {
	Description *string `json:"description,omitempty"`
	ID          *string `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type DeleteBulk struct {
	IDList common.IDList `json:"idList,omitempty"`
}

type ModifyTrustedCAChain struct {
	Description   *string  `json:"description,omitempty"`
	Information   *string  `json:"information,omitempty"`
	InterCertData []string `json:"interCertData,omitempty"`
	Name          *string  `json:"name,omitempty"`
	RootCertData  *string  `json:"rootCertData,omitempty"`
}

type ServiceCertificate struct {
	Certificate *common.GenericRef `json:"certificate,omitempty"`
	Service     *string            `json:"service,omitempty"`
}

type TrustedCAChain struct {
	Description   *string  `json:"description,omitempty"`
	ID            *string  `json:"id,omitempty"`
	InterCertData []string `json:"interCertData,omitempty"`
	Name          *string  `json:"name,omitempty"`
	RootCertData  *string  `json:"rootCertData,omitempty"`
}

type TrustedCAChainCertList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}

type TrustedCAChainCertListType struct {
	Description      *string  `json:"description,omitempty"`
	ID               *string  `json:"id,omitempty"`
	Information      *string  `json:"information,omitempty"`
	InterCertData    []string `json:"interCertData,omitempty"`
	ModifiedDateTime *string  `json:"modifiedDateTime,omitempty"`
	ModifierUsername *string  `json:"modifierUsername,omitempty"`
	Name             *string  `json:"name,omitempty"`
	RootCertData     *string  `json:"rootCertData,omitempty"`
}
