package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/TaisukeFujise/blog_api/apperrors"
	"github.com/TaisukeFujise/blog_api/controllers/services"
	"github.com/TaisukeFujise/blog_api/models"
)

type CommentController struct {
	service services.CommentServicer
}

func NewCommentController(s services.CommentServicer) *CommentController {
	return &CommentController{service: s}
}

// POST /comment: リクエストボディで受け取ったコメントを投稿する
func (c *CommentController) PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	var reqComment models.Comment

	if err := json.NewDecoder(req.Body).Decode(&reqComment); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	comment, err := c.service.PostCommentService(req.Context(), reqComment)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}
	json.NewEncoder(w).Encode(comment)
}
