package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/sci"
)

type WSGSCIService struct {
	apiClient *APIClient
}

func NewWSGSCIService(c *APIClient) *WSGSCIService {
	s := new(WSGSCIService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGSCIService() *WSGSCIService {
	serv := new(WSGSCIService)
	serv.apiClient = ss.apiClient
	return serv
}

// AddSciSciEventCode
//
// Use this API command to modify SciAcceptedEventCodes.
//
// Request Body:
//	 - body *sci.ModifyEventCode
func (s *WSGSCIService) AddSciSciEventCode(ctx context.Context, body *sci.ModifyEventCode) error {
}

// AddSciSciProfile
//
// Use this API command to create sciProfile.
//
// Request Body:
//	 - body *sci.CreateSciProfile
func (s *WSGSCIService) AddSciSciProfile(ctx context.Context, body *sci.CreateSciProfile) (*common.CreateResult, error) {
}

// AddSciSciProfileSciPriority
//
// Use this API command to modify sciPriorities.
//
// Request Body:
//	 - body *sci.ModifySciPriorityList
func (s *WSGSCIService) AddSciSciProfileSciPriority(ctx context.Context, body *sci.ModifySciPriorityList) error {
}

// DeleteSciSciProfile
//
// Use this API command to delete sciProfile list.
//
// Request Body:
//	 - body *sci.DeleteSciProfileList
func (s *WSGSCIService) DeleteSciSciProfile(ctx context.Context, body *sci.DeleteSciProfileList) error {
}

// DeleteSciSciProfileById
//
// Use this API command to delete sciProfile.
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGSCIService) DeleteSciSciProfileById(ctx context.Context, pId string) (*common.EmptyResult, error) {
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
//
// Path Parameters:
// - pId string
//		- required
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

// PartialUpdateSciSciProfileById
//
// Use this API command to modify sciProfile.
//
// Request Body:
//	 - body *sci.ModifySciProfile
//
// Path Parameters:
// - pId string
//		- required
func (s *WSGSCIService) PartialUpdateSciSciProfileById(ctx context.Context, body *sci.ModifySciProfile, pId string) error {
}
