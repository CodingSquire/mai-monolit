//Package httpclient http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"context"
	"github.com/CodingSquire/mai-monolit/pkg/api"

	"github.com/valyala/fasthttp"
)





// Service implements Service interface
type Service interface {
	CreateThesis(ctx context.Context, request *api.CreateThesisRequest)(response api.CreateThesisResponse, err error)
	//ChangeThesis(ctx context.Context, request *api.ChangeThesisRequest)(response api.ChangeThesisResponse, err error)
	//GetThesisByFilter(ctx context.Context, request *api.GetThesisByFilterRequest)(response api.GetThesisByFilterResponse, err error)
}

type client struct {
	cli *fasthttp.HostClient

	transportCreateThesis CreateThesisTransport
}

// GetBrandsByID ...
func (s *client) CreateThesis(ctx context.Context, request *api.CreateThesisRequest)(response api.CreateThesisResponse, err error){
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if err = s.transportCreateThesis.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportCreateThesis.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,

	transportCreateThesis CreateThesisTransport,
) Service {
	return &client{
		cli: cli,

		transportCreateThesis: transportCreateThesis,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(
	serverURL string,
	serverHost string,
	maxConns int,
	errorProcessor errorProcessor,
	errorCreator errorCreator,

	uriPathCreateThesis string,



	httpMethodGetBrandsByID string,

) Service {

	transportCreateThesis := NewCreateThesisTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathCreateThesis,
		httpMethodGetBrandsByID,
	)

	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},

		transportCreateThesis,
	)
}
