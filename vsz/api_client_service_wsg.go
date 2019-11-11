package vsz

type WSGService struct {
	apiClient *APIClient
}

func NewWSGService(c *APIClient) *WSGService {
	ws := new(WSGService)
	ws.apiClient = c
	return ws
}
