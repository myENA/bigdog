package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portaldetectionprofile"
)

type WSGPortalDetectionandSuppressionProfileService struct {
    c *Client
}

func NewWSGPortalDetectionandSuppressionProfileService (c *Client) *WSGPortalDetectionandSuppressionProfileService {
    s := new(WSGPortalDetectionandSuppressionProfileService)
    s.c = c
    return s
}

