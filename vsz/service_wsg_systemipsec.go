package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/systemipsec"
)

type WSGSystemIPsecService struct {
	apiClient *APIClient
}

func NewWSGSystemIPsecService(c *APIClient) *WSGSystemIPsecService {
	s := new(WSGSystemIPsecService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSystemIPsecService() *WSGSystemIPsecService {
	serv := new(WSGSystemIPsecService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindSystemIpsec
//
// Use this API command to retrieve the System IPSec.
func (s *WSGSystemIPsecService) FindSystemIpsec(ctx context.Context) (*systemipsec.GetResult, error) {
}

// UpdateSystemIpsec
//
// Use this API command to modify the System IPSec.
//
// Request Body:
//	 - body *systemipsec.Update
func (s *WSGSystemIPsecService) UpdateSystemIpsec(ctx context.Context, body *systemipsec.Update) (*common.CreateResult, error) {
}
