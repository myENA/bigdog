package vsz

// API Version: v8_0

import (
	"context"
	"github.com/myENA/ruckus-client/vsz/types/wsg/common"
	"github.com/myENA/ruckus-client/vsz/types/wsg/urlfiltering"
)

type WSGURLFilteringPolicyService struct {
    client *Client
}

func NewWSGURLFilteringPolicyService (client *Client) *WSGURLFilteringPolicyService {
    s := new(WSGURLFilteringPolicyService)
    s.client = client
    return s
}

func (ss *WSGService) WSGURLFilteringPolicyService () *WSGURLFilteringPolicyService {
    serv := new(WSGURLFilteringPolicyService)
    serv.client = ss.client
    return serv
}

func (s *WSGURLFilteringPolicyService) FindUrlFilteringBlockCategories (ctx context.Context) (*urlfiltering.UrlFilteringBlockCategoriesList, error) {
}

func (s *WSGURLFilteringPolicyService) FindUrlFilteringByQueryCriteria (ctx context.Context) (*urlfiltering.UrlFilteringPolicyList, error) {
}

func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicy (ctx context.Context) (*urlfiltering.UrlFilteringPolicyList, error) {
}

func (s *WSGURLFilteringPolicyService) FindUrlFilteringUrlFilteringPolicyById (ctx context.Context, id string) (*urlfiltering.UrlFilteringPolicy, error) {
}

