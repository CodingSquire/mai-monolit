//Package httpserver http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import "net/http"

const (
	// URIPrefix ...
	URIPrefix = "/api/v1/thesis"

	// URIPathClientCreateThesis ...
	URIPathClientCreateThesis = URIPrefix + "/create"
	// URIPathClientChangeThesis ...
	URIPathClientChangeThesis = URIPrefix + "/change"
	// URIPathClientGetThesisByFilter ...
	URIPathClientGetThesisByFilter = URIPrefix + "/get"

	// HTTPMethodCreateThesis ...
	HTTPMethodCreateThesis = http.MethodPost
	// HTTPMethodChangeThesis ...
	HTTPMethodChangeThesis = http.MethodPost
	// HTTPMethodGetThesisByFilter ...
	HTTPMethodGetThesisByFilter = http.MethodPost
)
