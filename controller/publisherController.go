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

type PublisherController struct {
	db *gorm.DB
}

func NewPublisherController(db *gorm.DB) *PublisherController {
	return &PublisherController{
		db: db,
	}
}

func (c *PublisherController) CreatePublisher(ctx *gin.Context) {
	publisherRequest := dto.PublisherRequest{}

	err := ctx.ShouldBindJSON(&publisherRequest)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	publisher := entity.Publisher{
		Name: publisherRequest.Name,
		City: publisherRequest.City,
	}

	_, err = govalidator.ValidateStruct(&publisher)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	err = c.db.Create(&publisher).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	helper.WriteJsonResponse(ctx, http.StatusCreated, dto.PublisherCreateResponse{
		Id:        publisher.Id,
		Name:      publisher.Name,
		City:      publisher.City,
		CreatedAt: publisher.CreatedAt,
		UpdatedAt: publisher.UpdatedAt,
	})
}

func (c *PublisherController) GetAllPublishers(ctx *gin.Context) {
	var publishers []entity.Publisher

	err := c.db.Find(&publishers).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	var response dto.PublisherGetResponse
	for _, publisher := range publishers {
		publisherResponse := dto.PublisherData{
			Id:        publisher.Id,
			Name:      publisher.Name,
			City:      publisher.City,
			CreatedAt: publisher.CreatedAt,
			UpdatedAt: publisher.UpdatedAt,
		}
		response.Publishers = append(response.Publishers, publisherResponse)
	}
	helper.WriteJsonResponse(ctx, http.StatusOK, response)
}

func (c *PublisherController) UpdatePublisher(ctx *gin.Context) {
	publisherId := ctx.Param("publisherId")
	var publisherUpdateRequest dto.PublisherRequest
	var publisher entity.Publisher

	err := ctx.ShouldBindJSON(&publisherUpdateRequest)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	updatePublisher := entity.Publisher{
		Name: publisherUpdateRequest.Name,
		City: publisherUpdateRequest.City,
	}

	err = c.db.First(&publisher, publisherId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, "Data Not Found")
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Model(&publisher).Updates(updatePublisher).Error
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	response := dto.PublisherCreateResponse{
		Id:        publisher.Id,
		Name:      publisher.Name,
		City:      publisher.City,
		CreatedAt: publisher.CreatedAt,
		UpdatedAt: publisher.UpdatedAt,
	}

	helper.WriteJsonResponse(ctx, http.StatusOK, response)
}

func (c *PublisherController) DeletePublisher(ctx *gin.Context) {
	publisherId := ctx.Param("publisherId")
	var publisher entity.Publisher

	err := c.db.First(&publisher, publisherId).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, "Data Not Found")
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	err = c.db.Delete(&publisher).Error
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
		"message": "Publisher successfully deleted",
	})
}
