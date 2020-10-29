package api

type Thesis struct{
	ID int					`json:"id"`
	AuthorID int			`json:"author_id"`
	SectionID int			`json:"section_id"`
	SubSectionsID int		`json:"subsection_id"`
	Originality float64		`json:"originality"`
	Subject string			`json:"subject"`
	Thesis string			`json:"thesis"`
	Fields string			`json:"custom_fields"`
	DateCreate string
	DateLastChange string
	Status string
	Changes string
}
