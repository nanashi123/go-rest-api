package testdata

import "github.com/Nanashi123/go-api-practice/models"

var ArticleTestData = []models.Article{
	models.Article{
		ID:       1,
		Title:    "firstPost",
		Contents: "This is my first blog",
		UserName: "saki",
		NiceNum:  3,
	},
	models.Article{
		ID:       2,
		Title:    "2nd",
		Contents: "Second blog post",
		UserName: "saki",
		NiceNum:  4,
	},
}

var CommentTestData = []models.Comment{
	models.Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment test",
	},
	models.Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "test comment test 2",
	},
}
