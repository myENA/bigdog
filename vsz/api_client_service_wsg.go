package vsz

type WSGService struct {
	apiClient *APIClient
}

func NewWSGService(c *APIClient) *WSGService {
	s := new(WSGService)
	s.apiClient = c
	return s
}
