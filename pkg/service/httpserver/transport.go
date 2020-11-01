package httpserver

import (
	"context"
	"github.com/CodingSquire/mai-monolit/pkg/api"
	"github.com/mailru/easyjson"
	"github.com/valyala/fasthttp"
	"net/http"
)

type errorCreator func(status int, format string, v ...interface{}) error



//type service interface {
//	CreateThesis(ctx context.Context, request *api.CreateThesisRequest)(response api.CreateThesisResponse, err error)
//	ChangeThesis(ctx context.Context, request *api.ChangeThesisRequest)(response api.ChangeThesisResponse, err error)
//	GetThesisByFilter(ctx context.Context, request *api.GetThesisByFilterRequest)(response api.GetThesisByFilterResponse, err error)
//}

// CreateThesisTransport transport interface
type CreateThesisTransport interface {
	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.CreateThesisRequest, err error)
	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.CreateThesisResponse) (err error)
}

type createThesisTransport struct {
	errorCreator errorCreator
}

// DecodeRequest method for decoding requests on server side
func (t *createThesisTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.CreateThesisRequest, err error) {
	err=request.UnmarshalJSON(r.Body())
	if err!=nil{
		return request,t.errorCreator(http.StatusInternalServerError, "failed to decode JSON response: %s", err)
	}
	return
}

// EncodeResponse method for encoding response on server side
func (t *createThesisTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.CreateThesisResponse) (err error) {
	r.Header.Set("Content-Type", "application/json")
	if _, err = easyjson.MarshalToWriter(response, r.BodyWriter()); err != nil {
		return t.errorCreator(http.StatusInternalServerError, "failed to encode JSON response: %s", err)
	}
	return

}

// NewCreateThesisTransport the transport creator for http requests
func NewCreateThesisTransport(errorCreator errorCreator) CreateThesisTransport {
	return &createThesisTransport{
		errorCreator: errorCreator,
	}
}
