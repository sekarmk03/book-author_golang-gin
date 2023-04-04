package routers

import (
	"book-author/common/database"
	"book-author/controller"
	"book-author/middleware"

	"github.com/gin-gonic/gin"
)

func StartApp() *gin.Engine {
	db := database.ConnectDB()
	router := gin.Default()
	user := controller.NewUserController(db)
	book := controller.NewBookController(db)
	author := controller.NewAuthorController(db)
	publisher := controller.NewPublisherController(db)

	userGroup := router.Group("/auth")
	{
		userGroup.POST("/login", user.Login)
		userGroup.POST("/register", user.CreateUser)
	}

	bookGroup := router.Group("/books")
	{
		bookGroup.GET("/", book.GetAllBooks)
		bookGroup.POST("/", middleware.Auth(), book.CreateBook)
		bookGroup.PUT("/:bookId", middleware.Auth(), book.UpdateBook)
		bookGroup.DELETE("/:bookId", middleware.Auth(), book.DeleteBook)
	}

	authorGroup := router.Group("/authors")
	{
		authorGroup.GET("/", author.GetAllAuthors)
		authorGroup.POST("/", middleware.Auth(), author.CreateAuthor)
		authorGroup.PUT("/authorId", middleware.Auth(), author.UpdateAuthor)
		authorGroup.DELETE("/authorId", middleware.Auth(), author.DeleteAuthor)
	}

	publisherGroup := router.Group("/publishers")
	{
		publisherGroup.GET("/", publisher.GetAllPublishers)
		publisherGroup.POST("/", middleware.Auth(), publisher.CreatePublisher)
		publisherGroup.PUT("/", middleware.Auth(), publisher.UpdatePublisher)
		publisherGroup.DELETE("/", middleware.Auth(), publisher.DeletePublisher)
	}

	return router
}
