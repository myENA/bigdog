package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
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
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureDownloadByApMac(ctx context.Context, apMac string) ([]byte, error) {
	var (
		req  *APIRequest
		resp []byte
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsApPacketCaptureDownloadByApMac, true)
}

// AddApsApPacketCaptureStartFileCaptureByApMac
//
// Use this API to start AP packet capture
//
// Request Body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartFileCaptureByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, apMac string) (*WSGAPPackCaptureApPacketCaptureRes, error) {
	var (
		req  *APIRequest
		resp *WSGAPPackCaptureApPacketCaptureRes
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsApPacketCaptureStartFileCaptureByApMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddApsApPacketCaptureStartStreamingByApMac
//
// Use this API to start AP packet streaming
//
// Request Body:
//	 - body *WSGAPPackCaptureApPacketCaptureReq
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStartStreamingByApMac(ctx context.Context, body *WSGAPPackCaptureApPacketCaptureReq, apMac string) (*WSGAPPackCaptureApPacketCaptureRes, error) {
	var (
		req  *APIRequest
		resp *WSGAPPackCaptureApPacketCaptureRes
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsApPacketCaptureStartStreamingByApMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddApsApPacketCaptureStopByApMac
//
// Use this API to stop AP packet capture or streaming
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsApPacketCaptureStopByApMac(ctx context.Context, apMac string) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsApPacketCaptureStopByApMac, true)
}

// AddApsOperationalBlinkLedByApMac
//
// use this API to make ap blink its led to show its position.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) AddApsOperationalBlinkLedByApMac(ctx context.Context, apMac string) (interface{}, error) {
	var (
		req  *APIRequest
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsOperationalBlinkLedByApMac, true)
}

// AddApsSwitchoverCluster
//
// Use this API command to switchover AP to another cluster
//
// Request Body:
//	 - body *WSGAPSwitchoverAP
func (s *WSGAccessPointOperationalService) AddApsSwitchoverCluster(ctx context.Context, body *WSGAPSwitchoverAP) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, body, "required"); err != nil {
		return resp, err
	} else if err = pkgValidator.StructCtx(ctx, body); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPost, RouteWSGAddApsSwitchoverCluster, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// FindApsApPacketCaptureByApMac
//
// Use this API to get AP packet capture status
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsApPacketCaptureByApMac(ctx context.Context, apMac string) (*WSGAPPackCaptureApPacketCaptureRes, error) {
	var (
		req  *APIRequest
		resp *WSGAPPackCaptureApPacketCaptureRes
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsApPacketCaptureByApMac, true)
}

// FindApsOperationalAlarmsByApMac
//
// Use this API command to retrieve the list of alarms on an AP.
//
// Required Parameters:
// - apMac string
//		- required
//
// Optional Parameters:
// - category string
//		- nullable
// - code float64
//		- nullable
// - endTime string
//		- nullable
// - index string
//		- nullable
// - listSize string
//		- nullable
// - severity string
//		- nullable
// - startTime string
//		- nullable
// - status string
//		- nullable
func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmsByApMac(ctx context.Context, apMac string, optionalParams map[string]interface{}) (*WSGAPAlarmList, error) {
	var (
		req  *APIRequest
		resp *WSGAPAlarmList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalAlarmsByApMac, true)
}

// FindApsOperationalAlarmSummaryByApMac
//
// Use this API command to retrieve the alarm summary of an AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalAlarmSummaryByApMac(ctx context.Context, apMac string) (*WSGAPAlarmSummary, error) {
	var (
		req  *APIRequest
		resp *WSGAPAlarmSummary
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalAlarmSummaryByApMac, true)
}

// FindApsOperationalEventSummaryByApMac
//
// Use this API command to retrieve the event summary of an AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalEventSummaryByApMac(ctx context.Context, apMac string) (*WSGAPEventSummary, error) {
	var (
		req  *APIRequest
		resp *WSGAPEventSummary
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalEventSummaryByApMac, true)
}

// FindApsOperationalNeighborByApMac
//
// Use this API command to retrieve a list of neighbor access points on mesh AP.
//
// Required Parameters:
// - apMac string
//		- required
//
// Optional Parameters:
// - index string
//		- nullable
// - listSize string
//		- nullable
func (s *WSGAccessPointOperationalService) FindApsOperationalNeighborByApMac(ctx context.Context, apMac string, optionalParams map[string]interface{}) (*WSGAPNeighborAPList, error) {
	var (
		req  *APIRequest
		resp *WSGAPNeighborAPList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalNeighborByApMac, true)
}

// FindApsOperationalSummaryByApMac
//
// Use this API command to retrieve the operational information of an AP.
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGAccessPointOperationalService) FindApsOperationalSummaryByApMac(ctx context.Context, apMac string) (*WSGAPOperationalSummary, error) {
	var (
		req  *APIRequest
		resp *WSGAPOperationalSummary
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, apMac, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindApsOperationalSummaryByApMac, true)
}
