package segment_svc

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

func (s *loggingMiddleware) CreateSegment(ctx context.Context, request *api.CreateSegmentRequest)(response api.CreateSegmentResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(ctx, err).
			Str("method", "CreateSegment").
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
	response, err = s.svc.CreateSegment(ctx, request)
	return response, err
}

func (s *loggingMiddleware) ChangeSegment(ctx context.Context, request *api.ChangeSegmentRequest)(response api.ChangeSegmentResponse, err error){
	defer func(begin time.Time) {
		s.wrap(ctx, err).
			Str("method", "ChangeSegment").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	s.logger.Debug().
		Str("method", "ChangeSegment").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.ChangeSegment(ctx, request)
	return response, err
}

func (s *loggingMiddleware) GetSegmentByFilter(ctx context.Context, request *api.GetSegmentByFilterRequest)(response api.GetSegmentByFilterResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(ctx, err).
			Str("method", "GetSegmentByFilter").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	s.logger.Debug().
		Str("method", "GetSegmentByFilter").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.GetSegmentByFilter(ctx, request)
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
