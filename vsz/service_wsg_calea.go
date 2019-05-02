package vsz

// API Version: v8_0

type WSGCALEAService struct {
    client *Client
}

func NewWSGCALEAService (client *Client) *WSGCALEAService {
    s := new(WSGCALEAService)
    s.client = client
    return s
}

func (ss *WSGService) WSGCALEAService () *WSGCALEAService {
    serv := new(WSGCALEAService)
    serv.client = ss.client
    return serv
}

