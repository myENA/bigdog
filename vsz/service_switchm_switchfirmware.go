package vsz

// API Version: v8_0

type SwitchMSwitchFirmwareService struct {
    c *Client
}

func NewSwitchMSwitchFirmwareService (c *Client) *SwitchMSwitchFirmwareService {
    s := new(SwitchMSwitchFirmwareService)
    s.c = c
    return s
}

