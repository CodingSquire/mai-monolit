package conference_svc

import (
	"context"
	"strconv"
	"time"

	"github.com/CodingSquire/mai-monolit/pkg/api"
)

// Service implements Service interface
type Service interface {
	CreateConference(ctx context.Context, request *api.CreateConferenceRequest)(response api.CreateConferenceResponse, err error)
	ChangeConference(ctx context.Context, request *api.ChangeConferenceRequest)(response api.ChangeConferenceResponse, err error)
	GetConferenceByFilter(ctx context.Context, request *api.GetConferenceByFilterRequest)(response api.GetConferenceByFilterResponse, err error)
}

type dataService interface {
	SaveConference(ctx context.Context, data api.Conference)(result int, err error)
	GetConference(ctx context.Context, id int)(result api.Conference, err error)
}

type service struct{
	dataService dataService
}

func (s *service)CreateConference(ctx context.Context, request *api.CreateConferenceRequest)(response api.CreateConferenceResponse, err error){
	conference:=api.Conference{
		ID:             request.ID,
		DateCreate:     time.Now().String(),
		DateLastChange: time.Now().String(),
	}
	id,err:=s.dataService.SaveConference(ctx,conference)
	if err!=nil{
		return response, err
	}
	response.Result=strconv.Itoa(id)
	return response, err
}

func (s *service)ChangeConference(ctx context.Context, request *api.ChangeConferenceRequest)(response api.ChangeConferenceResponse, err error){
	conference,err:=s.dataService.GetConference(ctx,request.ID)
	if err!=nil{
		return response,err
	}
	compareThesis(&conference,request)
	id,err:=s.dataService.SaveConference(ctx,conference)
	if err!=nil{
		return response, err
	}
	response.Result=strconv.Itoa(id)
	return response, err
}

func (s *service)GetConferenceByFilter(ctx context.Context, request *api.GetConferenceByFilterRequest)(response api.GetConferenceByFilterResponse, err error){
	conference,err:=s.dataService.GetConference(ctx,request.ID)
	if err!=nil{
		return response, err
	}
	response.Conference=conference
	return response, err
}

func compareThesis(old *api.Conference,data *api.ChangeConferenceRequest){
	if data.Status!=nil{
		old.Status=*data.Status
	}

	old.DateLastChange=time.Now().String()
}

//NewService ...
func NewService(dataService dataService) Service{
	return &service{
		dataService:dataService,
	}
}
