package main

import (
	httpserver2 "github.com/CodingSquire/mai-monolit/pkg/service/httpserver"
	"github.com/valyala/fasthttp"
	"os"

	"github.com/joeshaw/envdecode"

	"github.com/CodingSquire/mai-monolit/pkg/dataservice"
	"github.com/CodingSquire/mai-monolit/pkg/httpserver"
	"github.com/CodingSquire/mai-monolit/pkg/logger"
	"github.com/CodingSquire/mai-monolit/pkg/logger/log"
	"github.com/CodingSquire/mai-monolit/pkg/service"
)

type configuration struct {
	Logger                logger.Config
	Port  string `env:"PORT,default=true"`
	Debug bool   `env:"DEBUG,default=true"`
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
	//ctx := l.WithContext(context.Background())
	l.Info().Str("git_commit", gitCommit).Str("git_branch", gitBranch).Interface("config", cfg).Msg("The gathered config")



	dataservice:=dataservice.NewService(httpserver.NewError)
	svc := service.NewService(dataservice)
	svc =service.NewLoggingMiddleware(svc,l)

	//
	//
	//requestCreate:=api.CreateThesisRequest{
	//	ID:            11,
	//	AuthorID:      1,
	//	SectionID:     13,
	//	SubSectionsID: 14,
	//	Originality:   15,
	//	Subject:       "sdfsdf",
	//	Thesis:        "sdfsdfsdfsdf",
	//	Fields:        "sdfsdfsdfsdfsdfsdfsdf",
	//}
	//svc.CreateThesis(ctx,&requestCreate)
	//
	//
	//svc.GetThesisByFilter(ctx,&api.GetThesisByFilterRequest{ID: 11})
	//
	//
	//requestChange:=api.ChangeThesisRequest{
	//	ID:            11,
	//	AuthorID:      intptr(12),
	//	Subject:       strptr("New Autor"),
	//}
	//
	//svc.ChangeThesis(ctx,&requestChange)
	//
	//
	//svc.GetThesisByFilter(ctx,&api.GetThesisByFilterRequest{ID: 11})


	router := httpserver2.NewPreparedServer(svc)



		port:= ":"+os.Getenv("PORT")
		err:=fasthttp.ListenAndServe(port, router.Handler)
		if err!=nil{
			log.Fatal().Err(err).Msg("Crash service")
		}


}

//func intptr(v int) *int    { return &v }
//func strptr(v string) *string    { return &v }
