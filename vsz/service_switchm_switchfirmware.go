package vsz

// API Version: v8_0

type SwitchMSwitchFirmwareService struct {
    client *Client
}

func NewSwitchMSwitchFirmwareService (client *Client) *SwitchMSwitchFirmwareService {
    s := new(SwitchMSwitchFirmwareService)
    s.client = client
    return s
}

func (ss *SwitchMService) SwitchMSwitchFirmwareService () *SwitchMSwitchFirmwareService {
    serv := new(SwitchMSwitchFirmwareService)
    serv.client = ss.client
    return serv
}

