package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"errors"
	"io"
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

// SwitchMTroubleshootingRemoteClientConnectivityRequest
//
// Definition: troubleShooting_remoteClientConnectivityRequest
type SwitchMTroubleshootingRemoteClientConnectivityRequest struct {
	ClientMac *string `json:"clientMac,omitempty"`

	GroupId *string `json:"groupId,omitempty"`

	SwitchIds []string `json:"switchIds,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityRequest) UnmarshalJSON(b []byte) error {
	type _SwitchMTroubleshootingRemoteClientConnectivityRequest SwitchMTroubleshootingRemoteClientConnectivityRequest
	tmpType := new(_SwitchMTroubleshootingRemoteClientConnectivityRequest)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "clientMac")
	delete(tmpType.XAdditionalProperties, "groupId")
	delete(tmpType.XAdditionalProperties, "switchIds")
	*t = SwitchMTroubleshootingRemoteClientConnectivityRequest(*tmpType)
	return nil
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityRequest) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.ClientMac != nil {
		tmp["clientMac"] = t.ClientMac
	}
	if t.GroupId != nil {
		tmp["groupId"] = t.GroupId
	}
	if t.SwitchIds != nil {
		tmp["switchIds"] = t.SwitchIds
	}
	return json.Marshal(tmp)
}

func NewSwitchMTroubleshootingRemoteClientConnectivityRequest() *SwitchMTroubleshootingRemoteClientConnectivityRequest {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityRequest)
	return m
}

// SwitchMTroubleshootingRemoteClientConnectivityResponse
//
// Definition: troubleShooting_remoteClientConnectivityResponse
type SwitchMTroubleshootingRemoteClientConnectivityResponse struct {
	Data *SwitchMTroubleshootingRemoteClientConnectivityResponseDataType `json:"data,omitempty"`

	Error *SwitchMTroubleshootingRemoteClientConnectivityResponseErrorType `json:"error,omitempty"`

	Extra interface{} `json:"extra,omitempty"`

	MetaData interface{} `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityResponse) UnmarshalJSON(b []byte) error {
	type _SwitchMTroubleshootingRemoteClientConnectivityResponse SwitchMTroubleshootingRemoteClientConnectivityResponse
	tmpType := new(_SwitchMTroubleshootingRemoteClientConnectivityResponse)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "data")
	delete(tmpType.XAdditionalProperties, "error")
	delete(tmpType.XAdditionalProperties, "extra")
	delete(tmpType.XAdditionalProperties, "metaData")
	delete(tmpType.XAdditionalProperties, "success")
	*t = SwitchMTroubleshootingRemoteClientConnectivityResponse(*tmpType)
	return nil
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityResponse) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Data != nil {
		tmp["data"] = t.Data
	}
	if t.Error != nil {
		tmp["error"] = t.Error
	}
	if t.Extra != nil {
		tmp["extra"] = t.Extra
	}
	if t.MetaData != nil {
		tmp["metaData"] = t.MetaData
	}
	if t.Success != nil {
		tmp["success"] = t.Success
	}
	return json.Marshal(tmp)
}

type SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTroubleshootingRemoteClientConnectivityResponse
}

func newSwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTroubleshootingRemoteClientConnectivityResponse)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTroubleshootingRemoteClientConnectivityResponse() *SwitchMTroubleshootingRemoteClientConnectivityResponse {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponse)
	return m
}

// SwitchMTroubleshootingRemoteClientConnectivityResponseDataType
//
// Definition: troubleShooting_remoteClientConnectivityResponseDataType
type SwitchMTroubleshootingRemoteClientConnectivityResponseDataType struct {
	Ap *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType `json:"ap,omitempty"`

	Client *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType `json:"client,omitempty"`

	NetworkSwitches []*SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType `json:"networkSwitches,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseDataType() *SwitchMTroubleshootingRemoteClientConnectivityResponseDataType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseDataType)
	return m
}

// SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType
//
// Definition: troubleShooting_remoteClientConnectivityResponseDataTypeApType
type SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType struct {
	Connected *bool `json:"connected,omitempty"`

	Mac *string `json:"mac,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType() *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeApType)
	return m
}

// SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType
//
// Definition: troubleShooting_remoteClientConnectivityResponseDataTypeClientType
type SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType struct {
	Connected *bool `json:"connected,omitempty"`

	Mac *string `json:"mac,omitempty"`

	Name *string `json:"name,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType() *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeClientType)
	return m
}

// SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType
//
// Definition: troubleShooting_remoteClientConnectivityResponseDataTypeNetworkSwitchesType
type SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType struct {
	Connected *bool `json:"connected,omitempty"`

	Name *string `json:"name,omitempty"`

	PortName *string `json:"portName,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType) UnmarshalJSON(b []byte) error {
	type _SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType
	tmpType := new(_SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "connected")
	delete(tmpType.XAdditionalProperties, "name")
	delete(tmpType.XAdditionalProperties, "portName")
	*t = SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType(*tmpType)
	return nil
}

func (t *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Connected != nil {
		tmp["connected"] = t.Connected
	}
	if t.Name != nil {
		tmp["name"] = t.Name
	}
	if t.PortName != nil {
		tmp["portName"] = t.PortName
	}
	return json.Marshal(tmp)
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType() *SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseDataTypeNetworkSwitchesType)
	return m
}

// SwitchMTroubleshootingRemoteClientConnectivityResponseErrorType
//
// Definition: troubleShooting_remoteClientConnectivityResponseErrorType
type SwitchMTroubleshootingRemoteClientConnectivityResponseErrorType struct {
	Message *string `json:"message,omitempty"`

	MsgKey *string `json:"msgKey,omitempty"`

	MsgValues []string `json:"msgValues,omitempty"`
}

func NewSwitchMTroubleshootingRemoteClientConnectivityResponseErrorType() *SwitchMTroubleshootingRemoteClientConnectivityResponseErrorType {
	m := new(SwitchMTroubleshootingRemoteClientConnectivityResponseErrorType)
	return m
}

// SwitchMTroubleshootingRemoteCommandResponse
//
// Definition: troubleShooting_remoteCommandResponse
type SwitchMTroubleshootingRemoteCommandResponse struct {
	Data *string `json:"data,omitempty"`

	Error interface{} `json:"error,omitempty"`

	Extra interface{} `json:"extra,omitempty"`

	MetaData interface{} `json:"metaData,omitempty"`

	Success *bool `json:"success,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteCommandResponse) UnmarshalJSON(b []byte) error {
	type _SwitchMTroubleshootingRemoteCommandResponse SwitchMTroubleshootingRemoteCommandResponse
	tmpType := new(_SwitchMTroubleshootingRemoteCommandResponse)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "data")
	delete(tmpType.XAdditionalProperties, "error")
	delete(tmpType.XAdditionalProperties, "extra")
	delete(tmpType.XAdditionalProperties, "metaData")
	delete(tmpType.XAdditionalProperties, "success")
	*t = SwitchMTroubleshootingRemoteCommandResponse(*tmpType)
	return nil
}

func (t *SwitchMTroubleshootingRemoteCommandResponse) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.Data != nil {
		tmp["data"] = t.Data
	}
	if t.Error != nil {
		tmp["error"] = t.Error
	}
	if t.Extra != nil {
		tmp["extra"] = t.Extra
	}
	if t.MetaData != nil {
		tmp["metaData"] = t.MetaData
	}
	if t.Success != nil {
		tmp["success"] = t.Success
	}
	return json.Marshal(tmp)
}

type SwitchMTroubleshootingRemoteCommandResponseAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTroubleshootingRemoteCommandResponse
}

func newSwitchMTroubleshootingRemoteCommandResponseAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTroubleshootingRemoteCommandResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTroubleshootingRemoteCommandResponseAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTroubleshootingRemoteCommandResponse)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMTroubleshootingRemoteCommandResponse() *SwitchMTroubleshootingRemoteCommandResponse {
	m := new(SwitchMTroubleshootingRemoteCommandResponse)
	return m
}

// SwitchMTroubleshootingRemotePingRequest
//
// Definition: troubleShooting_remotePingRequest
type SwitchMTroubleshootingRemotePingRequest struct {
	PacketSize *float64 `json:"packetSize,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`

	SourceIp *string `json:"sourceIp,omitempty"`

	TargetIp *string `json:"targetIp,omitempty"`

	Ttl *float64 `json:"ttl,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemotePingRequest) UnmarshalJSON(b []byte) error {
	type _SwitchMTroubleshootingRemotePingRequest SwitchMTroubleshootingRemotePingRequest
	tmpType := new(_SwitchMTroubleshootingRemotePingRequest)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "packetSize")
	delete(tmpType.XAdditionalProperties, "sourceId")
	delete(tmpType.XAdditionalProperties, "sourceIp")
	delete(tmpType.XAdditionalProperties, "targetIp")
	delete(tmpType.XAdditionalProperties, "ttl")
	*t = SwitchMTroubleshootingRemotePingRequest(*tmpType)
	return nil
}

func (t *SwitchMTroubleshootingRemotePingRequest) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.PacketSize != nil {
		tmp["packetSize"] = t.PacketSize
	}
	if t.SourceId != nil {
		tmp["sourceId"] = t.SourceId
	}
	if t.SourceIp != nil {
		tmp["sourceIp"] = t.SourceIp
	}
	if t.TargetIp != nil {
		tmp["targetIp"] = t.TargetIp
	}
	if t.Ttl != nil {
		tmp["ttl"] = t.Ttl
	}
	return json.Marshal(tmp)
}

func NewSwitchMTroubleshootingRemotePingRequest() *SwitchMTroubleshootingRemotePingRequest {
	m := new(SwitchMTroubleshootingRemotePingRequest)
	return m
}

// SwitchMTroubleshootingRemoteTracerouteRequest
//
// Definition: troubleShooting_remoteTracerouteRequest
type SwitchMTroubleshootingRemoteTracerouteRequest struct {
	MaxTtl *float64 `json:"maxTtl,omitempty"`

	SourceId *string `json:"sourceId,omitempty"`

	SourceIp *string `json:"sourceIp,omitempty"`

	TargetIp *string `json:"targetIp,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMTroubleshootingRemoteTracerouteRequest) UnmarshalJSON(b []byte) error {
	type _SwitchMTroubleshootingRemoteTracerouteRequest SwitchMTroubleshootingRemoteTracerouteRequest
	tmpType := new(_SwitchMTroubleshootingRemoteTracerouteRequest)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "maxTtl")
	delete(tmpType.XAdditionalProperties, "sourceId")
	delete(tmpType.XAdditionalProperties, "sourceIp")
	delete(tmpType.XAdditionalProperties, "targetIp")
	*t = SwitchMTroubleshootingRemoteTracerouteRequest(*tmpType)
	return nil
}

func (t *SwitchMTroubleshootingRemoteTracerouteRequest) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.MaxTtl != nil {
		tmp["maxTtl"] = t.MaxTtl
	}
	if t.SourceId != nil {
		tmp["sourceId"] = t.SourceId
	}
	if t.SourceIp != nil {
		tmp["sourceIp"] = t.SourceIp
	}
	if t.TargetIp != nil {
		tmp["targetIp"] = t.TargetIp
	}
	return json.Marshal(tmp)
}

func NewSwitchMTroubleshootingRemoteTracerouteRequest() *SwitchMTroubleshootingRemoteTracerouteRequest {
	m := new(SwitchMTroubleshootingRemoteTracerouteRequest)
	return m
}

// SwitchMTroubleshootingSupportLogStatus
//
// Definition: troubleShooting_supportLogStatus
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

type SwitchMTroubleshootingSupportLogStatusAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMTroubleshootingSupportLogStatus
}

func newSwitchMTroubleshootingSupportLogStatusAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMTroubleshootingSupportLogStatusAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMTroubleshootingSupportLogStatusAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMTroubleshootingSupportLogStatus)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
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
// Operation ID: executeSwitchRemoteClientConnectivity
// Operation path: /switch/troubleshooting/connectiontracking
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMTroubleshootingRemoteClientConnectivityRequest
func (s *SwitchMTroubleshootingService) ExecuteSwitchRemoteClientConnectivity(ctx context.Context, body *SwitchMTroubleshootingRemoteClientConnectivityRequest, mutators ...RequestMutator) (*SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMExecuteSwitchRemoteClientConnectivity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse), err
}

// ExecuteSwitchRemotePing
//
// Switch remote ping
//
// Attempt to ping an address or hostname from a specific switch. Note: This API is quite slow, and may take > 5 seconds to respond
//
// Operation ID: executeSwitchRemotePing
// Operation path: /switch/troubleshooting/ping
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMTroubleshootingRemotePingRequest
func (s *SwitchMTroubleshootingService) ExecuteSwitchRemotePing(ctx context.Context, body *SwitchMTroubleshootingRemotePingRequest, mutators ...RequestMutator) (*SwitchMTroubleshootingRemoteCommandResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTroubleshootingRemoteCommandResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMTroubleshootingRemoteCommandResponseAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMExecuteSwitchRemotePing, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMTroubleshootingRemoteCommandResponseAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTroubleshootingRemoteCommandResponseAPIResponse), err
}

// ExecuteSwitchRemoteTraceroute
//
// Switch remote traceroute
//
// Attempt to execute a traceroute from a specific switch to a destination.  Note: This is a very slow API, taking > 30 seconds to respond.
//
// Operation ID: executeSwitchRemoteTraceroute
// Operation path: /switch/troubleshooting/traceroute
// Success code: 200 (OK)
//
// Request body:
//	 - body *SwitchMTroubleshootingRemoteTracerouteRequest
func (s *SwitchMTroubleshootingService) ExecuteSwitchRemoteTraceroute(ctx context.Context, body *SwitchMTroubleshootingRemoteTracerouteRequest, mutators ...RequestMutator) (*SwitchMTroubleshootingRemoteCommandResponseAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTroubleshootingRemoteCommandResponseAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMTroubleshootingRemoteCommandResponseAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodPost, RouteSwitchMExecuteSwitchRemoteTraceroute, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMTroubleshootingRemoteCommandResponseAPIResponse), err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTroubleshootingRemoteCommandResponseAPIResponse), err
}

// FindSupportLogBySwitchId
//
// Use this API to request ICX to prepare support log.
//
// Operation ID: findSupportLogBySwitchId
// Operation path: /supportLog/{switchId}
// Success code: 200 (OK)
//
// Required parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleshootingService) FindSupportLogBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMCommonCreateResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMCommonCreateResultAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindSupportLogBySwitchId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMCommonCreateResultAPIResponse), err
}

// FindSupportLogDownloadBySwitchId
//
// Use this API to download support log.
//
// Operation ID: findSupportLogDownloadBySwitchId
// Operation path: /supportLog/download/{switchId}
// Success code: 200 (OK)
//
// Required parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleshootingService) FindSupportLogDownloadBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newEmptyAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindSupportLogDownloadBySwitchId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindSupportLogStatusBySwitchId
//
// Use this API to get the status of current support log request.
//
// Operation ID: findSupportLogStatusBySwitchId
// Operation path: /supportLog/status/{switchId}
// Success code: 200 (OK)
//
// Required parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleshootingService) FindSupportLogStatusBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMTroubleshootingSupportLogStatusAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMTroubleshootingSupportLogStatusAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMTroubleshootingSupportLogStatusAPIResponse), err
	}
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteSwitchMFindSupportLogStatusBySwitchId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMTroubleshootingSupportLogStatusAPIResponse), err
}
