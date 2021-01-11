package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
	"net/http"
)

type SwitchMAAAServersService struct {
	apiClient *VSZClient
}

func NewSwitchMAAAServersService(c *VSZClient) *SwitchMAAAServersService {
	s := new(SwitchMAAAServersService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMAAAServersService() *SwitchMAAAServersService {
	return NewSwitchMAAAServersService(ss.apiClient)
}

// SwitchMAAAServersAAAServer
//
// Definition: aaaServers_AAAServer
type SwitchMAAAServersAAAServer struct {
	// AcctPort
	// AAA server accounting port
	AcctPort *int `json:"acctPort,omitempty"`

	AltoId *string `json:"altoId,omitempty"`

	// AuthPort
	// AAA server authentication port
	AuthPort *int `json:"authPort,omitempty"`

	// CreatedTime
	// The create time of the AAA server
	CreatedTime *int `json:"createdTime,omitempty"`

	// CreatorId
	// AAA server creator Id
	CreatorId *string `json:"creatorId,omitempty"`

	// CreatorUsername
	// AAA server creator name
	CreatorUsername *string `json:"creatorUsername,omitempty"`

	// GroupId
	// Group ID of the AAA server
	GroupId *string `json:"groupId,omitempty"`

	// Id
	// AAA server ID
	Id *string `json:"id,omitempty"`

	// Ip
	// AAA server IP address
	Ip *string `json:"ip,omitempty"`

	// Level
	// Access level of AAA server
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty"`

	// Name
	// Name of the AAA server
	Name *string `json:"name,omitempty"`

	// Password
	// Password for local user
	Password *string `json:"password,omitempty"`

	// Purpose
	// AAA server purpose
	// Constraints:
	//    - oneof:[DEFAULT,AUTHENTICATION_ONLY,AUTHORIZATION_ONLY,ACCOUNTING_ONLY]
	Purpose *string `json:"purpose,omitempty"`

	// Secret
	// AAA server secret
	Secret *string `json:"secret,omitempty"`

	// ServerType
	// The server tpye of the AAA server
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ServerType *string `json:"serverType,omitempty"`

	// UpdatedTime
	// The modify time of the AAA server
	UpdatedTime *int `json:"updatedTime,omitempty"`

	// UpdaterId
	// AAA server updater Id
	UpdaterId *string `json:"updaterId,omitempty"`

	// UpdaterUsername
	// AAA server updater name
	UpdaterUsername *string `json:"updaterUsername,omitempty"`

	// Username
	// Username for local user
	Username *string `json:"username,omitempty"`
}

type SwitchMAAAServersAAAServerAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMAAAServersAAAServer
}

func newSwitchMAAAServersAAAServerAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMAAAServersAAAServerAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMAAAServersAAAServerAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMAAAServersAAAServer)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMAAAServersAAAServer() *SwitchMAAAServersAAAServer {
	m := new(SwitchMAAAServersAAAServer)
	return m
}

// SwitchMAAAServersQueryResult
//
// Definition: aaaServers_aaaServersQueryResult
type SwitchMAAAServersQueryResult struct {
	// Extra
	// Any additional response data
	Extra interface{} `json:"extra,omitempty"`

	// FirstIndex
	// Index of the first AAA Server returned out of the complete AAA Server list
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Indicator of whether there are more AAA Servers after the current displayed list
	HasMore *bool `json:"hasMore,omitempty"`

	List []*SwitchMAAAServersAAAServer `json:"list,omitempty"`

	// RawDataTotalCount
	// Total AAA Servers count
	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	// TotalCount
	// Total AAA Servers count in this response
	TotalCount *int `json:"totalCount,omitempty"`
}

type SwitchMAAAServersQueryResultAPIResponse struct {
	*RawAPIResponse
	Data *SwitchMAAAServersQueryResult
}

func newSwitchMAAAServersQueryResultAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(SwitchMAAAServersQueryResultAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *SwitchMAAAServersQueryResultAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(SwitchMAAAServersQueryResult)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewSwitchMAAAServersQueryResult() *SwitchMAAAServersQueryResult {
	m := new(SwitchMAAAServersQueryResult)
	return m
}

// SwitchMAAAServersCreateAdminAAAServer
//
// Definition: aaaServers_createAdminAAAServer
type SwitchMAAAServersCreateAdminAAAServer struct {
	// AcctPort
	// AAA server accounting port
	AcctPort *int `json:"acctPort,omitempty"`

	AltoId *string `json:"altoId,omitempty"`

	// AuthPort
	// AAA server authentication port
	AuthPort *int `json:"authPort,omitempty"`

	// Ip
	// AAA server IP address
	Ip *string `json:"ip,omitempty"`

	// Level
	// Access level of AAA server
	// Constraints:
	//    - oneof:[READ_WRITE,PORT_CONFIG,READ_ONLY]
	Level *string `json:"level,omitempty"`

	// Name
	// Name of the AAA server
	Name *string `json:"name,omitempty"`

	// Password
	// Password for local user
	Password *string `json:"password,omitempty"`

	// Purpose
	// AAA server purpose
	// Constraints:
	//    - oneof:[DEFAULT,AUTHENTICATION_ONLY,AUTHORIZATION_ONLY,ACCOUNTING_ONLY]
	Purpose *string `json:"purpose,omitempty"`

	// Secret
	// AAA server secret
	Secret *string `json:"secret,omitempty"`

	// ServerType
	// The server tpye of the AAA server
	// Constraints:
	//    - oneof:[RADIUS,TACACS_PLUS,LOCAL]
	ServerType *string `json:"serverType,omitempty"`

	// Username
	// Username for local user
	Username *string `json:"username,omitempty"`
}

func NewSwitchMAAAServersCreateAdminAAAServer() *SwitchMAAAServersCreateAdminAAAServer {
	m := new(SwitchMAAAServersCreateAdminAAAServer)
	return m
}

// AddGroupAaaServersByGroupId
//
// Use this API command to create a new AAA server.
//
// Operation ID: addGroupAaaServersByGroupId
// Operation path: /group/{groupId}/aaaServers
// Success code: 201 (Created)
//
// Request body:
//	 - body *SwitchMAAAServersCreateAdminAAAServer
//
// Required parameters:
// - groupId string
//		- required
func (s *SwitchMAAAServersService) AddGroupAaaServersByGroupId(ctx context.Context, body *SwitchMAAAServersCreateAdminAAAServer, groupId string, mutators ...RequestMutator) (*SwitchMCommonCreateResultAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPost, RouteSwitchMAddGroupAaaServersByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*SwitchMCommonCreateResultAPIResponse), err
	}
	req.PathParams.Set("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusCreated, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMCommonCreateResultAPIResponse), err
}

// DeleteGroupAaaServersByGroupId
//
// Use this API command to delete AAA Servers.
//
// Operation ID: deleteGroupAaaServersByGroupId
// Operation path: /group/{groupId}/aaaServers
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMCommonBulkDeleteRequest
//
// Required parameters:
// - groupId string
//		- required
func (s *SwitchMAAAServersService) DeleteGroupAaaServersByGroupId(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, groupId string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteGroupAaaServersByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// DeleteGroupAaaServersById
//
// Use this API command to delete a AAA server.
//
// Operation ID: deleteGroupAaaServersById
// Operation path: /group/{groupId}/aaaServers/{id}
// Success code: 204 (No Content)
//
// Required parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMAAAServersService) DeleteGroupAaaServersById(ctx context.Context, groupId string, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodDelete, RouteSwitchMDeleteGroupAaaServersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, "*/*")
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}

// FindGroupAaaServersByGroupId
//
// Use this API command to retrieve a list of AAA server.
//
// Operation ID: findGroupAaaServersByGroupId
// Operation path: /group/{groupId}/aaaServers
// Success code: 200 (OK)
//
// Required parameters:
// - groupId string
//		- required
func (s *SwitchMAAAServersService) FindGroupAaaServersByGroupId(ctx context.Context, groupId string, mutators ...RequestMutator) (*SwitchMAAAServersQueryResultAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMAAAServersQueryResultAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMAAAServersQueryResultAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindGroupAaaServersByGroupId, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMAAAServersQueryResultAPIResponse), err
}

// FindGroupAaaServersById
//
// Use this API command to retrieve a AAA server.
//
// Operation ID: findGroupAaaServersById
// Operation path: /group/{groupId}/aaaServers/{id}
// Success code: 200 (OK)
//
// Required parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMAAAServersService) FindGroupAaaServersById(ctx context.Context, groupId string, id string, mutators ...RequestMutator) (*SwitchMAAAServersAAAServerAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newSwitchMAAAServersAAAServerAPIResponse
	)
	if err = ctx.Err(); err != nil {
		return resp.(*SwitchMAAAServersAAAServerAPIResponse), err
	}
	req = apiRequestFromPool(http.MethodGet, RouteSwitchMFindGroupAaaServersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*SwitchMAAAServersAAAServerAPIResponse), err
}

// UpdateGroupAaaServersById
//
// Use this API command to modify the basic information on a AAA server by complete attributes.
//
// Operation ID: updateGroupAaaServersById
// Operation path: /group/{groupId}/aaaServers/{id}
// Success code: 204 (No Content)
//
// Request body:
//	 - body *SwitchMAAAServersCreateAdminAAAServer
//
// Required parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMAAAServersService) UpdateGroupAaaServersById(ctx context.Context, body *SwitchMAAAServersCreateAdminAAAServer, groupId string, id string, mutators ...RequestMutator) (*EmptyAPIResponse, error) {
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
	req = apiRequestFromPool(http.MethodPut, RouteSwitchMUpdateGroupAaaServersById, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyContentType, headerValueApplicationJSON)
	req.Header.Set(headerKeyAccept, "*/*")
	if err = req.SetBody(body); err != nil {
		return resp.(*EmptyAPIResponse), err
	}
	req.PathParams.Set("groupId", groupId)
	req.PathParams.Set("id", id)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusNoContent, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*EmptyAPIResponse), err
}
