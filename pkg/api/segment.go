package api

type Segment struct{
	ID int					`json:"id"`
	DateCreate string
	DateLastChange string
	Status string
	Changes string
}
