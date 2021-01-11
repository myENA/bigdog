package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
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

func newWSGSystemIPsecGetResultAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSystemIPsecGetResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSystemIPsecGetResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSystemIPsecGetResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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
// Use this API command to retrieve the System IPSec.
//
// Operation ID: findSystemIpsec
// Operation path: /systemIpsec
// Success code: 200 (OK)
func (s *WSGSystemIPsecService) FindSystemIpsec(ctx context.Context, mutators ...RequestMutator) (*WSGSystemIPsecGetResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSystemIPsecGetResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*WSGSystemIPsecGetResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSystemIpsec, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSystemIPsecGetResultAPIResponse), err
}

// UpdateSystemIpsec
//
// Use this API command to modify the System IPSec.
//
// Operation ID: updateSystemIpsec
// Operation path: /systemIpsec
// Success code: 200 (OK)
//
// Request body:
//	 - body *WSGSystemIPsecUpdate
func (s *WSGSystemIPsecService) UpdateSystemIpsec(ctx context.Context, body *WSGSystemIPsecUpdate, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newRawAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPut, RouteWSGUpdateSystemIpsec, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
