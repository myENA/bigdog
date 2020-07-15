package bigdog

// API Version: 1.0.0

import (
	"context"
	"net/http"
)

type SCIFacetService struct {
	apiClient *SCIClient
}

func NewSCIFacetService(c *SCIClient) *SCIFacetService {
	s := new(SCIFacetService)
	s.apiClient = c
	return s
}

func (ss *SCIService) SCIFacetService() *SCIFacetService {
	return NewSCIFacetService(ss.apiClient)
}

// SCIFacetGetFacet200ResponseType
//
// Definition: facet.getFacet200ResponseType
type SCIFacetGetFacet200ResponseType []interface{}

func MakeSCIFacetGetFacet200ResponseType() SCIFacetGetFacet200ResponseType {
	m := make(SCIFacetGetFacet200ResponseType, 0)
	return m
}

// FacetGetFacet
//
// Operation ID: facet.getFacet
//
// For the <b><code>filter</code></b> field below, an example would be
// <pre>
//   <code class="json">
//     { "type": "and", "fields": [{ "type": "selector", "dimension": "system", "value": "sys1" }]}
//   </code>
// </pre>
//
// Request Body:
//	 - body *SCICommonQueryBody
//
// Required Parameters:
// - name string
//		- required
//		- oneof:[system,switchHierarchy,apmac,ssid,switches]
func (s *SCIFacetService) FacetGetFacet(ctx context.Context, body *SCICommonQueryBody, name string, mutators ...RequestMutator) (SCIFacetGetFacet200ResponseType, *APIResponseMeta, error) {
	var (
		req      *APIRequest
		rm       *APIResponseMeta
		resp     SCIFacetGetFacet200ResponseType
		httpResp *http.Response
		err      error
	)
	if err = ctx.Err(); err != nil {
		return resp, rm, err
	}
	req = NewAPIRequest(http.MethodPost, RouteSCIFacetGetFacet, true)
	req.SetHeader(headerKeyContentType, headerValueApplicationJSON)
	req.SetHeader(headerKeyAccept, headerValueApplicationJSON)
	if err = req.SetBody(body); err != nil {
		return resp, rm, err
	}
	req.SetPathParameter("name", name)
	httpResp, err = s.apiClient.Do(ctx, req, mutators...)
	resp = MakeSCIFacetGetFacet200ResponseType()
	rm, err = handleResponse(req, http.StatusOK, httpResp, &resp, err)
	return resp, rm, err
}
