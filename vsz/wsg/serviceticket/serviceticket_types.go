package serviceticket

// API Version: v8_0

type LoginRequest struct {
	// Password
	// Logon password
	Password *string `json:"password,omitempty"`

	// Username
	// Logon user name
	Username *string `json:"username,omitempty"`
}

type LoginResponse struct {
	ControllerVersion *string `json:"controllerVersion,omitempty"`

	// ServiceTicket
	// Logon authentication successful, the server generates a service ticket
	ServiceTicket *string `json:"serviceTicket,omitempty"`
}
