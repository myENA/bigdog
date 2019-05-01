package vsz

type SwitchMService struct {
	client *Client
}

func NewSwitchMService(c *Client) *SwitchMService {
	ss := new(SwitchMService)
	ss.client = c
	return ss
}
