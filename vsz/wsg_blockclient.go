package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGBlockClientService struct {
	apiClient *APIClient
}

func NewWSGBlockClientService(c *APIClient) *WSGBlockClientService {
	s := new(WSGBlockClientService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGBlockClientService() *WSGBlockClientService {
	return NewWSGBlockClientService(ss.apiClient)
}

// AddBlockClient
//
// Create new Block Clients by list.
//
// Request Body:
//	 - body *WSGProfileBulkBlockClient
func (s *WSGBlockClientService) AddBlockClient(ctx context.Context, body *WSGProfileBulkBlockClient) (WSGProfileCreateResultList, error) {
	var (
		req  *APIRequest
		resp WSGProfileCreateResultList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddBlockClient, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// AddBlockClientByApMacByApMac
//
// Create a new Block Client by AP MAC.
//
// Request Body:
//	 - body *WSGProfileBlockClient
//
// Required Parameters:
// - apMac string
//		- required
func (s *WSGBlockClientService) AddBlockClientByApMacByApMac(ctx context.Context, body *WSGProfileBlockClient, apMac string) (*WSGCommonCreateResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonCreateResult
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
	req = NewAPIRequest(http.MethodPost, RouteWSGAddBlockClientByApMacByApMac, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteBlockClient
//
// Delete Block Client List.
//
// Request Body:
//	 - body *WSGCommonBulkDeleteRequest
func (s *WSGBlockClientService) DeleteBlockClient(ctx context.Context, body *WSGCommonBulkDeleteRequest) (*WSGCommonEmptyResult, error) {
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
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteBlockClient, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// DeleteBlockClientById
//
// Delete a Block Client.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockClientService) DeleteBlockClientById(ctx context.Context, id string) (*WSGCommonEmptyResult, error) {
	var (
		req  *APIRequest
		resp *WSGCommonEmptyResult
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodDelete, RouteWSGDeleteBlockClientById, true)
}

// FindBlockClientById
//
// Retrieve a Block Client.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockClientService) FindBlockClientById(ctx context.Context, id string) (*WSGProfileBlockClient, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBlockClient
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindBlockClientById, true)
}

// FindBlockClientByQueryCriteria
//
// Retrieve a list of Block Client.
//
// Request Body:
//	 - body *WSGCommonQueryCriteriaSuperSet
func (s *WSGBlockClientService) FindBlockClientByQueryCriteria(ctx context.Context, body *WSGCommonQueryCriteriaSuperSet) (*WSGProfileBlockClientList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBlockClientList
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
	req = NewAPIRequest(http.MethodPost, RouteWSGFindBlockClientByQueryCriteria, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// FindBlockClientByZoneByZoneId
//
// Retrieve a list of Block Client.
//
// Required Parameters:
// - zoneId string
//		- required
func (s *WSGBlockClientService) FindBlockClientByZoneByZoneId(ctx context.Context, zoneId string) (*WSGProfileBlockClientList, error) {
	var (
		req  *APIRequest
		resp *WSGProfileBlockClientList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, zoneId, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodGet, RouteWSGFindBlockClientByZoneByZoneId, true)
}

// PartialUpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Request Body:
//	 - body *WSGProfileModifyBlockClient
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockClientService) PartialUpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, id string) (*WSGCommonEmptyResult, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPatch, RouteWSGPartialUpdateBlockClientById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}

// UpdateBlockClientById
//
// Modify a specific Block Client basic.
//
// Request Body:
//	 - body *WSGProfileModifyBlockClient
//
// Required Parameters:
// - id string
//		- required
func (s *WSGBlockClientService) UpdateBlockClientById(ctx context.Context, body *WSGProfileModifyBlockClient, id string) (*WSGCommonEmptyResult, error) {
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
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req = NewAPIRequest(http.MethodPut, RouteWSGUpdateBlockClientById, true)
	if err = req.SetBody(body); err != nil {
		return resp, err
	}
}
