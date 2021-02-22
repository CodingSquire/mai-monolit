package conference_svc

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

func (s *loggingMiddleware) CreateConference(ctx context.Context, request *api.CreateConferenceRequest)(response api.CreateConferenceResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(ctx, err).
			Str("method", "CreateConference").
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
	response, err = s.svc.CreateConference(ctx, request)
	return response, err
}

func (s *loggingMiddleware) ChangeConference(ctx context.Context, request *api.ChangeConferenceRequest)(response api.ChangeConferenceResponse, err error){
	defer func(begin time.Time) {
		s.wrap(ctx, err).
			Str("method", "ChangeConference").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	s.logger.Debug().
		Str("method", "ChangeConference").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.ChangeConference(ctx, request)
	return response, err
}

func (s *loggingMiddleware) GetConferenceByFilter(ctx context.Context, request *api.GetConferenceByFilterRequest)(response api.GetConferenceByFilterResponse, err error) {
	defer func(begin time.Time) {
		s.wrap(ctx, err).
			Str("method", "GetConferenceByFilter").
			Time("timestamp", time.Now()).
			Interface("response", response).
			Dur("elapsed", time.Since(begin)).
			Err(err).
			Msg("End.")
	}(time.Now())
	s.logger.Debug().
		Str("method", "GetConferenceByFilter").
		Time("timestamp", time.Now()).
		Interface("request", request).
		Msg("Start.")
	response, err = s.svc.GetConferenceByFilter(ctx, request)
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
