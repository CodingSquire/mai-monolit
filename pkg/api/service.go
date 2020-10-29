package api

type CreateThesisRequest struct{
	ID int					`json:"id"`
	AuthorID int			`json:"author_id"`
	SectionID int			`json:"section_id"`
	SubSectionsID int		`json:"subsection_id"`
	Originality float64		`json:"originality"`
	Subject string			`json:"subject"`
	Thesis string			`json:"thesis"`
	Fields string			`json:"custom_fields"`
}

type CreateThesisResponse struct{
	Result string			`json:"result"`
}

type ChangeThesisRequest struct{
	ID int					`json:"id"`
	AuthorID *int			`json:"author_id"`
	SectionID *int			`json:"section_id"`
	SubSectionsID *int		`json:"subsection_id"`
	Originality *float64		`json:"originality"`
	Subject *string			`json:"subject"`
	Thesis *string			`json:"thesis"`
	Fields *string			`json:"custom_fields"`
}

type ChangeThesisResponse struct{
	Result string			`json:"result"`
}

type GetThesisByFilterRequest struct{
	ID int					`json:"id"`
}

type GetThesisByFilterResponse struct{
	Thesis Thesis			`json:"thesis"`
}