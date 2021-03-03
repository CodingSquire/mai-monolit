package api

import "encoding/json"

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

func (c *CreateThesisRequest) Marshal() ([]byte, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return body, nil
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

func (c *ChangeThesisRequest) Marshal() ([]byte, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type ChangeThesisResponse struct{
	Result string			`json:"result"`
}

type GetThesisByFilterRequest struct{
	ID int					`json:"id"`
}

func (g *GetThesisByFilterRequest) Marshal() ([]byte, error) {
	body, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type GetThesisByFilterResponse struct{
	Thesis Thesis			`json:"thesis"`
}