package controllers_test

import (
	"testing"

	"github.com/Nanashi123/go-api-practice/controllers"
	"github.com/Nanashi123/go-api-practice/controllers/testdata"
)

// 1. テストで使うリソース（コントローラ構造体）を用意
var aCon *controllers.ArticleContoller

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleContoller(ser)

	m.Run()
}
