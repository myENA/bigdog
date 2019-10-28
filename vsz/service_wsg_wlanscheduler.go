package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/wlanscheduler"
)

type WSGWLANSchedulerService struct {
	client *Client
}

func NewWSGWLANSchedulerService(client *Client) *WSGWLANSchedulerService {
	s := new(WSGWLANSchedulerService)
	s.client = client
	return s
}

func (ss *WSGService) WSGWLANSchedulerService() *WSGWLANSchedulerService {
	serv := new(WSGWLANSchedulerService)
	serv.client = ss.client
	return serv
}

// FindRkszonesWlanSchedulersById
//
// Use this API command to retrieve a WLAN schedule.
//
// Path Parameters:
// - pId string
//		- required
// - pZoneId string
//		- required
func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersById(ctx context.Context, pId string, pZoneId string) (*wlanscheduler.WlanSchedule, error) {
}

// FindRkszonesWlanSchedulersByZoneId
//
// Use this API command to retrieve the list of WLAN schedule from a zone.
//
// Path Parameters:
// - pZoneId string
//		- required
//
// Query Parameters:
// - qIndex string
// - qListSize string
func (s *WSGWLANSchedulerService) FindRkszonesWlanSchedulersByZoneId(ctx context.Context, pZoneId string, qIndex string, qListSize string) (*wlanscheduler.WlanScheduleList, error) {
}
