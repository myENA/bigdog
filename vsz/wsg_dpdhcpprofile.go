package vsz

// API Version: v8_1

import (
	"context"
	"net/http"
)

type WSGDPDHCPProfileService struct {
	apiClient *APIClient
}

func NewWSGDPDHCPProfileService(c *APIClient) *WSGDPDHCPProfileService {
	s := new(WSGDPDHCPProfileService)
	s.apiClient = c
	return s
}

func (ss *WSGService) WSGDPDHCPProfileService() *WSGDPDHCPProfileService {
	return NewWSGDPDHCPProfileService(ss.apiClient)
}

// AddDpDhcpProfiles
//
// Use this API command to create basic DP DHCP profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileBasicBO
func (s *WSGDPDHCPProfileService) AddDpDhcpProfiles(ctx context.Context, body *WSGDPProfileDpDhcpProfileBasicBO) (*WSGDPProfileDpDhcpProfileBasicBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileBasicBO
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfiles, true)
}

// AddDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to create DP DHCP profile - host.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileHostBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *WSGDPProfileDpDhcpProfileHostBO, id string) (*WSGDPProfileDpDhcpProfileHostBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileHostBO
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileHostsById, true)
}

// AddDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to create DP DHCP profile - option43 space.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *WSGDPProfileDpDhcpProfileOptionSpaceBO, id string) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
}

// AddDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to create DP DHCP profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfilePoolBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) AddDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *WSGDPProfileDpDhcpProfilePoolBO, id string) (*WSGDPProfileDpDhcpProfilePoolBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfilePoolBO
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
	req := NewAPIRequest(http.MethodPost, RouteWSGAddDpDhcpProfilesDpDhcpProfilePoolsById, true)
}

// DeleteDpDhcpProfiles
//
// Use this API command to delete DP DHCP profiles.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfiles(ctx context.Context, body *WSGDPProfileBulkDelete) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfiles, true)
}

// DeleteDpDhcpProfilesById
//
// Use this API command to delete DP DHCP profile.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesById(ctx context.Context, id string) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesById, true)
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to delete DP DHCP profile - host.
//
// Required Parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, hostId string, id string) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, hostId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
}

// DeleteDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to delete DP DHCP profile - hosts.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileHostsById, true)
}

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to delete DP DHCP profile - option43 spaces.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, body *WSGDPProfileBulkDelete, id string) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
}

// DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to delete DP DHCP profile - option43 space.
//
// Required Parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, id string, spaceId string) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, spaceId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
}

// DeleteDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to delete DP DHCP profile - pools.
//
// Request Body:
//	 - body *WSGDPProfileBulkDelete
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, body *WSGDPProfileBulkDelete, id string) (interface{}, error) {
	var (
		resp interface{}
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
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsById, true)
}

// DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to delete DP DHCP profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) DeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, id string, poolId string) (interface{}, error) {
	var (
		resp interface{}
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, poolId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodDelete, RouteWSGDeleteDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
}

// FindDpDhcpProfiles
//
// Use this API command to retrieve DP profile - basic list.
func (s *WSGDPDHCPProfileService) FindDpDhcpProfiles(ctx context.Context) (*WSGDPProfileDpDhcpProfileBasicBOList, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileBasicBOList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfiles, true)
}

// FindDpDhcpProfilesById
//
// Use this API command to retrieve DP profile - basic.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesById(ctx context.Context, id string) (*WSGDPProfileDpDhcpProfileBasicBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileBasicBO
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesById, true)
}

// FindDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to retrieve DP DHCP profile - host.
//
// Required Parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, hostId string, id string) (*WSGDPProfileDpDhcpProfileHostBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileHostBO
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, hostId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
}

// FindDpDhcpProfilesDpDhcpProfileHostsById
//
// Use this API command to retrieve DP DHCP profile - host list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileHostsById(ctx context.Context, id string) (*WSGDPProfileDpDhcpProfileHostBOList, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileHostBOList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileHostsById, true)
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesById
//
// Use this API command to retrieve DP DHCP profile - option43 space list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesById(ctx context.Context, id string) (*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBOList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesById, true)
}

// FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to retrieve DP DHCP profile - option43 space.
//
// Required Parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, id string, spaceId string) (*WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileOptionSpaceApplyToBO
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, spaceId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
}

// FindDpDhcpProfilesDpDhcpProfilePoolsById
//
// Use this API command to retrieve DP DHCP profile - pool list.
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsById(ctx context.Context, id string) (*WSGDPProfileDpDhcpProfilePoolBOList, error) {
	var (
		resp *WSGDPProfileDpDhcpProfilePoolBOList
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsById, true)
}

// FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to retrieve DP DHCP profile - pool.
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) FindDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, id string, poolId string) (*WSGDPProfileDpDhcpProfilePoolBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfilePoolBO
		err  error
	)
	if err = ctx.Err(); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, id, "required"); err != nil {
		return resp, err
	}
	if err = pkgValidator.VarCtx(ctx, poolId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodGet, RouteWSGFindDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
}

// UpdateDpDhcpProfilesById
//
// Use this API command to modify DP DHCP profile - basic.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileBasicBO
//
// Required Parameters:
// - id string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesById(ctx context.Context, body *WSGDPProfileDpDhcpProfileBasicBO, id string) (*WSGDPProfileDpDhcpProfileBasicBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileBasicBO
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
	req := NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesById, true)
}

// UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId
//
// Use this API command to modify DP DHCP profile - host.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileHostBO
//
// Required Parameters:
// - hostId string
//		- required
// - id string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileHostsByHostId(ctx context.Context, body *WSGDPProfileDpDhcpProfileHostBO, hostId string, id string) (*WSGDPProfileDpDhcpProfileHostBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfileHostBO
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
	if err = pkgValidator.VarCtx(ctx, hostId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileHostsByHostId, true)
}

// UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId
//
// Use this API command to update DP DHCP profile - option43 space.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfileOptionSpaceBO
//
// Required Parameters:
// - id string
//		- required
// - spaceId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId(ctx context.Context, body *WSGDPProfileDpDhcpProfileOptionSpaceBO, id string, spaceId string) (interface{}, error) {
	var (
		resp interface{}
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
	if err = pkgValidator.VarCtx(ctx, spaceId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfileOptionSpacesBySpaceId, true)
}

// UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId
//
// Use this API command to modify DP DHCP profile - pool.
//
// Request Body:
//	 - body *WSGDPProfileDpDhcpProfilePoolBO
//
// Required Parameters:
// - id string
//		- required
// - poolId string
//		- required
func (s *WSGDPDHCPProfileService) UpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId(ctx context.Context, body *WSGDPProfileDpDhcpProfilePoolBO, id string, poolId string) (*WSGDPProfileDpDhcpProfilePoolBO, error) {
	var (
		resp *WSGDPProfileDpDhcpProfilePoolBO
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
	if err = pkgValidator.VarCtx(ctx, poolId, "required"); err != nil {
		return resp, err
	}
	req := NewAPIRequest(http.MethodPut, RouteWSGUpdateDpDhcpProfilesDpDhcpProfilePoolsByPoolId, true)
}