// articlesテーブルを操作する関数を実装

package repositories

import (
	"database/sql"

	"github.com/Nanashi123/go-api-practice/models"
)

// POST /article: リクエストボディで受け取った記事を投稿する
//   - 構造体 models.Articleで受け取って、それをデータベースに挿入する処理が必要
func InsertArticle(db *sql.DB, article models.Article) (models.Article, error) {
	const sqlStr = `
		insert into articles (title, contents, username, nice, created_at) values
		(?, ?, ?, 0, now())
	`

	//自分で書く
	result, err := db.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	if err != nil {
		return models.Article{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Article{}, err
	}

	newArticle := models.Article{
		ID:       int(id),
		Title:    article.Title,
		Contents: article.Contents,
		UserName: article.UserName,
		NiceNum:  0,
	}

	return newArticle, nil
}

// GET /article/list: クエリパラメータ page で指定されたページ(1ページに5個の記事を表示)に表示するための記事一覧を取得する
// - 指定された記事データをデータベースから取得して、それをmodels.Article構造体のスライス[]models.Articlesに詰めて返す処理が必要

func SelectArticleList(db *sql.DB, page int) ([]models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice, created_at
		from articles
		limit ?
		offset ?;

	`
	const articleNumPerPage = 5
	var offset int = (page - 1) * articleNumPerPage
	articleArray := make([]models.Article, 0)

	rows, err := db.Query(sqlStr, articleNumPerPage, offset)
	if err != nil {
		return articleArray, err
	}

	defer rows.Close()

	for rows.Next() {
		var article models.Article
		var createdTime sql.NullTime
		err := rows.Scan(
			&article.ID,
			&article.Title,
			&article.Contents,
			&article.UserName,
			&article.NiceNum,
			&createdTime,
		)

		if err != nil {
			return articleArray, err
		}

		if createdTime.Valid {
			article.CreatedAt = createdTime.Time
		}

		articleArray = append(articleArray, article)
	}

	if err := rows.Err(); err != nil {
		return articleArray, err
	}

	return articleArray, nil
}

// GET /article/{id}: 指定ＩＤの記事を取得する
// - 指定ＩＤの記事データをデータベースから取得して、それをmodels.Article構造体の形で返す処理が必要
// - 指定ＩＤの記事についてコメント一覧をデータベースから取得して、それをmodels.Comment構造体のスライス[]models.Commentに詰めて返す処理が必要

func SelectArticleDetail(db *sql.DB, articleID int) (models.Article, error) {
	const sqlStr = `
		select article_id, title, contents, username, nice, created_at
		from articles
		where article_id = ?;
	`
	row := db.QueryRow(sqlStr, articleID)
	if err := row.Err(); err != nil {
		return models.Article{}, err
	}

	var article models.Article
	var createdTime sql.NullTime

	err := row.Scan(
		&article.ID,
		&article.Title,
		&article.Contents,
		&article.UserName,
		&article.NiceNum,
		&createdTime,
	)

	if err != nil {
		return article, err
	}

	if createdTime.Valid {
		article.CreatedAt = createdTime.Time
	}

	return article, nil
}

//POST /article/nice: 記事にいいねをつける
// – 指定されたIDの記事のいいね数を+1するようにデータベースの中身を更新する処理が必要

func UpdateNiceNum(db *sql.DB, articleID int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	const sqlStr = `
		select nice
		from articles
		where article_id = ?;
	`

	row := tx.QueryRow(sqlStr, articleID)

	var nice int
	err = row.Scan(&nice)
	if err != nil {
		tx.Rollback()
		return err
	}

	const sqlUpdateNice = `
		update articles set nice = ? where article_id = ?
	`
	_, err = tx.Exec(sqlUpdateNice, nice+1, articleID)
	if err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil
}
