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
	CreateSegment(ctx context.Context, request *api.CreateSegmentRequest)(response api.CreateSegmentResponse, err error)
	ChangeSegment(ctx context.Context, request *api.ChangeSegmentRequest)(response api.ChangeSegmentResponse, err error)
	GetSegmentByFilter(ctx context.Context, request *api.GetSegmentByFilterRequest)(response api.GetSegmentByFilterResponse, err error)
}

type createSegmentServer struct {
	transport      CreateSegmentTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *createSegmentServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.CreateSegment(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewCreateSegmentServer the server creator
func  NewCreateSegmentServer(transport CreateSegmentTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := createSegmentServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}






type changeSegmentServer struct {
	transport      ChangeSegmentTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *changeSegmentServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.ChangeSegment(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewChangeSegmentServer the server creator
func  NewChangeSegmentServer(transport ChangeSegmentTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := changeSegmentServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}












type getSegmentByFilterServer struct {
	transport      GetSegmentByFilterTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getSegmentByFilterServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetSegmentByFilter(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetSegmentByFilterServer the server creator
func  NewGetSegmentByFilterServer(transport GetSegmentByFilterTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getSegmentByFilterServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}




// NewPreparedServer factory for server api handler
func NewPreparedServer(svc service) *fasthttprouter.Router {
	errorProcessor := httpserver.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	createSegmentTransport := NewCreateSegmentTransport(httpserver.NewError)
	changeSegmentTransport := NewChangeSegmentTransport(httpserver.NewError)
	getSegmentByFilterTransport := NewGetSegmentByFilterTransport(httpserver.NewError)
	return httpserver.MakeFastHTTPRouter(
		[]*httpserver.HandlerSettings{
			{
				Path:   URIPathClientCreateSegment,
				Method: HTTPMethodCreateSegment,
				Handler: NewCreateSegmentServer(
					createSegmentTransport,
					svc,
					errorProcessor,
				),
			},
			{
				Path:   URIPathClientChangeSegment,
				Method: HTTPMethodChangeSegment,
				Handler: NewChangeSegmentServer(
					changeSegmentTransport,
					svc,
					errorProcessor,
				),
			},
			{
				Path:   URIPathClientGetSegmentByFilter,
				Method: HTTPMethodGetSegmentByFilter,
				Handler: NewGetSegmentByFilterServer(
					getSegmentByFilterTransport,
					svc,
					errorProcessor,
				),
			},
		},
	)
}
