package httpclient

import (

	"context"


	"github.com/CodingSquire/mai-monolit/pkg/api"
	"github.com/valyala/fasthttp"

)

type errorCreator func(status int, format string, v ...interface{}) error

type errorProcessor interface {
	Encode(ctx context.Context, r *fasthttp.Response, err error)
	Decode(r *fasthttp.Response) error
}

// CreateThesisTransport transport interface
type CreateThesisTransport interface {
	EncodeRequest(ctx context.Context, r *fasthttp.Request, request *api.CreateThesisRequest) (err error)
	DecodeResponse(ctx context.Context, r *fasthttp.Response) (response api.CreateThesisResponse, err error)
}

type createThesisTransport struct {
	errorProcessor errorProcessor
	errorCreator   errorCreator
	pathTemplate   string
	method         string
}

// EncodeRequest method for encoding requests on client side
func (t *createThesisTransport) EncodeRequest(ctx context.Context, r *fasthttp.Request, request *api.CreateThesisRequest) (err error){
	return
}

// DecodeResponse method for decoding response on client side
func (t *createThesisTransport) DecodeResponse(ctx context.Context, r *fasthttp.Response) (response api.CreateThesisResponse, err error){
	return
}

// NewCreateThesisTransport the transport creator for http requests
func NewCreateThesisTransport(
	errorProcessor errorProcessor,
	errorCreator errorCreator,
	pathTemplate string,
	method string,
) CreateThesisTransport {
	return &createThesisTransport{
		errorProcessor: errorProcessor,
		errorCreator:   errorCreator,
		pathTemplate:   pathTemplate,
		method:         method,
	}
}

