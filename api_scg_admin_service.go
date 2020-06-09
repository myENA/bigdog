package bigdog

// this service is manually created and maintained

import (
	"context"
	"net/http"
)

const (
	RouteSCGAdminAAAServers = "/wsg/api/scg/aaaServers/admin"
	RouteSCGAdminAAAServer  = "/wsg/api/scg/aaaServers/admin/{id}"
)

type SCGAdminService struct {
	apiClient *VSZClient
}

func NewSCGAdminService(c *VSZClient) *SCGAdminService {
	s := new(SCGAdminService)
	s.apiClient = c
	return s
}

type SCGAdminAAAServer struct {
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

func NewSCGAdminAAAServer() *SCGAdminAAAServer {
	m := new(SCGAdminAAAServer)
	return m
}

type SCGAdminFindAAAServersResponse struct {
	Data *struct {
		HasMore           bool                 `json:"hasMore"`
		List              []*SCGAdminAAAServer `json:"list"`
		RawDataTotalCount float64              `json:"rawDataTotalCount"`
		TotalCount        float64              `json:"totalCount"`
	} `json:"data"`
	Error    interface{} `json:"error"`
	Extra    interface{} `json:"extra"`
	MetaData *struct {
		Fields          []map[string]interface{} `json:"fields"`
		IDProperty      string                   `json:"idProperty"`
		MessageProperty string                   `json:"messageProperty"`
		Root            string                   `json:"root"`
		SuccessProperty string                   `json:"successProperty"`
		TotalProperty   string                   `json:"totalProperty"`
	} `json:"metaData"`
	Success bool `json:"success"`
}

func NewSCGAdminFindAAAServersResponse() *SCGAdminFindAAAServersResponse {
	m := new(SCGAdminFindAAAServersResponse)
	return m
}

type SCGAdminLDAPServerDefinition struct {
	Type string `json:"type" validate:"required,oneof=LDAP"`

	// AdminDomainName
	// Constraints:
	//    - required
	AdminDomainName WSGCommonNormalName2to128 `json:"adminDomainName" validate:"required,max=128,min=2"`

	// DomainName
	// Constraints:
	//    - required
	DomainName WSGCommonNormalName2to64 `json:"domainName" validate:"required,max=64,min=2"`

	DefaultUserGroupEnabled bool `json:"defaultUserGroupEnabled"`

	Description WSGCommonDescription `json:"description"`

	// Ip
	// Constraints:
	//    - required
	Ip WSGCommonIpAddress `json:"ip" validate:"required"`

	// KeyAttribute
	// Constraints:
	//    - required
	KeyAttribute WSGCommonNormalName2to64 `json:"keyAttribute" validate:"required,max=64,min=2"`

	// Name
	// Constraints:
	//    - required
	Name WSGCommonNormalName `json:"name" validate:"required,max=32,min=2"`

	// Password
	// Admin password
	// Constraints:
	//    - required
	Password string `json:"password" validate:"required"`

	// Port
	// Port
	// Constraints:
	//    - required
	//    - default:389
	//    - min:1
	//    - max:65535
	Port int `json:"port" validate:"required,gte=1,lte=65535"`

	// SearchFilter
	// Constraints:
	//    - required
	SearchFilter WSGCommonNormalName2to64 `json:"searchFilter" validate:"required,max=64,min=2"`

	RadiusRealm string `json:"radiusRealm" validate:"required"`

	Key string `json:"key"`
}

func NewSCGAdminLDAPServerDefinition() *SCGAdminLDAPServerDefinition {
	m := new(SCGAdminLDAPServerDefinition)
	return m
}

type SCGAdminCreateLDAPAAAServerResponse struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error"`
	Data    *struct {
		Key         string      `json:"key"`
		Label       string      `json:"label"`
		ExtraValues interface{} `json:"extraValues"`
		Type        interface{} `json:"type"`
		ZoneName    *string     `json:"zone_name"`
	} `json:"data"`
	Extra interface{} `json:"extra"`
}

func NewSCGAdminCreateLDAPAAAServerResponse() *SCGAdminCreateLDAPAAAServerResponse {
	m := new(SCGAdminCreateLDAPAAAServerResponse)
	return m
}

type SCGAdminGenericResponse struct {
	Success bool        `json:"success"`
	Error   interface{} `json:"error"`
}

func NewSCGAdminGenericResponse() *SCGAdminGenericResponse {
	m := new(SCGAdminGenericResponse)
	return m
}

func (s *SCGAdminService) FindAAAServers(ctx context.Context) (*SCGAdminFindAAAServersResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCGAdminFindAAAServersResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSCGAdminAAAServers, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCGAdminFindAAAServersResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

func (s *SCGAdminService) CreateLDAPAAAServer(ctx context.Context, body *SCGAdminLDAPServerDefinition) (*SCGAdminCreateLDAPAAAServerResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCGAdminCreateLDAPAAAServerResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	// manually set this for now.
	if body != nil {
		body.Type = "LDAP"
	}
	req = NewAPIRequest(http.MethodPost, RouteSCGAdminAAAServers, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.Headers().Set(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCGAdminCreateLDAPAAAServerResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

func (s *SCGAdminService) UpdateLDAPAAAServer(ctx context.Context, id string, body *SCGAdminLDAPServerDefinition) (*SCGAdminGenericResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCGAdminGenericResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	// manually set this for now.
	if body != nil {
		body.Type = "LDAP"
	}
	req = NewAPIRequest(http.MethodPut, RouteSCGAdminAAAServer, true)
	req.SetPathParameter("id", id)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.Headers().Set(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCGAdminGenericResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

func (s *SCGAdminService) DeleteAAAServer(ctx context.Context, id string) (*SCGAdminGenericResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SCGAdminGenericResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCGAdminAAAServer, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCGAdminGenericResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
