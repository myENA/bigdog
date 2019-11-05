package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/system"
)

type WSGLWAPPTOSCGService struct {
	apiClient *APIClient
}

func NewWSGLWAPPTOSCGService(c *APIClient) *WSGLWAPPTOSCGService {
	s := new(WSGLWAPPTOSCGService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGLWAPPTOSCGService() *WSGLWAPPTOSCGService {
	serv := new(WSGLWAPPTOSCGService)
	serv.apiClient = ss.apiClient
	return serv
}

// FindLwapp2scg
//
// Use this API command to retrieve Lwapp Config.
func (s *WSGLWAPPTOSCGService) FindLwapp2scg(ctx context.Context) (*system.Lwapp2scgConfiguration, error) {
}

// PartialUpdateLwapp2scg
//
// Use this API command to modify the basic information of the Lwapp Config.
//
// Request Body:
//	 - body *system.ModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scg(ctx context.Context, body *system.ModifyLwapp2scg) (*common.EmptyResult, error) {
}

// PartialUpdateLwapp2scgApList
//
// Use this API command to modify the apList of the Lwapp Config.
//
// Request Body:
//	 - body *system.ModifyLwapp2scg
func (s *WSGLWAPPTOSCGService) PartialUpdateLwapp2scgApList(ctx context.Context, body *system.ModifyLwapp2scg) (*common.EmptyResult, error) {
}
