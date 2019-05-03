package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGHotspot20WiFiOperatorProfileService struct {
    client *Client
}

func NewWSGHotspot20WiFiOperatorProfileService (client *Client) *WSGHotspot20WiFiOperatorProfileService {
    s := new(WSGHotspot20WiFiOperatorProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGHotspot20WiFiOperatorProfileService () *WSGHotspot20WiFiOperatorProfileService {
    serv := new(WSGHotspot20WiFiOperatorProfileService)
    serv.client = ss.client
    return serv
}

func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20OperatorsCertificateById (ctx context.Context, id string) (*common.EmptyResult, error) {
}

func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20Operators (ctx context.Context) (*profile.Hs20OperatorList, error) {
}

func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsById (ctx context.Context, id string) (*profile.Hs20Operator, error) {
}

func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsByQueryCriteria (ctx context.Context) (*profile.Hs20OperatorList, error) {
}

