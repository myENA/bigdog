package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSessionmanagementService struct {
	apiClient *APIClient
}

func NewWSGSessionmanagementService(c *APIClient) *WSGSessionmanagementService {
	s := new(WSGSessionmanagementService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSessionmanagementService() *WSGSessionmanagementService {
	return NewWSGSessionmanagementService(ss.apiClient)
}

type WSGSessionmanagementRuckusSession struct {
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

func NewWSGSessionmanagementRuckusSession() *WSGSessionmanagementRuckusSession {
	m := new(WSGSessionmanagementRuckusSession)
	return m
}

type WSGSessionmanagementRuckusSessions struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSessionmanagementRuckusSession `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSessionmanagementRuckusSessions() *WSGSessionmanagementRuckusSessions {
	m := new(WSGSessionmanagementRuckusSessions)
	return m
}

// FindSessionManagement
//
// Use this API command to retrieve information about the current logon sessions.
func (s *WSGSessionmanagementService) FindSessionManagement(ctx context.Context) (*WSGSessionmanagementRuckusSessions, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     *WSGSessionmanagementRuckusSessions
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindSessionManagement, true)
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSessionmanagementRuckusSessions()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
