package main

import (
	"os"

	"github.com/valyala/fasthttp"

	httpserver2 "github.com/CodingSquire/mai-monolit/pkg/thesis-svc/httpserver"

	"github.com/joeshaw/envdecode"

	"github.com/CodingSquire/mai-monolit/pkg/dataservice"
	"github.com/CodingSquire/mai-monolit/pkg/httpserver"
	"github.com/CodingSquire/mai-monolit/pkg/logger"
	"github.com/CodingSquire/mai-monolit/pkg/logger/log"
	thesis "github.com/CodingSquire/mai-monolit/pkg/thesis-svc"
)

type configuration struct {
	Logger logger.Config
	Port   string `env:"PORT,default=true"`
	Debug  bool   `env:"DEBUG,default=true"`
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

	service := dataservice.NewService(httpserver.NewError)
	svc := thesis.NewService(service)
	svc = thesis.NewLoggingMiddleware(svc, l)

	router := httpserver2.NewPreparedServer(svc)

	port := ":" + os.Getenv("PORT")
	err := fasthttp.ListenAndServe(port, router.Handler)
	if err != nil {
		log.Fatal().Err(err).Msg("Crash service")
	}

}

