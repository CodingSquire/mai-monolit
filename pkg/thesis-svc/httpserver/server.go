//Package httpserver http server
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"
	"github.com/CodingSquire/mai-monolit/pkg/api"
	"github.com/CodingSquire/mai-monolit/pkg/httpserver"
	"net/http"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"



)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type service interface {
	CreateThesis(ctx context.Context, request *api.CreateThesisRequest)(response api.CreateThesisResponse, err error)
	ChangeThesis(ctx context.Context, request *api.ChangeThesisRequest)(response api.ChangeThesisResponse, err error)
	GetThesisByFilter(ctx context.Context, request *api.GetThesisByFilterRequest)(response api.GetThesisByFilterResponse, err error)
}

type createThesisServer struct {
	transport      CreateThesisTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *createThesisServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.CreateThesis(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewCreateThesisServer the server creator
func  NewCreateThesisServer(transport CreateThesisTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := createThesisServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}






type changeThesisServer struct {
	transport      ChangeThesisTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *changeThesisServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.ChangeThesis(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewChangeThesisServer the server creator
func  NewChangeThesisServer(transport ChangeThesisTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := changeThesisServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}












type getThesisByFilterServer struct {
	transport      GetThesisByFilterTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getThesisByFilterServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetThesisByFilter(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetThesisByFilterServer the server creator
func  NewGetThesisByFilterServer(transport GetThesisByFilterTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getThesisByFilterServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}




// NewPreparedServer factory for server api handler
func NewPreparedServer(svc service) *fasthttprouter.Router {
	errorProcessor := httpserver.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	createThesisTransport := NewCreateThesisTransport(httpserver.NewError)
	changeThesisTransport := NewChangeThesisTransport(httpserver.NewError)
	getThesisByFilterTransport := NewGetThesisByFilterTransport(httpserver.NewError)
	return httpserver.MakeFastHTTPRouter(
		[]*httpserver.HandlerSettings{
			{
				Path:   URIPathClientCreateThesis,
				Method: HTTPMethodCreateThesis,
				Handler: NewCreateThesisServer(
					createThesisTransport,
					svc,
					errorProcessor,
				),
			},
			{
				Path:   URIPathClientChangeThesis,
				Method: HTTPMethodChangeThesis,
				Handler: NewChangeThesisServer(
					changeThesisTransport,
					svc,
					errorProcessor,
				),
			},
			{
				Path:   URIPathClientGetThesisByFilter,
				Method: HTTPMethodGetThesisByFilter,
				Handler: NewGetThesisByFilterServer(
					getThesisByFilterTransport,
					svc,
					errorProcessor,
				),
			},
		},
	)
}
