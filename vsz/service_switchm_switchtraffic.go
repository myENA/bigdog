package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/switchm/common"
	"github.com/myENA/ruckus-client/vsz/types/switchm/traffic"
)

type SwitchMSwitchTrafficService struct {
	client *Client
}

func NewSwitchMSwitchTrafficService(client *Client) *SwitchMSwitchTrafficService {
	s := new(SwitchMSwitchTrafficService)
	s.client = client
	return s
}

func (ss *SwitchMService) SwitchMSwitchTrafficService() *SwitchMSwitchTrafficService {
	serv := new(SwitchMSwitchTrafficService)
	serv.client = ss.client
	return serv
}

// AddTrafficTopPoeutilization
//
// Use this API command retrieve the top 10 switches by the PoE utilization.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTopPoeutilization(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TopSwitchPoEUtilizationQueryResultList, error) {
}

// AddTrafficTopPorterror
//
// Use this API command to get the top 10 switches by the porterror.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTopPorterror(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TopPortErrorQueryResultList, error) {
}

// AddTrafficTopPortusage
//
// Use this API command to get the top 10 ports by the traffic.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTopPortusage(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TopPortTrafficUsageQueryResultList, error) {
}

// AddTrafficTopUsage
//
// Use this API command to retrieve Top Swich/Port usage data based on the time duration.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTopUsage(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TopTrafficUsageQueryResultList, error) {
}

// AddTrafficTotalTrend
//
// Use this API command to retrieve Swich/Port trend data based on the time duration.
//
// Request Body:
//	 - body *common.QueryCriteriaSuperSet
func (s *SwitchMSwitchTrafficService) AddTrafficTotalTrend(ctx context.Context, body *common.QueryCriteriaSuperSet) (*traffic.TrafficQueryResultList, error) {
}
