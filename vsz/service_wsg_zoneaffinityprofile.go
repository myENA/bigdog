package vsz

// API Version: v8_0

import (
	"github.com/myENA/ruckus-client/vsz/types/wsg/profile"
)

type WSGZoneAffinityProfileService struct {
    client *Client
}

func NewWSGZoneAffinityProfileService (client *Client) *WSGZoneAffinityProfileService {
    s := new(WSGZoneAffinityProfileService)
    s.client = client
    return s
}

func (ss *WSGService) WSGZoneAffinityProfileService () *WSGZoneAffinityProfileService {
    serv := new(WSGZoneAffinityProfileService)
    serv.client = ss.client
    return serv
}

