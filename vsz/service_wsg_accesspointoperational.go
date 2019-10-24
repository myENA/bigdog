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

func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureDownloadByApMac(ctx context.Context, pApMac string) (json.RawMessage, error) {
}

func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartFileCaptureByApMac(ctx context.Context, body *appackcapture.ApPacketCaptureReq, pApMac string) (*appackcapture.ApPacketCaptureRes, error) {
}

func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartStreamingByApMac(ctx context.Context, body *appackcapture.ApPacketCaptureReq, pApMac string) (*appackcapture.ApPacketCaptureRes, error) {
}

func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStopByApMac(ctx context.Context, pApMac string) error {
}

func (s *WSGAccessPointOperationalService) AddApsOperationalBlinkLedByApMac(ctx context.Context, pApMac string) error {
}

func (s *WSGAccessPointOperationalService) AddApsSwitchoverCluster(ctx context.Context, body *ap.SwitchoverAP) (*common.EmptyResult, error) {
}

func (s *WSGAccessPointOperationalService) FindApsApPacketCaptureByApMac(ctx context.Context, pApMac string) (*appackcapture.ApPacketCaptureRes, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmsByApMac(ctx context.Context, pApMac string, qCategory string, qCode float64, qEndTime string, qIndex string, qListSize string, qSeverity string, qStartTime string, qStatus string) (*ap.AlarmList, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmSummaryByApMac(ctx context.Context, pApMac string) (*ap.AlarmSummary, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalEventSummaryByApMac(ctx context.Context, pApMac string) (*ap.EventSummary, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalNeighborByApMac(ctx context.Context, pApMac string, qIndex string, qListSize string) (*ap.NeighborAPList, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalSummaryByApMac(ctx context.Context, pApMac string) (*ap.ApOperationalSummary, error) {
}
