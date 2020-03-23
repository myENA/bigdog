package vsz

type SwitchMService struct {
	apiClient *APIClient
}

func NewSwitchMService(c *APIClient) *SwitchMService {
	s := new(SwitchMService)
	s.apiClient = c
	return s
}
