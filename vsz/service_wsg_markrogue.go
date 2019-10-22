package vsz

// API Version: v8_1

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/ap"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
)

type WSGMarkRogueService struct {
    client *Client
}

func NewWSGMarkRogueService (client *Client) *WSGMarkRogueService {
    s := new(WSGMarkRogueService)
    s.client = client
    return s
}

func (ss *WSGService) WSGMarkRogueService () *WSGMarkRogueService {
    serv := new(WSGMarkRogueService)
    serv.client = ss.client
    return serv
}

func (s *WSGMarkRogueService) AddRogueMarkIgnore (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGMarkRogueService) AddRogueMarkMalicious (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGMarkRogueService) AddRogueMarkRogue (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGMarkRogueService) AddRogueUnMark (ctx context.Context) (*common.EmptyResult, error) {
}

func (s *WSGMarkRogueService) FindRogueMarkKnown (ctx context.Context) (*ap.ModifyRogueType, error) {
}

