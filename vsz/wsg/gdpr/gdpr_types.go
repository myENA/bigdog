package gdpr

// API Version: v8_0

type Ftp struct {
	FtpHost            *string `json:"ftpHost,omitempty"`
	FtpPassword        *string `json:"ftpPassword,omitempty"`
	FtpPort            *int    `json:"ftpPort,omitempty"`
	FtpProtocol        *string `json:"ftpProtocol,omitempty"`
	FtpRemoteDirectory *string `json:"ftpRemoteDirectory,omitempty"`
	FtpUserName        *string `json:"ftpUserName,omitempty"`
}

type Report struct {
	Action    *string `json:"action,omitempty"`
	ClientMac *string `json:"clientMac,omitempty"`
	Ftp       *Ftp    `json:"ftp,omitempty"`
}
