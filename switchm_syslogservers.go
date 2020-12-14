package bigdog

// API Version: v9_1

import (
	"context"
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
func (s *SwitchMSyslogServersService) AddGroupSyslogServersByGroupId(ctx context.Context, body *SwitchMSyslogServersCreateSyslogServerConfig, groupId string, mutators ...RequestMutator) (*SwitchMCommonCreateResult, *APIResponseMeta, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddGroupSyslogServersByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.PathParams.Set("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMCommonCreateResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SwitchMSyslogServersService) DeleteGroupSyslogServersById(ctx context.Context, groupId string, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteGroupSyslogServersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
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
func (s *SwitchMSyslogServersService) FindGroupSyslogServersByGroupId(ctx context.Context, groupId string, mutators ...RequestMutator) (*SwitchMSyslogServersQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMSyslogServersQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindGroupSyslogServersByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	req.PathParams.Set("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewSwitchMSyslogServersQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
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
func (s *SwitchMSyslogServersService) UpdateGroupSyslogServersById(ctx context.Context, body *SwitchMSyslogServersSyslogServer, groupId string, id string, mutators ...RequestMutator) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateGroupSyslogServersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}