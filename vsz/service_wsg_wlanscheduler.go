package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlanscheduler"
)

type WSGWLANSchedulerService struct {
    client *Client
}

func NewWSGWLANSchedulerService (client *Client) *WSGWLANSchedulerService {
    s := new(WSGWLANSchedulerService)
    s.client = client
    return s
}

func (ss *WSGService) WSGWLANSchedulerService () *WSGWLANSchedulerService {
    serv := new(WSGWLANSchedulerService)
    serv.client = ss.client
    return serv
}

func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersById (ctx context.Context, id string, zoneId string) (*wlanscheduler.WlanSchedule, error) {
}

func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersByZoneId (ctx context.Context, zoneId string) (*wlanscheduler.WlanScheduleList, error) {
}

