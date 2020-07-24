package bigdog

// API Version: 1.0.0

import (
	"encoding/json"
)

// SCIModelsFilter
//
// Definition: filter
type SCIModelsFilter struct {
	// Filter
	// {"type":"and","fields": [{"type":"or","fields": [{"type":"selector","dimension":"radio","value":"2.4"}]}]}}
	// Constraints:
	//    - required
	//    - max:1000000
	Filter *string `json:"filter"`

	Id *float64 `json:"id,omitempty"`

	Name *string `json:"name,omitempty"`

	ReportId *float64 `json:"reportId,omitempty"`

	UrlSegmentName *string `json:"urlSegmentName,omitempty"`

	UserId *float64 `json:"userId,omitempty"`
}

func NewSCIModelsFilter() *SCIModelsFilter {
	m := new(SCIModelsFilter)
	return m
}

// SCIModelsMigration
//
// Definition: Migration
//
// View and run pending migrations.
type SCIModelsMigration struct {
	Id *float64 `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// RunDtTm
	// Constraints:
	//    - required
	RunDtTm *string `json:"runDtTm"`
}

func NewSCIModelsMigration() *SCIModelsMigration {
	m := new(SCIModelsMigration)
	return m
}

// SCIModelsMigrationMap
//
// Definition: MigrationMap
//
// Migration Mappings.
type SCIModelsMigrationMap struct {
	Data interface{} `json:"data,omitempty"`

	// From
	// Constraints:
	//    - required
	From *string `json:"from"`

	Id *float64 `json:"id,omitempty"`

	// To
	// Constraints:
	//    - required
	To *string `json:"to"`

	// Type
	// Constraints:
	//    - required
	Type *string `json:"type"`
}

func NewSCIModelsMigrationMap() *SCIModelsMigrationMap {
	m := new(SCIModelsMigrationMap)
	return m
}

// SCIModelsPciProfile
//
// Definition: pciProfile
type SCIModelsPciProfile struct {
	// Answers
	// list of selected ssids
	// Constraints:
	//    - required
	Answers interface{} `json:"answers"`

	Id *float64 `json:"id,omitempty"`

	// Name
	// Name of the profile
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// Ssids
	// list of selected ssids
	Ssids []interface{} `json:"ssids,omitempty"`
}

func NewSCIModelsPciProfile() *SCIModelsPciProfile {
	m := new(SCIModelsPciProfile)
	return m
}

// SCIModelsPciReport
//
// Definition: pciReport
type SCIModelsPciReport struct {
	// Date
	// Date of report generation
	// Constraints:
	//    - required
	//    - default:'$now'
	Date *string `json:"date"`

	Id *float64 `json:"id,omitempty"`

	ProfileId *float64 `json:"profileId,omitempty"`

	// Statuses
	// Status of the controllers in this report
	// Constraints:
	//    - required
	Statuses []interface{} `json:"statuses"`
}

func NewSCIModelsPciReport() *SCIModelsPciReport {
	m := new(SCIModelsPciReport)
	return m
}

// SCIModelsReport
//
// Definition: report
type SCIModelsReport struct {
	Component *string `json:"component,omitempty"`

	DatasourcesUsed []interface{} `json:"datasourcesUsed,omitempty"`

	ExcludedFilters []string `json:"excludedFilters,omitempty"`

	FilterDataSource *string `json:"filterDataSource,omitempty"`

	Headers []interface{} `json:"headers,omitempty"`

	Id *float64 `json:"id,omitempty"`

	Layout []*SCIModelsReportLayout `json:"layout,omitempty"`

	RouteParameters interface{} `json:"routeParameters,omitempty"`

	// Title
	// Constraints:
	//    - required
	Title *string `json:"title"`

	// UrlSegmentName
	// Constraints:
	//    - required
	UrlSegmentName *string `json:"urlSegmentName"`
}

func NewSCIModelsReport() *SCIModelsReport {
	m := new(SCIModelsReport)
	return m
}

// SCIModelsReportLayout
//
// Definition: reportLayout
//
// Layout descriptor of a report
type SCIModelsReportLayout struct {
	DesiredWidth *string `json:"desiredWidth,omitempty"`

	Layout []*SCIModelsReportLayout `json:"layout,omitempty"`

	Section *int `json:"section,omitempty"`

	XAdditionalProperties map[string]interface{} `json:"-"`
}

func (t *SCIModelsReportLayout) UnmarshalJSON(b []byte) error {
	type _SCIModelsReportLayout SCIModelsReportLayout
	tmpType := new(_SCIModelsReportLayout)
	if err := json.Unmarshal(b, tmpType); err != nil {
		return err
	}
	tmpType.XAdditionalProperties = make(map[string]interface{})
	if err := json.Unmarshal(b, &tmpType.XAdditionalProperties); err != nil {
		return err
	}
	delete(tmpType.XAdditionalProperties, "desiredWidth")
	delete(tmpType.XAdditionalProperties, "layout")
	delete(tmpType.XAdditionalProperties, "section")
	*t = SCIModelsReportLayout(*tmpType)
	return nil
}

func (t *SCIModelsReportLayout) MarshalJSON() ([]byte, error) {
	if t == nil {
		return nil, nil
	}
	var tmp map[string]interface{}
	if t.XAdditionalProperties == nil {
		tmp = make(map[string]interface{})
	} else {
		tmp = t.XAdditionalProperties
	}
	if t.DesiredWidth != nil {
		tmp["desiredWidth"] = t.DesiredWidth
	}
	if t.Layout != nil {
		tmp["layout"] = t.Layout
	}
	if t.Section != nil {
		tmp["section"] = t.Section
	}
	return json.Marshal(tmp)
}

func NewSCIModelsReportLayout() *SCIModelsReportLayout {
	m := new(SCIModelsReportLayout)
	return m
}

// SCIModelsResourceGroup
//
// Definition: resourceGroup
type SCIModelsResourceGroup struct {
	Description *string `json:"description,omitempty"`

	// Filter
	// {"type":"and","fields": [{"type":"or","fields": [{"type":"selector","dimension":"radio","value":"2.4"}]}]}}
	// Constraints:
	//    - required
	//    - max:1000000
	Filter *string `json:"filter"`

	Id *float64 `json:"id,omitempty"`

	// IsDefault
	// Constraints:
	//    - required
	//    - default:false
	IsDefault *bool `json:"isDefault"`

	LastUpdated *string `json:"lastUpdated,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *string `json:"name"`
}

func NewSCIModelsResourceGroup() *SCIModelsResourceGroup {
	m := new(SCIModelsResourceGroup)
	return m
}

// SCIModelsSchedule
//
// Definition: schedule
type SCIModelsSchedule struct {
	// Day
	// 1 for 1st or monday
	Day *float64 `json:"day,omitempty"`

	// Enabled
	// 0 for disabled, 1 for enabled
	// Constraints:
	//    - required
	Enabled *bool `json:"enabled"`

	FilterId *float64 `json:"filterId,omitempty"`

	// Format
	// pdf, csv
	// Constraints:
	//    - required
	Format *string `json:"format"`

	// Frequency
	// daily, weekly, monthly, quarterly, yearly
	// Constraints:
	//    - required
	Frequency *string `json:"frequency"`

	// Hour
	// 0 for midnight, 13 for 1pm, UTC
	Hour *float64 `json:"hour,omitempty"`

	Id *float64 `json:"id,omitempty"`

	// Name
	// Constraints:
	//    - required
	Name *string `json:"name"`

	// Recipients
	// one@domain.com
	// two@domain.com
	// Constraints:
	//    - required
	Recipients []interface{} `json:"recipients"`

	ReportId *float64 `json:"reportId,omitempty"`

	// Timezone
	// Timezone
	// Constraints:
	//    - required
	Timezone *string `json:"timezone"`

	UserId *float64 `json:"userId,omitempty"`
}

func NewSCIModelsSchedule() *SCIModelsSchedule {
	m := new(SCIModelsSchedule)
	return m
}

// SCIModelsSection
//
// Definition: section
type SCIModelsSection struct {
	Component *string `json:"component,omitempty"`

	DefaultParameters interface{} `json:"defaultParameters,omitempty"`

	Id *float64 `json:"id,omitempty"`

	Layout interface{} `json:"layout,omitempty"`

	// QueryName
	// Constraints:
	//    - required
	QueryName *string `json:"queryName"`

	SystemOwnerOnly *bool `json:"systemOwnerOnly,omitempty"`

	// Title
	// Constraints:
	//    - required
	Title *string `json:"title"`

	Url *string `json:"url,omitempty"`
}

func NewSCIModelsSection() *SCIModelsSection {
	m := new(SCIModelsSection)
	return m
}

// SCIModelsSetting
//
// Definition: setting
type SCIModelsSetting struct {
	// Key
	// Constraints:
	//    - required
	Key *string `json:"key"`

	// Values
	// Constraints:
	//    - required
	Values interface{} `json:"values"`
}

func NewSCIModelsSetting() *SCIModelsSetting {
	m := new(SCIModelsSetting)
	return m
}

// SCIModelsSystem
//
// Definition: system
type SCIModelsSystem struct {
	BackupLocation *string `json:"backupLocation,omitempty"`

	// Format
	// Constraints:
	//    - required
	Format *string `json:"format"`

	// Id
	// Constraints:
	//    - required
	Id *string `json:"id"`

	// LastContact
	// last contacted/aggregated time
	LastContact *string `json:"lastContact,omitempty"`

	// Location
	// Constraints:
	//    - required
	Location *string `json:"location"`

	PciData interface{} `json:"pciData,omitempty"`

	// Type
	// Constraints:
	//    - required
	Type *string `json:"type"`

	// User
	// Constraints:
	//    - required
	User *string `json:"user"`

	Version *string `json:"version,omitempty"`
}

func NewSCIModelsSystem() *SCIModelsSystem {
	m := new(SCIModelsSystem)
	return m
}

// SCIModelsUser
//
// Definition: user
type SCIModelsUser struct {
	// Email
	// Constraints:
	//    - required
	Email *string `json:"email"`

	// FirstName
	// Constraints:
	//    - required
	FirstName *string `json:"firstName"`

	Id *float64 `json:"id,omitempty"`

	IsExternal *bool `json:"isExternal,omitempty"`

	LastLogin *string `json:"lastLogin,omitempty"`

	// LastName
	// Constraints:
	//    - required
	LastName *string `json:"lastName"`

	Username *string `json:"username,omitempty"`
}

func NewSCIModelsUser() *SCIModelsUser {
	m := new(SCIModelsUser)
	return m
}
