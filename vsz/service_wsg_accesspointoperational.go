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

func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureDownloadByApMac(ctx context.Context, apMac string) (json.RawMessage, error) {
}

func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartFileCaptureByApMac(ctx context.Context, apMac string) (*appackcapture.ApPacketCaptureRes, error) {
}

func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartStreamingByApMac(ctx context.Context, apMac string) (*appackcapture.ApPacketCaptureRes, error) {
}

func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStopByApMac(ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointOperationalService) AddApsOperationalBlinkLedByApMac(ctx context.Context, apMac string) error {
}

func (s *WSGAccessPointOperationalService) AddApsSwitchoverCluster(ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGAccessPointOperationalService) FindApsApPacketCaptureByApMac(ctx context.Context, apMac string) (*appackcapture.ApPacketCaptureRes, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmsByApMac(ctx context.Context, apMac string) (*ap.AlarmList, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmSummaryByApMac(ctx context.Context, apMac string) (*ap.AlarmSummary, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalEventSummaryByApMac(ctx context.Context, apMac string) (*ap.EventSummary, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalNeighborByApMac(ctx context.Context, apMac string) (*ap.NeighborAPList, error) {
}

func (s *WSGAccessPointOperationalService) FindApsOperationalSummaryByApMac(ctx context.Context, apMac string) (*ap.ApOperationalSummary, error) {
}
