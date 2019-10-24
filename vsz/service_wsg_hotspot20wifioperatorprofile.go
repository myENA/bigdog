package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGHotspot20WiFiOperatorProfileService struct {
	client *Client
}

func NewWSGHotspot20WiFiOperatorProfileService(client *Client) *WSGHotspot20WiFiOperatorProfileService {
	s := new(WSGHotspot20WiFiOperatorProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGHotspot20WiFiOperatorProfileService() *WSGHotspot20WiFiOperatorProfileService {
	serv := new(WSGHotspot20WiFiOperatorProfileService)
	serv.client = ss.client
	return serv
}

// DeleteProfilesHs20OperatorsCertificateById
//
// Use this API command to disable certificate of a Hotspot 2.0 Wi-Fi operator.
func (s *WSGHotspot20WiFiOperatorProfileService) DeleteProfilesHs20OperatorsCertificateById(ctx context.Context, pId string) (*common.EmptyResult, error) {
}

// FindProfilesHs20Operators
//
// Use this API command to retrieve list of Hotspot 2.0 Wi-Fi Operators.
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20Operators(ctx context.Context, qIndex string, qListSize string) (*profile.Hs20OperatorList, error) {
}

// FindProfilesHs20OperatorsById
//
// Use this API command to retrieve a Hotspot 2.0 Wi-Fi operator.
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsById(ctx context.Context, pId string) (*profile.Hs20Operator, error) {
}

// FindProfilesHs20OperatorsByQueryCriteria
//
// Query hotspot 2.0 Wi-Fi operators.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGHotspot20WiFiOperatorProfileService) FindProfilesHs20OperatorsByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.Hs20OperatorList, error) {
}
