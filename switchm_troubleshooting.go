package bigdog

// API Version: v9_1

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

func newSwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse) Hydrate() error {
	r.Data = new(SwitchMTroubleshootingRemoteClientConnectivityResponse)
	return json.NewDecoder(r).Decode(r.Data)
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

func newSwitchMTroubleshootingRemoteCommandResponseAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMTroubleshootingRemoteCommandResponseAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMTroubleshootingRemoteCommandResponseAPIResponse) Hydrate() error {
	r.Data = new(SwitchMTroubleshootingRemoteCommandResponse)
	return json.NewDecoder(r).Decode(r.Data)
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

func newSwitchMTroubleshootingSupportLogStatusAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMTroubleshootingSupportLogStatusAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMTroubleshootingSupportLogStatusAPIResponse) Hydrate() error {
	r.Data = new(SwitchMTroubleshootingSupportLogStatus)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMTroubleshootingSupportLogStatus() *SwitchMTroubleshootingSupportLogStatus {
	m := new(SwitchMTroubleshootingSupportLogStatus)
	return m
}

// ExecuteSwitchRemoteClientConnectivity
//
// Operation ID: executeSwitchRemoteClientConnectivity
//
// Switch trace client connectivity
//
// Trace a specific client's path of connectivity to a given switch
//
// Request Body:
//	 - body *SwitchMTroubleshootingRemoteClientConnectivityRequest
func (s *SwitchMTroubleshootingService) ExecuteSwitchRemoteClientConnectivity(ctx context.Context, body *SwitchMTroubleshootingRemoteClientConnectivityRequest, mutators ...RequestMutator) (*SwitchMTroubleshootingRemoteClientConnectivityResponseAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMExecuteSwitchRemoteClientConnectivity, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTroubleshootingRemoteClientConnectivityResponse()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ExecuteSwitchRemotePing
//
// Operation ID: executeSwitchRemotePing
//
// Switch remote ping
//
// Attempt to ping an address or hostname from a specific switch. Note: This API is quite slow, and may take > 5 seconds to respond
//
// Request Body:
//	 - body *SwitchMTroubleshootingRemotePingRequest
func (s *SwitchMTroubleshootingService) ExecuteSwitchRemotePing(ctx context.Context, body *SwitchMTroubleshootingRemotePingRequest, mutators ...RequestMutator) (*SwitchMTroubleshootingRemoteCommandResponseAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMExecuteSwitchRemotePing, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTroubleshootingRemoteCommandResponse()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// ExecuteSwitchRemoteTraceroute
//
// Operation ID: executeSwitchRemoteTraceroute
//
// Switch remote traceroute
//
// Attempt to execute a traceroute from a specific switch to a destination.  Note: This is a very slow API, taking > 30 seconds to respond.
//
// Request Body:
//	 - body *SwitchMTroubleshootingRemoteTracerouteRequest
func (s *SwitchMTroubleshootingService) ExecuteSwitchRemoteTraceroute(ctx context.Context, body *SwitchMTroubleshootingRemoteTracerouteRequest, mutators ...RequestMutator) (*SwitchMTroubleshootingRemoteCommandResponseAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMExecuteSwitchRemoteTraceroute, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "text/plain;charset=UTF-8")
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTroubleshootingRemoteCommandResponse()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSupportLogBySwitchId
//
// Operation ID: findSupportLogBySwitchId
//
// Use this API to request ICX to prepare support log.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleshootingService) FindSupportLogBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindSupportLogBySwitchId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMCommonCreateResult()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindSupportLogDownloadBySwitchId
//
// Operation ID: findSupportLogDownloadBySwitchId
//
// Use this API to download support log.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleshootingService) FindSupportLogDownloadBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*RawAPIResponse, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindSupportLogDownloadBySwitchId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, nil, err)
	return rm, err
}

// FindSupportLogStatusBySwitchId
//
// Operation ID: findSupportLogStatusBySwitchId
//
// Use this API to get the status of current support log request.
//
// Required Parameters:
// - switchId string
//		- required
func (s *SwitchMTroubleshootingService) FindSupportLogStatusBySwitchId(ctx context.Context, switchId string, mutators ...RequestMutator) (*SwitchMTroubleshootingSupportLogStatusAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindSupportLogStatusBySwitchId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("switchId", switchId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMTroubleshootingSupportLogStatus()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
