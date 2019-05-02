package vsz

// API Version: v8_0

type WSGLWAPPTOSCGService struct {
    c *Client
}

func NewWSGLWAPPTOSCGService (c *Client) *WSGLWAPPTOSCGService {
    s := new(WSGLWAPPTOSCGService)
    s.c = c
    return s
}

