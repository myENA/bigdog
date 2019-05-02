package vsz

// API Version: v8_0

type WSGCALEAService struct {
    c *Client
}

func NewWSGCALEAService (c *Client) *WSGCALEAService {
    s := new(WSGCALEAService)
    s.c = c
    return s
}

