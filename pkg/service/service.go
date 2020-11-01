package service

import (
	"context"
	"strconv"
	"time"

	"github.com/CodingSquire/mai-monolit/pkg/api"
)

// Service implements Service interface
type Service interface {
	CreateThesis(ctx context.Context, request *api.CreateThesisRequest)(response api.CreateThesisResponse, err error)
	ChangeThesis(ctx context.Context, request *api.ChangeThesisRequest)(response api.ChangeThesisResponse, err error)
	GetThesisByFilter(ctx context.Context, request *api.GetThesisByFilterRequest)(response api.GetThesisByFilterResponse, err error)
}

type dataService interface {
	SaveThesis(ctx context.Context, data api.Thesis)(result int, err error)
	GetThesis(ctx context.Context, id int)(result api.Thesis, err error)
}

type service struct{
	dataService dataService
}

func (s *service)CreateThesis(ctx context.Context, request *api.CreateThesisRequest)(response api.CreateThesisResponse, err error){
	thesis:=api.Thesis{
		ID:             request.ID,
		AuthorID:       request.AuthorID,
		SectionID:      request.SectionID,
		SubSectionsID:  request.SubSectionsID,
		Originality:    request.Originality,
		Subject:        request.Subject,
		Thesis:         request.Thesis,
		Fields:         request.Fields,
		DateCreate:     time.Now().String(),
		DateLastChange: time.Now().String(),
	}
	id,err:=s.dataService.SaveThesis(ctx,thesis)
	if err!=nil{
		return response, err
	}
	response.Result=strconv.Itoa(id)
	return response, err
}

func (s *service)ChangeThesis(ctx context.Context, request *api.ChangeThesisRequest)(response api.ChangeThesisResponse, err error){
	thesis,err:=s.dataService.GetThesis(ctx,request.ID)
	if err!=nil{
		return response,err
	}
	compareThesis(&thesis,request)
	id,err:=s.dataService.SaveThesis(ctx,thesis)
	if err!=nil{
		return response, err
	}
	response.Result=strconv.Itoa(id)
	return response, err
}

func (s *service)GetThesisByFilter(ctx context.Context, request *api.GetThesisByFilterRequest)(response api.GetThesisByFilterResponse, err error){
	thesis,err:=s.dataService.GetThesis(ctx,request.ID)
	if err!=nil{
		return response, err
	}
	response.Thesis=thesis
	return response, err
}

func compareThesis(old *api.Thesis,data *api.ChangeThesisRequest){
	if data.Fields!=nil{
		old.Fields=*data.Fields
	}
	if data.Thesis!=nil{
		old.Thesis=*data.Thesis
	}
	if data.Subject!=nil{
		old.Subject=*data.Subject
	}
	if data.Originality!=nil{
		old.Originality=*data.Originality
	}
	if data.SubSectionsID!=nil{
		old.SubSectionsID=*data.SubSectionsID
	}
	if data.SectionID!=nil{
		old.SectionID=*data.SectionID
	}
	if data.AuthorID!=nil{
		old.AuthorID=*data.AuthorID
	}
	old.DateLastChange=time.Now().String()
}

//NewService ...
func NewService(dataService dataService) Service{
	return &service{
		dataService:dataService,
	}
}
