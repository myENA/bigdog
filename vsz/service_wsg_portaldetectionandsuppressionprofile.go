package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/portaldetectionprofile"
)

type WSGPortalDetectionandSuppressionProfileService struct {
    client *Client
}

func NewWSGPortalDetectionandSuppressionProfileService (client *Client) *WSGPortalDetectionandSuppressionProfileService {
    s := new(WSGPortalDetectionandSuppressionProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGPortalDetectionandSuppressionProfileService () *WSGPortalDetectionandSuppressionProfileService {
    serv := new(WSGPortalDetectionandSuppressionProfileService)
    serv.client = ss.client
    return serv
}

