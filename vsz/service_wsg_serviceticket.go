package vsz

// API Version: v8_0

type WSGServiceTicketService struct {
    client *Client
}

func NewWSGServiceTicketService (client *Client) *WSGServiceTicketService {
    s := new(WSGServiceTicketService)
    s.client = client
    return s
}

func (ss *WSGService) WSGServiceTicketService () *WSGServiceTicketService {
    serv := new(WSGServiceTicketService)
    serv.client = ss.client
    return serv
}

