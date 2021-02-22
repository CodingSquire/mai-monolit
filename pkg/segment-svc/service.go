package segment_svc

import (
	"context"
	"strconv"
	"time"

	"github.com/CodingSquire/mai-monolit/pkg/api"
)

// Service implements Service interface
type Service interface {
	CreateSegment(ctx context.Context, request *api.CreateSegmentRequest)(response api.CreateSegmentResponse, err error)
	ChangeSegment(ctx context.Context, request *api.ChangeSegmentRequest)(response api.ChangeSegmentResponse, err error)
	GetSegmentByFilter(ctx context.Context, request *api.GetSegmentByFilterRequest)(response api.GetSegmentByFilterResponse, err error)
}

type dataService interface {
	SaveSegment(ctx context.Context, data api.Segment)(result int, err error)
	GetSegment(ctx context.Context, id int)(result api.Segment, err error)
}

type service struct{
	dataService dataService
}

func (s *service)CreateSegment(ctx context.Context, request *api.CreateSegmentRequest)(response api.CreateSegmentResponse, err error){
	segment:=api.Segment{
		ID:             request.ID,
		DateCreate:     time.Now().String(),
		DateLastChange: time.Now().String(),
	}
	id,err:=s.dataService.SaveSegment(ctx,segment)
	if err!=nil{
		return response, err
	}
	response.Result=strconv.Itoa(id)
	return response, err
}

func (s *service)ChangeSegment(ctx context.Context, request *api.ChangeSegmentRequest)(response api.ChangeSegmentResponse, err error){
	segment,err:=s.dataService.GetSegment(ctx,request.ID)
	if err!=nil{
		return response,err
	}
	compareThesis(&segment,request)
	id,err:=s.dataService.SaveSegment(ctx,segment)
	if err!=nil{
		return response, err
	}
	response.Result=strconv.Itoa(id)
	return response, err
}

func (s *service)GetSegmentByFilter(ctx context.Context, request *api.GetSegmentByFilterRequest)(response api.GetSegmentByFilterResponse, err error){
	segment,err:=s.dataService.GetSegment(ctx,request.ID)
	if err!=nil{
		return response, err
	}
	response.Segment=segment
	return response, err
}

func compareThesis(old *api.Segment,data *api.ChangeSegmentRequest){
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
