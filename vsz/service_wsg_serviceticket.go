package vsz

// API Version: v8_0

type WSGServiceTicketService struct {
    c *Client
}

func NewWSGServiceTicketService (c *Client) *WSGServiceTicketService {
    s := new(WSGServiceTicketService)
    s.c = c
    return s
}

