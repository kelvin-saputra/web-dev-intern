package dto

type CreateCommentRequest struct {
	UserId int `json:"userId" binding:"required"`
	Title string `json:"title" binding:"required"`
	Body string `json:"body" binding:"required"`
}

type DeleteCommentRequest struct {
	Id []int `json:"id" binding:"required"`
}

type DeleteCommentByUserIdRequest struct {
	UserId int `json:"userId" binding:"required"`
}

type CommentResponse struct { 
	UserId int `json:"userId"`
	Id int `json:"id"`
	Title string `json:"title"`
	Body string `json:"body"`
}