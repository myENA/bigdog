package vsz

type WSGService struct {
	client *Client
}

func NewWSGService(c *Client) *WSGService {
	ws := new(WSGService)
	ws.client = c
	return ws
}
