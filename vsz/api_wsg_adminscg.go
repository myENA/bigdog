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

	ClientCertId *string `json:"clientCertId,omitempty"`

	CnSanIdentity *string `json:"cnSanIdentity,omitempty"`

	CreateOn *int `json:"createOn,omitempty"`

	CreateUserName *string `json:"createUserName,omitempty"`

	CreatorUUID *string `json:"creatorUUID,omitempty"`

	DefaultAdmin *string `json:"defaultAdmin,omitempty"`

	DefaultUserGroup *string `json:"defaultUserGroup,omitempty"`

	DefaultUserGroupEnabled *bool `json:"defaultUserGroupEnabled,omitempty"`

	Description *string `json:"description,omitempty"`

	DomainId *string `json:"domainId,omitempty"`

	DomainName *string `json:"domainName,omitempty"`

	EnableSecondaryRadius *int `json:"enableSecondaryRadius,omitempty"`

	GlobalCatalog *string `json:"globalCatalog,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	IpMode *string `json:"ipMode,omitempty"`

	Ipv6 *string `json:"ipv6,omitempty"`

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

	RadiusIPv6 *string `json:"radiusIPv6,omitempty"`

	RadiusPort *int `json:"radiusPort,omitempty"`

	RadiusProtocol *string `json:"radiusProtocol,omitempty"`

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

	SecondaryRadiusIP *string `json:"secondaryRadiusIP,omitempty"`

	SecondaryRadiusIPv6 *string `json:"secondaryRadiusIPv6,omitempty"`

	SecondaryRadiusPort *string `json:"secondaryRadiusPort,omitempty"`

	SecondaryRadiusProtocol *string `json:"secondaryRadiusProtocol,omitempty"`

	SecondaryRadiusShareSecret *string `json:"secondaryRadiusShareSecret,omitempty"`

	StandbyAdminDomainName *string `json:"standbyAdminDomainName,omitempty"`

	StandbyDomainName *string `json:"standbyDomainName,omitempty"`

	StandbyEnableSecondaryRadius *bool `json:"standbyEnableSecondaryRadius,omitempty"`

	StandbyGlobalCatalog *string `json:"standbyGlobalCatalog,omitempty"`

	StandbyIp *string `json:"standbyIp,omitempty"`

	StandbyIpv6 *string `json:"standbyIpv6,omitempty"`

	StandbyKeyAttribute *string `json:"standbyKeyAttribute,omitempty"`

	StandbyPassword *string `json:"standbyPassword,omitempty"`

	StandbyPort *string `json:"standbyPort,omitempty"`

	StandbyRadiusIP *string `json:"standbyRadiusIP,omitempty"`

	StandbyRadiusIPv6 *string `json:"standbyRadiusIPv6,omitempty"`

	StandbyRadiusPort *string `json:"standbyRadiusPort,omitempty"`

	StandbyRadiusShareSecret *string `json:"standbyRadiusShareSecret,omitempty"`

	StandbySearchFilter *string `json:"standbySearchFilter,omitempty"`

	StandbySecondaryRadiusIP *string `json:"standbySecondaryRadiusIP,omitempty"`

	StandbySecondaryRadiusIPv6 *string `json:"standbySecondaryRadiusIPv6,omitempty"`

	StandbySecondaryRadiusPort *int `json:"standbySecondaryRadiusPort,omitempty"`

	StandbySecondaryRadiusShareSecret *string `json:"standbySecondaryRadiusShareSecret,omitempty"`

	StandbyServerEnabled *bool `json:"standbyServerEnabled,omitempty"`

	StandbyWindowsDomainName *string `json:"standbyWindowsDomainName,omitempty"`

	TacacsService *string `json:"tacacsService,omitempty"`

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

	WindowsDomainName *string `json:"windowsDomainName,omitempty"`

	ZombiePeriod *int `json:"zombiePeriod,omitempty"`
}

func NewWSGAdminSCGScgAaaServerDataTypeListType() *WSGAdminSCGScgAaaServerDataTypeListType {
	m := new(WSGAdminSCGScgAaaServerDataTypeListType)
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
