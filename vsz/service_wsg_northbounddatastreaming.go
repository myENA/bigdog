package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/northbounddatastreaming"
)

type WSGNorthboundDataStreamingService struct {
	client *Client
}

func NewWSGNorthboundDataStreamingService(client *Client) *WSGNorthboundDataStreamingService {
	s := new(WSGNorthboundDataStreamingService)
	s.client = client
	return s
}

func (ss *WSGService) WSGNorthboundDataStreamingService() *WSGNorthboundDataStreamingService {
	serv := new(WSGNorthboundDataStreamingService)
	serv.client = ss.client
	return serv
}

// AddNorthboundDataStreamingProfile
//
// Use this API command to create northbound Data Streaming Profile
//
// Request Body:
//	 - body *northbounddatastreaming.CreateNorthboundDataStreamingProfile
func (s *WSGNorthboundDataStreamingService) AddNorthboundDataStreamingProfile(ctx context.Context, body *northbounddatastreaming.CreateNorthboundDataStreamingProfile) (*common.CreateResult, error) {
}

// FindNorthboundDataStreamingEventCodes
//
// Use this API command to retrieve NorthboundDataStreamingEventCodes.
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingEventCodes(ctx context.Context) (*northbounddatastreaming.NorthboundDataStreamingEventCodes, error) {
}

// FindNorthboundDataStreamingProfileById
//
// Use this API command to retrieve northbound Data Streaming Profile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileById(ctx context.Context, pId string) (*northbounddatastreaming.NorthboundDataStreamingProfile, error) {
}

// FindNorthboundDataStreamingProfileList
//
// Use this API command to retrieve northbound Data Streaming Profile List
func (s *WSGNorthboundDataStreamingService) FindNorthboundDataStreamingProfileList(ctx context.Context) (*northbounddatastreaming.NorthboundDataStreamingProfileList, error) {
}

// UpdateNorthboundDataStreamingSettings
//
// Use this API command to modify Northbound Data Streaming Settings.
//
// Request Body:
//	 - body *northbounddatastreaming.NorthboundDataStreamingSettings
func (s *WSGNorthboundDataStreamingService) UpdateNorthboundDataStreamingSettings(ctx context.Context, body *northbounddatastreaming.NorthboundDataStreamingSettings) error {
}
