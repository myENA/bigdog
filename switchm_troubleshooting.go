package ruckus

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSwitchTroubleshootingService struct {
	apiClient *VSZClient
}

func NewSwitchMSwitchTroubleshootingService(c *VSZClient) *SwitchMSwitchTroubleshootingService {
	s := new(SwitchMSwitchTroubleshootingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSwitchTroubleshootingService() *SwitchMSwitchTroubleshootingService {
	return NewSwitchMSwitchTroubleshootingService(ss.apiClient)
}

type SwitchMSwitchTroubleshootingRemoteClientConnectivityRequest struct {
	ClientMac *string `json:"clientMac,omitempty"`

	GroupId *string `json:"groupId,omitempty"`

	SwitchIds []string `json:"switchIds,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemoteClientConnectivityRequest() *SwitchMSwitchTroubleshootingRemoteClientConnectivityRequest {
	m := new(SwitchMSwitchTroubleshootingRemoteClientConnectivityRequest)
	return m
}

type SwitchMSwitchTroubleshootingRemoteClientConnectivityResponse struct {
	Data **SwitchMSwitchTroubleshootingRemoteClientConnectivityResponse `json:"data,omitempty"`

	Error **SwitchMSwitchTroubleshootingRemoteClientConnectivityResponse `json:"error,omitempty"`

	Extra **SwitchMSwitchTroubleshootingRemoteClientConnectivityResponse `json:"extra,omitempty"`

	MetaData **SwitchMSwitchTroubleshootingRemoteClientConnectivityResponse `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemoteClientConnectivityResponse() *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponse {
	m := new(SwitchMSwitchTroubleshootingRemoteClientConnectivityResponse)
	return m
}

type SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataType struct {
	Ap **SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataType `json:"ap,omitempty"`

	Client **SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataType `json:"client,omitempty"`

	NetworkSwitches []**SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataType `json:"networkSwitches,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataType() *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataType {
	m := new(SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataType)
	return m
}

type SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeApType struct {
	Connected *bool `json:"connected,omitempty"`

	Mac *string `json:"mac,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeApType() *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeApType {
	m := new(SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeApType)
	return m
}

type SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeClientType struct {
	Connected *bool `json:"connected,omitempty"`

	Mac *string `json:"mac,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeClientType() *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeClientType {
	m := new(SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeClientType)
	return m
}

type SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType struct {
	Connected *bool `json:"connected,omitempty"`

	Name *string `json:"name,omitempty"`

	PortName *string `json:"portName,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType() *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType {
	m := new(SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType)
	return m
}

type SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseErrorType struct {
	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`

	MsgValues []string `json:"msgValues,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemoteClientConnectivityResponseErrorType() *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseErrorType {
	m := new(SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseErrorType)
	return m
}

type SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTroubleshootingRemoteClientConnectivityResponseExtraType() *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseExtraType {
	m := new(SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseExtraType)
	return m
}

type SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTroubleshootingRemoteClientConnectivityResponseMetaDataType() *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseMetaDataType {
	m := new(SwitchMSwitchTroubleshootingRemoteClientConnectivityResponseMetaDataType)
	return m
}

type SwitchMSwitchTroubleshootingRemoteCommandResponse struct {
	Data *string `json:"data,omitempty"`

	Error **SwitchMSwitchTroubleshootingRemoteCommandResponse `json:"error,omitempty"`

	Extra **SwitchMSwitchTroubleshootingRemoteCommandResponse `json:"extra,omitempty"`

	MetaData **SwitchMSwitchTroubleshootingRemoteCommandResponse `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemoteCommandResponse() *SwitchMSwitchTroubleshootingRemoteCommandResponse {
	m := new(SwitchMSwitchTroubleshootingRemoteCommandResponse)
	return m
}

type SwitchMSwitchTroubleshootingRemoteCommandResponseErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTroubleshootingRemoteCommandResponseErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTroubleshootingRemoteCommandResponseErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTroubleshootingRemoteCommandResponseErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTroubleshootingRemoteCommandResponseErrorType() *SwitchMSwitchTroubleshootingRemoteCommandResponseErrorType {
	m := new(SwitchMSwitchTroubleshootingRemoteCommandResponseErrorType)
	return m
}

type SwitchMSwitchTroubleshootingRemoteCommandResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTroubleshootingRemoteCommandResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTroubleshootingRemoteCommandResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTroubleshootingRemoteCommandResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTroubleshootingRemoteCommandResponseExtraType() *SwitchMSwitchTroubleshootingRemoteCommandResponseExtraType {
	m := new(SwitchMSwitchTroubleshootingRemoteCommandResponseExtraType)
	return m
}

type SwitchMSwitchTroubleshootingRemoteCommandResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMSwitchTroubleshootingRemoteCommandResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMSwitchTroubleshootingRemoteCommandResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMSwitchTroubleshootingRemoteCommandResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMSwitchTroubleshootingRemoteCommandResponseMetaDataType() *SwitchMSwitchTroubleshootingRemoteCommandResponseMetaDataType {
	m := new(SwitchMSwitchTroubleshootingRemoteCommandResponseMetaDataType)
	return m
}

type SwitchMSwitchTroubleshootingRemotePingRequest struct {
	PacketSize *int `json:"packetSize,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`

	SourceIp *string `json:"sourceIp,omitempty"`

	TargetIp *string `json:"targetIp,omitempty"`

	Ttl *int `json:"ttl,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemotePingRequest() *SwitchMSwitchTroubleshootingRemotePingRequest {
	m := new(SwitchMSwitchTroubleshootingRemotePingRequest)
	return m
}

type SwitchMSwitchTroubleshootingRemoteTracerouteRequest struct {
	MaxTtl *int `json:"maxTtl,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`

	SourceIp *string `json:"sourceIp,omitempty"`

	TargetIp *string `json:"targetIp,omitempty"`
}

func NewSwitchMSwitchTroubleshootingRemoteTracerouteRequest() *SwitchMSwitchTroubleshootingRemoteTracerouteRequest {
	m := new(SwitchMSwitchTroubleshootingRemoteTracerouteRequest)
	return m
}

type SwitchMSwitchTroubleshootingSupportLogStatus struct {
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

func NewSwitchMSwitchTroubleshootingSupportLogStatus() *SwitchMSwitchTroubleshootingSupportLogStatus {
	m := new(SwitchMSwitchTroubleshootingSupportLogStatus)
	return m
}

// ExecuteSwitchRemoteClientConnectivity
//
// Switch trace client connectivity
//
// Trace a specific client's path of connectivity to a given switch
//
// Request Body:
//	 - body *SwitchMSwitchTroubleshootingRemoteClientConnectivityRequest
func (s *SwitchMSwitchTroubleshootingService) ExecuteSwitchRemoteClientConnectivity(ctx context.Context, body *SwitchMSwitchTroubleshootingRemoteClientConnectivityRequest) (*SwitchMSwitchTroubleshootingRemoteClientConnectivityResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTroubleshootingRemoteClientConnectivityResponse
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
	req.SetHeader(headerKeyContentType, "text/plain;charset=UTF-8")
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTroubleshootingRemoteClientConnectivityResponse()
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
//	 - body *SwitchMSwitchTroubleshootingRemotePingRequest
func (s *SwitchMSwitchTroubleshootingService) ExecuteSwitchRemotePing(ctx context.Context, body *SwitchMSwitchTroubleshootingRemotePingRequest) (*SwitchMSwitchTroubleshootingRemoteCommandResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTroubleshootingRemoteCommandResponse
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
	req.SetHeader(headerKeyContentType, "text/plain;charset=UTF-8")
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTroubleshootingRemoteCommandResponse()
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
//	 - body *SwitchMSwitchTroubleshootingRemoteTracerouteRequest
func (s *SwitchMSwitchTroubleshootingService) ExecuteSwitchRemoteTraceroute(ctx context.Context, body *SwitchMSwitchTroubleshootingRemoteTracerouteRequest) (*SwitchMSwitchTroubleshootingRemoteCommandResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTroubleshootingRemoteCommandResponse
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
	req.SetHeader(headerKeyContentType, "text/plain;charset=UTF-8")
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTroubleshootingRemoteCommandResponse()
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
func (s *SwitchMSwitchTroubleshootingService) FindSupportLogBySwitchId(ctx context.Context, switchId string) (*SwitchMCommonCreateResult, *APIResponseMeta, error) {
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
func (s *SwitchMSwitchTroubleshootingService) FindSupportLogDownloadBySwitchId(ctx context.Context, switchId string) (*APIResponseMeta, error) {
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
func (s *SwitchMSwitchTroubleshootingService) FindSupportLogStatusBySwitchId(ctx context.Context, switchId string) (*SwitchMSwitchTroubleshootingSupportLogStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSwitchTroubleshootingSupportLogStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSupportLogStatusBySwitchId, true)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMSwitchTroubleshootingSupportLogStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
