package vsz

// API Version: v8_0

type WSGDataPlaneOperationalService struct {
    c *Client
}

func NewWSGDataPlaneOperationalService (c *Client) *WSGDataPlaneOperationalService {
    s := new(WSGDataPlaneOperationalService)
    s.c = c
    return s
}

