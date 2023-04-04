package dto

import "time"

type BookRequest struct {
	Isbn        string `json:"isbn"`
	Title       string `json:"title"`
	AuthorId    uint   `json:"author_id"`
	PublisherId uint   `json:"publisher_id"`
}

type BookCreateResponse struct {
	Id          uint       `json:"id"`
	Isbn        string     `json:"isbn"`
	Title       string     `json:"title"`
	AuthorId    uint       `json:"author_id"`
	PublisherId uint       `json:"publisher_id"`
	CreatedAt   *time.Time `json:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at,omitempty"`
}

type BookGetResponse struct {
	Books []BookData `json:"books"`
}

type BookData struct {
	Id          uint                  `json:"id"`
	Isbn        string                `json:"isbn"`
	Title       string                `json:"title"`
	AuthorId    uint                  `json:"author_id"`
	PublisherId uint                  `json:"publisher_id"`
	Author      AuthorBookResponse    `json:"author"`
	Publisher   PublisherBookResponse `json:"publisher"`
	CreatedAt   *time.Time            `json:"created_at,omitempty"`
	UpdatedAt   *time.Time            `json:"updated_at,omitempty"`
}

type AuthorBookResponse struct {
	Id    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

type PublisherBookResponse struct {
	Id   uint   `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}
