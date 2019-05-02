package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGRogueClassificationPolicyService struct {
    c *Client
}

func NewWSGRogueClassificationPolicyService (c *Client) *WSGRogueClassificationPolicyService {
    s := new(WSGRogueClassificationPolicyService)
    s.c = c
    return s
}

