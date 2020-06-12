package bigdog

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMTroubleshootingService struct {
	apiClient *VSZClient
}

func NewSwitchMTroubleshootingService(c *VSZClient) *SwitchMTroubleshootingService {
	s := new(SwitchMTroubleshootingService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTroubleshootingService() *SwitchMTroubleshootingService {
	return NewSwitchMTroubleshootingService(ss.apiClient)
}

type SwitchMTroubleshootingRemoteClientConnectivityRequest struct {
	ClientMac *string `json:"clientMac,omitempty"`

	GroupId *string `json:"groupId,omitempty"`

	SwitchIds []string `json:"switchIds,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityRequest() *SwitchMTroubleshootingRemoteClientConnectivityRequest {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityRequest)
	return m
}

type SwitchMTroubleshootingRemoteClientConnectivityResponse struct {
	Data *SwitchMTroubleshootingRemoteClientConnectivityResponseDataType `json:"data,omitempty"`

	Error *SwitchMTroubleshootingRemoteClientConnectivityResponseErrorType `json:"error,omitempty"`

	Extra *SwitchMTroubleshootingRemoteClientConnectivityResponseExtraType `json:"extra,omitempty"`

	MetaData *SwitchMTroubleshootingRemoteClientConnectivityResponseMetaDataType `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponse() *SwitchMTroubleshootingRemoteClientConnectivityResponse {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponse)
	return m
}

type SwitchMTroubleshootingRemoteClientConnectivityResponseDataType struct {
	Ap *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType `json:"ap,omitempty"`

	Client *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType `json:"client,omitempty"`

	NetworkSwitches []*SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType `json:"networkSwitches,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseDataType() *SwitchMTroubleshootingRemoteClientConnectivityResponseDataType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseDataType)
	return m
}

type SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType struct {
	Connected *bool `json:"connected,omitempty"`

	Mac *string `json:"mac,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType() *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType)
	return m
}

type SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType struct {
	Connected *bool `json:"connected,omitempty"`

	Mac *string `json:"mac,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType() *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType)
	return m
}

type SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType struct {
	Connected *bool `json:"connected,omitempty"`

	Name *string `json:"name,omitempty"`

	PortName *string `json:"portName,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType() *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType)
	return m
}

type SwitchMTroubleshootingRemoteClientConnectivityResponseErrorType struct {
	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`

	MsgValues []string `json:"msgValues,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseErrorType() *SwitchMTroubleshootingRemoteClientConnectivityResponseErrorType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseErrorType)
	return m
}

type SwitchMTroubleshootingRemoteClientConnectivityResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleshootingRemoteClientConnectivityResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseExtraType() *SwitchMTroubleshootingRemoteClientConnectivityResponseExtraType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseExtraType)
	return m
}

type SwitchMTroubleshootingRemoteClientConnectivityResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleshootingRemoteClientConnectivityResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseMetaDataType() *SwitchMTroubleshootingRemoteClientConnectivityResponseMetaDataType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseMetaDataType)
	return m
}

type SwitchMTroubleshootingRemoteCommandResponse struct {
	Data *string `json:"data,omitempty"`

	Error *SwitchMTroubleshootingRemoteCommandResponseErrorType `json:"error,omitempty"`

	Extra *SwitchMTroubleshootingRemoteCommandResponseExtraType `json:"extra,omitempty"`

	MetaData *SwitchMTroubleshootingRemoteCommandResponseMetaDataType `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`
}

func NewSwitchMTroubleshootingRemoteCommandResponse() *SwitchMTroubleshootingRemoteCommandResponse {
	m := new(SwitchMTroubleshootingRemoteCommandResponse)
	return m
}

type SwitchMTroubleshootingRemoteCommandResponseErrorType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteCommandResponseErrorType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleshootingRemoteCommandResponseErrorType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleshootingRemoteCommandResponseErrorType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleshootingRemoteCommandResponseErrorType() *SwitchMTroubleshootingRemoteCommandResponseErrorType {
	m := new(SwitchMTroubleshootingRemoteCommandResponseErrorType)
	return m
}

type SwitchMTroubleshootingRemoteCommandResponseExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteCommandResponseExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleshootingRemoteCommandResponseExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleshootingRemoteCommandResponseExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleshootingRemoteCommandResponseExtraType() *SwitchMTroubleshootingRemoteCommandResponseExtraType {
	m := new(SwitchMTroubleshootingRemoteCommandResponseExtraType)
	return m
}

type SwitchMTroubleshootingRemoteCommandResponseMetaDataType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteCommandResponseMetaDataType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMTroubleshootingRemoteCommandResponseMetaDataType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMTroubleshootingRemoteCommandResponseMetaDataType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMTroubleshootingRemoteCommandResponseMetaDataType() *SwitchMTroubleshootingRemoteCommandResponseMetaDataType {
	m := new(SwitchMTroubleshootingRemoteCommandResponseMetaDataType)
	return m
}

type SwitchMTroubleshootingRemotePingRequest struct {
	PacketSize *int `json:"packetSize,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`

	SourceIp *string `json:"sourceIp,omitempty"`

	TargetIp *string `json:"targetIp,omitempty"`

	Ttl *int `json:"ttl,omitempty"`
}

func NewSwitchMTroubleshootingRemotePingRequest() *SwitchMTroubleshootingRemotePingRequest {
	m := new(SwitchMTroubleshootingRemotePingRequest)
	return m
}

type SwitchMTroubleshootingRemoteTracerouteRequest struct {
	MaxTtl *int `json:"maxTtl,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`

	SourceIp *string `json:"sourceIp,omitempty"`

	TargetIp *string `json:"targetIp,omitempty"`
}

func NewSwitchMTroubleshootingRemoteTracerouteRequest() *SwitchMTroubleshootingRemoteTracerouteRequest {
	m := new(SwitchMTroubleshootingRemoteTracerouteRequest)
	return m
}

type SwitchMTroubleshootingSupportLogStatus struct {
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

func NewSwitchMTroubleshootingSupportLogStatus() *SwitchMTroubleshootingSupportLogStatus {
	m := new(SwitchMTroubleshootingSupportLogStatus)
	return m
}

// ExecuteSwitchRemoteClientConnectivity
//
// Switch trace client connectivity
//
// Trace a specific client's path of connectivity to a given switch
//
// Request Body:
//	 - body *SwitchMTroubleshootingRemoteClientConnectivityRequest
func (s *SwitchMTroubleshootingService) ExecuteSwitchRemoteClientConnectivity(ctx context.Context, body *SwitchMTroubleshootingRemoteClientConnectivityRequest, mutators ...RequestMutator) (*SwitchMTroubleshootingRemoteClientConnectivityResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTroubleshootingRemoteClientConnectivityResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMExecuteSwitchRemoteClientConnectivity, true)
	req.SetHeader(headerKeyContentType, "text/plain;charset=UTF-8")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTroubleshootingRemoteClientConnectivityResponse()
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
//	 - body *SwitchMTroubleshootingRemotePingRequest
func (s *SwitchMTroubleshootingService) ExecuteSwitchRemotePing(ctx context.Context, body *SwitchMTroubleshootingRemotePingRequest, mutators ...RequestMutator) (*SwitchMTroubleshootingRemoteCommandResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTroubleshootingRemoteCommandResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMExecuteSwitchRemotePing, true)
	req.SetHeader(headerKeyContentType, "text/plain;charset=UTF-8")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTroubleshootingRemoteCommandResponse()
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
//	 - body *SwitchMTroubleshootingRemoteTracerouteRequest
func (s *SwitchMTroubleshootingService) ExecuteSwitchRemoteTraceroute(ctx context.Context, body *SwitchMTroubleshootingRemoteTracerouteRequest, mutators ...RequestMutator) (*SwitchMTroubleshootingRemoteCommandResponse, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTroubleshootingRemoteCommandResponse
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMExecuteSwitchRemoteTraceroute, true)
	req.SetHeader(headerKeyContentType, "text/plain;charset=UTF-8")
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTroubleshootingRemoteCommandResponse()
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
func (s *SwitchMTroubleshootingService) FindSupportLogBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMCommonCreateResult, *APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *SwitchMTroubleshootingService) FindSupportLogDownloadBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*APIResponseMeta, error) {
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
	req.SetHeader(headerKeyAccept, "*/*")
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
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
func (s *SwitchMTroubleshootingService) FindSupportLogStatusBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMTroubleshootingSupportLogStatus, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMTroubleshootingSupportLogStatus
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindSupportLogStatusBySwitchId, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	req.SetPathParameter("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTroubleshootingSupportLogStatus()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
