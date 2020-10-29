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
}

type service struct{
	data map[int] api.Thesis
	errorCreator errorCreator
}

func (s *service)SaveThesis(ctx context.Context,data api.Thesis)(result int, err error){
	s.data[data.ID]=data
	return data.ID,nil
}

func (s *service)GetThesis(ctx context.Context,id int)(result api.Thesis, err error){
	if d,ok:=s.data[id];!ok{
		return result, s.errorCreator(404,"wrong ID: %d",id)
	}else{
		return d,nil
	}
}


//NewService ...
func NewService(creator errorCreator) Service{
	return &service{
		data:         make(map[int]api.Thesis),
		errorCreator: creator,
	}
}