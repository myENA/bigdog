package vsz

import (
	"context"
	"net/http"
)

type CustomService struct {
	apiClient *APIClient
}

func NewCustomService(c *APIClient) *CustomService {
	s := new(CustomService)
	s.apiClient = c
	return s
}

type AdminScgAaaServersDataList struct {
	AdminDomainName                   *string       `json:"adminDomainName,omitempty"`
	AuditLabel                        *string       `json:"auditLabel,omitempty"`
	AuthType                          *string       `json:"authType,omitempty"`
	ClientCertID                      *string       `json:"clientCertId,omitempty"`
	CnSanIdentity                     *string       `json:"cnSanIdentity,omitempty"`
	CreateOn                          *int          `json:"createOn,omitempty"`
	CreateUserName                    *string       `json:"createUserName,omitempty"`
	CreatorUUID                       *string       `json:"creatorUUID,omitempty"`
	DefaultAdmin                      *string       `json:"defaultAdmin,omitempty"`
	DefaultUserGroup                  *string       `json:"defaultUserGroup,omitempty"`
	DefaultUserGroupEnabled           *bool         `json:"defaultUserGroupEnabled,omitempty"`
	Description                       *string       `json:"description,omitempty"`
	DomainID                          *string       `json:"domainId,omitempty"`
	DomainName                        *string       `json:"domainName,omitempty"`
	EnableSecondaryRadius             *int          `json:"enableSecondaryRadius,omitempty"`
	GlobalCatalog                     *bool         `json:"globalCatalog,omitempty"`
	ID                                *string       `json:"id,omitempty"`
	IP                                *string       `json:"ip,omitempty"`
	IPMode                            *string       `json:"ipMode,omitempty"`
	Ipv6                              *string       `json:"ipv6,omitempty"`
	IsConflict                        *int          `json:"isConflict,omitempty"`
	Key                               *string       `json:"key,omitempty"`
	KeyAttribute                      *string       `json:"keyAttribute,omitempty"`
	MaxRetries                        *int          `json:"maxRetries,omitempty"`
	Name                              *string       `json:"name,omitempty"`
	NullValueColumnNames              []interface{} `json:"nullValueColumnNames,omitempty"`
	OcspURL                           *string       `json:"ocspUrl,omitempty"`
	Password                          *string       `json:"password,omitempty"`
	Port                              *int          `json:"port,omitempty"`
	ProxyAgentEnabled                 *bool         `json:"proxyAgentEnabled,omitempty"`
	ProxyUserPassword                 *string       `json:"proxyUserPassword,omitempty"`
	ProxyUserPrincipalName            *string       `json:"proxyUserPrincipalName,omitempty"`
	RadiusIP                          *string       `json:"radiusIP,omitempty"`
	RadiusIPv6                        *string       `json:"radiusIPv6,omitempty"`
	RadiusPort                        *int          `json:"radiusPort,omitempty"`
	RadiusProtocol                    *string       `json:"radiusProtocol,omitempty"`
	RadiusRealm                       *string       `json:"radiusRealm,omitempty"`
	RadiusShareSecret                 *string       `json:"radiusShareSecret,omitempty"`
	Realms                            []string      `json:"realms,omitempty"`
	RequestTimeout                    *float64      `json:"requestTimeout,omitempty"`
	RespWindow                        *float64      `json:"respWindow,omitempty"`
	ResponseFail                      *bool         `json:"responseFail,omitempty"`
	RetryPriInvl                      *float64      `json:"retryPriInvl,omitempty"`
	ReviveInterval                    *float64      `json:"reviveInterval,omitempty"`
	SearchFilter                      *string       `json:"searchFilter,omitempty"`
	SecondaryRadiusEnabled            *bool         `json:"secondaryRadiusEnabled,omitempty"`
	SecondaryRadiusIP                 *string       `json:"secondaryRadiusIP,omitempty"`
	SecondaryRadiusIPv6               *string       `json:"secondaryRadiusIPv6,omitempty"`
	SecondaryRadiusPort               *int          `json:"secondaryRadiusPort,omitempty"`
	SecondaryRadiusProtocol           *string       `json:"secondaryRadiusProtocol,omitempty"`
	SecondaryRadiusShareSecret        *string       `json:"secondaryRadiusShareSecret,omitempty"`
	StandbyAdminDomainName            *string       `json:"standbyAdminDomainName,omitempty"`
	StandbyDomainName                 *string       `json:"standbyDomainName,omitempty"`
	StandbyEnableSecondaryRadius      *string       `json:"standbyEnableSecondaryRadius,omitempty"`
	StandbyGlobalCatalog              *bool         `json:"standbyGlobalCatalog,omitempty"`
	StandbyIP                         *string       `json:"standbyIp,omitempty"`
	StandbyIpv6                       *string       `json:"standbyIpv6,omitempty"`
	StandbyKeyAttribute               interface{}   `json:"standbyKeyAttribute,omitempty"`
	StandbyPassword                   *string       `json:"standbyPassword,omitempty"`
	StandbyPort                       *int          `json:"standbyPort,omitempty"`
	StandbyRadiusIP                   *string       `json:"standbyRadiusIP,omitempty"`
	StandbyRadiusIPv6                 *string       `json:"standbyRadiusIPv6,omitempty"`
	StandbyRadiusPort                 *int          `json:"standbyRadiusPort,omitempty"`
	StandbyRadiusShareSecret          *string       `json:"standbyRadiusShareSecret,omitempty"`
	StandbySearchFilter               *string       `json:"standbySearchFilter,omitempty"`
	StandbySecondaryRadiusIP          *string       `json:"standbySecondaryRadiusIP,omitempty"`
	StandbySecondaryRadiusIPv6        *string       `json:"standbySecondaryRadiusIPv6,omitempty"`
	StandbySecondaryRadiusPort        *int          `json:"standbySecondaryRadiusPort,omitempty"`
	StandbySecondaryRadiusShareSecret *string       `json:"standbySecondaryRadiusShareSecret,omitempty"`
	StandbyServerEnabled              *bool         `json:"standbyServerEnabled,omitempty"`
	StandbyWindowsDomainName          *string       `json:"standbyWindowsDomainName,omitempty"`
	TacacsService                     *string       `json:"tacacsService,omitempty"`
	TenantUUID                        *string       `json:"tenantUUID,omitempty"`
	TLSEnabled                        *bool         `json:"tlsEnabled,omitempty"`
	Type                              *string       `json:"type,omitempty"`
	ValidAAAServer                    *bool         `json:"validAAAServer,omitempty"`
	ValidIP                           *bool         `json:"validIP,omitempty"`
	ValidIPv6                         *bool         `json:"validIPv6,omitempty"`
	ValidRadiusIP                     *bool         `json:"validRadiusIp,omitempty"`
	ValidRadiusIpv6                   *bool         `json:"validRadiusIpv6,omitempty"`
	ValidSecondaryRadiusIP            *bool         `json:"validSecondaryRadiusIP,omitempty"`
	ValidSecondaryRadiusIPv6          *bool         `json:"validSecondaryRadiusIPv6,omitempty"`
	WindowsDomainName                 *string       `json:"windowsDomainName,omitempty"`
	ZombiePeriod                      *int          `json:"zombiePeriod,omitempty"`
}

func NewAdminScgAaaServersDataList() *AdminScgAaaServersDataList {
	m := new(AdminScgAaaServersDataList)
	return m
}

type AdminScgAaaServersData struct {
	HasMore           *bool                         `json:"hasMore,omitempty"`
	List              []*AdminScgAaaServersDataList `json:"list,omitempty"`
	RawDataTotalCount *float64                      `json:"rawDataTotalCount,omitempty"`
	TotalCount        *float64                      `json:"totalCount,omitempty"`
}

func NewAdminScgAaaServersData() *AdminScgAaaServersData {
	m := new(AdminScgAaaServersData)
	return m
}

type AdminScgAaaServersMetaData struct {
	Fields          []map[string]interface{} `json:"fields,omitempty"`
	IDProperty      *string                  `json:"idProperty,omitempty"`
	MessageProperty *string                  `json:"messageProperty,omitempty"`
	Root            *string                  `json:"root,omitempty"`
	SuccessProperty *string                  `json:"successProperty,omitempty"`
	TotalProperty   *string                  `json:"totalProperty,omitempty"`
}

func NewAdminScgAaaServersMetaData() *AdminScgAaaServersMetaData {
	m := new(AdminScgAaaServersMetaData)
	return m
}

type AdminScgAaaServers struct {
	Data     *AdminScgAaaServersData     `json:"data,omitempty"`
	Error    interface{}                 `json:"error,omitempty"`
	Extra    interface{}                 `json:"extra,omitempty"`
	MetaData *AdminScgAaaServersMetaData `json:"metaData,omitempty"`
	Success  *bool                       `json:"success,omitempty"`
}

func NewAdminScgAaaServers() *AdminScgAaaServers {
	m := new(AdminScgAaaServers)
	return m
}

func (s *CustomService) FindAdminScgAaaServers(ctx context.Context) (*AdminScgAaaServers, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *AdminScgAaaServers
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteCustomFindAdminScgAaaServers, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewAdminScgAaaServers()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
