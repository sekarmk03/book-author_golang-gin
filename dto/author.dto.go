package dto

import "time"

type AuthorRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type AuthorCreateResponse struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type AuthorGetResponse struct {
	Authors []AuthorData `json:"authors"`
}

type AuthorData struct {
	Id        uint       `json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

/*
Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InNla2FybWFkdTk5QGdtYWlsLmNvbSIsImlkIjoxfQ.s3HP2N5Yc1TkDHfADQXRW87_Kb_sqyNQEkYr1LDDnzs
*/
