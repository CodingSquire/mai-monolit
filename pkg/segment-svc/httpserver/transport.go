package httpserver

import (
	"context"
	"github.com/CodingSquire/mai-monolit/pkg/api"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"net/http"
)

type errorCreator func(status int, format string, v ...interface{}) error

// CreateSegmentTransport transport interface
type CreateSegmentTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.CreateSegmentRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.CreateSegmentResponse) (err error)
}

type createSegmentTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *createSegmentTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.CreateSegmentRequest, err error) {
	err=request.UnmarshalJSON(r.Body())
	if err!=nil{
		return request,t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *createSegmentTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.CreateSegmentResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return

}

// NewCreateSegmentTransport the transport creator for http requests
func NewCreateSegmentTransport(errorCreator errorCreator) CreateSegmentTransport {
	return &createSegmentTransport{
		errorCreator: errorCreator,
	}
}

// ChangeSegmentTransport transport interface
type ChangeSegmentTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.ChangeSegmentRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.ChangeSegmentResponse) (err error)
}

type changeSegmentTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *changeSegmentTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.ChangeSegmentRequest, err error) {
	err=request.UnmarshalJSON(r.Body())
	if err!=nil{
		return request,t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *changeSegmentTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.ChangeSegmentResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewCreateSegmentTransport the transport creator for http requests
func NewChangeSegmentTransport(errorCreator errorCreator) ChangeSegmentTransport {
	return &changeSegmentTransport{
		errorCreator: errorCreator,
	}
}

// GetSegmentByFilterTransport transport interface
type GetSegmentByFilterTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetSegmentByFilterRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetSegmentByFilterResponse) (err error)
}

type getSegmentByFilterTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getSegmentByFilterTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetSegmentByFilterRequest, err error) {
	err=request.UnmarshalJSON(r.Body())
	if err!=nil{
		return request,t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getSegmentByFilterTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetSegmentByFilterResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetSegmentByFilterTransport the transport creator for http requests
func NewGetSegmentByFilterTransport(errorCreator errorCreator) GetSegmentByFilterTransport {
	return &getSegmentByFilterTransport{
		errorCreator: errorCreator,
	}
}

