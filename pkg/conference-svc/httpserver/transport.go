package httpserver

import (
	"context"
	"github.com/CodingSquire/mai-monolit/pkg/api"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"net/http"
)

type errorCreator func(status int, format string, v ...interface{}) error

// CreateConferenceTransport transport interface
type CreateConferenceTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.CreateConferenceRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.CreateConferenceResponse) (err error)
}

type createConferenceTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *createConferenceTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.CreateConferenceRequest, err error) {
	err=request.UnmarshalJSON(r.Body())
	if err!=nil{
		return request,t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *createConferenceTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.CreateConferenceResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return

}

// NewCreateConferenceTransport the transport creator for http requests
func NewCreateConferenceTransport(errorCreator errorCreator) CreateConferenceTransport {
	return &createConferenceTransport{
		errorCreator: errorCreator,
	}
}

// ChangeConferenceTransport transport interface
type ChangeConferenceTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.ChangeConferenceRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.ChangeConferenceResponse) (err error)
}

type changeConferenceTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *changeConferenceTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.ChangeConferenceRequest, err error) {
	err=request.UnmarshalJSON(r.Body())
	if err!=nil{
		return request,t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *changeConferenceTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.ChangeConferenceResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewCreateConferenceTransport the transport creator for http requests
func NewChangeConferenceTransport(errorCreator errorCreator) ChangeConferenceTransport {
	return &changeConferenceTransport{
		errorCreator: errorCreator,
	}
}

// GetConferenceByFilterTransport transport interface
type GetConferenceByFilterTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetConferenceByFilterRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetConferenceByFilterResponse) (err error)
}

type getConferenceByFilterTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *getConferenceByFilterTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetConferenceByFilterRequest, err error) {
	err=request.UnmarshalJSON(r.Body())
	if err!=nil{
		return request,t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *getConferenceByFilterTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetConferenceByFilterResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return
}

// NewGetConferenceByFilterTransport the transport creator for http requests
func NewGetConferenceByFilterTransport(errorCreator errorCreator) GetConferenceByFilterTransport {
	return &getConferenceByFilterTransport{
		errorCreator: errorCreator,
	}
}

