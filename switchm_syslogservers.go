package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMSyslogServersService struct {
	apiClient *VSZClient
}

func NewSwitchMSyslogServersService(c *VSZClient) *SwitchMSyslogServersService {
	s := new(SwitchMSyslogServersService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMSyslogServersService() *SwitchMSyslogServersService {
	return NewSwitchMSyslogServersService(ss.apiClient)
}

// SwitchMSyslogServersCreateSyslogServerConfig
//
// Definition: syslogServers_createSyslogServerConfig
type SwitchMSyslogServersCreateSyslogServerConfig struct {
	Ip *string `json:"ip,omitempty"`

	UdpPort *int `json:"udpPort,omitempty"`
}

func NewSwitchMSyslogServersCreateSyslogServerConfig() *SwitchMSyslogServersCreateSyslogServerConfig {
	m := new(SwitchMSyslogServersCreateSyslogServerConfig)
	return m
}

// SwitchMSyslogServersSyslogServer
//
// Definition: syslogServers_syslogServer
type SwitchMSyslogServersSyslogServer struct {
	CreatedTime *int `json:"createdTime,omitempty"`

	CreatorId *string `json:"creatorId,omitempty"`

	CreatorUsername *string `json:"creatorUsername,omitempty"`

	GroupId *string `json:"groupId,omitempty"`

	Id *string `json:"id,omitempty"`

	Ip *string `json:"ip,omitempty"`

	UdpPort *int `json:"udpPort,omitempty"`

	UpdatedTime *int `json:"updatedTime,omitempty"`

	UpdaterId *string `json:"updaterId,omitempty"`

	UpdaterUsername *string `json:"updaterUsername,omitempty"`
}

func NewSwitchMSyslogServersSyslogServer() *SwitchMSyslogServersSyslogServer {
	m := new(SwitchMSyslogServersSyslogServer)
	return m
}

// SwitchMSyslogServersQueryResult
//
// Definition: syslogServers_syslogServersQueryResult
type SwitchMSyslogServersQueryResult struct {
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Indicator of whether there are more configs after the current displayed list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more configs after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMSyslogServersSyslogServer `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total records count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMSyslogServersQueryResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMSyslogServersQueryResult
}

func newSwitchMSyslogServersQueryResultAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(SwitchMSyslogServersQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *SwitchMSyslogServersQueryResultAPIResponse) Hydrate() error {
	r.Data = new(SwitchMSyslogServersQueryResult)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewSwitchMSyslogServersQueryResult() *SwitchMSyslogServersQueryResult {
	m := new(SwitchMSyslogServersQueryResult)
	return m
}

// AddGroupSyslogServersByGroupId
//
// Operation ID: addGroupSyslogServersByGroupId
//
// Use this API command to create a new Syslog server.
//
// Request Body:
//	 - body *SwitchMSyslogServersCreateSyslogServerConfig
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMSyslogServersService) AddGroupSyslogServersByGroupId(ctx context.Context, body *SwitchMSyslogServersCreateSyslogServerConfig, groupId string, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddGroupSyslogServersByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMCommonCreateResultAPIResponse), err
}

// DeleteGroupSyslogServersById
//
// Operation ID: deleteGroupSyslogServersById
//
// Use this API command to delete a Syslog server.
//
// Required Parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMSyslogServersService) DeleteGroupSyslogServersById(ctx context.Context, groupId string, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteGroupSyslogServersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}

// FindGroupSyslogServersByGroupId
//
// Operation ID: findGroupSyslogServersByGroupId
//
// Use this API command to get Syslog servers.
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMSyslogServersService) FindGroupSyslogServersByGroupId(ctx context.Context, groupId string, mutators ...RequestMutator) (*SwitchMSyslogServersQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMSyslogServersQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMSyslogServersQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindGroupSyslogServersByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMSyslogServersQueryResultAPIResponse), err
}

// UpdateGroupSyslogServersById
//
// Operation ID: updateGroupSyslogServersById
//
// Use this API command to update a Syslog server.
//
// Request Body:
//	 - body *SwitchMSyslogServersSyslogServer
//
// Required Parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMSyslogServersService) UpdateGroupSyslogServersById(ctx context.Context, body *SwitchMSyslogServersSyslogServer, groupId string, id string, mutators ...RequestMutator) (*RawAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateGroupSyslogServersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*RawAPIResponse), err
	}
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*RawAPIResponse), err
}
