package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMTroubleShootingService struct {
	apiClient *APIClient
}

func NewSwitchMTroubleShootingService(c *APIClient) *SwitchMTroubleShootingService {
	s := new(SwitchMTroubleShootingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTroubleShootingService() *SwitchMTroubleShootingService {
	return NewSwitchMTroubleShootingService(ss.apiClient)
}

type SwitchMTroubleShootingRemoteClientConnectivityRequest struct {
	ClientMac *string `json:"clientMac,omitempty"`

	GroupId *string `json:"groupId,omitempty"`

	SwitchIds []string `json:"switchIds,omitempty"`
}

func NewSwitchMTroubleShootingRemoteClientConnectivityRequest() *SwitchMTroubleShootingRemoteClientConnectivityRequest {
	m := new(SwitchMTroubleShootingRemoteClientConnectivityRequest)
	return m
}

type SwitchMTroubleShootingRemoteClientConnectivityResponse struct {
	Data *SwitchMTroubleShootingRemoteClientConnectivityResponseDataType `json:"data,omitempty"`

	Error *SwitchMTroubleShootingRemoteClientConnectivityResponseErrorType `json:"error,omitempty"`

	Extra *SwitchMTroubleShootingRemoteClientConnectivityResponseExtraType `json:"extra,omitempty"`

	MetaData *SwitchMTroubleShootingRemoteClientConnectivityResponseMetaDataType `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func NewSwitchMTroubleShootingRemoteClientConnectivityResponse() *SwitchMTroubleShootingRemoteClientConnectivityResponse {
	m := new(SwitchMTroubleShootingRemoteClientConnectivityResponse)
	return m
}

type SwitchMTroubleShootingRemoteClientConnectivityResponseDataType struct {
	Ap *SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeApType `json:"ap,omitempty"`

	Client *SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeClientType `json:"client,omitempty"`

	NetworkSwitches []*SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType `json:"networkSwitches,omitempty"`
}

func NewSwitchMTroubleShootingRemoteClientConnectivityResponseDataType() *SwitchMTroubleShootingRemoteClientConnectivityResponseDataType {
	m := new(SwitchMTroubleShootingRemoteClientConnectivityResponseDataType)
	return m
}

type SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeApType struct {
	Connected *bool `json:"connected,omitempty"`

	Mac *string `json:"mac,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeApType() *SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeApType {
	m := new(SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeApType)
	return m
}

type SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeClientType struct {
	Connected *bool `json:"connected,omitempty"`

	Mac *string `json:"mac,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeClientType() *SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeClientType {
	m := new(SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeClientType)
	return m
}

type SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType struct {
	Connected *bool `json:"connected,omitempty"`

	Name *string `json:"name,omitempty"`

	PortName *string `json:"portName,omitempty"`
}

func NewSwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType() *SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType {
	m := new(SwitchMTroubleShootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType)
	return m
}

type SwitchMTroubleShootingRemoteClientConnectivityResponseErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleShootingRemoteClientConnectivityResponseErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleShootingRemoteClientConnectivityResponseErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleShootingRemoteClientConnectivityResponseErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleShootingRemoteClientConnectivityResponseErrorType() *SwitchMTroubleShootingRemoteClientConnectivityResponseErrorType {
	m := new(SwitchMTroubleShootingRemoteClientConnectivityResponseErrorType)
	return m
}

type SwitchMTroubleShootingRemoteClientConnectivityResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleShootingRemoteClientConnectivityResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleShootingRemoteClientConnectivityResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleShootingRemoteClientConnectivityResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleShootingRemoteClientConnectivityResponseExtraType() *SwitchMTroubleShootingRemoteClientConnectivityResponseExtraType {
	m := new(SwitchMTroubleShootingRemoteClientConnectivityResponseExtraType)
	return m
}

type SwitchMTroubleShootingRemoteClientConnectivityResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleShootingRemoteClientConnectivityResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleShootingRemoteClientConnectivityResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleShootingRemoteClientConnectivityResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleShootingRemoteClientConnectivityResponseMetaDataType() *SwitchMTroubleShootingRemoteClientConnectivityResponseMetaDataType {
	m := new(SwitchMTroubleShootingRemoteClientConnectivityResponseMetaDataType)
	return m
}

type SwitchMTroubleShootingRemoteCommandResponse struct {
	Data *string `json:"data,omitempty"`

	Error *SwitchMTroubleShootingRemoteCommandResponseErrorType `json:"error,omitempty"`

	Extra *SwitchMTroubleShootingRemoteCommandResponseExtraType `json:"extra,omitempty"`

	MetaData *SwitchMTroubleShootingRemoteCommandResponseMetaDataType `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func NewSwitchMTroubleShootingRemoteCommandResponse() *SwitchMTroubleShootingRemoteCommandResponse {
	m := new(SwitchMTroubleShootingRemoteCommandResponse)
	return m
}

type SwitchMTroubleShootingRemoteCommandResponseErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleShootingRemoteCommandResponseErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleShootingRemoteCommandResponseErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleShootingRemoteCommandResponseErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleShootingRemoteCommandResponseErrorType() *SwitchMTroubleShootingRemoteCommandResponseErrorType {
	m := new(SwitchMTroubleShootingRemoteCommandResponseErrorType)
	return m
}

type SwitchMTroubleShootingRemoteCommandResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleShootingRemoteCommandResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleShootingRemoteCommandResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleShootingRemoteCommandResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleShootingRemoteCommandResponseExtraType() *SwitchMTroubleShootingRemoteCommandResponseExtraType {
	m := new(SwitchMTroubleShootingRemoteCommandResponseExtraType)
	return m
}

type SwitchMTroubleShootingRemoteCommandResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleShootingRemoteCommandResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleShootingRemoteCommandResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleShootingRemoteCommandResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleShootingRemoteCommandResponseMetaDataType() *SwitchMTroubleShootingRemoteCommandResponseMetaDataType {
	m := new(SwitchMTroubleShootingRemoteCommandResponseMetaDataType)
	return m
}

type SwitchMTroubleShootingRemotePingRequest struct {
	PacketSize *int `json:"packetSize,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`

	SourceIp *string `json:"sourceIp,omitempty"`

	TargetIp *string `json:"targetIp,omitempty"`

	Ttl *int `json:"ttl,omitempty"`
}

func NewSwitchMTroubleShootingRemotePingRequest() *SwitchMTroubleShootingRemotePingRequest {
	m := new(SwitchMTroubleShootingRemotePingRequest)
	return m
}

type SwitchMTroubleShootingRemoteTracerouteRequest struct {
	MaxTtl *int `json:"maxTtl,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`

	SourceIp *string `json:"sourceIp,omitempty"`

	TargetIp *string `json:"targetIp,omitempty"`
}

func NewSwitchMTroubleShootingRemoteTracerouteRequest() *SwitchMTroubleShootingRemoteTracerouteRequest {
	m := new(SwitchMTroubleShootingRemoteTracerouteRequest)
	return m
}

type SwitchMTroubleShootingSupportLogStatus struct {
	// CreatedTime
	// Created Time of this SupportLog Request
	CreatedTime *string `json:"createdTime,omitempty"`

	// DownloadStatus
	// SupportLog Download Status (DOWNLOADING, DONE, FAILED)
	// Constraints:
	//    - oneof:[DOWNLOADING,DONE,TIMEOUT,FAILED]
	DownloadStatus *string `json:"downloadStatus,omitempty"`

	// SerialNumber
	// Switch Serial Number
	SerialNumber *string `json:"serialNumber,omitempty"`

	// SwitchId
	// Switch MAC Address
	SwitchId *string `json:"switchId,omitempty"`
}

func NewSwitchMTroubleShootingSupportLogStatus() *SwitchMTroubleShootingSupportLogStatus {
	m := new(SwitchMTroubleShootingSupportLogStatus)
	return m
}

// ExecuteSwitchRemoteClientConnectivity
//
// Switch trace client connectivity
//

// Trace a specific client's path of connectivity to a given switch
//
// Request Body:
//	 - body *SwitchMTroubleShootingRemoteClientConnectivityRequest
func (s *SwitchMTroubleShootingService) ExecuteSwitchRemoteClientConnectivity(ctx context.Context, body *SwitchMTroubleShootingRemoteClientConnectivityRequest) (*SwitchMTroubleShootingRemoteClientConnectivityResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTroubleShootingRemoteClientConnectivityResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMExecuteSwitchRemoteClientConnectivity, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTroubleShootingRemoteClientConnectivityResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ExecuteSwitchRemotePing
//
// Switch remote ping
//

// Attempt to ping an address or hostname from a specific switch. Note: This API is quite slow, and may take > 5 seconds to respond
//
// Request Body:
//	 - body *SwitchMTroubleShootingRemotePingRequest
func (s *SwitchMTroubleShootingService) ExecuteSwitchRemotePing(ctx context.Context, body *SwitchMTroubleShootingRemotePingRequest) (*SwitchMTroubleShootingRemoteCommandResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTroubleShootingRemoteCommandResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMExecuteSwitchRemotePing, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTroubleShootingRemoteCommandResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ExecuteSwitchRemoteTraceroute
//
// Switch remote traceroute
//

// Attempt to execute a traceroute from a specific switch to a destination.  Note: This is a very slow API, taking > 30 seconds to respond.
//
// Request Body:
//	 - body *SwitchMTroubleShootingRemoteTracerouteRequest
func (s *SwitchMTroubleShootingService) ExecuteSwitchRemoteTraceroute(ctx context.Context, body *SwitchMTroubleShootingRemoteTracerouteRequest) (*SwitchMTroubleShootingRemoteCommandResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTroubleShootingRemoteCommandResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMExecuteSwitchRemoteTraceroute, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTroubleShootingRemoteCommandResponse()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSupportLogBySwitchId
//
// Use this API to request ICX to prepare support log.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleShootingService) FindSupportLogBySwitchId(ctx context.Context, switchId string) (*SwitchMCommonCreateResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMCommonCreateResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSupportLogBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSupportLogDownloadBySwitchId
//
// Use this API to download support log.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleShootingService) FindSupportLogDownloadBySwitchId(ctx context.Context, switchId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSupportLogDownloadBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}

// FindSupportLogStatusBySwitchId
//
// Use this API to get the status of current support log request.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleShootingService) FindSupportLogStatusBySwitchId(ctx context.Context, switchId string) (*SwitchMTroubleShootingSupportLogStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTroubleShootingSupportLogStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSupportLogStatusBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMTroubleShootingSupportLogStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
