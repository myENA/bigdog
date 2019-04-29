package syslog

// API Version: v8_0

type ModifySyslogSettings struct {
	AppLogFacility               *string          `json:"appLogFacility,omitempty"`
	AppLogSeverity               *string          `json:"appLogSeverity,omitempty"`
	AuditLogFacility             *string          `json:"auditLogFacility,omitempty"`
	AuditLogSeverity             *string          `json:"auditLogSeverity,omitempty"`
	Enabled                      *bool            `json:"enabled,omitempty"`
	EventFilter                  *int             `json:"eventFilter,omitempty"`
	EventFilterSeverity          *string          `json:"eventFilterSeverity,omitempty"`
	EventLogFacility             *string          `json:"eventLogFacility,omitempty"`
	ForwardUEEventsMsgFormatType *string          `json:"forwardUEEventsMsgFormatType,omitempty"`
	OtherLogSeverity             *string          `json:"otherLogSeverity,omitempty"`
	PrimaryServer                *PrimaryServer   `json:"primaryServer,omitempty"`
	Priority                     *Priority        `json:"priority,omitempty"`
	SecondaryServer              *SecondaryServer `json:"secondaryServer,omitempty"`
}

type PrimaryServer struct {
	Host     *string `json:"host,omitempty"`
	Port     *int    `json:"port,omitempty"`
	Protocol *string `json:"protocol,omitempty"`
}

type Priority struct {
	Critical      *string `json:"critical,omitempty"`
	Debug         *string `json:"debug,omitempty"`
	Informational *string `json:"informational,omitempty"`
	Major         *string `json:"major,omitempty"`
	Minor         *string `json:"minor,omitempty"`
	Warning       *string `json:"warning,omitempty"`
}

type SecondaryServer struct {
	Host           *string `json:"host,omitempty"`
	Port           *int    `json:"port,omitempty"`
	Protocol       *string `json:"protocol,omitempty"`
	RedundancyMode *string `json:"redundancyMode,omitempty"`
}

type SyslogServerSetting struct {
	AppLogFacility               *string          `json:"appLogFacility,omitempty"`
	AppLogSeverity               *string          `json:"appLogSeverity,omitempty"`
	AuditLogFacility             *string          `json:"auditLogFacility,omitempty"`
	AuditLogSeverity             *string          `json:"auditLogSeverity,omitempty"`
	Enabled                      *bool            `json:"enabled,omitempty"`
	EventFilter                  *int             `json:"eventFilter,omitempty"`
	EventFilterSeverity          *string          `json:"eventFilterSeverity,omitempty"`
	EventLogFacility             *string          `json:"eventLogFacility,omitempty"`
	ForwardUEEventsMsgFormatType *string          `json:"forwardUEEventsMsgFormatType,omitempty"`
	OtherLogSeverity             *string          `json:"otherLogSeverity,omitempty"`
	PrimaryServer                *PrimaryServer   `json:"primaryServer,omitempty"`
	Priority                     *Priority        `json:"priority,omitempty"`
	SecondaryServer              *SecondaryServer `json:"secondaryServer,omitempty"`
}
