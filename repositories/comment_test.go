package repositories_test

import (
	"testing"

	"github.com/TaisukeFujise/blog_api/models"
	"github.com/TaisukeFujise/blog_api/repositories"
)

func TestSelectCommentList(t *testing.T) {
	expectedNum := 2
	got, err := repositories.SelectCommentList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d comments\n", expectedNum, num)
	}
}

func TestInsertComment(t *testing.T) {
	comment := models.Comment{
		CommentID: 3,
		ArticleID: 1,
		Message:   "additional comment",
	}

	expectedCommentNum := 3
	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}
	if newComment.CommentID != expectedCommentNum {
		t.Errorf("new comment id is expected %d but got %d\n", expectedCommentNum, newComment.CommentID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where comment_id = ? and article_id = ? and message = ?
		`
		testDB.Exec(sqlStr, comment.CommentID, comment.ArticleID, comment.Message)
	})
}
