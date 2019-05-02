package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRogueClassificationPolicyService struct {
    client *Client
}

func NewWSGRogueClassificationPolicyService (client *Client) *WSGRogueClassificationPolicyService {
    s := new(WSGRogueClassificationPolicyService)
    s.client = client
    return s
}

func (ss *WSGService) WSGRogueClassificationPolicyService () *WSGRogueClassificationPolicyService {
    serv := new(WSGRogueClassificationPolicyService)
    serv.client = ss.client
    return serv
}

