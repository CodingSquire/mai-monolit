package httpserver

import (
	"context"
	"github.com/CodingSquire/mai-monolit/pkg/api"
	"github.com/valyala/fasthttp"
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
	return
}

// EncodeResponse method for encoding response on server side
func (t *createThesisTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.CreateThesisResponse) (err error) {
	return
}

// NewCreateThesisTransport the transport creator for http requests
func NewCreateThesisTransport(errorCreator errorCreator) CreateThesisTransport {
	return &createThesisTransport{
		errorCreator: errorCreator,
	}
}
//
//// ChangeThesisTransport transport interface
//type ChangeThesisTransport interface {
//	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.ChangeThesisRequest, err error)
//	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.ChangeThesisResponse) (err error)
//}
//
//type changeThesisTransport struct {
//	errorCreator errorCreator
//}
//
//// DecodeRequest method for decoding requests on server side
//func (t *changeThesisTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.ChangeThesisRequest, err error) {
//	return
//}
//
//// EncodeResponse method for encoding response on server side
//func (t *changeThesisTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.ChangeThesisResponse) (err error) {
//	return
//}
//
//// NewCreateThesisTransport the transport creator for http requests
//func NewChangeThesisTransport(errorCreator errorCreator) ChangeThesisTransport {
//	return &changeThesisTransport{
//		errorCreator: errorCreator,
//	}
//}
//
//// GetThesisByFilterTransport transport interface
//type GetThesisByFilterTransport interface {
//	DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetThesisByFilterRequest, err error)
//	EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetThesisByFilterResponse) (err error)
//}
//
//type getThesisByFilterTransport struct {
//	errorCreator errorCreator
//}
//
//// DecodeRequest method for decoding requests on server side
//func (t *getThesisByFilterTransport) DecodeRequest(ctx context.Context, r *fasthttp.Request) (request api.GetThesisByFilterRequest, err error) {
//	return
//}
//
//// EncodeResponse method for encoding response on server side
//func (t *getThesisByFilterTransport) EncodeResponse(ctx context.Context, r *fasthttp.Response, response *api.GetThesisByFilterResponse) (err error) {
//	return
//}
//
//// NewGetThesisByFilterTransport the transport creator for http requests
//func NewGetThesisByFilterTransport(errorCreator errorCreator) GetThesisByFilterTransport {
//	return &getThesisByFilterTransport{
//		errorCreator: errorCreator,
//	}
//}

