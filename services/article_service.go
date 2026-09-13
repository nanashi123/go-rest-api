// ArticleDetailHandler: 指定IDの記事をデータベースから取得する
// PostArticleHandler内: 記事データをデータベース内に挿入し、その値を返す
// ArticleListHandler内: クエリパラメータで指定されたページの記事一覧をデータベースから取得する
// PostNiceHandler内: 指定記事にいいねを+1する更新作業をデータベースに保存し、その結果を返す
// PostCommentHandler内: コメントデータをデータベース内に挿入し、その値を返す

package services

import (
	"database/sql"
	"errors"

	"github.com/Nanashi123/go-api-practice/apperrors"
	"github.com/Nanashi123/go-api-practice/models"
	"github.com/Nanashi123/go-api-practice/repositories"
)

// ArticleDetailHandler: 指定IDの記事をデータベースから取得する
func (s *MyAppService) GetArticleService(articleID int) (models.Article, error) {
	var article models.Article
	var commentList []models.Comment
	var articleGetErr, commentGetErr error

	// Article型とerror型を同時に扱う構造体
	type articleResult struct {
		article models.Article
		err     error
	}
	// articleResult型のチャネルを定義
	articleChan := make(chan articleResult)
	defer close(articleChan)

	// articleChanを通じて、SelectArticleDetail関数の結果を送信
	go func(ch chan<- articleResult, db *sql.DB, articleID int) {
		article, err := repositories.SelectArticleDetail(db, articleID)
		ch <- articleResult{article: article, err: err}
	}(articleChan, s.db, articleID)

	// Comment型のスライスとerror型を同時に扱う構造体
	type commentResult struct {
		commentList *[]models.Comment
		err         error
	}
	// commentResult型のチャネルを定義
	commentChan := make(chan commentResult)
	defer close(commentChan)

	// commentChanを通じて、SelectCommentList関数の結果を送信
	go func(ch chan<- commentResult, db *sql.DB, articleID int) {
		commentList, err := repositories.SelectCommentList(db, articleID)
		ch <- commentResult{commentList: &commentList, err: err}
	}(commentChan, s.db, articleID)

	for i := 0; i < 2; i++ {
		select {
		case ar := <-articleChan:
			article, articleGetErr = ar.article, ar.err
		case cr := <-commentChan:
			commentList, commentGetErr = *cr.commentList, cr.err
		}
	}

	if articleGetErr != nil {
		if errors.Is(articleGetErr, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(articleGetErr, "no data")
			return models.Article{}, err
		}
		err := apperrors.GetDataFailed.Wrap(articleGetErr, "failed to get article")
		return models.Article{}, err
	}

	if commentGetErr != nil {
		err := apperrors.GetDataFailed.Wrap(commentGetErr, "failed to get comment list")
		return models.Article{}, err
	}

	article.CommentList = append(article.CommentList, commentList...)

	return article, nil
}

// PostArticleHandler内: 記事データをデータベース内に挿入し、その値を返す
func (s *MyAppService) PostArticleService(article models.Article) (models.Article, error) {
	newArticle, err := repositories.InsertArticle(s.db, article)
	if err != nil {
		// 新しくMyAppError構造体を作り、各フィールドに適切な値をセットしていく
		//err = &MyAppError{ErrCode: [エラーコード], Message: [メッセージ内容], Err: err}
		err = apperrors.InsertDataFailed.Wrap(err, "failed to create article")
		return models.Article{}, err
	}
	return newArticle, nil
}

// ArticleListHandler内: クエリパラメータで指定されたページの記事一覧をデータベースから取得する
func (s *MyAppService) GetArticleListService(page int) ([]models.Article, error) {
	articleList, err := repositories.SelectArticleList(s.db, page)
	if err != nil {
		err = apperrors.GetDataFailed.Wrap(err, "failed to get article list")
		return nil, err
	}

	if len(articleList) == 0 {
		err := apperrors.NAData.Wrap(ErrNoData, "no data")
		return nil, err
	}

	return articleList, nil
}

// PostNiceHandler内: 指定記事にいいねを+1する更新作業をデータベースに保存し、その結果を返す
func (s *MyAppService) PostNiceService(article models.Article) (models.Article, error) {
	err := repositories.UpdateNiceNum(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err := apperrors.NiceTargetFailed.Wrap(err, "target article does not exist")
			return models.Article{}, err
		}
		err = apperrors.UpdateDataFailed.Wrap(err, "failed to update nice count")
		return models.Article{}, err
	}

	updatedArticle, err := repositories.SelectArticleDetail(s.db, article.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			err := apperrors.NAData.Wrap(err, "no data")
			return models.Article{}, err
		}
		err = apperrors.GetDataFailed.Wrap(err, "failed to get article after nice update")
		return models.Article{}, err
	}
	return updatedArticle, nil
}
