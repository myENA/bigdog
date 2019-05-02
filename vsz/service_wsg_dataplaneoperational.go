package vsz

// API Version: v8_0

type WSGDataPlaneOperationalService struct {
    client *Client
}

func NewWSGDataPlaneOperationalService (client *Client) *WSGDataPlaneOperationalService {
    s := new(WSGDataPlaneOperationalService)
    s.client = client
    return s
}

func (ss *WSGService) WSGDataPlaneOperationalService () *WSGDataPlaneOperationalService {
    serv := new(WSGDataPlaneOperationalService)
    serv.client = ss.client
    return serv
}

