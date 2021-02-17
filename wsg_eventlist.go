package bigdog

// API Version: v9_1

import (
	"errors"
	"io"
)

// WSGEventListEventQueryResultList
//
// Definition: eventList_eventQueryResultList
type WSGEventListEventQueryResultList struct {
	FirstIndex *int `json:"firstIndex,omitempty"`

	HasMore *bool `json:"hasMore,omitempty"`

	List []*WSGEventListSingleEvent `json:"list,omitempty"`

	RawDataTotalCount *int `json:"rawDataTotalCount,omitempty"`

	TotalCount *int `json:"totalCount,omitempty"`
}

type WSGEventListEventQueryResultListAPIResponse struct {
	*RawAPIResponse
	Data *WSGEventListEventQueryResultList
}

func newWSGEventListEventQueryResultListAPIResponse(src APISource, meta APIResponseMeta, body io.ReadCloser) APIResponse {
	r := new(WSGEventListEventQueryResultListAPIResponse)
	r.RawAPIResponse = newRawAPIResponse(src, meta, body).(*RawAPIResponse)
	return r
}

func (r *WSGEventListEventQueryResultListAPIResponse) Hydrate() (interface{}, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	if r.err != nil {
		if errors.Is(r.err, ErrResponseHydrated) {
			return r.Data, nil
		}
		return nil, r.err
	}
	data := new(WSGEventListEventQueryResultList)
	if err := r.doHydrate(data); err != nil {
		return nil, err
	}
	r.Data = data
	return r.Data, nil
}
func NewWSGEventListEventQueryResultList() *WSGEventListEventQueryResultList {
	m := new(WSGEventListEventQueryResultList)
	return m
}

// WSGEventListSingleEvent
//
// Definition: eventList_singleEvent
type WSGEventListSingleEvent struct {
	Activity *string `json:"activity,omitempty"`

	Category *string `json:"category,omitempty"`

	EventCode *int `json:"eventCode,omitempty"`

	EventType *string `json:"eventType,omitempty"`

	Id *string `json:"id,omitempty"`

	InsertionTime *int `json:"insertionTime,omitempty"`

	Severity *string `json:"severity,omitempty"`
}

func NewWSGEventListSingleEvent() *WSGEventListSingleEvent {
	m := new(WSGEventListSingleEvent)
	return m
}
