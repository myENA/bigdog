package ruckus

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSystemipsecService struct {
	apiClient *VSZClient
}

func NewWSGSystemipsecService(c *VSZClient) *WSGSystemipsecService {
	s := new(WSGSystemipsecService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSystemipsecService() *WSGSystemipsecService {
	return NewWSGSystemipsecService(ss.apiClient)
}

type WSGSystemipsecGetResult struct {
	// AuthType
	// Authentication type
	AuthType *string `json:"authType,omitempty"`

	// CertSubject
	// Assign specific cert subject
	CertSubject *string `json:"certSubject,omitempty"`

	// EspProposals
	// Proposal of Encapsulating Security Payload
	EspProposals []**WSGSystemipsecGetResult `json:"espProposals,omitempty"`

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
	IkeProposals []**WSGSystemipsecGetResult `json:"ikeProposals,omitempty"`

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

func NewWSGSystemipsecGetResult() *WSGSystemipsecGetResult {
	m := new(WSGSystemipsecGetResult)
	return m
}

type WSGSystemipsecProposal struct {
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

func NewWSGSystemipsecProposal() *WSGSystemipsecProposal {
	m := new(WSGSystemipsecProposal)
	return m
}

type WSGSystemipsecUpdate struct {
	// AuthType
	// Authentication type
	AuthType *string `json:"authType,omitempty"`

	// CertSubject
	// Assign specific cert subject
	CertSubject *string `json:"certSubject,omitempty"`

	// EspProposals
	// Proposal of Encapsulating Security Payload
	EspProposals []**WSGSystemipsecUpdate `json:"espProposals,omitempty"`

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
	IkeProposals []**WSGSystemipsecUpdate `json:"ikeProposals,omitempty"`

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

func NewWSGSystemipsecUpdate() *WSGSystemipsecUpdate {
	m := new(WSGSystemipsecUpdate)
	return m
}

// FindSystemIpsec
//
// Use this API command to retrieve the System IPSec.
func (s *WSGSystemipsecService) FindSystemIpsec(ctx context.Context) (*WSGSystemipsecGetResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSystemipsecGetResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSystemIpsec, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSystemipsecGetResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateSystemIpsec
//
// Use this API command to modify the System IPSec.
//
// Request Body:
//	 - body *WSGSystemipsecUpdate
func (s *WSGSystemipsecService) UpdateSystemIpsec(ctx context.Context, body *WSGSystemipsecUpdate) (interface{}, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     interface{}
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateSystemIpsec, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = new(interface{})
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
