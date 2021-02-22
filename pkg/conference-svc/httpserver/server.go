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
	CreateConference(ctx context.Context, request *api.CreateConferenceRequest)(response api.CreateConferenceResponse, err error)
	ChangeConference(ctx context.Context, request *api.ChangeConferenceRequest)(response api.ChangeConferenceResponse, err error)
	GetConferenceByFilter(ctx context.Context, request *api.GetConferenceByFilterRequest)(response api.GetConferenceByFilterResponse, err error)
}

type createConferenceServer struct {
	transport      CreateConferenceTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *createConferenceServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.CreateConference(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewCreateConferenceServer the server creator
func  NewCreateConferenceServer(transport CreateConferenceTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := createConferenceServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}






type changeConferenceServer struct {
	transport      ChangeConferenceTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *changeConferenceServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.ChangeConference(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewChangeConferenceServer the server creator
func  NewChangeConferenceServer(transport ChangeConferenceTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := changeConferenceServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}












type getConferenceByFilterServer struct {
	transport      GetConferenceByFilterTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getConferenceByFilterServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetConferenceByFilter(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetConferenceByFilterServer the server creator
func  NewGetConferenceByFilterServer(transport GetConferenceByFilterTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getConferenceByFilterServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}




// NewPreparedServer factory for server api handler
func NewPreparedServer(svc service) *fasthttprouter.Router {
	errorProcessor := httpserver.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	createConferenceTransport := NewCreateConferenceTransport(httpserver.NewError)
	changeConferenceTransport := NewChangeConferenceTransport(httpserver.NewError)
	getConferenceByFilterTransport := NewGetConferenceByFilterTransport(httpserver.NewError)
	return httpserver.MakeFastHTTPRouter(
		[]*httpserver.HandlerSettings{
			{
				Path:   URIPathClientCreateConference,
				Method: HTTPMethodCreateConference,
				Handler: NewCreateConferenceServer(
					createConferenceTransport,
					svc,
					errorProcessor,
				),
			},
			{
				Path:   URIPathClientChangeConference,
				Method: HTTPMethodChangeConference,
				Handler: NewChangeConferenceServer(
					changeConferenceTransport,
					svc,
					errorProcessor,
				),
			},
			{
				Path:   URIPathClientGetConferenceByFilter,
				Method: HTTPMethodGetConferenceByFilter,
				Handler: NewGetConferenceByFilterServer(
					getConferenceByFilterTransport,
					svc,
					errorProcessor,
				),
			},
		},
	)
}
