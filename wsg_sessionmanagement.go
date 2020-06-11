package bigdog

// API Version: v9_0

import (
	"context"
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

type WSGSessionManagementRuckusSessions struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSessionManagementRuckusSession `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSessionManagementRuckusSessions() *WSGSessionManagementRuckusSessions {
	m := new(WSGSessionManagementRuckusSessions)
	return m
}

// FindSessionManagement
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
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSessionManagement, true)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = NewWSGSessionManagementRuckusSessions()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
