package service

import (
	"context"
	"github.com/rs/zerolog"
	"time"

	"github.com/CodingSquire/mai-monolit/pkg/api"
)

// loggingMiddleware wraps Service and logs request information to the provided logger
type loggingMiddleware struct {
	svc    Service
	logger *api.Logger
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
	s.logger.Debug().
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
	s.logger.Debug().
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
	s.logger.Debug().
		Str("method", "GetThesisByFilter").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.GetThesisByFilter(ctx, request)
	return response, err
}

func (s *loggingMiddleware) wrap(ctx context.Context, err error) *zerolog.Event {
	lvl := s.logger.Info()
	if err != nil {
		lvl = s.logger.Error()
	}
	return lvl
}

// NewLoggingMiddleware ...
func NewLoggingMiddleware(svc Service,logger *api.Logger ) Service {
	return &loggingMiddleware{
		svc:    svc,
		logger: logger,
	}
}
