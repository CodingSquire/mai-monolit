//Package httpserver http server
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpserver

import (
	"context"
	"github.com/CodingSquire/mai-monolit/pkg/api"
	"net/http"

	"github.com/buaazp/fasthttprouter"
	"github.com/valyala/fasthttp"

	v1 "git.wildberries.ru/portals/analytics-back/pkg/api/bansvc/v1"
	"git.wildberries.ru/portals/analytics-back/pkg/httpserver"
)

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
}

type service interface {
	CreateThesis(ctx context.Context, request *api.CreateThesisRequest)(response api.CreateThesisResponse, err error)
	ChangeThesis(ctx context.Context, request *api.ChangeThesisRequest)(response api.ChangeThesisResponse, err error)
	GetThesisByFilter(ctx context.Context, request *api.GetThesisByFilterRequest)(response api.GetThesisByFilterResponse, err error)
}

type getBrandsByIDServer struct {
	transport      GetBrandsByIDTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getBrandsByIDServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetBrandsByID(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetBrandsByIDServer the server creator
func NewGetBrandsByIDServer(transport GetBrandsByIDTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getBrandsByIDServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getByFilterServer struct {
	transport      GetByFilterTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getByFilterServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetByFilter(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetByFilterServer the server creator
func NewGetByFilterServer(transport GetByFilterTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getByFilterServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

type getExcelReportServer struct {
	transport      GetExcelReportTransport
	service        service
	errorProcessor errorProcessor
}

// ServeHTTP implements http.Handler.
func (s *getExcelReportServer) ServeHTTP(ctx *fasthttp.RequestCtx) {
	request, err := s.transport.DecodeRequest(ctx, &ctx.Request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	response, err := s.service.GetExcelReport(ctx, &request)
	if err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}

	if err := s.transport.EncodeResponse(ctx, &ctx.Response, &response); err != nil {
		s.errorProcessor.Encode(ctx, &ctx.Response, err)
		return
	}
}

// NewGetExcelReportServer the server creator
func NewGetExcelReportServer(transport GetExcelReportTransport, service service, errorProcessor errorProcessor) fasthttp.RequestHandler {
	ls := getExcelReportServer{
		transport:      transport,
		service:        service,
		errorProcessor: errorProcessor,
	}
	return ls.ServeHTTP
}

// NewPreparedServer factory for server api handler
func NewPreparedServer(svc service) *fasthttprouter.Router {
	errorProcessor := httpserver.NewErrorProcessor(http.StatusInternalServerError, "internal error")
	getExcelReportTransport := NewGetExcelReportTransport(httpserver.NewError)
	getBrandsByIDTransport := NewGetBrandsByIDTransport(httpserver.NewError)
	getByFilterTransport := NewGetByFilterTransport(httpserver.NewError)

	return httpserver.MakeFastHTTPRouter(
		[]*httpserver.HandlerSettings{
			{
				Path:   URIPathClientGetExcelReport,
				Method: HTTPMethodGetExcelReport,
				Handler: NewGetExcelReportServer(
					getExcelReportTransport,
					svc,
					errorProcessor,
				),
			},
			{
				Path:   URIPathClientGetBrandsByID,
				Method: HTTPMethodGetBrandsByID,
				Handler: NewGetBrandsByIDServer(
					getBrandsByIDTransport,
					svc,
					errorProcessor,
				),
			}, {
				Path:   URIPathClientGetByFilter,
				Method: HTTPMethodGetByFilter,
				Handler: NewGetByFilterServer(
					getByFilterTransport,
					svc,
					errorProcessor,
				),
			},
		},
	)
}
