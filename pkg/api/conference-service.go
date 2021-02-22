package api

import "encoding/json"

type CreateConferenceRequest struct {
	ID             int `json:"id"`
	DateCreate     *string
	DateLastChange *string
	Status         *string `json:"status"`
	Changes        *string
}

func (c *CreateConferenceRequest) Marshal() ([]byte, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type CreateConferenceResponse struct {
	Result string `json:"result"`
}

type ChangeConferenceRequest struct {
	ID             int `json:"id"`
	DateCreate     *string
	DateLastChange *string
	Status         *string `json:"status"`
	Changes        *string
}

func (c *ChangeConferenceRequest) Marshal() ([]byte, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type ChangeConferenceResponse struct {
	Result string `json:"result"`
}

type GetConferenceByFilterRequest struct {
	ID int `json:"id"`
}

func (g *GetConferenceByFilterRequest) Marshal() ([]byte, error) {
	body, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type GetConferenceByFilterResponse struct {
	Conference Conference `json:"conference"`
}
