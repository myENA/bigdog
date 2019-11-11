package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
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

type WSGSessionManagementRuckusSessions struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGSessionManagementRuckusSession `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

// FindSessionManagement
//
// Use this API command to retrieve information about the current logon sessions.
func (s *WSGSessionManagementService) FindSessionManagement(ctx context.Context) (*WSGSessionManagementRuckusSessions, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
