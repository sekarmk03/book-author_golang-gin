package dto

import "time"

type PublisherRequest struct {
	Name string `json:"name"`
	City string `json:"city"`
}

type PublisherCreateResponse struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	City      string     `json:"city"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type PublisherGetResponse struct {
	Publishers []PublisherData `json:"authors"`
}

type PublisherData struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	City      string     `json:"city"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}
