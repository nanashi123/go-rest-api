package repositories

import (
	"database/sql"

	"github.com/Nanashi123/go-api-practice/models"
)

// 指定IDの記事についたコメント一覧を取得する関数
// -> 取得したコメントデータと、発生したエラーを返り値にする
func SelectCommentList(db *sql.DB, articleID int) ([]models.Comment, error) {
	const sqlStr = `
		select comment_id, article_id, message, created_at
		from comments
		where article_id = ?;
	`

	commentArray := make([]models.Comment, 0)

	rows, err := db.Query(sqlStr, articleID)
	if err != nil {
		return commentArray, err
	}

	defer rows.Close()

	for rows.Next() {
		var comment models.Comment
		var createdTime sql.NullTime
		err := rows.Scan(&comment.CommentID, &comment.ArticleID, &comment.Message, &createdTime)

		if err != nil {
			return commentArray, err
		}

		if createdTime.Valid {
			comment.CreatedAt = createdTime.Time
		}

		commentArray = append(commentArray, comment)
	}

	if err := rows.Err(); err != nil {
		return commentArray, err
	}

	return commentArray, nil
}

//POST /comment: リクエストボディで受け取ったコメントを投稿する
// – 構造体models.Commentを受け取って、それをデータベースに挿入する処理が必要

func InsertComment(db *sql.DB, comment models.Comment) (models.Comment, error) {
	const sqlStr = `
		insert into comments (article_id, message, created_at) values
		(?, ?, now());
	`

	result, err := db.Exec(sqlStr, comment.ArticleID, comment.Message)
	if err != nil {
		return models.Comment{}, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return models.Comment{}, err
	}

	newComment := models.Comment{
		CommentID: int(id),
		ArticleID: comment.ArticleID,
		Message:   comment.Message,
	}

	return newComment, nil
}
