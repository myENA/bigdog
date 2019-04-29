package serviceticket

// API Version: v8_0

type LoginRequest struct {
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`
}

type LoginResponse struct {
	ControllerVersion *string `json:"controllerVersion,omitempty"`
	ServiceTicket     *string `json:"serviceTicket,omitempty"`
}
