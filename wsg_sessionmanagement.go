package bigdog

// API Version: v9_1

import (
	"context"
	"errors"
	"io"
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

func newWSGSessionManagementRuckusSessionsAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSessionManagementRuckusSessionsAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSessionManagementRuckusSessionsAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGSessionManagementRuckusSessions)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGSessionManagementRuckusSessions() *WSGSessionManagementRuckusSessions {
	m := new(WSGSessionManagementRuckusSessions)
	return m
}

// FindSessionManagement
//
// Use this API command to retrieve information about the current logon sessions.
//
// Operation ID: findSessionManagement
// Operation path: /sessionManagement
// Success code: 200 (OK)
func (s *WSGSessionManagementService) FindSessionManagement(ctx context.Context, mutators ...RequestMutator) (*WSGSessionManagementRuckusSessionsAPIResponse, error) {
	var (
		req      *APIRequest
		httpResp *http.Response
		resp     APIResponse
		err      error

		respFn = newWSGSessionManagementRuckusSessionsAPIResponse
	)
	req = apiRequestFromPool(APISourceVSZ, http.MethodGet, RouteWSGFindSessionManagement, true)
	defer recycleAPIRequest(req)
	req.Header.Set(headerKeyAccept, "*/*")
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp, err = handleAPIResponse(req, http.StatusOK, httpResp, respFn, s.apiClient.autoHydrate, err)
	return resp.(*WSGSessionManagementRuckusSessionsAPIResponse), err
}
