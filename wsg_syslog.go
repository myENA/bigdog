package bigdog

// API Version: v9_1

import (
	"encoding/json"
	"io"
)

// WSGSyslogModifySyslogSettings
//
// Definition: syslog_modifySyslogSettings
type WSGSyslogModifySyslogSettings struct {
	// AppLogFacility
	// appLogFacility: Local0-Local7
	AppLogFacility *string `json:"appLogFacility,omitempty"`

	// AppLogSeverity
	// appLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	AppLogSeverity *string `json:"appLogSeverity,omitempty"`

	// AuditLogFacility
	// auditLogFacility: Local0-Local7
	AuditLogFacility *string `json:"auditLogFacility,omitempty"`

	// AuditLogSeverity
	// auditLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	AuditLogSeverity *string `json:"auditLogSeverity,omitempty"`

	// Enabled
	// enable logging to remote syslog server
	Enabled *bool `json:"enabled,omitempty"`

	// EventFilter
	// Event Filter, 0 : All events, 1 : All events except client association/disassociation events, 2 : All events above a severity
	// Constraints:
	//    - min:0
	//    - max:2
	EventFilter *int `json:"eventFilter,omitempty"`

	// EventFilterSeverity
	// Event Filter Severity, This only applies when the Event Filter is set to 2
	// Constraints:
	//    - default:'Debug'
	//    - oneof:[Critical,Major,Minor,Warning,Informational,Debug]
	EventFilterSeverity *string `json:"eventFilterSeverity,omitempty"`

	// EventLogFacility
	// Facility for Event
	EventLogFacility *string `json:"eventLogFacility,omitempty"`

	// ForwardUEEventsMsgFormatType
	// forwardUEEventsMsgFormatType : COMMON, ZD
	ForwardUEEventsMsgFormatType *string `json:"forwardUEEventsMsgFormatType,omitempty"`

	// OtherLogSeverity
	// otherLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	OtherLogSeverity *string `json:"otherLogSeverity,omitempty"`

	PrimaryServer *WSGSyslogPrimaryServer `json:"primaryServer,omitempty"`

	Priority *WSGSyslogPriority `json:"priority,omitempty"`

	SecondaryServer *WSGSyslogSecondaryServer `json:"secondaryServer,omitempty"`
}

func NewWSGSyslogModifySyslogSettings() *WSGSyslogModifySyslogSettings {
	m := new(WSGSyslogModifySyslogSettings)
	return m
}

// WSGSyslogPrimaryServer
//
// Definition: syslog_primaryServer
type WSGSyslogPrimaryServer struct {
	// Host
	// address of the syslog server.
	Host *string `json:"host,omitempty"`

	// Port
	// port number of the syslog server
	Port *int `json:"port,omitempty"`

	// Protocol
	// protocol of the syslog server
	// Constraints:
	//    - oneof:[UDP,TCP]
	Protocol *string `json:"protocol,omitempty"`
}

func NewWSGSyslogPrimaryServer() *WSGSyslogPrimaryServer {
	m := new(WSGSyslogPrimaryServer)
	return m
}

// WSGSyslogPriority
//
// Definition: syslog_priority
type WSGSyslogPriority struct {
	// Critical
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Critical *string `json:"critical,omitempty"`

	// Debug
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Debug *string `json:"debug,omitempty"`

	// Informational
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Informational *string `json:"informational,omitempty"`

	// Major
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Major *string `json:"major,omitempty"`

	// Minor
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Minor *string `json:"minor,omitempty"`

	// Warning
	// Event severity
	// Constraints:
	//    - oneof:[ERROR,WARN,INFO,DEBUG]
	Warning *string `json:"warning,omitempty"`
}

func NewWSGSyslogPriority() *WSGSyslogPriority {
	m := new(WSGSyslogPriority)
	return m
}

// WSGSyslogSecondaryServer
//
// Definition: syslog_secondaryServer
type WSGSyslogSecondaryServer struct {
	// Host
	// address of the syslog server.
	Host *string `json:"host,omitempty"`

	// Port
	// port number of the syslog server
	Port *int `json:"port,omitempty"`

	// Protocol
	// protocol of the syslog server
	// Constraints:
	//    - oneof:[UDP,TCP]
	Protocol *string `json:"protocol,omitempty"`

	// RedundancyMode
	// The redundancy mode of syslog server
	// Constraints:
	//    - oneof:[active_active,primary_backup]
	RedundancyMode *string `json:"redundancyMode,omitempty"`
}

func NewWSGSyslogSecondaryServer() *WSGSyslogSecondaryServer {
	m := new(WSGSyslogSecondaryServer)
	return m
}

// WSGSyslogServerSetting
//
// Definition: syslog_syslogServerSetting
type WSGSyslogServerSetting struct {
	// AppLogFacility
	// appLogFacility: Local0-Local7
	AppLogFacility *string `json:"appLogFacility,omitempty"`

	// AppLogSeverity
	// appLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	AppLogSeverity *string `json:"appLogSeverity,omitempty"`

	// AuditLogFacility
	// auditLogFacility: Local0-Local7
	AuditLogFacility *string `json:"auditLogFacility,omitempty"`

	// AuditLogSeverity
	// auditLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	AuditLogSeverity *string `json:"auditLogSeverity,omitempty"`

	// Enabled
	// enable logging to remote syslog server
	Enabled *bool `json:"enabled,omitempty"`

	// EventFilter
	// Event Filter, 0 : All events, 1 : All events except client association/disassociation events, 2 : All events above a severity
	// Constraints:
	//    - min:0
	//    - max:2
	EventFilter *int `json:"eventFilter,omitempty"`

	// EventFilterSeverity
	// Event Filter Severity, This only applies when the Event Filter is set to 2
	// Constraints:
	//    - oneof:[Critical,Major,Minor,Warning,Informational,Debug]
	EventFilterSeverity *string `json:"eventFilterSeverity,omitempty"`

	// EventLogFacility
	// auditLogFacility: Local0-Local7
	EventLogFacility *string `json:"eventLogFacility,omitempty"`

	// ForwardUEEventsMsgFormatType
	// forwardUEEventsMsgFormatType : COMMON, ZD
	ForwardUEEventsMsgFormatType *string `json:"forwardUEEventsMsgFormatType,omitempty"`

	// OtherLogSeverity
	// otherLogSeverity: Emerg, Alert, Crit, Error, Warning, Notice, Info, Debug
	OtherLogSeverity *string `json:"otherLogSeverity,omitempty"`

	PrimaryServer *WSGSyslogPrimaryServer `json:"primaryServer,omitempty"`

	Priority *WSGSyslogPriority `json:"priority,omitempty"`

	SecondaryServer *WSGSyslogSecondaryServer `json:"secondaryServer,omitempty"`
}

type WSGSyslogServerSettingAPIResponse struct {
	*RawAPIResponse
	Data *WSGSyslogServerSetting
}

func newWSGSyslogServerSettingAPIResponse(meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGSyslogServerSettingAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGSyslogServerSettingAPIResponse) Hydrate() error {
	r.Data = new(WSGSyslogServerSetting)
	return json.NewDecoder(r).Decode(r.Data)
}
func NewWSGSyslogServerSetting() *WSGSyslogServerSetting {
	m := new(WSGSyslogServerSetting)
	return m
}
