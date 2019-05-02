package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/northbounddatastreaming"
)

type WSGNorthboundDataStreamingService struct {
    client *Client
}

func NewWSGNorthboundDataStreamingService (client *Client) *WSGNorthboundDataStreamingService {
    s := new(WSGNorthboundDataStreamingService)
    s.client = client
    return s
}

func (ss *WSGService) WSGNorthboundDataStreamingService () *WSGNorthboundDataStreamingService {
    serv := new(WSGNorthboundDataStreamingService)
    serv.client = ss.client
    return serv
}

func (s *WSGNorthboundDataStreamingService) AddNorthboundDataStreamingProfile (ctx context.Context) (common.CreateResult, error) {
}

func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingEventCodes (ctx context.Context) (northbounddatastreaming.NorthboundDataStreamingEventCodes, error) {
}

func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileById (ctx context.Context, id string) (northbounddatastreaming.NorthboundDataStreamingProfile, error) {
}

func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileList (ctx context.Context) (northbounddatastreaming.NorthboundDataStreamingProfileList, error) {
}

func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingSettings (ctx context.Context) error {
}

