package vsz

// API Version: v8_1

import (
	"context"
	"errors"
	"fmt"
	"gopkg.in/go-playground/validator.v9"
)

type WSGAccessPointOperationalService struct {
	apiClient *APIClient
}

func NewWSGAccessPointOperationalService(c *APIClient) *WSGAccessPointOperationalService {
	s := new(WSGAccessPointOperationalService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGAccessPointOperationalService() *WSGAccessPointOperationalService {
	return NewWSGAccessPointOperationalService(ss.apiClient)
}

// AddApsApPacketCaptureDownloadByApMac
//
// Use this API to download AP packet capture file
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureDownloadByApMac(ctx context.Context, pApMac string) ([]byte, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddApsApPacketCaptureStartFileCaptureByApMac
//
// Use this API to start AP packet capture
//
// Request Body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartFileCaptureByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, pApMac string) (*WSGAPPackCaptureApPacketCaptureRes, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddApsApPacketCaptureStartStreamingByApMac
//
// Use this API to start AP packet streaming
//
// Request Body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartStreamingByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, pApMac string) (*WSGAPPackCaptureApPacketCaptureRes, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddApsApPacketCaptureStopByApMac
//
// Use this API to stop AP packet capture or streaming
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStopByApMac(ctx context.Context, pApMac string) (interface{}, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddApsOperationalBlinkLedByApMac
//
// use this API to make ap blink its led to show its position.
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsOperationalBlinkLedByApMac(ctx context.Context, pApMac string) (interface{}, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// AddApsSwitchoverCluster
//
// Use this API command to switchover AP to another cluster
//
// Request Body:
//	 - body *WSGAPSwitchoverAP
func (s *WSGAccessPointOperationalService) AddApsSwitchoverCluster(ctx context.Context, body *WSGAPSwitchoverAP) (*WSGCommonEmptyResult, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApsApPacketCaptureByApMac
//
// Use this API to get AP packet capture status
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsApPacketCaptureByApMac(ctx context.Context, pApMac string) (*WSGAPPackCaptureApPacketCaptureRes, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApsOperationalAlarmsByApMac
//
// Use this API command to retrieve the list of alarms on an AP.
//
// Path Parameters:
// - pApMac string
//		- required
//
// Query Parameters:
// - qCategory string
//		- nullable
// - qCode float64
//		- nullable
// - qEndTime string
//		- nullable
// - qIndex string
//		- nullable
// - qListSize string
//		- nullable
// - qSeverity string
//		- nullable
// - qStartTime string
//		- nullable
// - qStatus string
//		- nullable
func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmsByApMac(ctx context.Context, pApMac string, qCategory string, qCode float64, qEndTime string, qIndex string, qListSize string, qSeverity string, qStartTime string, qStatus string) (*WSGAPAlarmList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApsOperationalAlarmSummaryByApMac
//
// Use this API command to retrieve the alarm summary of an AP.
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmSummaryByApMac(ctx context.Context, pApMac string) (*WSGAPAlarmSummary, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApsOperationalEventSummaryByApMac
//
// Use this API command to retrieve the event summary of an AP.
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalEventSummaryByApMac(ctx context.Context, pApMac string) (*WSGAPEventSummary, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApsOperationalNeighborByApMac
//
// Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Path Parameters:
// - pApMac string
//		- required
//
// Query Parameters:
// - qIndex string
//		- nullable
// - qListSize string
//		- nullable
func (s *WSGAccessPointOperationalService) FindApsOperationalNeighborByApMac(ctx context.Context, pApMac string, qIndex string, qListSize string) (*WSGAPNeighborAPList, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}

// FindApsOperationalSummaryByApMac
//
// Use this API command to retrieve the operational information of an AP.
//
// Path Parameters:
// - pApMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalSummaryByApMac(ctx context.Context, pApMac string) (*WSGAPOperationalSummary, error) {
	if ctx == nil {
		return nil, errors.New("ctx cannot be empty")
	}
	if err := ctx.Err(); err != nil {
		return nil, fmt.Errorf("provided context is done: %s", err)
	}
}
