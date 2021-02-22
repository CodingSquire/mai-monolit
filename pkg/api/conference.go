package api

type Conference struct{
	ID int					`json:"id"`
	DateCreate string
	DateLastChange string
	Status string
	Changes string
}
