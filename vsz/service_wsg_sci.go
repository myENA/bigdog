package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/sci"
)

type WSGSCIService struct {
	client *Client
}

func NewWSGSCIService(client *Client) *WSGSCIService {
	s := new(WSGSCIService)
	s.client = client
	return s
}

func (ss *WSGService) WSGSCIService() *WSGSCIService {
	serv := new(WSGSCIService)
	serv.client = ss.client
	return serv
}

// AddSciSciProfileSciPriority
//
// Use this API command to modify sciPriorities.
//
// Request Body:
//	 - body *sci.ModifySciPriorityList
func (s *WSGSCIService) AddSciSciProfileSciPriority(ctx context.Context, body *sci.ModifySciPriorityList) error {
}

// FindSciSciEventCode
//
// Use this API command to retrieve SciAcceptedEventCodes.
func (s *WSGSCIService) FindSciSciEventCode(ctx context.Context) (*sci.SciEventCode, error) {
}

// FindSciSciProfile
//
// Use this API command to retrieve sciProfile list.
func (s *WSGSCIService) FindSciSciProfile(ctx context.Context) (*sci.SciProfileList, error) {
}

// FindSciSciProfileById
//
// Use this API command to retrieve sciProfile.
func (s *WSGSCIService) FindSciSciProfileById(ctx context.Context, pId string) (*sci.SciProfile, error) {
}

// PartialUpdateSciSciEnabled
//
// Use this API command to modify SCI settings is enabled or not.
//
// Request Body:
//	 - body *sci.ModifySciEnabled
func (s *WSGSCIService) PartialUpdateSciSciEnabled(ctx context.Context, body *sci.ModifySciEnabled) error {
}
