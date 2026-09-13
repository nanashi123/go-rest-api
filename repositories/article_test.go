package repositories_test

import (
	"testing"

	"github.com/Nanashi123/go-api-practice/models"
	"github.com/Nanashi123/go-api-practice/repositories"
	"github.com/Nanashi123/go-api-practice/repositories/testdata"
)

func TestSelectArticleDetail(t *testing.T) {
	// 1.「テストケース名」と「テストデータ」セットのスライスを作成
	tests := []struct {
		testTitle string         //テストのタイトル
		expected  models.Article //テストで期待する値
	}{
		{
			// 記事 ID1 番のテストデータ
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		},
		{
			// 記事 ID2 番のテストデータ
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	// 2. 1で作ったものをfor文で回す
	for _, test := range tests {
		// testを使ってサブテストを書く
		t.Run(test.testTitle, func(t *testing.T) {
			// test.expected.IDの記事IDのデータを取得して、結果をgotに格納
			got, err := repositories.SelectArticleDetail(testDB, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			// テスト対象の関数の結果と1の結果を比べる
			if got.ID != test.expected.ID {
				// 不一致だった場合にはテスト失敗
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
				// 不一致だった場合にはテスト失敗
				t.Errorf("Title: get %s but want %s\n", got.Title, test.expected.Title)
			}
			if got.Contents != test.expected.Contents {
				t.Errorf("Content: get %s but want %s\n", got.Contents, test.expected.Contents)
			}
			if got.UserName != test.expected.UserName {
				t.Errorf("UserName: get %s but want %s\n", got.UserName, test.expected.UserName)
			}
			if got.NiceNum != test.expected.NiceNum {
				t.Errorf("NiceNum: get %d but want %d\n", got.NiceNum, test.expected.NiceNum)
			}
		})
	}

	// t.Fatal もt.Errorfも実行されずに関数が終わった場合にはテスト成功
}

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)
	got, err := repositories.SelectArticleList(testDB, 1)
	if err != nil {
		t.Fatal(err)
	}

	// あとで理解する
	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

// got = 関数を実行して実際に返ってきた結果
// expected = この結果になってほしい、という期待値

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}

	// go test時にidが増え続けるのでコメントアウト
	//expectedArticleNum := 3
	newArticle, err := repositories.InsertArticle(testDB, article)
	if err != nil {
		t.Fatal(err)
	}

	if newArticle.ID == 0 {
		//t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
		t.Errorf("new article id is 0")
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ? and article_id = ?
		`
		_, err := testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName, newArticle.ID)
		if err != nil {
			t.Error(err)
		}
	})
}

func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	before, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(func() {
		const sqlStr = `
			update articles
			set nice = ?
			where article_id = ?
			`

		_, err := testDB.Exec(sqlStr, before.NiceNum, articleID)
		if err != nil {
			t.Error(err)
		}
	})

	err = repositories.UpdateNiceNum(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repositories.SelectArticleDetail(testDB, articleID)
	if err != nil {
		t.Fatal(err)
	}

	if after.NiceNum-before.NiceNum != 1 {
		t.Errorf("fail to update nice num: after %d before %d", after.NiceNum, before.NiceNum)
	}
}
