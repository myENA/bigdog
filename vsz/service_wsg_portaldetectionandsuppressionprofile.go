package vsz

// API Version: v8_0

import (
	"context"
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

func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesById (ctx context.Context, id string, zoneId string) (portaldetectionprofile.PortalDetectionProfile, error) {
}

func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByQueryCriteria (ctx context.Context) (portaldetectionprofile.PortalDetectionProfileList, error) {
}

func (s *WSGPortalDetectionandSuppressionProfileService) FindRkszonesPortalDetectionProfilesByZoneId (ctx context.Context, zoneId string) (portaldetectionprofile.PortalDetectionProfileList, error) {
}

