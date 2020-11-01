package httpclient

import (
	"context"
	"net/http"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/valyala/fasthttp"

	v1 "git.wildberries.ru/portals/analytics-back/pkg/api/bansvc/v1"
	"git.wildberries.ru/portals/analytics-back/pkg/banreportsvc"
	httpserver2 "git.wildberries.ru/portals/analytics-back/pkg/banreportsvc/httpserver"
	"git.wildberries.ru/portals/analytics-back/pkg/httpserver"
	jwt2 "git.wildberries.ru/portals/analytics-back/pkg/jwt"
)

const (
	getBrandsByIDSuccessTest = "get brands by supplierID success test"
	getBrandsByIDFailTest    = "get brands by supplierID fail test"

	getByFilterShkSuccessTest   = "get banned products by filter with SHK success test"
	getByFilterNoShkSuccessTest = "get banned products by filter success test"

	getByFilterShkFailTest   = "get banned products by filter with SHK fail test"
	getByFilterNoShkFailTest = "get banned products by filter fail test"

	getExcelReportFailTest    = "get excel report fail test"
	getExcelReportSuccessTest = "get excel report success test"
)

const (
	supplierID          = 1637
	errorSupplierID     = -1
	brandIDAdidas       = 21
	brandIDColambetta   = 1334
	brandIDAugustinWelz = 7314
	defLimit            = 50
	defOffset           = 0
	tokenSuccess        = "success token"
	tokenFail           = "fail token"
)

const (
	protocol                 = "http://"
	serverBanService         = "localhost:8080"
	maxConns                 = 4096
	defaultErrorCode         = http.StatusInternalServerError
	defaultErrorMessage      = "default error"
	maxRequestBodySize       = 25 * 1024 * 1024
	serverLaunchingWaitSleep = 2 * time.Second
	serverReadTimeout        = 2 * time.Second
)

const (
	methodGetBrandsByID    = "GetBrandsByID"
	methodGetByFilter      = "GetByFilter"
	methodParseTokenString = "ParseTokenString"
	methodGetExcelReport   = "GetExcelReport"
)

var (
	someErrorCode    = http.StatusNotFound
	someErrorMessage = "some error"
	someServiceError = httpserver.NewError(someErrorCode, someErrorMessage)
	allowedIssuers   = []string{"suppliers-portal-new-dataline", "portal-purchase-admin-dataline"}
)

func TestClient_GetBrandsByIDSuccess(t *testing.T) {
	getBrandsByIDRequest := v1.GetBrandsByIDRequest{
		AuthToken:  tokenSuccess,
		SupplierID: supplierID,
	}
	expectedGetBrandsByIDResponse := v1.GetBrandsByIDResponse{
		Data: v1.GetBrandsByIDResponseData{
			Brands: []*v1.Brand{
				{
					HumanFriendlyName: "/brands/adidas",
					ID:                intptr(brandIDAdidas),
					Name:              "adidas",
				},

				{
					HumanFriendlyName: "/brands/colambetta",
					ID:                intptr(brandIDColambetta),
					Name:              "Colambetta",
				},
				{
					HumanFriendlyName: "/brands/augustin-welz",
					ID:                intptr(brandIDAugustinWelz),
					Name:              "Augustin Welz",
				},
			},
		},
	}

	t.Run(getBrandsByIDSuccessTest, func(t *testing.T) {
		serviceMock := new(banreportsvc.MockService)
		serviceMock.On(
			methodGetBrandsByID,
			context.Background(),
			&getBrandsByIDRequest,
		).Return(expectedGetBrandsByIDResponse, nil).Once()
		client := NewPreparedClient(
			protocol+serverBanService,
			serverBanService,
			maxConns,
			nil,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
			httpserver.NewError,
			httpserver2.URIPathClientGetExcelReport,
			httpserver2.URIPathClientGetBrandsByID,
			httpserver2.URIPathClientGetByFilter,
			httpserver2.HTTPMethodGetExcelReport,
			httpserver2.HTTPMethodGetBrandsByID,
			httpserver2.HTTPMethodGetByFilter,
		)

		jwt := new(jwt2.MockJwt)
		jwt.On(
			methodParseTokenString,
			tokenSuccess,
		).Return(supplierID, nil)

		service := banreportsvc.NewValidationMiddleware(
			jwt,
			httpserver.NewError,
			serviceMock,
		)

		handlerFunc := httpserver2.NewGetBrandsByIDServer(
			httpserver2.NewGetBrandsByIDTransport(httpserver.NewError),
			service,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
		)
		router := httpserver.MakeFastHTTPRouter(
			[]*httpserver.HandlerSettings{
				{
					Path:    httpserver2.URIPathClientGetBrandsByID,
					Method:  httpserver2.HTTPMethodGetBrandsByID,
					Handler: handlerFunc,
				},
			},
		)
		server := fasthttp.Server{
			Handler:            router.Handler,
			MaxRequestBodySize: maxRequestBodySize,
			ReadTimeout:        serverReadTimeout,
		}

		go server.ListenAndServe(serverBanService)
		defer server.Shutdown()
		time.Sleep(serverLaunchingWaitSleep)

		actualGetBrandsByIDResponse, err := client.GetBrandsByID(context.Background(), &getBrandsByIDRequest)
		assert.NoError(t, err, "unexpected error", err)
		assert.Equal(t, expectedGetBrandsByIDResponse, actualGetBrandsByIDResponse)
	})
}

func TestClient_GetByFilterShkSuccessTest(t *testing.T) {
	hasShk := true
	getByFilterShkRequest := v1.GetByFilterRequest{
		AuthToken: tokenSuccess,
		Filter: v1.Filter{
			BrandID:    intptr(brandIDColambetta),
			SupplierID: supplierID,
			HasShk:     boolptr(hasShk),
			Limit:      intptr(defLimit),
			Offset:     intptr(defOffset),
		},
	}
	expectedGetByFilterShkResponse := v1.GetByFilterResponse{
		Data: v1.GetByFilterResponseData{
			Report: []*v1.ReportElement{
				{
					Brand:  "Colambetta",
					Nm:     5896025,
					Sa:     "MS40504/зеленый",
					Src:    "CNT",
					Reason: "Карточка ожидает заполнения",
				},
				{
					Brand:  "Colambetta",
					Nm:     5896025,
					Sa:     "MS40504/зеленый",
					Src:    "AJP",
					Reason: "Не установлена цена",
				},
			},
		},
	}
	t.Run(getByFilterShkSuccessTest, func(t *testing.T) {
		serviceMock := new(banreportsvc.MockService)
		serviceMock.On(
			methodGetByFilter,
			context.Background(),
			&getByFilterShkRequest,
		).Return(expectedGetByFilterShkResponse, nil)
		jwt := new(jwt2.MockJwt)
		jwt.On(
			methodParseTokenString,
			tokenSuccess,
		).Return(supplierID, nil)

		service := banreportsvc.NewValidationMiddleware(
			jwt,
			httpserver.NewError,
			serviceMock,
		)

		client := NewPreparedClient(
			protocol+serverBanService,
			serverBanService,
			maxConns,
			nil,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
			httpserver.NewError,
			httpserver2.URIPathClientGetExcelReport,
			httpserver2.URIPathClientGetBrandsByID,
			httpserver2.URIPathClientGetByFilter,
			httpserver2.HTTPMethodGetExcelReport,
			httpserver2.HTTPMethodGetBrandsByID,
			httpserver2.HTTPMethodGetByFilter,
		)
		handlerFunc := httpserver2.NewGetByFilterServer(
			httpserver2.NewGetByFilterTransport(httpserver.NewError),
			service,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
		)
		router := httpserver.MakeFastHTTPRouter(
			[]*httpserver.HandlerSettings{
				{
					Path:    httpserver2.URIPathClientGetByFilter,
					Method:  httpserver2.HTTPMethodGetByFilter,
					Handler: handlerFunc,
				},
			},
		)
		server := fasthttp.Server{
			Handler:            router.Handler,
			MaxRequestBodySize: maxRequestBodySize,
			ReadTimeout:        serverReadTimeout,
		}

		go server.ListenAndServe(serverBanService)
		defer server.Shutdown()
		time.Sleep(serverLaunchingWaitSleep)

		actualGetByFilterShkResponse, err := client.GetByFilter(context.Background(), &getByFilterShkRequest)
		assert.NoError(t, err, "unexpected error", err)
		assert.Equal(t, expectedGetByFilterShkResponse, actualGetByFilterShkResponse)
	})
}

func TestClient_GetByFilterNoShkSuccessTest(t *testing.T) {
	getByFilterNoShkRequest := v1.GetByFilterRequest{
		AuthToken: tokenSuccess,
		Filter: v1.Filter{
			BrandID:    intptr(brandIDColambetta),
			SupplierID: supplierID,
			Limit:      intptr(defLimit),
			Offset:     intptr(defOffset),
		},
	}
	expectedGetByFilterNoShkResponse := v1.GetByFilterResponse{
		Data: v1.GetByFilterResponseData{
			Report: []*v1.ReportElement{
				{
					Brand:  "Colambetta",
					Nm:     5896025,
					Sa:     "MS40504/зеленый",
					Src:    "CNT",
					Reason: "Карточка ожидает заполнения",
				},
				{
					Brand:  "Colambetta",
					Nm:     5896025,
					Sa:     "MS40504/зеленый",
					Src:    "AJP",
					Reason: "Не установлена цена",
				},
			},
		},
	}
	t.Run(getByFilterNoShkSuccessTest, func(t *testing.T) {
		serviceMock := new(banreportsvc.MockService)
		serviceMock.On(
			methodGetByFilter,
			context.Background(),
			&getByFilterNoShkRequest,
		).Return(expectedGetByFilterNoShkResponse, nil)
		jwt := new(jwt2.MockJwt)
		jwt.On(
			methodParseTokenString,
			tokenSuccess,
		).Return(supplierID, nil)

		service := banreportsvc.NewValidationMiddleware(
			jwt,
			httpserver.NewError,
			serviceMock,
		)

		client := NewPreparedClient(
			protocol+serverBanService,
			serverBanService,
			maxConns,
			nil,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
			httpserver.NewError,
			httpserver2.URIPathClientGetExcelReport,
			httpserver2.URIPathClientGetBrandsByID,
			httpserver2.URIPathClientGetByFilter,
			httpserver2.HTTPMethodGetExcelReport,
			httpserver2.HTTPMethodGetBrandsByID,
			httpserver2.HTTPMethodGetByFilter,
		)
		handlerFunc := httpserver2.NewGetByFilterServer(
			httpserver2.NewGetByFilterTransport(httpserver.NewError),
			service,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
		)
		router := httpserver.MakeFastHTTPRouter(
			[]*httpserver.HandlerSettings{
				{
					Path:    httpserver2.URIPathClientGetByFilter,
					Method:  httpserver2.HTTPMethodGetByFilter,
					Handler: handlerFunc,
				},
			},
		)
		server := fasthttp.Server{
			Handler:            router.Handler,
			MaxRequestBodySize: maxRequestBodySize,
			ReadTimeout:        serverReadTimeout,
		}

		go server.ListenAndServe(serverBanService)
		defer server.Shutdown()
		time.Sleep(serverLaunchingWaitSleep)

		actualGetByFilterShkResponse, err := client.GetByFilter(context.Background(), &getByFilterNoShkRequest)
		assert.NoError(t, err, "unexpected error", err)
		assert.Equal(t, expectedGetByFilterNoShkResponse, actualGetByFilterShkResponse)
	})
}

func TestClient_GetExcelReportSuccessTest(t *testing.T) {
	hasShk := true
	getExcelReportRequest := v1.GetExcelReportRequest{
		AuthToken: tokenSuccess,
		Filter: v1.Filter{
			BrandID:    intptr(brandIDColambetta),
			SupplierID: supplierID,
			HasShk:     boolptr(hasShk),
			Limit:      intptr(defLimit),
			Offset:     intptr(defOffset),
		},
	}
	expectedExcelFile := []byte{'e'}
	expectedGetExcelReportResponse := v1.GetExcelReportResponse{
		Data: v1.GetExcelReportResponseData{
			Body: expectedExcelFile,
		},
	}

	t.Run(getExcelReportSuccessTest, func(t *testing.T) {
		serviceMock := new(banreportsvc.MockService)
		serviceMock.On(
			methodGetExcelReport,
			context.Background(),
			&getExcelReportRequest,
		).Return(expectedGetExcelReportResponse, nil)
		jwt := new(jwt2.MockJwt)
		jwt.On(
			methodParseTokenString,
			tokenSuccess,
		).Return(supplierID, nil)

		service := banreportsvc.NewValidationMiddleware(
			jwt,
			httpserver.NewError,
			serviceMock,
		)

		client := NewPreparedClient(
			protocol+serverBanService,
			serverBanService,
			maxConns,
			nil,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
			httpserver.NewError,
			httpserver2.URIPathClientGetExcelReport,
			httpserver2.URIPathClientGetBrandsByID,
			httpserver2.URIPathClientGetByFilter,
			httpserver2.HTTPMethodGetExcelReport,
			httpserver2.HTTPMethodGetBrandsByID,
			httpserver2.HTTPMethodGetByFilter,
		)
		handlerFunc := httpserver2.NewGetExcelReportServer(
			httpserver2.NewGetExcelReportTransport(httpserver.NewError),
			service,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
		)
		router := httpserver.MakeFastHTTPRouter(
			[]*httpserver.HandlerSettings{
				{
					Path:    httpserver2.URIPathClientGetExcelReport,
					Method:  httpserver2.HTTPMethodGetExcelReport,
					Handler: handlerFunc,
				},
			},
		)
		server := fasthttp.Server{
			Handler:            router.Handler,
			MaxRequestBodySize: maxRequestBodySize,
			ReadTimeout:        serverReadTimeout,
		}

		go server.ListenAndServe(serverBanService)
		defer server.Shutdown()
		time.Sleep(serverLaunchingWaitSleep)

		actualGetExcelReportResponse, err := client.GetExcelReport(context.Background(), &getExcelReportRequest)
		assert.NoError(t, err, "unexpected error", err)
		assert.Equal(t, expectedGetExcelReportResponse, actualGetExcelReportResponse)
	})
}

func TestClient_GetBrandsByIDFail(t *testing.T) {
	getBrandsByIDRequest := v1.GetBrandsByIDRequest{
		AuthToken:  tokenFail,
		SupplierID: errorSupplierID,
	}
	expectedGetBrandsByIDResponse := v1.GetBrandsByIDResponse{
		Data:         v1.GetBrandsByIDResponseData{},
		Error:        true,
		ErrorText:    someErrorMessage,
		CustomErrors: nil,
	}
	expectedError := someServiceError
	t.Run(getBrandsByIDFailTest, func(t *testing.T) {
		serviceMock := new(banreportsvc.MockService)
		serviceMock.On(
			"GetBrandsByID",
			context.Background(),
			&getBrandsByIDRequest,
		).Return(expectedGetBrandsByIDResponse, someServiceError)
		jwt := new(jwt2.MockJwt)
		jwt.On(
			methodParseTokenString,
			tokenFail,
		).Return(errorSupplierID, nil)
		service := banreportsvc.NewValidationMiddleware(
			jwt,
			httpserver.NewError,
			serviceMock,
		)

		client := NewPreparedClient(
			protocol+serverBanService,
			serverBanService,
			maxConns,
			nil,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
			httpserver.NewError,
			httpserver2.URIPathClientGetExcelReport,
			httpserver2.URIPathClientGetBrandsByID,
			httpserver2.URIPathClientGetByFilter,
			httpserver2.HTTPMethodGetExcelReport,
			httpserver2.HTTPMethodGetBrandsByID,
			httpserver2.HTTPMethodGetByFilter,
		)
		handlerFunc := httpserver2.NewGetBrandsByIDServer(
			httpserver2.NewGetBrandsByIDTransport(httpserver.NewError),
			service,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
		)
		router := httpserver.MakeFastHTTPRouter(
			[]*httpserver.HandlerSettings{
				{
					Path:    httpserver2.URIPathClientGetBrandsByID,
					Method:  httpserver2.HTTPMethodGetBrandsByID,
					Handler: handlerFunc,
				},
			},
		)
		server := fasthttp.Server{
			Handler:            router.Handler,
			MaxRequestBodySize: maxRequestBodySize,
			ReadTimeout:        serverReadTimeout,
		}

		go server.ListenAndServe(serverBanService)
		defer server.Shutdown()
		time.Sleep(serverLaunchingWaitSleep)

		actualGetBrandsByIDResponse, err := client.GetBrandsByID(context.Background(), &getBrandsByIDRequest)
		assert.Equal(t, err, expectedError)
		assert.Equal(t, expectedGetBrandsByIDResponse, actualGetBrandsByIDResponse)
	})
}

func TestClient_GetByFilterShkFailTest(t *testing.T) {
	hasShk := true
	getByFilterShkRequest := v1.GetByFilterRequest{
		AuthToken: tokenFail,
		Filter: v1.Filter{
			BrandID:    intptr(brandIDColambetta),
			SupplierID: errorSupplierID,
			HasShk:     boolptr(hasShk),
			Limit:      intptr(defLimit),
			Offset:     intptr(defOffset),
		},
	}
	expectedGetByFilterShkResponse := v1.GetByFilterResponse{
		Data:         v1.GetByFilterResponseData{},
		Error:        true,
		ErrorText:    someErrorMessage,
		CustomErrors: nil,
	}
	expectedError := someServiceError
	t.Run(getByFilterShkFailTest, func(t *testing.T) {
		serviceMock := new(banreportsvc.MockService)
		serviceMock.On(
			methodGetByFilter,
			context.Background(),
			&getByFilterShkRequest,
		).Return(expectedGetByFilterShkResponse, someServiceError)
		jwt := new(jwt2.MockJwt)
		jwt.On(
			methodParseTokenString,
			tokenFail,
		).Return(errorSupplierID, nil)
		service := banreportsvc.NewValidationMiddleware(
			jwt,
			httpserver.NewError,
			serviceMock,
		)

		client := NewPreparedClient(
			protocol+serverBanService,
			serverBanService,
			maxConns,
			nil,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
			httpserver.NewError,
			httpserver2.URIPathClientGetExcelReport,
			httpserver2.URIPathClientGetBrandsByID,
			httpserver2.URIPathClientGetByFilter,
			httpserver2.HTTPMethodGetExcelReport,
			httpserver2.HTTPMethodGetBrandsByID,
			httpserver2.HTTPMethodGetByFilter,
		)
		handlerFunc := httpserver2.NewGetByFilterServer(
			httpserver2.NewGetByFilterTransport(httpserver.NewError),
			service,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
		)
		router := httpserver.MakeFastHTTPRouter(
			[]*httpserver.HandlerSettings{
				{
					Path:    httpserver2.URIPathClientGetByFilter,
					Method:  httpserver2.HTTPMethodGetByFilter,
					Handler: handlerFunc,
				},
			},
		)
		server := fasthttp.Server{
			Handler:            router.Handler,
			MaxRequestBodySize: maxRequestBodySize,
			ReadTimeout:        serverReadTimeout,
		}

		go server.ListenAndServe(serverBanService)
		defer server.Shutdown()
		time.Sleep(serverLaunchingWaitSleep)

		actualGetByFilterShkResponse, err := client.GetByFilter(context.Background(), &getByFilterShkRequest)
		assert.Equal(t, err, expectedError)
		assert.Equal(t, expectedGetByFilterShkResponse, actualGetByFilterShkResponse)
	})
}

func TestClient_GetByFilterNoShkFailTest(t *testing.T) {
	getByFilterNoShkRequest := v1.GetByFilterRequest{
		AuthToken: tokenFail,
		Filter: v1.Filter{
			BrandID:    intptr(brandIDColambetta),
			SupplierID: errorSupplierID,
			Limit:      intptr(defLimit),
			Offset:     intptr(defOffset),
		},
	}
	expectedGetByFilterNoShkResponse := v1.GetByFilterResponse{
		Data:         v1.GetByFilterResponseData{},
		Error:        true,
		ErrorText:    someErrorMessage,
		CustomErrors: nil,
	}
	expectedError := someServiceError
	t.Run(getByFilterNoShkFailTest, func(t *testing.T) {
		serviceMock := new(banreportsvc.MockService)
		serviceMock.On(
			methodGetByFilter,
			context.Background(),
			&getByFilterNoShkRequest,
		).Return(expectedGetByFilterNoShkResponse, someServiceError)
		jwt := new(jwt2.MockJwt)
		jwt.On(
			methodParseTokenString,
			tokenFail,
		).Return(errorSupplierID, nil)
		service := banreportsvc.NewValidationMiddleware(
			jwt,
			httpserver.NewError,
			serviceMock,
		)
		client := NewPreparedClient(
			protocol+serverBanService,
			serverBanService,
			maxConns,
			nil,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
			httpserver.NewError,
			httpserver2.URIPathClientGetExcelReport,
			httpserver2.URIPathClientGetBrandsByID,
			httpserver2.URIPathClientGetByFilter,
			httpserver2.HTTPMethodGetExcelReport,
			httpserver2.HTTPMethodGetBrandsByID,
			httpserver2.HTTPMethodGetByFilter,
		)
		handlerFunc := httpserver2.NewGetByFilterServer(
			httpserver2.NewGetByFilterTransport(httpserver.NewError),
			service,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
		)
		router := httpserver.MakeFastHTTPRouter(
			[]*httpserver.HandlerSettings{
				{
					Path:    httpserver2.URIPathClientGetByFilter,
					Method:  httpserver2.HTTPMethodGetByFilter,
					Handler: handlerFunc,
				},
			},
		)
		server := fasthttp.Server{
			Handler:            router.Handler,
			MaxRequestBodySize: maxRequestBodySize,
			ReadTimeout:        serverReadTimeout,
		}

		go server.ListenAndServe(serverBanService)
		defer server.Shutdown()
		time.Sleep(serverLaunchingWaitSleep)

		actualGetByFilterNoShkResponse, err := client.GetByFilter(context.Background(), &getByFilterNoShkRequest)
		assert.Equal(t, err, expectedError)
		assert.Equal(t, expectedGetByFilterNoShkResponse, actualGetByFilterNoShkResponse)
	})
}

func TestClient_GetExcelReportFailTest(t *testing.T) {
	hasShk := true
	getExcelReportRequest := v1.GetExcelReportRequest{
		AuthToken: tokenSuccess,
		Filter: v1.Filter{
			BrandID:    intptr(brandIDColambetta),
			SupplierID: supplierID,
			HasShk:     boolptr(hasShk),
			Limit:      intptr(defLimit),
			Offset:     intptr(defOffset),
		},
	}
	expectedError := someServiceError
	expectedGetExcelReportResponse := v1.GetExcelReportResponse{
		Error:     true,
		ErrorText: someErrorMessage,
	}

	t.Run(getExcelReportFailTest, func(t *testing.T) {
		serviceMock := new(banreportsvc.MockService)
		serviceMock.On(
			methodGetExcelReport,
			context.Background(),
			&getExcelReportRequest,
		).Return(expectedGetExcelReportResponse, someServiceError)
		jwt := new(jwt2.MockJwt)
		jwt.On(
			methodParseTokenString,
			tokenSuccess,
		).Return(supplierID, nil)

		service := banreportsvc.NewValidationMiddleware(
			jwt,
			httpserver.NewError,
			serviceMock,
		)

		client := NewPreparedClient(
			protocol+serverBanService,
			serverBanService,
			maxConns,
			nil,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
			httpserver.NewError,
			httpserver2.URIPathClientGetExcelReport,
			httpserver2.URIPathClientGetBrandsByID,
			httpserver2.URIPathClientGetByFilter,
			httpserver2.HTTPMethodGetExcelReport,
			httpserver2.HTTPMethodGetBrandsByID,
			httpserver2.HTTPMethodGetByFilter,
		)
		handlerFunc := httpserver2.NewGetExcelReportServer(
			httpserver2.NewGetExcelReportTransport(httpserver.NewError),
			service,
			httpserver.NewErrorProcessor(defaultErrorCode, defaultErrorMessage),
		)
		router := httpserver.MakeFastHTTPRouter(
			[]*httpserver.HandlerSettings{
				{
					Path:    httpserver2.URIPathClientGetExcelReport,
					Method:  httpserver2.HTTPMethodGetExcelReport,
					Handler: handlerFunc,
				},
			},
		)
		server := fasthttp.Server{
			Handler:            router.Handler,
			MaxRequestBodySize: maxRequestBodySize,
			ReadTimeout:        serverReadTimeout,
		}

		go server.ListenAndServe(serverBanService)
		defer server.Shutdown()
		time.Sleep(serverLaunchingWaitSleep)

		actualGetExcelReportResponse, err := client.GetExcelReport(context.Background(), &getExcelReportRequest)
		assert.Equal(t, err, expectedError)
		assert.Equal(t, expectedGetExcelReportResponse, actualGetExcelReportResponse)
	})
}

func intptr(v int) *int    { return &v }
func boolptr(v bool) *bool { return &v }
