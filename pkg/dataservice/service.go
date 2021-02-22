package dataservice

import (
	"context"
	"github.com/CodingSquire/mai-monolit/pkg/api"
)

type errorCreator func(status int, format string, v ...interface{}) error


// Service implements Service interface
type Service interface {
	SaveThesis(ctx context.Context, data api.Thesis)(result int, err error)
	GetThesis(ctx context.Context, id int)(result api.Thesis, err error)
	SaveConference(ctx context.Context, data api.Conference)(result int, err error)
	GetConference(ctx context.Context, id int)(result api.Conference, err error)
	SaveSegment(ctx context.Context, data api.Segment)(result int, err error)
	GetSegment(ctx context.Context, id int)(result api.Segment, err error)
}

type service struct{
	dataThesis map[int] api.Thesis
	dataConference map[int] api.Conference
	dataSegment map[int] api.Segment
	errorCreator errorCreator
}

func (s *service)SaveThesis(ctx context.Context,data api.Thesis)(result int, err error){
	s.dataThesis[data.ID]=data
	return data.ID,nil
}

func (s *service)GetThesis(ctx context.Context,id int)(result api.Thesis, err error){
	if d,ok:=s.dataThesis[id];!ok{
		return result, s.errorCreator(404,"wrong ID: %d",id)
	}else{
		return d,nil
	}
}

func (s *service)SaveConference(ctx context.Context,data api.Conference)(result int, err error){
	s.dataConference[data.ID]=data
	return data.ID,nil
}

func (s *service)GetConference(ctx context.Context,id int)(result api.Conference, err error){
	if d,ok:=s.dataConference[id];!ok{
		return result, s.errorCreator(404,"wrong ID: %d",id)
	}else{
		return d,nil
	}
}

func (s *service)SaveSegment(ctx context.Context,data api.Segment)(result int, err error){
	s.dataSegment[data.ID]=data
	return data.ID,nil
}

func (s *service)GetSegment(ctx context.Context,id int)(result api.Segment, err error){
	if d,ok:=s.dataSegment[id];!ok{
		return result, s.errorCreator(404,"wrong ID: %d",id)
	}else{
		return d,nil
	}
}



//NewService ...
func NewService(creator errorCreator) Service{
	return &service{
		dataThesis:         make(map[int]api.Thesis),
		dataConference:         make(map[int]api.Conference),
		dataSegment:         make(map[int]api.Segment),
		errorCreator: creator,
	}
}