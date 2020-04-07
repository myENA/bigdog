package vsz

// API Version: v9_0

import (
	"context"
	"net/http"
)

type WSGSessionManagementService struct {
	apiClient *APIClient
}

func NewWSGSessionManagementService(c *APIClient) *WSGSessionManagementService {
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
	// Constraints:
	//    - nullable
	AuthType *string `json:"authType,omitempty"`

	// LastAccessTime
	// The last access time
	// Constraints:
	//    - nullable
	LastAccessTime *string `json:"lastAccessTime,omitempty"`

	// LastAccessURI
	// The last access URI
	// Constraints:
	//    - nullable
	LastAccessURI *string `json:"lastAccessURI,omitempty"`

	// SessionId
	// The user session ID
	// Constraints:
	//    - nullable
	SessionId *string `json:"sessionId,omitempty"`

	// SourceIp
	// The source IP address
	// Constraints:
	//    - nullable
	SourceIp *string `json:"sourceIp,omitempty"`

	// UserName
	// Logon user name
	// Constraints:
	//    - nullable
	UserName *string `json:"userName,omitempty"`

	// UserUUID
	// The user UUID
	// Constraints:
	//    - nullable
	UserUUID *string `json:"userUUID,omitempty"`
}

func NewWSGSessionManagementRuckusSession() *WSGSessionManagementRuckusSession {
	m := new(WSGSessionManagementRuckusSession)
	return m
}

type WSGSessionManagementRuckusSessions struct {
	// FirstIndex
	// Constraints:
	//    - nullable
	FirstIndex *int `json:"firstIndex,omitempty"`

	// HasMore
	// Constraints:
	//    - nullable
	HasMore *bool `json:"hasMore,omitempty"`

	// List
	// Constraints:
	//    - nullable
	List []*WSGSessionManagementRuckusSession `json:"list,omitempty" validate:"omitempty,dive"`

	// TotalCount
	// Constraints:
	//    - nullable
	TotalCount *int `json:"totalCount,omitempty"`
}

func NewWSGSessionManagementRuckusSessions() *WSGSessionManagementRuckusSessions {
	m := new(WSGSessionManagementRuckusSessions)
	return m
}

// FindSessionManagement
//
// Use this API command to retrieve information about the current logon sessions.
func (s *WSGSessionManagementService) FindSessionManagement(ctx context.Context) (*WSGSessionManagementRuckusSessions, *APIResponseMeta, error) {
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
	httpResp, err = s.apiClient.Do(ctx, req)
	resp = NewWSGSessionManagementRuckusSessions()
	rm, err = handleResponse(req, http.StatusOK, httpResp, resp, err)
	return resp, rm, err
}
