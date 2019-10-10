package gdpr

// API Version: v8_1

type Ftp struct {
	// FtpHost
	// IP/DN of FTP
	FtpHost *string `json:"ftpHost,omitempty"`

	// FtpPassword
	// Password for FTP login
	FtpPassword *string `json:"ftpPassword,omitempty"`

	// FtpPort
	// Port used by FTP
	FtpPort *int `json:"ftpPort,omitempty" validate:"gte=21,lte=65535"`

	// FtpProtocol
	// Protocol used
	FtpProtocol *string `json:"ftpProtocol,omitempty" validate:"oneof=FTP SFTP"`

	// FtpRemoteDirectory
	// Destination directory used for file upload
	FtpRemoteDirectory *string `json:"ftpRemoteDirectory,omitempty"`

	// FtpUserName
	// Username for FTP login
	FtpUserName *string `json:"ftpUserName,omitempty"`
}

type Report struct {
	// Action
	// Request action
	Action *string `json:"action,omitempty" validate:"oneof=SEARCH DELETE INTERRUPT PROGRESS"`

	// ClientMac
	// Client mac
	ClientMac *string `json:"clientMac,omitempty"`

	Ftp *Ftp `json:"ftp,omitempty"`
}
