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
	// Constraints:
	//    - nullable
	//    - min:21
	//    - max:65535
	FtpPort *int `json:"ftpPort,omitempty" validate:"omitempty,gte=21,lte=65535"`

	// FtpProtocol
	// Protocol used
	// Constraints:
	//    - nullable
	//    - oneof:[FTP,SFTP]
	FtpProtocol *string `json:"ftpProtocol,omitempty" validate:"omitempty,oneof=FTP SFTP"`

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
	// Constraints:
	//    - nullable
	//    - oneof:[SEARCH,DELETE,INTERRUPT,PROGRESS]
	Action *string `json:"action,omitempty" validate:"omitempty,oneof=SEARCH DELETE INTERRUPT PROGRESS"`

	// ClientMac
	// Client mac
	ClientMac *string `json:"clientMac,omitempty"`

	Ftp *Ftp `json:"ftp,omitempty"`
}
