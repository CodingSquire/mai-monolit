//Package httpserver http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import "net/http"

const (
	// URIPrefix ...
	URIPrefix = "/api/v1/conference"

	// URIPathClientCreateConference ...
	URIPathClientCreateConference = URIPrefix + "/create"
	// URIPathClientChangeConference ...
	URIPathClientChangeConference = URIPrefix + "/change"
	// URIPathClientGetConferenceByFilter ...
	URIPathClientGetConferenceByFilter = URIPrefix + "/get"

	// HTTPMethodCreateConference ...
	HTTPMethodCreateConference = http.MethodPost
	// HTTPMethodChangeConference ...
	HTTPMethodChangeConference = http.MethodPost
	// HTTPMethodGetConferenceByFilter ...
	HTTPMethodGetConferenceByFilter = http.MethodPost
)
