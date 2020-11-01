package service

import (
	"context"
	"github.com/rs/zerolog"

	"time"

	"github.com/CodingSquire/mai-monolit/pkg/api"
	"github.com/CodingSquire/mai-monolit/pkg/logger"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	svc    Service
}

func (s *loggingMiddleware) CreateThesis(ctx context.Context, request *api.CreateThesisRequest)(response api.CreateThesisResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(ctx, err).
			Str("method", "CreateThesis").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	logger.Ctx(ctx).Debug().
		Str("method", "CreateThesis").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.CreateThesis(ctx, request)
	return response, err
}

func (s *loggingMiddleware) ChangeThesis(ctx context.Context, request *api.ChangeThesisRequest)(response api.ChangeThesisResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(ctx, err).
			Str("method", "ChangeThesis").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	logger.Ctx(ctx).Debug().
		Str("method", "ChangeThesis").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.ChangeThesis(ctx, request)
	return response, err
}

func (s *loggingMiddleware) GetThesisByFilter(ctx context.Context, request *api.GetThesisByFilterRequest)(response api.GetThesisByFilterResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(ctx, err).
			Str("method", "GetThesisByFilter").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	logger.Ctx(ctx).Debug().
		Str("method", "GetThesisByFilter").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.GetThesisByFilter(ctx, request)
	return response, err
}

func (s *loggingMiddleware) wrap(ctx context.Context, err error) *zerolog.Event {
	lvl := logger.Ctx(ctx).Debug()
	if err != nil {
		lvl = logger.Ctx(ctx).Error()
	}
	return lvl
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware( svc Service) Service {
	return &loggingMiddleware{
		svc:    svc,
	}
}
