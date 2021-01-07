package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGSystemIPsecService struct {
	apiClient *VSZClient
}

func NewWSGSystemIPsecService(c *VSZClient) *WSGSystemIPsecService {
	s := new(WSGSystemIPsecService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSystemIPsecService() *WSGSystemIPsecService {
	return NewWSGSystemIPsecService(ss.apiClient)
}

// WSGSystemIPsecGetResult
//
// Definition: systemIPsec_getResult
type WSGSystemIPsecGetResult struct {
	// AuthType
	// Authentication type
	AuthType *string `json:"authType,omitempty"`

	// CertSubject
	// Assign specific cert subject
	CertSubject *string `json:"certSubject,omitempty"`

	// EspProposals
	// Proposal of Encapsulating Security Payload
	EspProposals []*WSGSystemIPsecProposal `json:"espProposals,omitempty"`

	// EspRekeyDisabled
	// Disable rekey mechanisam of Encapsulating Security Payload
	EspRekeyDisabled *bool `json:"espRekeyDisabled,omitempty"`

	// EspRekeyTime
	// Rekey time of Encapsulating Security Payload
	EspRekeyTime *int `json:"espRekeyTime,omitempty"`

	// EspRekeyTimeUnit
	// Rekey time unit of Encapsulating Security Payload
	EspRekeyTimeUnit *string `json:"espRekeyTimeUnit,omitempty"`

	// IkeProposals
	// Proposal of Internet Key Exchange
	IkeProposals []*WSGSystemIPsecProposal `json:"ikeProposals,omitempty"`

	// IkeRekeyDisabled
	// Disable rekey mechanisam of Internet Key Exchange
	IkeRekeyDisabled *bool `json:"ikeRekeyDisabled,omitempty"`

	// IkeRekeyTime
	// Rekey time of Internet Key Exchange
	IkeRekeyTime *int `json:"ikeRekeyTime,omitempty"`

	// IkeRekeyTimeUnit
	// Rekey time unit of Internet Key Exchange
	IkeRekeyTimeUnit *string `json:"ikeRekeyTimeUnit,omitempty"`

	// IpSecEnabled
	// Enable System IPSec
	IpSecEnabled *bool `json:"ipSecEnabled,omitempty"`

	// OcspEnabled
	// Enable OCSP
	OcspEnabled *bool `json:"ocspEnabled,omitempty"`

	// OcspServerUri
	// The URI of OCSP server
	OcspServerUri *string `json:"ocspServerUri,omitempty"`

	// PreSharedKey
	// Pre-shared key
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// RemoteId
	// The IP of IPSec receiver
	RemoteId *string `json:"remoteId,omitempty"`

	// ScgCertId
	// SCG client certification id.
	ScgCertId *string `json:"scgCertId,omitempty"`

	// SecurityGateway
	// Security gateway IP
	SecurityGateway *string `json:"securityGateway,omitempty"`

	// SubnetMask
	// Subnet Mask of security gateway
	SubnetMask *string `json:"subnetMask,omitempty"`

	// TrustChainProfileId
	// Assign trust chain profile id
	TrustChainProfileId *string `json:"trustChainProfileId,omitempty"`
}

type WSGSystemIPsecGetResultAPIResponse struct {
	*RawAPIResponse
	Data *WSGSystemIPsecGetResult
}

func newWSGSystemIPsecGetResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGSystemIPsecGetResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGSystemIPsecGetResultAPIResponse) Hydrate() error {
	r.Data = new(WSGSystemIPsecGetResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSystemIPsecGetResult() *WSGSystemIPsecGetResult {
	m := new(WSGSystemIPsecGetResult)
	return m
}

// WSGSystemIPsecProposal
//
// Definition: systemIPsec_proposal
type WSGSystemIPsecProposal struct {
	// AuthAlg
	// Authentication algorithm
	// Constraints:
	//    - required
	AuthAlg *string `json:"authAlg"`

	// EncAlg
	// Encrytion algorithm
	// Constraints:
	//    - required
	EncAlg *string `json:"encAlg"`
}

func NewWSGSystemIPsecProposal() *WSGSystemIPsecProposal {
	m := new(WSGSystemIPsecProposal)
	return m
}

// WSGSystemIPsecUpdate
//
// Definition: systemIPsec_update
type WSGSystemIPsecUpdate struct {
	// AuthType
	// Authentication type
	AuthType *string `json:"authType,omitempty"`

	// CertSubject
	// Assign specific cert subject
	CertSubject *string `json:"certSubject,omitempty"`

	// EspProposals
	// Proposal of Encapsulating Security Payload
	EspProposals []*WSGSystemIPsecProposal `json:"espProposals,omitempty"`

	// EspRekeyDisabled
	// Disable rekey mechanisam of Encapsulating Security Payload
	EspRekeyDisabled *bool `json:"espRekeyDisabled,omitempty"`

	// EspRekeyTime
	// Rekey time of Encapsulating Security Payload
	EspRekeyTime *int `json:"espRekeyTime,omitempty"`

	// EspRekeyTimeUnit
	// Rekey time unit of Encapsulating Security Payload
	EspRekeyTimeUnit *string `json:"espRekeyTimeUnit,omitempty"`

	// IkeProposals
	// Proposal of Internet Key Exchange
	IkeProposals []*WSGSystemIPsecProposal `json:"ikeProposals,omitempty"`

	// IkeRekeyDisabled
	// Disable rekey mechanisam of Internet Key Exchange
	IkeRekeyDisabled *bool `json:"ikeRekeyDisabled,omitempty"`

	// IkeRekeyTime
	// Rekey time of Internet Key Exchange
	IkeRekeyTime *int `json:"ikeRekeyTime,omitempty"`

	// IkeRekeyTimeUnit
	// Rekey time unit of Internet Key Exchange
	IkeRekeyTimeUnit *string `json:"ikeRekeyTimeUnit,omitempty"`

	// IpSecEnabled
	// Enable System IPSec
	// Constraints:
	//    - required
	IpSecEnabled *bool `json:"ipSecEnabled"`

	// OcspEnabled
	// Enable OCSP
	OcspEnabled *bool `json:"ocspEnabled,omitempty"`

	// OcspServerUri
	// The URI of OCSP server
	OcspServerUri *string `json:"ocspServerUri,omitempty"`

	// PreSharedKey
	// Pre-shared key
	PreSharedKey *string `json:"preSharedKey,omitempty"`

	// RemoteId
	// The IP of IPSec receiver
	RemoteId *string `json:"remoteId,omitempty"`

	// ScgCertId
	// SCG client certification id.
	ScgCertId *string `json:"scgCertId,omitempty"`

	// SecurityGateway
	// Security gateway IP
	SecurityGateway *string `json:"securityGateway,omitempty"`

	// SubnetMask
	// Subnet Mask of security gateway
	SubnetMask *string `json:"subnetMask,omitempty"`

	// TrustChainProfileId
	// Assign trust chain profile id
	TrustChainProfileId *string `json:"trustChainProfileId,omitempty"`
}

func NewWSGSystemIPsecUpdate() *WSGSystemIPsecUpdate {
	m := new(WSGSystemIPsecUpdate)
	return m
}

// FindSystemIpsec
//
// Operation ID: findSystemIpsec
//
// Use this API command to retrieve the System IPSec.
func (s *WSGSystemIPsecService) FindSystemIpsec(ctx context.Context, mutators ...RequestMutator) (*WSGSystemIPsecGetResultAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemIPsecGetResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSystemIpsec, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSystemIPsecGetResult()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateSystemIpsec
//
// Operation ID: updateSystemIpsec
//
// Use this API command to modify the System IPSec.
//
// Request Body:
//	 - body *WSGSystemIPsecUpdate
func (s *WSGSystemIPsecService) UpdateSystemIpsec(ctx context.Context, body *WSGSystemIPsecUpdate, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *RawAPIResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteWSGUpdateSystemIpsec, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = new(RawAPIResponse)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
