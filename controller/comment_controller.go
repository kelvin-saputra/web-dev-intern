package controller

import (
	"fmt"
	"log"
	"net/http"
	"primeskill/dto"
	"primeskill/service"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var request dto.CreateCommentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Printf("Error while binding JSON: %s", err)
		return
	}

	msg, result, err := service.CreateCommentService(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}

	response := dto.CommentResponse{
		UserId: result.UserId,
		Id:     result.Id,
		Title:  result.Title,
		Body:   result.Body,
	}
	c.JSON(http.StatusCreated, response)
}

// For Testing only, to insert all example comment to the database
func CreateCommentBulkMethod(c *gin.Context) {
	var request []dto.CreateCommentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body request is not valid or empty"})
		log.Printf("Error while binding JSON: %s", err)
		return
	}

	msg, results, err := service.CreateCommentBulkService(request)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}

	var response []dto.CommentResponse
	for _, comment := range results {
		response = append(response, dto.CommentResponse{
			UserId: comment.UserId,
			Id:     comment.Id,
			Title:  comment.Title,
			Body:   comment.Body,
		})
	}

	c.JSON(http.StatusCreated, response)
}

func GetAllComments(c *gin.Context) {
	msg, result, err := service.GetAllCommentsService()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}
	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}

	var response []dto.CommentResponse
	for _, comment := range result {
		response = append(response, dto.CommentResponse{
			UserId: comment.UserId,
			Id:     comment.Id,
			Title:  comment.Title,
			Body:   comment.Body,
		})
	}

	c.JSON(http.StatusOK, response)
}

func GetCommentById(c *gin.Context) {
	msg, result, err := service.GetCommentByIdService(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}

	response := dto.CommentResponse{
		UserId: result.UserId,
		Id:     result.Id,
		Title:  result.Title,
		Body:   result.Body,
	}

	c.JSON(http.StatusOK, response)
}

func GetCommentByUserId(c *gin.Context) {
	msg, result, err := service.GetCommentByUserIdService(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
	}

	if len(result) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}

	var response []dto.CommentResponse
	for _, comment := range result {
		response = append(response, dto.CommentResponse{
			UserId: comment.UserId,
			Id:     comment.Id,
			Title:  comment.Title,
			Body:   comment.Body,
		})
	}

	c.JSON(http.StatusOK, response)
}

func DeleteCommentById(c *gin.Context) {
	status, msg, err := service.DeleteCommentByIdService(c.Param("id"))

	if status == 404 {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Comment with ID %s deleted", c.Param("id"))})
}

func DeleteCommentByUserId(c *gin.Context) {
	status, msg, err := service.DeleteCommentByUserIdService(c.Param("id"))

	if err != nil && status == 500 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}

	if status == 404 {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func DeleteAllComments(c *gin.Context) {
	status, msg, err := service.DeleteAllCommentsService()
	if err != nil && status == 500 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}
	if status == 404 {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": msg})
}

func DeleteCommentByIdBulkMethod(c *gin.Context) {
	var request dto.DeleteCommentRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Body request is not valid or empty"})
		log.Printf("Error while binding JSON: %s", err)
		return
	}
	status, msg, err := service.DeleteCommentByIdBulkService(request)

	if err != nil && status == 500 {
		c.JSON(http.StatusInternalServerError, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}

	if status == 404 {
		c.JSON(http.StatusNotFound, gin.H{"error": msg})
		log.Printf("%s: %s", msg, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": msg})
}
