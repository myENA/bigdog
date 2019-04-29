package wlanscheduler

// API Version: v8_0

type CreateWLANScheduler struct {
	Description *string  `json:"description,omitempty"`
	Fri         []string `json:"fri,omitempty"`
	Mon         []string `json:"mon,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Sat         []string `json:"sat,omitempty"`
	Sun         []string `json:"sun,omitempty"`
	Thu         []string `json:"thu,omitempty"`
	Tue         []string `json:"tue,omitempty"`
	Wed         []string `json:"wed,omitempty"`
}

type ModifyWLANScheduler struct {
	Description *string  `json:"description,omitempty"`
	Fri         []string `json:"fri,omitempty"`
	Mon         []string `json:"mon,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Sat         []string `json:"sat,omitempty"`
	Sun         []string `json:"sun,omitempty"`
	Thu         []string `json:"thu,omitempty"`
	Tue         []string `json:"tue,omitempty"`
	Wed         []string `json:"wed,omitempty"`
}

type WLANSchedule struct {
	Description *string  `json:"description,omitempty"`
	Fri         []string `json:"fri,omitempty"`
	ID          *string  `json:"id,omitempty"`
	Mon         []string `json:"mon,omitempty"`
	Name        *string  `json:"name,omitempty"`
	Sat         []string `json:"sat,omitempty"`
	Sun         []string `json:"sun,omitempty"`
	Thu         []string `json:"thu,omitempty"`
	Tue         []string `json:"tue,omitempty"`
	Wed         []string `json:"wed,omitempty"`
	ZoneID      *string  `json:"zoneId,omitempty"`
}

type WLANScheduleList struct {
	FirstIndex *int    `json:"firstIndex,omitempty"`
	HasMore    *bool   `json:"hasMore,omitempty"`
	List       []*List `json:"list,omitempty"`
	TotalCount *int    `json:"totalCount,omitempty"`
}
