package routes

import (
	"primeskill/controller"

	"github.com/gin-gonic/gin"
)

func CommentRoute(router *gin.RouterGroup) {
	// [POST]
	CommentRouter := router.Group("/comment")
	CommentRouter.POST("/create", controller.CreateComment)
	// For Testing only, to insert all example comment to the database
	CommentRouter.POST("/create/bulk", controller.CreateCommentBulkMethod)

	// [GET]
	CommentRouter.GET("", controller.GetAllComments)
	CommentRouter.GET("/:id", controller.GetCommentById)
	CommentRouter.GET("/user/:id", controller.GetCommentByUserId)

	// [DELETE]
	CommentRouter.DELETE("/delete/:id", controller.DeleteCommentById)
	// In case comment needs to be deleted by the user id
	CommentRouter.DELETE("/delete/user/:id", controller.DeleteCommentByUserId)
	// In case all comments need to be deleted cause of some reason
	CommentRouter.DELETE("/delete/all", controller.DeleteAllComments)
	// In case multiple comments need to be deleted
	CommentRouter.DELETE("/delete/bulk", controller.DeleteCommentByIdBulkMethod)
}