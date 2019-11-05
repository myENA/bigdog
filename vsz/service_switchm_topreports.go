package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/switchmswitch"
)

type SwitchMTopReportsService struct {
	apiClient *APIClient
}

func NewSwitchMTopReportsService(c *APIClient) *SwitchMTopReportsService {
	s := new(SwitchMTopReportsService)
	s.apiClient = c
	return s
}

func (ss *SwitchMService) SwitchMTopReportsService() *SwitchMTopReportsService {
	serv := new(SwitchMTopReportsService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddSwitchTopByFirmware
//
// Use this API command to retrieves top N switch count based on firmware version.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByFirmware(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.TopSwitchesByFirmwareQueryResultList, error) {
}

// AddSwitchTopByModel
//
// Use this API command to retrieve top N switch count based on switch model.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMTopReportsService) AddSwitchTopByModel(ctx context.Context, body *common.QueryCriteriaSuperSet) (*switchmswitch.TopSwitchesByModelQueryResultList, error) {
}
