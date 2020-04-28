package vsz

// API Version: v9_0

import (
	"context"
	"encoding/json"
	"net/http"
)

type SwitchMAAAServersService struct {
	apiClient *APIClient
}

func NewSwitchMAAAServersService(c *APIClient) *SwitchMAAAServersService {
	s := new(SwitchMAAAServersService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMAAAServersService() *SwitchMAAAServersService {
	return NewSwitchMAAAServersService(ss.apiClient)
}

type SwitchMAAAServersAAAServer struct {
	// AcctPort
	// AAA server accounting port
	AcctPort *int `json:"acctPort,omitempty"`

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

func NewSwitchMAAAServersAAAServer() *SwitchMAAAServersAAAServer {
	m := new(SwitchMAAAServersAAAServer)
	return m
}

type SwitchMAAAServersQueryResult struct {
	// Extra
	// Any additional response data
	Extra *SwitchMAAAServersQueryResultExtraType `json:"extra,omitempty"`

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

func NewSwitchMAAAServersQueryResult() *SwitchMAAAServersQueryResult {
	m := new(SwitchMAAAServersQueryResult)
	return m
}

// SwitchMAAAServersQueryResultExtraType
//
// Any additional response data
type SwitchMAAAServersQueryResultExtraType struct {
	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SwitchMAAAServersQueryResultExtraType) UnmarshalJSON(b []byte) error {
	tmp := make(map[string]interface{})
	if err := json.Unmarshal(b, &tmp); err != nil {
		return err
	}
	*t = SwitchMAAAServersQueryResultExtraType{XAdditionalProperties: tmp}
	return nil
}

func (t *SwitchMAAAServersQueryResultExtraType) MarshalJSON() ([]byte, error) {
	if t == nil || t.XAdditionalProperties == nil {
		return nil, nil
	}
	return json.Marshal(t.XAdditionalProperties)
}

func NewSwitchMAAAServersQueryResultExtraType() *SwitchMAAAServersQueryResultExtraType {
	m := new(SwitchMAAAServersQueryResultExtraType)
	return m
}

type SwitchMAAAServersCreateAdminAAAServer struct {
	// AcctPort
	// AAA server accounting port
	AcctPort *int `json:"acctPort,omitempty"`

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
// Request Body:
//	 - body *SwitchMAAAServersCreateAdminAAAServer
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMAAAServersService) AddGroupAaaServersByGroupId(ctx context.Context, body *SwitchMAAAServersCreateAdminAAAServer, groupId string) (*SwitchMCommonCreateResult, *APIResponseMeta, error) {
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
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSwitchMAddGroupAaaServersByGroupId, true)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMCommonCreateResult()
	rm, err = handleResponse(req, http.StatusCreated, httpResp, resp, err)
	return resp, rm, err
}

// DeleteGroupAaaServersByGroupId
//
// Use this API command to delete AAA Servers.
//
// Request Body:
//	 - body *SwitchMCommonBulkDeleteRequest
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMAAAServersService) DeleteGroupAaaServersByGroupId(ctx context.Context, body *SwitchMCommonBulkDeleteRequest, groupId string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteGroupAaaServersByGroupId, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// DeleteGroupAaaServersById
//
// Use this API command to delete a AAA server.
//
// Required Parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMAAAServersService) DeleteGroupAaaServersById(ctx context.Context, groupId string, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteSwitchMDeleteGroupAaaServersById, true)
	req.SetPathParameter("groupId", groupId)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}

// FindGroupAaaServersByGroupId
//
// Use this API command to retrieve a list of AAA server.
//
// Required Parameters:
// - groupId string
//		- required
func (s *SwitchMAAAServersService) FindGroupAaaServersByGroupId(ctx context.Context, groupId string) (*SwitchMAAAServersQueryResult, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMAAAServersQueryResult
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindGroupAaaServersByGroupId, true)
	req.SetPathParameter("groupId", groupId)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMAAAServersQueryResult()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// FindGroupAaaServersById
//
// Use this API command to retrieve a AAA server.
//
// Required Parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMAAAServersService) FindGroupAaaServersById(ctx context.Context, groupId string, id string) (*SwitchMAAAServersAAAServer, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *SwitchMAAAServersAAAServer
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return resp, rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteSwitchMFindGroupAaaServersById, true)
	req.SetPathParameter("groupId", groupId)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewSwitchMAAAServersAAAServer()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}

// UpdateGroupAaaServersById
//
// Use this API command to modify the basic information on a AAA server by complete attributes.
//
// Request Body:
//	 - body *SwitchMAAAServersCreateAdminAAAServer
//
// Required Parameters:
// - groupId string
//		- required
// - id string
//		- required
func (s *SwitchMAAAServersService) UpdateGroupAaaServersById(ctx context.Context, body *SwitchMAAAServersCreateAdminAAAServer, groupId string, id string) (*APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return rm, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, groupId, "required"); err != nil {
		return rm, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return rm, err
	}
	req = NewAPIRequest(http.MethodPut, RouteSwitchMUpdateGroupAaaServersById, true)
	if err = req.SetBody(body); err != nil {
		return rm, err
	}
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetPathParameter("groupId", groupId)
	req.SetPathParameter("id", id)
	httpResp, err = s.apiClient.Do(ctx, req)
	rm, err = handleResponse(req, http.StatusNoContent, httpResp, nil, err)
	return rm, err
}
