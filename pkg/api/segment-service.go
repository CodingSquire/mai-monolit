package api

import "encoding/json"

type CreateSegmentRequest struct {
	ID             int `json:"id"`
	DateCreate     *string
	DateLastChange *string
	Status         *string `json:"status"`
	Changes        *string
}

func (c *CreateSegmentRequest) Marshal() ([]byte, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type CreateSegmentResponse struct {
	Result string `json:"result"`
}

type ChangeSegmentRequest struct {
	ID             int `json:"id"`
	DateCreate     *string
	DateLastChange *string
	Status         *string `json:"status"`
	Changes        *string
}

func (c *ChangeSegmentRequest) Marshal() ([]byte, error) {
	body, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type ChangeSegmentResponse struct {
	Result string `json:"result"`
}

type GetSegmentByFilterRequest struct {
	ID int `json:"id"`
}

func (g *GetSegmentByFilterRequest) Marshal() ([]byte, error) {
	body, err := json.Marshal(g)
	if err != nil {
		return nil, err
	}
	return body, nil
}

type GetSegmentByFilterResponse struct {
	Segment Segment `json:"segment"`
}
