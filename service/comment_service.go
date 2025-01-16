package service

import (
	"fmt"

	"primeskill/dto"
	"primeskill/model"
	"primeskill/config"
)


func CreateCommentService(request dto.CreateCommentRequest) (string, model.Comment, error) {
	comment := model.Comment{
		UserId: request.UserId,
		Title:  request.Title,
		Body:   request.Body,
	}

	
	result := config.DB.Create(&comment)
	if result.Error != nil {
		return "Error while sending the comment to the server", model.Comment{}, result.Error
	}
	return "", comment, nil
}

func CreateCommentBulkService(request []dto.CreateCommentRequest) (string, []model.Comment, error) {
	var comments []model.Comment

	for _, req := range request {
		comment := model.Comment{
			UserId: req.UserId,
			Title:  req.Title,
			Body:   req.Body,
		}
		comments = append(comments, comment)
	}

	result := config.DB.Create(&comments)
	if result.Error != nil {
		return "Error while sending the comment to the server", nil, result.Error
	}
	return "", comments, nil
}

// As I know could be improved the performance by using Caching Mechanism with Redis 
// But for now, I will use the basic way to retrieve the data from the database directly because it's not too much data
func GetAllCommentsService() (string, []model.Comment, error) {
	var comments []model.Comment
	result := config.DB.Find(&comments)
	if result.Error != nil {
		return "Error while retrieve comments from server", nil, result.Error
	}
	if len(comments) == 0 {
		return "No comments found in Database", nil, nil
	}
	return "", comments, nil
}

func GetCommentByIdService(id string) (string, model.Comment, error) {
	var comment model.Comment

	result := config.DB.First(&comment, id)
	if result.Error != nil {
		return fmt.Sprintf("Comment with ID %s not found", id), model.Comment{}, result.Error
	}
	return "", comment, nil
}

func GetCommentByUserIdService(id string) (string, []model.Comment, error) {
	var comments []model.Comment

	result := config.DB.Where("user_id = ?", id).Find(&comments)
	if result.Error != nil {
		return fmt.Sprintf("Error while retrieve user comment with ID %s", id), nil, result.Error
	}
	if len(comments) == 0 {
		return fmt.Sprintf("No comments found for user with ID %s", id), nil, nil
	}
	return "", comments, nil
}

func DeleteCommentByIdService(id string) (int, string, error) {
	var comment model.Comment

	result := config.DB.First(&comment, id)
	if result.Error != nil {
		return 404, fmt.Sprintf("Comment with ID %s not found", id), result.Error
	}

	result = config.DB.Delete(&comment)
	if result.Error != nil {
		return 500, fmt.Sprintf("Error while deleting comment with ID %s", id), result.Error
	}
	return 200, "", nil
}

func DeleteCommentByUserIdService(id string) (int, string, error) {
	var comments []model.Comment

	if err := config.DB.Where("user_id = ?", id).Find(&comments).Error; err != nil {
		return 500, fmt.Sprintf("Error while retrieve comments for user with ID %s", id), err
	}

	if len(comments) == 0 {
		return 404, fmt.Sprintf("No comments found for user with ID %s", id), nil
	}

	if err := config.DB.Delete(&comments).Error; err != nil {
		return 500, fmt.Sprintf("Error while deleting comments for user with ID %s", id), err
	}
	return 200, "Comment Deleted Successfully", nil
}

func DeleteAllCommentsService() (int, string, error) {
	var comments []model.Comment

	if err := config.DB.Find(&comments).Error; err != nil {
		return 500, "Error while retrieve all comments", err
	}
	if len(comments) == 0 {
		return 404, "No comments found", nil
	}

	if err := config.DB.Delete(&comments).Error; err != nil {
		return 500, "Error while deleting all comments", err
	}
	return 200, "All Comments Deleted Successfully", nil
}

func DeleteCommentByIdBulkService(request dto.DeleteCommentRequest) (int, string, error) {
	var comments []model.Comment

	if err := config.DB.Find(&comments, request.Id).Error; err != nil {
		return 500, "Error while retrieve comments", err
	}

	if len(comments) == 0 {
		return 404, "No comments found", nil
	}

	if err := config.DB.Delete(&comments).Error; err != nil {
		return 500, "Error while deleting comments", err
	}
	return 200, "Comments Deleted Successfully", nil
}