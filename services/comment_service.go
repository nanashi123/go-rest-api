package services

import (
	"github.com/Nanashi123/go-api-practice/apperrors"
	"github.com/Nanashi123/go-api-practice/models"
	"github.com/Nanashi123/go-api-practice/repositories"
)

// PostCommentHandler内: コメントデータをデータベース内に挿入し、その値を返す
func (s *MyAppService) PostCommentService(comment models.Comment) (models.Comment, error) {
	newComment, err := repositories.InsertComment(s.db, comment)
	if err != nil {
		err := apperrors.InsertDataFailed.Wrap(err, "failed to create comment")
		return models.Comment{}, err
	}
	return newComment, nil
}
