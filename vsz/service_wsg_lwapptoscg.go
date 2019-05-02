package vsz

// API Version: v8_0

type WSGLWAPPTOSCGService struct {
    client *Client
}

func NewWSGLWAPPTOSCGService (client *Client) *WSGLWAPPTOSCGService {
    s := new(WSGLWAPPTOSCGService)
    s.client = client
    return s
}

func (ss *WSGService) WSGLWAPPTOSCGService () *WSGLWAPPTOSCGService {
    serv := new(WSGLWAPPTOSCGService)
    serv.client = ss.client
    return serv
}

