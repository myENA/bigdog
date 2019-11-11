package vsz

type SwitchMService struct {
	apiClient *APIClient
}

func NewSwitchMService(c *APIClient) *SwitchMService {
	ss := new(SwitchMService)
	ss.apiClient = c
	return ss
}
