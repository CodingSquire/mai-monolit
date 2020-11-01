//Package httpclient http client
//CODE GENERATED AUTOMATICALLY
//THIS FILE COULD BE EDITED BY HANDS
package httpclient

import (
	"context"

	"github.com/valyala/fasthttp"

	v1 "git.wildberries.ru/portals/analytics-back/pkg/api/bansvc/v1"
)

var (
	// GetBrandsByID ...
	GetBrandsByID = option{}
	// GetByFilter ...
	GetByFilter = option{}
	// GetExcelReport ...
	GetExcelReport = option{}
)

type option struct{}

// Option ...
type Option interface {
	Prepare(ctx context.Context, r *fasthttp.Request)
}

// Service implements Service interface
type Service interface {
	GetBrandsByID(ctx context.Context, request *v1.GetBrandsByIDRequest) (response v1.GetBrandsByIDResponse, err error)
	GetByFilter(ctx context.Context, request *v1.GetByFilterRequest) (response v1.GetByFilterResponse, err error)
	GetExcelReport(ctx context.Context, request *v1.GetExcelReportRequest) (response v1.GetExcelReportResponse, err error)
}

type client struct {
	cli *fasthttp.HostClient

	transportGetExcelReport GetExcelReportClientTransport
	transportGetBrandsByID  GetBrandsByIDClientTransport
	transportGetByFilter    GetByFilterClientTransport
	options                 map[interface{}]Option
}

// GetBrandsByID ...
func (s *client) GetBrandsByID(ctx context.Context, request *v1.GetBrandsByIDRequest) (response v1.GetBrandsByIDResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetBrandsByID]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetBrandsByID.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetBrandsByID.DecodeResponse(ctx, res)
}

// GetByFilter ...
func (s *client) GetByFilter(ctx context.Context, request *v1.GetByFilterRequest) (response v1.GetByFilterResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetByFilter]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetByFilter.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetByFilter.DecodeResponse(ctx, res)
}

// GetExcelReport ...
func (s *client) GetExcelReport(ctx context.Context, request *v1.GetExcelReportRequest) (response v1.GetExcelReportResponse, err error) {
	req, res := fasthttp.AcquireRequest(), fasthttp.AcquireResponse()
	defer func() {
		fasthttp.ReleaseRequest(req)
		fasthttp.ReleaseResponse(res)
	}()
	if opt, ok := s.options[GetExcelReport]; ok {
		opt.Prepare(ctx, req)
	}
	if err = s.transportGetExcelReport.EncodeRequest(ctx, req, request); err != nil {
		return
	}
	err = s.cli.Do(req, res)
	if err != nil {
		return
	}
	return s.transportGetExcelReport.DecodeResponse(ctx, res)
}

// NewClient the client creator
func NewClient(
	cli *fasthttp.HostClient,

	transportGetExcelReport GetExcelReportClientTransport,
	transportGetBrandsByID GetBrandsByIDClientTransport,
	transportGetByFilter GetByFilterClientTransport,
	options map[interface{}]Option,
) Service {
	return &client{
		cli: cli,

		transportGetExcelReport: transportGetExcelReport,
		transportGetBrandsByID:  transportGetBrandsByID,
		transportGetByFilter:    transportGetByFilter,
		options:                 options,
	}
}

// NewPreparedClient create and set up http client
func NewPreparedClient(
	serverURL string,
	serverHost string,
	maxConns int,
	options map[interface{}]Option,
	errorProcessor errorProcessor,
	errorCreator errorCreator,

	uriPathGetExcelReport string,
	uriPathGetBrandsByID string,
	uriPathGetByFilter string,

	httpMethodGetExcelReport string,
	httpMethodGetBrandsByID string,
	httpMethodGetByFilter string,
) Service {

	transportGetBrandsByID := NewGetBrandsByIDClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetBrandsByID,
		httpMethodGetBrandsByID,
	)

	transportGetByFilter := NewGetByFilterClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetByFilter,
		httpMethodGetByFilter,
	)

	transportGetExcelReport := NewGetExcelReportClientTransport(
		errorProcessor,
		errorCreator,
		serverURL+uriPathGetExcelReport,
		httpMethodGetExcelReport,
	)
	return NewClient(
		&fasthttp.HostClient{
			Addr:     serverHost,
			MaxConns: maxConns,
		},

		transportGetExcelReport,
		transportGetBrandsByID,
		transportGetByFilter,
		options,
	)
}
