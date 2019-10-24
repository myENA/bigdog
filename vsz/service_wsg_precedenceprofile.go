package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGPrecedenceProfileService struct {
	client *Client
}

func NewWSGPrecedenceProfileService(client *Client) *WSGPrecedenceProfileService {
	s := new(WSGPrecedenceProfileService)
	s.client = client
	return s
}

func (ss *WSGService) WSGPrecedenceProfileService() *WSGPrecedenceProfileService {
	serv := new(WSGPrecedenceProfileService)
	serv.client = ss.client
	return serv
}

// FindPrecedence
//
// Use this API command to Get Precedence Profile list.
func (s *WSGPrecedenceProfileService) FindPrecedence(ctx context.Context, qIndex string, qListSize string) (*profile.PrecedenceList, error) {
}

// FindPrecedenceById
//
// Use this API command to Get Precedence Profile by profile's ID.
func (s *WSGPrecedenceProfileService) FindPrecedenceById(ctx context.Context, pId string) (*profile.CreatePrecedenceProfile, error) {
}

// FindPrecedenceByQueryCriteria
//
// Use this API command to query Precedence Profile.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *WSGPrecedenceProfileService) FindPrecedenceByQueryCriteria(ctx context.Context, body *common.QueryCriteriaSuperSet) (*profile.PrecedenceList, error) {
}
