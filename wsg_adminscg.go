package bigdog

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGAdminSCGService struct {
	apiClient *VSZClient
}

func NewWSGAdminSCGService(c *VSZClient) *WSGAdminSCGService {
	s := new(WSGAdminSCGService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAdminSCGService() *WSGAdminSCGService {
	return NewWSGAdminSCGService(ss.apiClient)
}

// WSGAdminSCGScgAaaServer
//
// Definition: adminscg_scgAaaServer
type WSGAdminSCGScgAaaServer struct {
	Data *WSGAdminSCGScgAaaServerDataType `json:"data,omitempty"`

	Error interface{} `json:"error,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func NewWSGAdminSCGScgAaaServer() *WSGAdminSCGScgAaaServer {
	m := new(WSGAdminSCGScgAaaServer)
	return m
}

// WSGAdminSCGScgAaaServerDataType
//
// Definition: adminscg_scgAaaServerDataType
type WSGAdminSCGScgAaaServerDataType struct {
	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGAdminSCGScgAaaServerEntry `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGAdminSCGScgAaaServerDataType() *WSGAdminSCGScgAaaServerDataType {
	m := new(WSGAdminSCGScgAaaServerDataType)
	return m
}

// WSGAdminSCGScgAaaServerEntry
//
// Definition: adminscg_scgAaaServerEntry
type WSGAdminSCGScgAaaServerEntry struct {
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

	NullValueColumnNames []interface{} `json:"nullValueColumnNames,omitempty"`

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

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *WSGAdminSCGScgAaaServerEntry) UnmarshalJSON(b []byte) error {
	type _WSGAdminSCGScgAaaServerEntry WSGAdminSCGScgAaaServerEntry
	tmpType := new(_WSGAdminSCGScgAaaServerEntry)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "adminDomainName")
	delete(tmpType.XAdditionalProperties, "auditLabel")
	delete(tmpType.XAdditionalProperties, "authType")
	delete(tmpType.XAdditionalProperties, "clientCertId")
	delete(tmpType.XAdditionalProperties, "cnSanIdentity")
	delete(tmpType.XAdditionalProperties, "createOn")
	delete(tmpType.XAdditionalProperties, "createUserName")
	delete(tmpType.XAdditionalProperties, "creatorUUID")
	delete(tmpType.XAdditionalProperties, "defaultAdmin")
	delete(tmpType.XAdditionalProperties, "defaultUserGroup")
	delete(tmpType.XAdditionalProperties, "defaultUserGroupEnabled")
	delete(tmpType.XAdditionalProperties, "description")
	delete(tmpType.XAdditionalProperties, "domainId")
	delete(tmpType.XAdditionalProperties, "domainName")
	delete(tmpType.XAdditionalProperties, "enableSecondaryRadius")
	delete(tmpType.XAdditionalProperties, "globalCatalog")
	delete(tmpType.XAdditionalProperties, "id")
	delete(tmpType.XAdditionalProperties, "ip")
	delete(tmpType.XAdditionalProperties, "ipMode")
	delete(tmpType.XAdditionalProperties, "ipv6")
	delete(tmpType.XAdditionalProperties, "isConflict")
	delete(tmpType.XAdditionalProperties, "key")
	delete(tmpType.XAdditionalProperties, "keyAttribute")
	delete(tmpType.XAdditionalProperties, "maxRetries")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "nullValueColumnNames")
	delete(tmpType.XAdditionalProperties, "ocspUrl")
	delete(tmpType.XAdditionalProperties, "password")
	delete(tmpType.XAdditionalProperties, "port")
	delete(tmpType.XAdditionalProperties, "proxyAgentEnabled")
	delete(tmpType.XAdditionalProperties, "proxyUserPassword")
	delete(tmpType.XAdditionalProperties, "proxyUserPrincipalName")
	delete(tmpType.XAdditionalProperties, "radiusIP")
	delete(tmpType.XAdditionalProperties, "radiusIPv6")
	delete(tmpType.XAdditionalProperties, "radiusPort")
	delete(tmpType.XAdditionalProperties, "radiusProtocol")
	delete(tmpType.XAdditionalProperties, "radiusRealm")
	delete(tmpType.XAdditionalProperties, "radiusShareSecret")
	delete(tmpType.XAdditionalProperties, "realms")
	delete(tmpType.XAdditionalProperties, "requestTimeout")
	delete(tmpType.XAdditionalProperties, "responseFail")
	delete(tmpType.XAdditionalProperties, "respWindow")
	delete(tmpType.XAdditionalProperties, "retryPriInvl")
	delete(tmpType.XAdditionalProperties, "reviveInterval")
	delete(tmpType.XAdditionalProperties, "searchFilter")
	delete(tmpType.XAdditionalProperties, "secondaryRadiusEnabled")
	delete(tmpType.XAdditionalProperties, "secondaryRadiusIP")
	delete(tmpType.XAdditionalProperties, "secondaryRadiusIPv6")
	delete(tmpType.XAdditionalProperties, "secondaryRadiusPort")
	delete(tmpType.XAdditionalProperties, "secondaryRadiusProtocol")
	delete(tmpType.XAdditionalProperties, "secondaryRadiusShareSecret")
	delete(tmpType.XAdditionalProperties, "standbyAdminDomainName")
	delete(tmpType.XAdditionalProperties, "standbyDomainName")
	delete(tmpType.XAdditionalProperties, "standbyEnableSecondaryRadius")
	delete(tmpType.XAdditionalProperties, "standbyGlobalCatalog")
	delete(tmpType.XAdditionalProperties, "standbyIp")
	delete(tmpType.XAdditionalProperties, "standbyIpv6")
	delete(tmpType.XAdditionalProperties, "standbyKeyAttribute")
	delete(tmpType.XAdditionalProperties, "standbyPassword")
	delete(tmpType.XAdditionalProperties, "standbyPort")
	delete(tmpType.XAdditionalProperties, "standbyRadiusIP")
	delete(tmpType.XAdditionalProperties, "standbyRadiusIPv6")
	delete(tmpType.XAdditionalProperties, "standbyRadiusPort")
	delete(tmpType.XAdditionalProperties, "standbyRadiusShareSecret")
	delete(tmpType.XAdditionalProperties, "standbySearchFilter")
	delete(tmpType.XAdditionalProperties, "standbySecondaryRadiusIP")
	delete(tmpType.XAdditionalProperties, "standbySecondaryRadiusIPv6")
	delete(tmpType.XAdditionalProperties, "standbySecondaryRadiusPort")
	delete(tmpType.XAdditionalProperties, "standbySecondaryRadiusShareSecret")
	delete(tmpType.XAdditionalProperties, "standbyServerEnabled")
	delete(tmpType.XAdditionalProperties, "standbyWindowsDomainName")
	delete(tmpType.XAdditionalProperties, "tacacsService")
	delete(tmpType.XAdditionalProperties, "tenantUUID")
	delete(tmpType.XAdditionalProperties, "tlsEnabled")
	delete(tmpType.XAdditionalProperties, "type")
	delete(tmpType.XAdditionalProperties, "validAAAServer")
	delete(tmpType.XAdditionalProperties, "validIP")
	delete(tmpType.XAdditionalProperties, "validIPv6")
	delete(tmpType.XAdditionalProperties, "validRadiusIp")
	delete(tmpType.XAdditionalProperties, "validRadiusIpv6")
	delete(tmpType.XAdditionalProperties, "validSecondaryRadiusIP")
	delete(tmpType.XAdditionalProperties, "validSecondaryRadiusIPv6")
	delete(tmpType.XAdditionalProperties, "windowsDomainName")
	delete(tmpType.XAdditionalProperties, "zombiePeriod")
	*t = WSGAdminSCGScgAaaServerEntry(*tmpType)
	return nil
}

func (t *WSGAdminSCGScgAaaServerEntry) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.AdminDomainName != nil {
		tmp["adminDomainName"] = t.AdminDomainName
	}
	if t.AuditLabel != nil {
		tmp["auditLabel"] = t.AuditLabel
	}
	if t.AuthType != nil {
		tmp["authType"] = t.AuthType
	}
	if t.ClientCertId != nil {
		tmp["clientCertId"] = t.ClientCertId
	}
	if t.CnSanIdentity != nil {
		tmp["cnSanIdentity"] = t.CnSanIdentity
	}
	if t.CreateOn != nil {
		tmp["createOn"] = t.CreateOn
	}
	if t.CreateUserName != nil {
		tmp["createUserName"] = t.CreateUserName
	}
	if t.CreatorUUID != nil {
		tmp["creatorUUID"] = t.CreatorUUID
	}
	if t.DefaultAdmin != nil {
		tmp["defaultAdmin"] = t.DefaultAdmin
	}
	if t.DefaultUserGroup != nil {
		tmp["defaultUserGroup"] = t.DefaultUserGroup
	}
	if t.DefaultUserGroupEnabled != nil {
		tmp["defaultUserGroupEnabled"] = t.DefaultUserGroupEnabled
	}
	if t.Description != nil {
		tmp["description"] = t.Description
	}
	if t.DomainId != nil {
		tmp["domainId"] = t.DomainId
	}
	if t.DomainName != nil {
		tmp["domainName"] = t.DomainName
	}
	if t.EnableSecondaryRadius != nil {
		tmp["enableSecondaryRadius"] = t.EnableSecondaryRadius
	}
	if t.GlobalCatalog != nil {
		tmp["globalCatalog"] = t.GlobalCatalog
	}
	if t.Id != nil {
		tmp["id"] = t.Id
	}
	if t.Ip != nil {
		tmp["ip"] = t.Ip
	}
	if t.IpMode != nil {
		tmp["ipMode"] = t.IpMode
	}
	if t.Ipv6 != nil {
		tmp["ipv6"] = t.Ipv6
	}
	if t.IsConflict != nil {
		tmp["isConflict"] = t.IsConflict
	}
	if t.Key != nil {
		tmp["key"] = t.Key
	}
	if t.KeyAttribute != nil {
		tmp["keyAttribute"] = t.KeyAttribute
	}
	if t.MaxRetries != nil {
		tmp["maxRetries"] = t.MaxRetries
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.NullValueColumnNames != nil {
		tmp["nullValueColumnNames"] = t.NullValueColumnNames
	}
	if t.OcspUrl != nil {
		tmp["ocspUrl"] = t.OcspUrl
	}
	if t.Password != nil {
		tmp["password"] = t.Password
	}
	if t.Port != nil {
		tmp["port"] = t.Port
	}
	if t.ProxyAgentEnabled != nil {
		tmp["proxyAgentEnabled"] = t.ProxyAgentEnabled
	}
	if t.ProxyUserPassword != nil {
		tmp["proxyUserPassword"] = t.ProxyUserPassword
	}
	if t.ProxyUserPrincipalName != nil {
		tmp["proxyUserPrincipalName"] = t.ProxyUserPrincipalName
	}
	if t.RadiusIP != nil {
		tmp["radiusIP"] = t.RadiusIP
	}
	if t.RadiusIPv6 != nil {
		tmp["radiusIPv6"] = t.RadiusIPv6
	}
	if t.RadiusPort != nil {
		tmp["radiusPort"] = t.RadiusPort
	}
	if t.RadiusProtocol != nil {
		tmp["radiusProtocol"] = t.RadiusProtocol
	}
	if t.RadiusRealm != nil {
		tmp["radiusRealm"] = t.RadiusRealm
	}
	if t.RadiusShareSecret != nil {
		tmp["radiusShareSecret"] = t.RadiusShareSecret
	}
	if t.Realms != nil {
		tmp["realms"] = t.Realms
	}
	if t.RequestTimeout != nil {
		tmp["requestTimeout"] = t.RequestTimeout
	}
	if t.ResponseFail != nil {
		tmp["responseFail"] = t.ResponseFail
	}
	if t.RespWindow != nil {
		tmp["respWindow"] = t.RespWindow
	}
	if t.RetryPriInvl != nil {
		tmp["retryPriInvl"] = t.RetryPriInvl
	}
	if t.ReviveInterval != nil {
		tmp["reviveInterval"] = t.ReviveInterval
	}
	if t.SearchFilter != nil {
		tmp["searchFilter"] = t.SearchFilter
	}
	if t.SecondaryRadiusEnabled != nil {
		tmp["secondaryRadiusEnabled"] = t.SecondaryRadiusEnabled
	}
	if t.SecondaryRadiusIP != nil {
		tmp["secondaryRadiusIP"] = t.SecondaryRadiusIP
	}
	if t.SecondaryRadiusIPv6 != nil {
		tmp["secondaryRadiusIPv6"] = t.SecondaryRadiusIPv6
	}
	if t.SecondaryRadiusPort != nil {
		tmp["secondaryRadiusPort"] = t.SecondaryRadiusPort
	}
	if t.SecondaryRadiusProtocol != nil {
		tmp["secondaryRadiusProtocol"] = t.SecondaryRadiusProtocol
	}
	if t.SecondaryRadiusShareSecret != nil {
		tmp["secondaryRadiusShareSecret"] = t.SecondaryRadiusShareSecret
	}
	if t.StandbyAdminDomainName != nil {
		tmp["standbyAdminDomainName"] = t.StandbyAdminDomainName
	}
	if t.StandbyDomainName != nil {
		tmp["standbyDomainName"] = t.StandbyDomainName
	}
	if t.StandbyEnableSecondaryRadius != nil {
		tmp["standbyEnableSecondaryRadius"] = t.StandbyEnableSecondaryRadius
	}
	if t.StandbyGlobalCatalog != nil {
		tmp["standbyGlobalCatalog"] = t.StandbyGlobalCatalog
	}
	if t.StandbyIp != nil {
		tmp["standbyIp"] = t.StandbyIp
	}
	if t.StandbyIpv6 != nil {
		tmp["standbyIpv6"] = t.StandbyIpv6
	}
	if t.StandbyKeyAttribute != nil {
		tmp["standbyKeyAttribute"] = t.StandbyKeyAttribute
	}
	if t.StandbyPassword != nil {
		tmp["standbyPassword"] = t.StandbyPassword
	}
	if t.StandbyPort != nil {
		tmp["standbyPort"] = t.StandbyPort
	}
	if t.StandbyRadiusIP != nil {
		tmp["standbyRadiusIP"] = t.StandbyRadiusIP
	}
	if t.StandbyRadiusIPv6 != nil {
		tmp["standbyRadiusIPv6"] = t.StandbyRadiusIPv6
	}
	if t.StandbyRadiusPort != nil {
		tmp["standbyRadiusPort"] = t.StandbyRadiusPort
	}
	if t.StandbyRadiusShareSecret != nil {
		tmp["standbyRadiusShareSecret"] = t.StandbyRadiusShareSecret
	}
	if t.StandbySearchFilter != nil {
		tmp["standbySearchFilter"] = t.StandbySearchFilter
	}
	if t.StandbySecondaryRadiusIP != nil {
		tmp["standbySecondaryRadiusIP"] = t.StandbySecondaryRadiusIP
	}
	if t.StandbySecondaryRadiusIPv6 != nil {
		tmp["standbySecondaryRadiusIPv6"] = t.StandbySecondaryRadiusIPv6
	}
	if t.StandbySecondaryRadiusPort != nil {
		tmp["standbySecondaryRadiusPort"] = t.StandbySecondaryRadiusPort
	}
	if t.StandbySecondaryRadiusShareSecret != nil {
		tmp["standbySecondaryRadiusShareSecret"] = t.StandbySecondaryRadiusShareSecret
	}
	if t.StandbyServerEnabled != nil {
		tmp["standbyServerEnabled"] = t.StandbyServerEnabled
	}
	if t.StandbyWindowsDomainName != nil {
		tmp["standbyWindowsDomainName"] = t.StandbyWindowsDomainName
	}
	if t.TacacsService != nil {
		tmp["tacacsService"] = t.TacacsService
	}
	if t.TenantUUID != nil {
		tmp["tenantUUID"] = t.TenantUUID
	}
	if t.TlsEnabled != nil {
		tmp["tlsEnabled"] = t.TlsEnabled
	}
	if t.Type != nil {
		tmp["type"] = t.Type
	}
	if t.ValidAAAServer != nil {
		tmp["validAAAServer"] = t.ValidAAAServer
	}
	if t.ValidIP != nil {
		tmp["validIP"] = t.ValidIP
	}
	if t.ValidIPv6 != nil {
		tmp["validIPv6"] = t.ValidIPv6
	}
	if t.ValidRadiusIp != nil {
		tmp["validRadiusIp"] = t.ValidRadiusIp
	}
	if t.ValidRadiusIpv6 != nil {
		tmp["validRadiusIpv6"] = t.ValidRadiusIpv6
	}
	if t.ValidSecondaryRadiusIP != nil {
		tmp["validSecondaryRadiusIP"] = t.ValidSecondaryRadiusIP
	}
	if t.ValidSecondaryRadiusIPv6 != nil {
		tmp["validSecondaryRadiusIPv6"] = t.ValidSecondaryRadiusIPv6
	}
	if t.WindowsDomainName != nil {
		tmp["windowsDomainName"] = t.WindowsDomainName
	}
	if t.ZombiePeriod != nil {
		tmp["zombiePeriod"] = t.ZombiePeriod
	}
	return json.Marshal(tmp)
}

func NewWSGAdminSCGScgAaaServerEntry() *WSGAdminSCGScgAaaServerEntry {
	m := new(WSGAdminSCGScgAaaServerEntry)
	return m
}

// FindScgAdminAaaServers
//
// Operation ID: findScgAdminAaaServers
//
// Use this API command to retrieve a list of registered AAA servers
func (s *WSGAdminSCGService) FindScgAdminAaaServers(ctx context.Context, mutators ...RequestMutator) (*WSGAdminSCGScgAaaServer, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGAdminSCGScgAaaServer()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
