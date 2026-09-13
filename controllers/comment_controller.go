package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/Nanashi123/go-api-practice/apperrors"
	"github.com/Nanashi123/go-api-practice/controllers/services"
	"github.com/Nanashi123/go-api-practice/models"
)

type CommentContoller struct {
	service services.CommentServicer
}

func NewCommentContoller(s services.CommentServicer) *CommentContoller {
	return &CommentContoller{service: s}
}

// /commentのハンドラ
func (c *CommentContoller) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "post comment bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := c.service.PostCommentService(reqComment)
	if err != nil {
		http.Error(w, "fail to internal exec\n", http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ResponseBodyEncodeFailed.Wrap(err, "failed to encode comment response")
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
}
