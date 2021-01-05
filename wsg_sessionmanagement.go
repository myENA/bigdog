package bigdog

// API Version: v9_1

import (
	"context"
	"encoding/json"
	"net/http"
)

type WSGSessionManagementService struct {
	apiClient *VSZClient
}

func NewWSGSessionManagementService(c *VSZClient) *WSGSessionManagementService {
	s := new(WSGSessionManagementService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSessionManagementService() *WSGSessionManagementService {
	return NewWSGSessionManagementService(ss.apiClient)
}

// WSGSessionManagementRuckusSession
//
// Definition: sessionManagement_ruckusSession
type WSGSessionManagementRuckusSession struct {
	// AuthType
	// The authentication type of logon
	AuthType *string `json:"authType,omitempty"`

	// LastAccessTime
	// The last access time
	LastAccessTime *string `json:"lastAccessTime,omitempty"`

	// LastAccessURI
	// The last access URI
	LastAccessURI *string `json:"lastAccessURI,omitempty"`

	// SessionId
	// The user session ID
	SessionId *string `json:"sessionId,omitempty"`

	// SourceIp
	// The source IP address
	SourceIp *string `json:"sourceIp,omitempty"`

	// UserName
	// Logon user name
	UserName *string `json:"userName,omitempty"`

	// UserUUID
	// The user UUID
	UserUUID *string `json:"userUUID,omitempty"`
}

func NewWSGSessionManagementRuckusSession() *WSGSessionManagementRuckusSession {
	m := new(WSGSessionManagementRuckusSession)
	return m
}

// WSGSessionManagementRuckusSessions
//
// Definition: sessionManagement_ruckusSessions
type WSGSessionManagementRuckusSessions struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSessionManagementRuckusSession `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGSessionManagementRuckusSessionsAPIResponse struct {
	*RawAPIResponse
	Data *WSGSessionManagementRuckusSessions
}

func newWSGSessionManagementRuckusSessionsAPIResponse(req *APIRequest, successCode int, httpResp *http.Response) APIResponse {
	r := new(WSGSessionManagementRuckusSessionsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(req, successCode, httpResp).(*RawAPIResponse)
	return r
}

func (r *WSGSessionManagementRuckusSessionsAPIResponse) Hydrate() error {
	r.Data = new(WSGSessionManagementRuckusSessions)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSessionManagementRuckusSessions() *WSGSessionManagementRuckusSessions {
	m := new(WSGSessionManagementRuckusSessions)
	return m
}

// FindSessionManagement
//
// Operation ID: findSessionManagement
//
// Use this API command to retrieve information about the current logon sessions.
func (s *WSGSessionManagementService) FindSessionManagement(ctx context.Context, mutators ...RequestMutator) (*WSGSessionManagementRuckusSessions, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSessionManagementRuckusSessions
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = apiRequestFromPool(http.MethodGet, RouteWSGFindSessionManagement, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSessionManagementRuckusSessions()
	rm, err = handleAPIResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
