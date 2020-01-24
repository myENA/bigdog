package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGMarkRogueService struct {
	apiClient *APIClient
}

func NewWSGMarkRogueService(c *APIClient) *WSGMarkRogueService {
	s := new(WSGMarkRogueService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGMarkRogueService() *WSGMarkRogueService {
	return NewWSGMarkRogueService(ss.apiClient)
}

// AddRogueMarkIgnore
//
// Mark a rogue AP as ignore.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkIgnore(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkIgnore, true)
}

// AddRogueMarkKnown
//
// Mark a rogue AP as know.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkKnown(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkKnown, true)
}

// AddRogueMarkMalicious
//
// Mark a rogue AP as malicious.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkMalicious(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkMalicious, true)
}

// AddRogueMarkRogue
//
// Mark a rogue AP as rogue.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueMarkRogue(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRogueMarkRogue, true)
}

// AddRogueUnMark
//
// Use this API command to remove the manual admin classification marking.
//
// Request Body:
//	 - body *WSGAPModifyRogueType
func (s *WSGMarkRogueService) AddRogueUnMark(ctx context.Context, body *WSGAPModifyRogueType) (*WSGCommonEmptyResult, error) {
	var (
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddRogueUnMark, true)
}

// FindRogueMarkKnown
//
// Get Known Rogue AP list.
func (s *WSGMarkRogueService) FindRogueMarkKnown(ctx context.Context) (*WSGAPModifyRogueType, error) {
	var (
		resp *WSGAPModifyRogueType
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindRogueMarkKnown, true)
}
