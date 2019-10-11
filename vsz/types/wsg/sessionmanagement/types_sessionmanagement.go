package sessionmanagement

// API Version: v8_1

type RuckusSession struct {
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

func NewRuckusSession() *RuckusSession {
	ruckusSessionType := new(RuckusSession)
	return ruckusSessionType
}

func NewDefaultRuckusSession() *RuckusSession {
	ruckusSessionType := new(RuckusSession)
	return ruckusSessionType
}

type RuckusSessions struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*RuckusSession `json:"list,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

func NewRuckusSessions() *RuckusSessions {
	ruckusSessionsType := new(RuckusSessions)
	return ruckusSessionsType
}

func NewDefaultRuckusSessions() *RuckusSessions {
	ruckusSessionsType := new(RuckusSessions)
	return ruckusSessionsType
}
