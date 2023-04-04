package controller

import (
	"book-author/common/helper"
	"book-author/dto"
	"book-author/entity"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/asaskevich/govalidator"

	"gorm.io/gorm"
)

type BookController struct {
	db *gorm.DB
}

func NewBookController(db *gorm.DB) *BookController {
	return &BookController{
		db: db,
	}
}

func (c *BookController) CreateBook(ctx *gin.Context) {
	bookRequest := dto.BookRequest{}

	err := ctx.ShouldBindJSON(&bookRequest)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	book := entity.Book{
		Isbn:        bookRequest.Isbn,
		Title:       bookRequest.Title,
		AuthorId:    bookRequest.AuthorId,
		PublisherId: bookRequest.PublisherId,
	}

	_, err = govalidator.ValidateStruct(&book)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	err = c.db.Create(&book).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	/*
		res := gin.H {
			"code": http.StatusCreated,
			"message": "Success create new book",
			"data": dto.BookCreateResponse{
				Id:          book.Id,
				Isbn:        book.Isbn,
				Title:       book.Title,
				AuthorId:    book.AuthorId,
				PublisherId: book.PublisherId,
				CreatedAt:   book.CreatedAt,
				UpdatedAt:   book.UpdatedAt,
			},
		}*/

	// helper.WriteJsonResponse(ctx, http.StatusCreated, res)
	helper.WriteJsonResponse(ctx, http.StatusCreated, dto.BookCreateResponse{
		Id:          book.Id,
		Isbn:        book.Isbn,
		Title:       book.Title,
		AuthorId:    book.AuthorId,
		PublisherId: book.PublisherId,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	})
}

func (c *BookController) GetAllBooks(ctx *gin.Context) {
	var books []entity.Book

	err := c.db.Preload("Publisher").Preload("Author").Find(&books).Error

	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	var response dto.BookGetResponse
	for _, book := range books {
		var authorData dto.AuthorBookResponse
		if book.Author != nil {
			authorData = dto.AuthorBookResponse{
				Id:    book.Author.Id,
				Name:  book.Author.Name,
				Email: book.Author.Email,
			}
		}

		var publisherData dto.PublisherBookResponse
		if book.Publisher != nil {
			publisherData = dto.PublisherBookResponse{
				Id:   book.Publisher.Id,
				Name: book.Publisher.Name,
				City: book.Publisher.City,
			}
		}

		bookResponse := dto.BookData{
			Id:          book.Id,
			Isbn:        book.Isbn,
			Title:       book.Title,
			AuthorId:    book.AuthorId,
			PublisherId: book.PublisherId,
			CreatedAt:   book.CreatedAt,
			UpdatedAt:   book.UpdatedAt,
			Author:      authorData,
			Publisher:   publisherData,
		}
		response.Books = append(response.Books, bookResponse)
	}
	helper.WriteJsonResponse(ctx, http.StatusOK, response)
}

func (c *BookController) UpdateBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	var bookUpdateRequest dto.BookRequest
	var book entity.Book

	err := ctx.ShouldBindJSON(&bookUpdateRequest)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	updateBook := entity.Book{
		Isbn:        bookUpdateRequest.Isbn,
		Title:       bookUpdateRequest.Title,
		AuthorId:    bookUpdateRequest.AuthorId,
		PublisherId: bookUpdateRequest.PublisherId,
	}

	err = c.db.First(&book, bookId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, "Data Not Found")
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Model(&book).Updates(updateBook).Error
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	response := dto.BookCreateResponse{
		Id:          book.Id,
		Isbn:        book.Isbn,
		Title:       book.Title,
		AuthorId:    book.AuthorId,
		PublisherId: book.PublisherId,
		CreatedAt:   book.CreatedAt,
		UpdatedAt:   book.UpdatedAt,
	}

	helper.WriteJsonResponse(ctx, http.StatusOK, response)
}

func (c *BookController) DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("bookId")
	var book entity.Book

	err := c.db.First(&book, bookId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, "Data Not Found")
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Delete(&book).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	helper.WriteJsonResponse(ctx, http.StatusOK, gin.H{
		"error":   false,
		"message": "Book successfully deleted",
	})
}
