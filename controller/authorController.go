package controller

import (
	"book-author/common/helper"
	"book-author/dto"
	"book-author/entity"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type AuthorController struct {
	db *gorm.DB
}

func NewAuthorController(db *gorm.DB) *AuthorController {
	return &AuthorController{
		db: db,
	}
}

func (c *AuthorController) CreateAuthor(ctx *gin.Context) {
	authorRequest := dto.AuthorRequest{}

	err := ctx.ShouldBindJSON(&authorRequest)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	author := entity.Author{
		Name:  authorRequest.Name,
		Email: authorRequest.Email,
	}

	_, err = govalidator.ValidateStruct(&author)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	err = c.db.Create(&author).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	helper.WriteJsonResponse(ctx, http.StatusCreated, dto.AuthorCreateResponse{
		Id:        author.Id,
		Name:      author.Name,
		Email:     author.Email,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	})
}

func (c *AuthorController) GetAllAuthors(ctx *gin.Context) {
	var authors []entity.Author

	err := c.db.Find(&authors).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	var response dto.AuthorGetResponse
	for _, author := range authors {
		authorResponse := dto.AuthorData{
			Id:        author.Id,
			Name:      author.Name,
			Email:     author.Email,
			CreatedAt: author.CreatedAt,
			UpdatedAt: author.UpdatedAt,
		}
		response.Authors = append(response.Authors, authorResponse)
	}
	helper.WriteJsonResponse(ctx, http.StatusOK, response)
}

func (c *AuthorController) UpdateAuthor(ctx *gin.Context) {
	authorId := ctx.Param("authorId")
	var authorUpdateRequest dto.AuthorRequest
	var author entity.Author

	err := ctx.ShouldBindJSON(&authorUpdateRequest)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	updateAuthor := entity.Author{
		Name:  authorUpdateRequest.Name,
		Email: authorUpdateRequest.Email,
	}

	err = c.db.First(&author, authorId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, "Data Not Found")
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Model(&author).Updates(updateAuthor).Error
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	response := dto.AuthorCreateResponse{
		Id:        author.Id,
		Name:      author.Name,
		Email:     author.Email,
		CreatedAt: author.CreatedAt,
		UpdatedAt: author.UpdatedAt,
	}

	helper.WriteJsonResponse(ctx, http.StatusOK, response)
}

func (c *AuthorController) DeleteAuthor(ctx *gin.Context) {
	authorId := ctx.Param("authorId")
	var author entity.Author

	err := c.db.First(&author, authorId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, "Data Not Found")
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Delete(&author).Error
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
		"message": "Author successfully deleted",
	})
}
