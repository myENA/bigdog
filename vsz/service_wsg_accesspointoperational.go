package vsz

// API Version: v8_1

import (
	"context"
	"encoding/json"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/appackcapture"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGAccessPointOperationalService struct {
	client *Client
}

func NewWSGAccessPointOperationalService(client *Client) *WSGAccessPointOperationalService {
	s := new(WSGAccessPointOperationalService)
	s.client = client
	return s
}

func (ss *WSGService) WSGAccessPointOperationalService() *WSGAccessPointOperationalService {
	serv := new(WSGAccessPointOperationalService)
	serv.client = ss.client
	return serv
}

// AddApsApPacketCaptureDownloadByApMac
//
// Use this API to download AP packet capture file
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureDownloadByApMac(ctx context.Context, pApMac string) (json.RawMessage, error) {
}

// AddApsApPacketCaptureStartFileCaptureByApMac
//
// Use this API to start AP packet capture
//
// Request Body:
//	 - body *appackcapture.ApPacketCaptureReq
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartFileCaptureByApMac(ctx context.Context, body *appackcapture.ApPacketCaptureReq, pApMac string) (*appackcapture.ApPacketCaptureRes, error) {
}

// AddApsApPacketCaptureStartStreamingByApMac
//
// Use this API to start AP packet streaming
//
// Request Body:
//	 - body *appackcapture.ApPacketCaptureReq
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartStreamingByApMac(ctx context.Context, body *appackcapture.ApPacketCaptureReq, pApMac string) (*appackcapture.ApPacketCaptureRes, error) {
}

// AddApsApPacketCaptureStopByApMac
//
// Use this API to stop AP packet capture or streaming
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStopByApMac(ctx context.Context, pApMac string) error {
}

// AddApsOperationalBlinkLedByApMac
//
// use this API to make ap blink its led to show its position.
func (s *WSGAccessPointOperationalService) AddApsOperationalBlinkLedByApMac(ctx context.Context, pApMac string) error {
}

// AddApsSwitchoverCluster
//
// Use this API command to switchover AP to another cluster
//
// Request Body:
//	 - body *ap.SwitchoverAP
func (s *WSGAccessPointOperationalService) AddApsSwitchoverCluster(ctx context.Context, body *ap.SwitchoverAP) (*common.EmptyResult, error) {
}

// FindApsApPacketCaptureByApMac
//
// Use this API to get AP packet capture status
func (s *WSGAccessPointOperationalService) FindApsApPacketCaptureByApMac(ctx context.Context, pApMac string) (*appackcapture.ApPacketCaptureRes, error) {
}

// FindApsOperationalAlarmsByApMac
//
// Use this API command to retrieve the list of alarms on an AP.
func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmsByApMac(ctx context.Context, pApMac string, qCategory string, qCode float64, qEndTime string, qIndex string, qListSize string, qSeverity string, qStartTime string, qStatus string) (*ap.AlarmList, error) {
}

// FindApsOperationalAlarmSummaryByApMac
//
// Use this API command to retrieve the alarm summary of an AP.
func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmSummaryByApMac(ctx context.Context, pApMac string) (*ap.AlarmSummary, error) {
}

// FindApsOperationalEventSummaryByApMac
//
// Use this API command to retrieve the event summary of an AP.
func (s *WSGAccessPointOperationalService) FindApsOperationalEventSummaryByApMac(ctx context.Context, pApMac string) (*ap.EventSummary, error) {
}

// FindApsOperationalNeighborByApMac
//
// Use this API command to retrieve a list of neighbor access points on mesh AP.
func (s *WSGAccessPointOperationalService) FindApsOperationalNeighborByApMac(ctx context.Context, pApMac string, qIndex string, qListSize string) (*ap.NeighborAPList, error) {
}

// FindApsOperationalSummaryByApMac
//
// Use this API command to retrieve the operational information of an AP.
func (s *WSGAccessPointOperationalService) FindApsOperationalSummaryByApMac(ctx context.Context, pApMac string) (*ap.ApOperationalSummary, error) {
}
