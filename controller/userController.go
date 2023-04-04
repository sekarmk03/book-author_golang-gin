package controller

import (
	"book-author/common/auth"
	"book-author/common/helper"
	"book-author/dto"
	"book-author/entity"
	"fmt"
	"net/http"

	"github.com/asaskevich/govalidator"

	"github.com/gin-gonic/gin"

	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	user := entity.User{}

	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	_, err = govalidator.ValidateStruct(&user)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	err = c.db.Create(&user).Error
	if err != nil {
		if err.Error() == gorm.ErrRecordNotFound.Error() {
			helper.NotFoundResponse(ctx, err.Error())
			return
		}
		helper.InternalServerJsonResponse(ctx, err.Error())
		return
	}

	helper.WriteJsonResponse(ctx, http.StatusCreated, dto.UserCreateResponse{
		Id:       user.Id,
		Username: user.Username,
		Email:    user.Email,
	})
}

func (c *UserController) Login(ctx *gin.Context) {
	user := entity.User{}
	err := ctx.ShouldBindJSON(&user)
	if err != nil {
		helper.BadRequestResponse(ctx, err.Error())
		return
	}

	password := user.Password
	err = c.db.Debug().Where("email = ?", user.Email).Take(&user).Error
	if err != nil {
		helper.WriteJsonResponse(ctx, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "email / password is not match",
		})
		return
	}

	comparePass := auth.ComparePassword(user.Password, password)
	if !comparePass {
		helper.WriteJsonResponse(ctx, http.StatusUnauthorized, gin.H{
			"error":   true,
			"message": "email / password is not match",
		})
		return
	}
	token := auth.GenerateToken(user.Id, user.Email)
	fmt.Println(token)
	ctx.JSON(http.StatusOK, dto.UserLoginResponse{
		Token: token,
	})
}
