package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGAdminSCGService struct {
	apiClient *APIClient
}

func NewWSGAdminSCGService(c *APIClient) *WSGAdminSCGService {
	s := new(WSGAdminSCGService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAdminSCGService() *WSGAdminSCGService {
	return NewWSGAdminSCGService(ss.apiClient)
}

type WSGAdminSCGScgAaaServer struct {
	Data *WSGAdminSCGScgAaaServerDataType `json:"data,omitempty"`

	Error *WSGAdminSCGScgAaaServerErrorType `json:"error,omitempty"`

	Extra *WSGAdminSCGScgAaaServerExtraType `json:"extra,omitempty"`

	MetaData *WSGAdminSCGScgAaaServerMetaDataType `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func NewWSGAdminSCGScgAaaServer() *WSGAdminSCGScgAaaServer {
	m := new(WSGAdminSCGScgAaaServer)
	return m
}

type WSGAdminSCGScgAaaServerDataType struct {
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdminSCGScgAaaServerDataTypeListType `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdminSCGScgAaaServerDataType() *WSGAdminSCGScgAaaServerDataType {
	m := new(WSGAdminSCGScgAaaServerDataType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListType struct {
	AdminDomainName *string `json:"adminDomainName,omitempty"`

	AuditLabel *string `json:"auditLabel,omitempty"`

	AuthType *string `json:"authType,omitempty"`

	ClientCertId *WSGAdminSCGScgAaaServerDataTypeListTypeClientCertIdType `json:"clientCertId,omitempty"`

	CnSanIdentity *string `json:"cnSanIdentity,omitempty"`

	CreateOn *int `json:"createOn,omitempty"`

	CreateUserName *string `json:"createUserName,omitempty"`

	CreatorUUID *string `json:"creatorUUID,omitempty"`

	DefaultAdmin *string `json:"defaultAdmin,omitempty"`

	DefaultUserGroup *string `json:"defaultUserGroup,omitempty"`

	DefaultUserGroupEnabled *bool `json:"defaultUserGroupEnabled,omitempty"`

	Description *WSGAdminSCGScgAaaServerDataTypeListTypeDescriptionType `json:"description,omitempty"`

	DomainId *WSGAdminSCGScgAaaServerDataTypeListTypeDomainIdType `json:"domainId,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	EnableSecondaryRadius *int `json:"enableSecondaryRadius,omitempty"`

	GlobalCatalog *WSGAdminSCGScgAaaServerDataTypeListTypeGlobalCatalogType `json:"globalCatalog,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	IpMode *string `json:"ipMode,omitempty"`

	Ipv6 *WSGAdminSCGScgAaaServerDataTypeListTypeIpv6Type `json:"ipv6,omitempty"`

	IsConflict *int `json:"isConflict,omitempty"`

	Key *string `json:"key,omitempty"`

	KeyAttribute *string `json:"keyAttribute,omitempty"`

	MaxRetries *int `json:"maxRetries,omitempty"`

	Name *string `json:"name,omitempty"`

	NullValueColumnNames []*WSGAdminSCGScgAaaServerDataTypeListTypeNullValueColumnNamesType `json:"nullValueColumnNames,omitempty"`

	OcspUrl *string `json:"ocspUrl,omitempty"`

	Password *string `json:"password,omitempty"`

	Port *int `json:"port,omitempty"`

	ProxyAgentEnabled *bool `json:"proxyAgentEnabled,omitempty"`

	ProxyUserPassword *string `json:"proxyUserPassword,omitempty"`

	ProxyUserPrincipalName *string `json:"proxyUserPrincipalName,omitempty"`

	RadiusIP *string `json:"radiusIP,omitempty"`

	RadiusIPv6 *WSGAdminSCGScgAaaServerDataTypeListTypeRadiusIPv6Type `json:"radiusIPv6,omitempty"`

	RadiusPort *int `json:"radiusPort,omitempty"`

	RadiusProtocol *WSGAdminSCGScgAaaServerDataTypeListTypeRadiusProtocolType `json:"radiusProtocol,omitempty"`

	RadiusRealm *string `json:"radiusRealm,omitempty"`

	RadiusShareSecret *string `json:"radiusShareSecret,omitempty"`

	Realms []string `json:"realms,omitempty"`

	RequestTimeout *int `json:"requestTimeout,omitempty"`

	ResponseFail *bool `json:"responseFail,omitempty"`

	RespWindow *int `json:"respWindow,omitempty"`

	RetryPriInvl *int `json:"retryPriInvl,omitempty"`

	ReviveInterval *int `json:"reviveInterval,omitempty"`

	SearchFilter *string `json:"searchFilter,omitempty"`

	SecondaryRadiusEnabled *bool `json:"secondaryRadiusEnabled,omitempty"`

	SecondaryRadiusIP *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPType `json:"secondaryRadiusIP,omitempty"`

	SecondaryRadiusIPv6 *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPv6Type `json:"secondaryRadiusIPv6,omitempty"`

	SecondaryRadiusPort *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusPortType `json:"secondaryRadiusPort,omitempty"`

	SecondaryRadiusProtocol *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusProtocolType `json:"secondaryRadiusProtocol,omitempty"`

	SecondaryRadiusShareSecret *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusShareSecretType `json:"secondaryRadiusShareSecret,omitempty"`

	StandbyAdminDomainName *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyAdminDomainNameType `json:"standbyAdminDomainName,omitempty"`

	StandbyDomainName *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyDomainNameType `json:"standbyDomainName,omitempty"`

	StandbyEnableSecondaryRadius *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyEnableSecondaryRadiusType `json:"standbyEnableSecondaryRadius,omitempty"`

	StandbyGlobalCatalog *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyGlobalCatalogType `json:"standbyGlobalCatalog,omitempty"`

	StandbyIp *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpType `json:"standbyIp,omitempty"`

	StandbyIpv6 *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpv6Type `json:"standbyIpv6,omitempty"`

	StandbyKeyAttribute *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyKeyAttributeType `json:"standbyKeyAttribute,omitempty"`

	StandbyPassword *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPasswordType `json:"standbyPassword,omitempty"`

	StandbyPort *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPortType `json:"standbyPort,omitempty"`

	StandbyRadiusIP *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPType `json:"standbyRadiusIP,omitempty"`

	StandbyRadiusIPv6 *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPv6Type `json:"standbyRadiusIPv6,omitempty"`

	StandbyRadiusPort *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusPortType `json:"standbyRadiusPort,omitempty"`

	StandbyRadiusShareSecret *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusShareSecretType `json:"standbyRadiusShareSecret,omitempty"`

	StandbySearchFilter *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySearchFilterType `json:"standbySearchFilter,omitempty"`

	StandbySecondaryRadiusIP *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPType `json:"standbySecondaryRadiusIP,omitempty"`

	StandbySecondaryRadiusIPv6 *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPv6Type `json:"standbySecondaryRadiusIPv6,omitempty"`

	StandbySecondaryRadiusPort *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusPortType `json:"standbySecondaryRadiusPort,omitempty"`

	StandbySecondaryRadiusShareSecret *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusShareSecretType `json:"standbySecondaryRadiusShareSecret,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	StandbyWindowsDomainName *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyWindowsDomainNameType `json:"standbyWindowsDomainName,omitempty"`

	TacacsService *WSGAdminSCGScgAaaServerDataTypeListTypeTacacsServiceType `json:"tacacsService,omitempty"`

	TenantUUID *string `json:"tenantUUID,omitempty"`

	TlsEnabled *bool `json:"tlsEnabled,omitempty"`

	Type *string `json:"type,omitempty"`

	ValidAAAServer *bool `json:"validAAAServer,omitempty"`

	ValidIP *bool `json:"validIP,omitempty"`

	ValidIPv6 *bool `json:"validIPv6,omitempty"`

	ValidRadiusIp *bool `json:"validRadiusIp,omitempty"`

	ValidRadiusIpv6 *bool `json:"validRadiusIpv6,omitempty"`

	ValidSecondaryRadiusIP *bool `json:"validSecondaryRadiusIP,omitempty"`

	ValidSecondaryRadiusIPv6 *bool `json:"validSecondaryRadiusIPv6,omitempty"`

	WindowsDomainName *WSGAdminSCGScgAaaServerDataTypeListTypeWindowsDomainNameType `json:"windowsDomainName,omitempty"`

	ZombiePeriod *int `json:"zombiePeriod,omitempty"`
}

func NewWSGAdminSCGScgAaaServerDataTypeListType() *WSGAdminSCGScgAaaServerDataTypeListType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeClientCertIdType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeClientCertIdType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeClientCertIdType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeClientCertIdType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeClientCertIdType() *WSGAdminSCGScgAaaServerDataTypeListTypeClientCertIdType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeClientCertIdType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeDescriptionType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeDescriptionType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeDescriptionType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeDescriptionType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeDescriptionType() *WSGAdminSCGScgAaaServerDataTypeListTypeDescriptionType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeDescriptionType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeDomainIdType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeDomainIdType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeDomainIdType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeDomainIdType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeDomainIdType() *WSGAdminSCGScgAaaServerDataTypeListTypeDomainIdType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeDomainIdType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeGlobalCatalogType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeGlobalCatalogType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeGlobalCatalogType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeGlobalCatalogType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeGlobalCatalogType() *WSGAdminSCGScgAaaServerDataTypeListTypeGlobalCatalogType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeGlobalCatalogType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeIpv6Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeIpv6Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeIpv6Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeIpv6Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeIpv6Type() *WSGAdminSCGScgAaaServerDataTypeListTypeIpv6Type {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeIpv6Type)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeNullValueColumnNamesType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeNullValueColumnNamesType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeNullValueColumnNamesType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeNullValueColumnNamesType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeNullValueColumnNamesType() *WSGAdminSCGScgAaaServerDataTypeListTypeNullValueColumnNamesType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeNullValueColumnNamesType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeRadiusIPv6Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeRadiusIPv6Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeRadiusIPv6Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeRadiusIPv6Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeRadiusIPv6Type() *WSGAdminSCGScgAaaServerDataTypeListTypeRadiusIPv6Type {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeRadiusIPv6Type)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeRadiusProtocolType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeRadiusProtocolType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeRadiusProtocolType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeRadiusProtocolType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeRadiusProtocolType() *WSGAdminSCGScgAaaServerDataTypeListTypeRadiusProtocolType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeRadiusProtocolType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPType() *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPv6Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPv6Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPv6Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPv6Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPv6Type() *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPv6Type {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusIPv6Type)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusPortType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusPortType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusPortType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusPortType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusPortType() *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusPortType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusPortType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusProtocolType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusProtocolType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusProtocolType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusProtocolType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusProtocolType() *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusProtocolType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusProtocolType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusShareSecretType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusShareSecretType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusShareSecretType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusShareSecretType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusShareSecretType() *WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusShareSecretType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeSecondaryRadiusShareSecretType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyAdminDomainNameType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyAdminDomainNameType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyAdminDomainNameType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyAdminDomainNameType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyAdminDomainNameType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyAdminDomainNameType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyAdminDomainNameType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyDomainNameType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyDomainNameType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyDomainNameType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyDomainNameType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyDomainNameType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyDomainNameType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyDomainNameType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyEnableSecondaryRadiusType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyEnableSecondaryRadiusType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyEnableSecondaryRadiusType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyEnableSecondaryRadiusType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyEnableSecondaryRadiusType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyEnableSecondaryRadiusType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyEnableSecondaryRadiusType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyGlobalCatalogType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyGlobalCatalogType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyGlobalCatalogType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyGlobalCatalogType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyGlobalCatalogType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyGlobalCatalogType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyGlobalCatalogType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpv6Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpv6Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpv6Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpv6Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpv6Type() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpv6Type {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyIpv6Type)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyKeyAttributeType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyKeyAttributeType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyKeyAttributeType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyKeyAttributeType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyKeyAttributeType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyKeyAttributeType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyKeyAttributeType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPasswordType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPasswordType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPasswordType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPasswordType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyPasswordType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPasswordType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPasswordType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPortType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPortType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPortType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPortType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyPortType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPortType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyPortType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPv6Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPv6Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPv6Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPv6Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPv6Type() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPv6Type {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusIPv6Type)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusPortType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusPortType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusPortType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusPortType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusPortType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusPortType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusPortType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusShareSecretType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusShareSecretType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusShareSecretType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusShareSecretType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusShareSecretType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusShareSecretType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyRadiusShareSecretType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbySearchFilterType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySearchFilterType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbySearchFilterType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySearchFilterType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbySearchFilterType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySearchFilterType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbySearchFilterType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPv6Type struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPv6Type) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPv6Type{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPv6Type) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPv6Type() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPv6Type {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusIPv6Type)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusPortType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusPortType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusPortType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusPortType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusPortType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusPortType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusPortType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusShareSecretType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusShareSecretType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusShareSecretType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusShareSecretType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusShareSecretType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusShareSecretType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbySecondaryRadiusShareSecretType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeStandbyWindowsDomainNameType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyWindowsDomainNameType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeStandbyWindowsDomainNameType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyWindowsDomainNameType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeStandbyWindowsDomainNameType() *WSGAdminSCGScgAaaServerDataTypeListTypeStandbyWindowsDomainNameType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeStandbyWindowsDomainNameType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeTacacsServiceType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeTacacsServiceType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeTacacsServiceType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeTacacsServiceType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeTacacsServiceType() *WSGAdminSCGScgAaaServerDataTypeListTypeTacacsServiceType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeTacacsServiceType)
	return m
}

type WSGAdminSCGScgAaaServerDataTypeListTypeWindowsDomainNameType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeWindowsDomainNameType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerDataTypeListTypeWindowsDomainNameType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerDataTypeListTypeWindowsDomainNameType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerDataTypeListTypeWindowsDomainNameType() *WSGAdminSCGScgAaaServerDataTypeListTypeWindowsDomainNameType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListTypeWindowsDomainNameType)
	return m
}

type WSGAdminSCGScgAaaServerErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerErrorType() *WSGAdminSCGScgAaaServerErrorType {
	m := new(WSGAdminSCGScgAaaServerErrorType)
	return m
}

type WSGAdminSCGScgAaaServerExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = WSGAdminSCGScgAaaServerExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *WSGAdminSCGScgAaaServerExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewWSGAdminSCGScgAaaServerExtraType() *WSGAdminSCGScgAaaServerExtraType {
	m := new(WSGAdminSCGScgAaaServerExtraType)
	return m
}

type WSGAdminSCGScgAaaServerMetaDataType struct {
	Fields []*WSGAdminSCGScgAaaServerMetaDataTypeFieldsType `json:"fields,omitempty"`

	IdProperty *string `json:"idProperty,omitempty"`

	MessageProperty *string `json:"messageProperty,omitempty"`

	Root *string `json:"root,omitempty"`

	SuccessProperty *string `json:"successProperty,omitempty"`

	TotalProperty *string `json:"totalProperty,omitempty"`
}

func NewWSGAdminSCGScgAaaServerMetaDataType() *WSGAdminSCGScgAaaServerMetaDataType {
	m := new(WSGAdminSCGScgAaaServerMetaDataType)
	return m
}

type WSGAdminSCGScgAaaServerMetaDataTypeFieldsType struct {
	Name *string `json:"name,omitempty"`
}

func NewWSGAdminSCGScgAaaServerMetaDataTypeFieldsType() *WSGAdminSCGScgAaaServerMetaDataTypeFieldsType {
	m := new(WSGAdminSCGScgAaaServerMetaDataTypeFieldsType)
	return m
}

// FindScgAdminAaaServers
//
// Use this API command to retrieve a list of registered AAA servers
func (s *WSGAdminSCGService) FindScgAdminAaaServers(ctx context.Context) (*WSGAdminSCGScgAaaServer, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGAdminSCGScgAaaServer
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindScgAdminAaaServers, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGAdminSCGScgAaaServer()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
