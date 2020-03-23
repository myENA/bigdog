package vsz

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
	apiClient *APIClient
}

func NewSCGAdminService(c *APIClient) *SCGAdminService {
	s := new(SCGAdminService)
	s.apiClient = c
	return s
}

type SCGAdminAAAServer struct {
	AdminDomainName                   string        `json:"adminDomainName"`
	AuditLabel                        string        `json:"auditLabel"`
	AuthType                          string        `json:"authType"`
	ClientCertID                      string        `json:"clientCertId"`
	CnSanIdentity                     string        `json:"cnSanIdentity"`
	CreateOn                          int           `json:"createOn"`
	CreateUserName                    string        `json:"createUserName"`
	CreatorUUID                       string        `json:"creatorUUID"`
	DefaultAdmin                      string        `json:"defaultAdmin"`
	DefaultUserGroup                  string        `json:"defaultUserGroup"`
	DefaultUserGroupEnabled           bool          `json:"defaultUserGroupEnabled"`
	Description                       string        `json:"description"`
	DomainID                          string        `json:"domainId"`
	DomainName                        string        `json:"domainName"`
	EnableSecondaryRadius             int           `json:"enableSecondaryRadius"`
	GlobalCatalog                     bool          `json:"globalCatalog"`
	ID                                string        `json:"id"`
	IP                                string        `json:"ip"`
	IPMode                            string        `json:"ipMode"`
	Ipv6                              string        `json:"ipv6"`
	IsConflict                        int           `json:"isConflict"`
	Key                               string        `json:"key"`
	KeyAttribute                      string        `json:"keyAttribute"`
	MaxRetries                        int           `json:"maxRetries"`
	Name                              string        `json:"name"`
	NullValueColumnNames              []interface{} `json:"nullValueColumnNames"`
	OcspURL                           string        `json:"ocspUrl"`
	Password                          string        `json:"password"`
	Port                              int           `json:"port"`
	ProxyAgentEnabled                 bool          `json:"proxyAgentEnabled"`
	ProxyUserPassword                 string        `json:"proxyUserPassword"`
	ProxyUserPrincipalName            string        `json:"proxyUserPrincipalName"`
	RadiusIP                          string        `json:"radiusIP"`
	RadiusIPv6                        string        `json:"radiusIPv6"`
	RadiusPort                        int           `json:"radiusPort"`
	RadiusProtocol                    string        `json:"radiusProtocol"`
	RadiusRealm                       string        `json:"radiusRealm"`
	RadiusShareSecret                 string        `json:"radiusShareSecret"`
	Realms                            []string      `json:"realms"`
	RequestTimeout                    float64       `json:"requestTimeout"`
	RespWindow                        float64       `json:"respWindow"`
	ResponseFail                      bool          `json:"responseFail"`
	RetryPriInvl                      float64       `json:"retryPriInvl"`
	ReviveInterval                    float64       `json:"reviveInterval"`
	SearchFilter                      string        `json:"searchFilter"`
	SecondaryRadiusEnabled            bool          `json:"secondaryRadiusEnabled"`
	SecondaryRadiusIP                 string        `json:"secondaryRadiusIP"`
	SecondaryRadiusIPv6               string        `json:"secondaryRadiusIPv6"`
	SecondaryRadiusPort               int           `json:"secondaryRadiusPort"`
	SecondaryRadiusProtocol           string        `json:"secondaryRadiusProtocol"`
	SecondaryRadiusShareSecret        string        `json:"secondaryRadiusShareSecret"`
	StandbyAdminDomainName            string        `json:"standbyAdminDomainName"`
	StandbyDomainName                 string        `json:"standbyDomainName"`
	StandbyEnableSecondaryRadius      string        `json:"standbyEnableSecondaryRadius"`
	StandbyGlobalCatalog              bool          `json:"standbyGlobalCatalog"`
	StandbyIP                         string        `json:"standbyIp"`
	StandbyIpv6                       string        `json:"standbyIpv6"`
	StandbyKeyAttribute               interface{}   `json:"standbyKeyAttribute"`
	StandbyPassword                   string        `json:"standbyPassword"`
	StandbyPort                       int           `json:"standbyPort"`
	StandbyRadiusIP                   string        `json:"standbyRadiusIP"`
	StandbyRadiusIPv6                 string        `json:"standbyRadiusIPv6"`
	StandbyRadiusPort                 int           `json:"standbyRadiusPort"`
	StandbyRadiusShareSecret          string        `json:"standbyRadiusShareSecret"`
	StandbySearchFilter               string        `json:"standbySearchFilter"`
	StandbySecondaryRadiusIP          string        `json:"standbySecondaryRadiusIP"`
	StandbySecondaryRadiusIPv6        string        `json:"standbySecondaryRadiusIPv6"`
	StandbySecondaryRadiusPort        int           `json:"standbySecondaryRadiusPort"`
	StandbySecondaryRadiusShareSecret string        `json:"standbySecondaryRadiusShareSecret"`
	StandbyServerEnabled              bool          `json:"standbyServerEnabled"`
	StandbyWindowsDomainName          string        `json:"standbyWindowsDomainName"`
	TacacsService                     string        `json:"tacacsService"`
	TenantUUID                        string        `json:"tenantUUID"`
	TLSEnabled                        bool          `json:"tlsEnabled"`
	Type                              string        `json:"type"`
	ValidAAAServer                    bool          `json:"validAAAServer"`
	ValidIP                           bool          `json:"validIP"`
	ValidIPv6                         bool          `json:"validIPv6"`
	ValidRadiusIP                     bool          `json:"validRadiusIp"`
	ValidRadiusIpv6                   bool          `json:"validRadiusIpv6"`
	ValidSecondaryRadiusIP            bool          `json:"validSecondaryRadiusIP"`
	ValidSecondaryRadiusIPv6          bool          `json:"validSecondaryRadiusIPv6"`
	WindowsDomainName                 string        `json:"windowsDomainName"`
	ZombiePeriod                      int           `json:"zombiePeriod"`
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
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	// manually set this for now.
	body.Type = "LDAP"
	req = NewAPIRequest(http.MethodPost, RouteSCGAdminAAAServers, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCGAdminCreateLDAPAAAServerResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	// manually set this for now.
	body.Type = "LDAP"
	req = NewAPIRequest(http.MethodPut, RouteSCGAdminAAAServer, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCGAdminGenericResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSCGAdminAAAServer, true)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSCGAdminGenericResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
