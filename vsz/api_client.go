package vsz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"time"

	"gopkg.in/go-playground/validator.v9"
)

const (
	ResourceNameCluster       = "CLUSTER_CATEGORY"
	ResourceNameAP            = "AP_CATEGORY"
	ResourceNameWLAN          = "WLAN_CATEGORY"
	ResourceNameDevice        = "DEVICE_CATEGORY"
	ResourceNameAdministrator = "ADMINISTRATOR_CATEGORY"
	ResourceNameTenant        = "TENANT_CATEGORY"
	ResourceNameICX           = "ICX_CATEGORY"

	ResourceDisplayCluster       = "SZ"
	ResourceDisplayAP            = "AP"
	ResourceDisplayWLAN          = "WLAN"
	ResourceDisplayDevice        = "User/Device/App"
	ResourceDisplayAdministrator = "Admin"
	ResourceDisplayTenant        = "MVNO"
	ResourceDisplayICX           = "ICX Switch"

	ResourceAccessFullAccess = "FULL_ACCESS"
	ResourceAccessModify     = "MODIFY"
	ResourceAccessReadOnly   = "READ_ONLY"
	ResourceAccessNoAccess   = "NO_ACCESS"

	RoleNameSuperAdmin           = "SUPER_ADMIN"
	RoleNameSystemAdmin          = "SYSTEM_ADMIN"
	RoleNameReadOnlySystemAdmin  = "RO_SYSTEM_ADMIN"
	RoleNameNetworkAdmin         = "NETWORK_ADMIN"
	RoleNameReadOnlyNetworkAdmin = "RO_NETWORK_ADMIN"
	RoleNameAPAdmin              = "AP_ADMIN"
	RoleNameGuestPassAdmin       = "GUEST_PASS_ADMIN"
	RoleNameCustom               = "CUSTOM"

	RoleLabelSuperAdmin           = "Super Admin"
	RoleLabelSystemAdmin          = "System Admin"
	RoleLabelReadOnlySystemAdmin  = "Read-Only System Admin"
	RoleLabelNetworkAdmin         = "Network Admin"
	RoleLabelReadOnlyNetworkAdmin = "Read-Only Network Admin"
	RoleLabelAPAdmin              = "AP Admin"
	RoleLabelGuestPassAdmin       = "Guest Pass Admin"
	RoleLabelCustom               = "SCGAdmin"

	SecurityProfileNameDefault    = "Default"
	SecurityProfileNameMoreSecure = "More Secure"

	logDebugAPIRequestPrepFormat     = "Preparing api request #%d \"%s %s\""
	logDebugAPIRequestNoBodyFormat   = "%s without body"
	logDebugAPIRequestWithBodyFormat = "%s with body: %s"

	serviceTicketQueryParameter = "serviceTicket"
)

var (
	pkgValidator *validator.Validate
)

func init() {
	pkgValidator = validator.New()
}

// PackageValidator returns the global validator instance used by this client
func PackageValidator() *validator.Validate {
	return pkgValidator
}

type APIConfig struct {
	// Address [required]
	//
	// Full address of VSZ, including scheme and port
	Address string `json:"address"`

	// PathPrefix [optional]
	//
	// Custom path prefix to prepend to all api calls.  Default is to leave this blank.
	PathPrefix string `json:"pathPrefix"`

	// Debug [optional]
	//
	// Set to true to enable debug logging
	Debug bool `json:"debug"`

	// ServiceTicketProvider [required]
	//
	// ServiceTicketProvider to use to handle request auth session
	ServiceTicketProvider ServiceTicketProvider `json:"-"`

	// Logger [optional]
	//
	// Logger to use.  Leave blank for no logging
	Logger *log.Logger `json:"-"`

	// HTTPClient [optional]
	HTTPClient *http.Client `json:"-"`
}

type APIClient struct {
	log   *log.Logger
	debug bool

	addr       string
	pathPrefix string
	stp        ServiceTicketProvider

	client *http.Client
}

func NewAPIClient(config *APIConfig) (*APIClient, error) {
	if config.Address == "" {
		return nil, errors.New("address must be defined")
	}
	if config.ServiceTicketProvider == nil {
		return nil, errors.New("authenticator must be defined")
	}

	if _, err := url.Parse(config.Address); err != nil {
		return nil, fmt.Errorf("invalid address specified: %w", err)
	}

	c := new(APIClient)
	c.addr = config.Address
	c.stp = config.ServiceTicketProvider
	c.pathPrefix = config.PathPrefix
	c.debug = config.Debug

	if config.Logger != nil {
		c.log = config.Logger
	} else {
		c.log = log.New(ioutil.Discard, "", 0)
	}

	if config.HTTPClient != nil {
		c.client = config.HTTPClient
	} else {
		// pooled transport config shamelessly borrowed from
		// https://github.com/hashicorp/go-cleanhttp/blob/v0.5.1/cleanhttp.go
		c.client = &http.Client{
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
				DialContext: (&net.Dialer{
					Timeout:   30 * time.Second,
					KeepAlive: 30 * time.Second,
				}).DialContext,
				MaxIdleConns:          100,
				IdleConnTimeout:       90 * time.Second,
				TLSHandshakeTimeout:   10 * time.Second,
				ExpectContinueTimeout: 1 * time.Second,
				MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
			},
		}
	}

	return c, nil
}

// Address returns the address of the VSZ node currently being connected to
func (c *APIClient) Address() string {
	return c.addr
}

// PathPrefix returns the prefix that will be applied to all api calls.  It is valid for this to be blank, and is by
// default
func (c *APIClient) PathPrefix() string {
	return c.pathPrefix
}

// ServiceTicketProvider returns the provider used by this client
func (c *APIClient) ServiceTicketProvider() ServiceTicketProvider {
	return c.stp
}

func (c *APIClient) WSG() *WSGService {
	return NewWSGService(c)
}

func (c *APIClient) SwitchM() *SwitchMService {
	return NewSwitchMService(c)
}

func (c *APIClient) SCGAdmin() *SCGAdminService {
	return NewSCGAdminService(c)
}

func (c *APIClient) Do(ctx context.Context, request *APIRequest) (*http.Response, error) {
	_, httpResponse, err := c.do(ctx, request)
	return httpResponse, err
}

func (c *APIClient) do(ctx context.Context, request *APIRequest) (ServiceTicketCAS, *http.Response, error) {
	var (
		httpRequest   *http.Request
		httpResponse  *http.Response
		cas           ServiceTicketCAS
		serviceTicket string
		err           error
	)
	if request.authenticated {
		if cas, serviceTicket, err = c.stp.Current(); err != nil {
			if cas, err = c.stp.Refresh(ctx, c, cas); err != nil {
				return cas, nil, err
			} else if cas, serviceTicket, err = c.stp.Current(); err != nil {
				return cas, nil, err
			}
		}
	}

	if c.debug {
		logMsg := fmt.Sprintf(logDebugAPIRequestPrepFormat, request.ID(), request.Method(), request.CompiledURI())

		body := request.Body()

		if len(body) == 0 {
			logMsg = fmt.Sprintf(logDebugAPIRequestNoBodyFormat, logMsg)
		} else {
			logMsg = fmt.Sprintf(logDebugAPIRequestWithBodyFormat, logMsg, string(body))
		}

		c.log.Print(logMsg)
	}
	if httpRequest, err = request.toHTTP(ctx, c.addr, c.pathPrefix, serviceTicket); err != nil {
		return cas, nil, err
	}
	httpResponse, err = c.client.Do(httpRequest)
	return cas, httpResponse, err
}

type APIResponseMeta struct {
	RequestMethod string `json:"requestMethod"`
	RequestURI    string `json:"requestURI"`

	SuccessCode int `json:"successCode"`

	ResponseCode   int    `json:"responseCode"`
	ResponseStatus string `json:"responseStatus"`
}

func newAPIResponseMeta(req *APIRequest, successCode int, httpResp *http.Response) *APIResponseMeta {
	rm := new(APIResponseMeta)
	rm.RequestMethod = req.Method()
	rm.RequestURI = req.CompiledURI()
	rm.SuccessCode = successCode
	if httpResp != nil {
		rm.ResponseCode = httpResp.StatusCode
		rm.ResponseStatus = http.StatusText(httpResp.StatusCode)
	}
	return rm
}

func (rm *APIResponseMeta) String() string {
	var msg string
	if rm.SuccessCode == rm.ResponseCode {
		msg = "Successful"
	} else {
		msg = "Failed"
	}
	return fmt.Sprintf("%s response from request %s %s", msg, rm.RequestMethod, rm.RequestURI)
}

func cleanupHTTPResponseBody(hp *http.Response, drain bool) {
	if hp == nil {
		return
	}
	if drain {
		_, _ = io.Copy(ioutil.Discard, hp.Body)
	}
	_ = hp.Body.Close()
}

func handleResponse(req *APIRequest, successCode int, httpResp *http.Response, modelPtr interface{}, sourceErr error) (*APIResponseMeta, error) {
	var finalErr error

	if sourceErr != nil {
		// just in case
		defer cleanupHTTPResponseBody(httpResp, true)

		// if the incoming error is from a service ticket provider, return as-is
		if sterr, ok := sourceErr.(*ServiceTicketProviderError); ok {
			return sterr.ResponseMeta(), sterr
		}

		// otherwise, build response meta and return source error
		return newAPIResponseMeta(req, successCode, httpResp), sourceErr
	}

	if httpResp == nil {
		panic(fmt.Sprintf("severe problem: nil *http.Response seen with nil error. meta: %v", newAPIResponseMeta(req, successCode, httpResp)))
	}

	if httpResp.StatusCode == successCode {
		// if the response code matches the expected "success" code...
		if modelPtr != nil {
			// ... and this query has a modeled response, attempt to unmarshal into that type

			defer cleanupHTTPResponseBody(httpResp, false)

			if err := json.NewDecoder(httpResp.Body).Decode(modelPtr); err != nil {
				finalErr = fmt.Errorf("error unmarshalling response body into %T: %w", modelPtr, err)
			}
		} else {
			defer cleanupHTTPResponseBody(httpResp, true)
		}
	} else {
		defer cleanupHTTPResponseBody(httpResp, false)

		apiErr := new(APIError)
		if err := json.NewDecoder(httpResp.Body).Decode(apiErr); err != nil {
			finalErr = fmt.Errorf("error unmarshalling error body into %T: %w", finalErr, err)
		} else {
			finalErr = apiErr
		}
	}

	// build response.
	return newAPIResponseMeta(req, successCode, httpResp), finalErr
}
