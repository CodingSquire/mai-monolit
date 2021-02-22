package main

import (
	"github.com/valyala/fasthttp"

	"github.com/joeshaw/envdecode"

	httpserver_conference "github.com/CodingSquire/mai-monolit/pkg/conference-svc/httpserver"
	httpserver_thesis "github.com/CodingSquire/mai-monolit/pkg/thesis-svc/httpserver"
	httpserver_segment "github.com/CodingSquire/mai-monolit/pkg/segment-svc/httpserver"

	segment "github.com/CodingSquire/mai-monolit/pkg/segment-svc"
	conference "github.com/CodingSquire/mai-monolit/pkg/conference-svc"
	"github.com/CodingSquire/mai-monolit/pkg/dataservice"
	"github.com/CodingSquire/mai-monolit/pkg/httpserver"
	"github.com/CodingSquire/mai-monolit/pkg/logger"
	"github.com/CodingSquire/mai-monolit/pkg/logger/log"
	thesis "github.com/CodingSquire/mai-monolit/pkg/thesis-svc"
)

type configuration struct {
	Logger         logger.Config
	PortThesis     string `env:"PORT_THESIS,default=8081"`
	PortConference string `env:"PORT_CONFERENCE,default=8082"`
	PortSegment    string `env:"PORT_SEGMENT,default=8083"`
	Debug          bool   `env:"DEBUG,default=true"`
}

var (
	gitCommit = "undefined"
	gitBranch = "undefined"
)

func main() {
	cfg := &configuration{}
	if err := envdecode.StrictDecode(cfg); err != nil {
		log.Fatal().Err(err).Str("git_commit", gitCommit).Str("git_branch", gitBranch).Msg("Cannot decode config envs")
	}

	l := logger.NewLogger(&cfg.Logger)
	l.Info().Str("git_commit", gitCommit).Str("git_branch", gitBranch).Interface("config", cfg).Msg("The gathered config")

	serviceThesis := dataservice.NewService(httpserver.NewError)
	svcThesis := thesis.NewService(serviceThesis)
	svcThesis = thesis.NewLoggingMiddleware(svcThesis, l)

	routerThesis := httpserver_thesis.NewPreparedServer(svcThesis)
	portThesis := ":" + cfg.PortThesis
	err := fasthttp.ListenAndServe(portThesis, routerThesis.Handler)
	if err != nil {
		log.Fatal().Err(err).Msg("Crash service")
	}

	serviceConference := dataservice.NewService(httpserver.NewError)
	svcConference := conference.NewService(serviceConference)
	svcConference = conference.NewLoggingMiddleware(svcConference, l)

	routerConference := httpserver_conference.NewPreparedServer(svcConference)
	portConference := ":" + cfg.PortConference
	err2 := fasthttp.ListenAndServe(portConference, routerConference.Handler)
	if err2 != nil {
		log.Fatal().Err(err).Msg("Crash service")
	}

	serviceSegment := dataservice.NewService(httpserver.NewError)
	svcSegment := segment.NewService(serviceSegment)
	svcSegment = segment.NewLoggingMiddleware(svcSegment, l)

	routerSegment := httpserver_segment.NewPreparedServer(svcSegment)
	portSegment := ":" + cfg.PortSegment
	err3 := fasthttp.ListenAndServe(portSegment, routerSegment.Handler)
	if err3 != nil {
		log.Fatal().Err(err).Msg("Crash service")
	}

}
