package repositories_test

import (
	"testing"

	"github.com/Nanashi123/go-api-practice/repositories"
	"github.com/Nanashi123/go-api-practice/repositories/testdata"
)

func TestSelectCommentList(t *testing.T) {
	articleID := testdata.CommentTestData[0].ArticleID
	expectedcommentNum := len(testdata.CommentTestData)
	got, err := repositories.SelectCommentList(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	if commentNum := len(got); commentNum != expectedcommentNum {
		t.Errorf("want %d but got %d commentlist num\n", expectedcommentNum, commentNum)
	}
}

func TestInsertComment(t *testing.T) {
	comment := testdata.CommentTestData[0]

	newComment, err := repositories.InsertComment(testDB, comment)
	if err != nil {
		t.Fatal(err)
	}

	if newComment.CommentID == 0 {
		t.Error("new comment id is 0")
	}

	if newComment.ArticleID != comment.ArticleID {
		t.Errorf("comment.ArticleID: get %d but want %d\n", newComment.ArticleID, comment.ArticleID)
	}

	if newComment.Message != comment.Message {
		t.Errorf("comment.Message: get %s but want %s\n", newComment.Message, comment.Message)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from comments
			where comment_id = ?
		`
		_, err := testDB.Exec(sqlStr, newComment.CommentID)
		if err != nil {
			t.Error(err)
		}
	})
}
