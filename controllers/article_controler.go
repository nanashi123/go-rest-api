package controllers

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/Nanashi123/go-api-practice/apperrors"
	"github.com/Nanashi123/go-api-practice/common"
	"github.com/Nanashi123/go-api-practice/controllers/services"
	"github.com/Nanashi123/go-api-practice/models"
	"github.com/gorilla/mux"
)

// Article用のコントローラ構造体
type ArticleContoller struct {
	service services.ArticleServicer
}

// コンストラクタ関数
func NewArticleContoller(s services.ArticleServicer) *ArticleContoller {
	return &ArticleContoller{service: s}
}

// /helloのハンドラ
func (c *ArticleContoller) HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}

// /articleのハンドラ
func (c *ArticleContoller) PostArticleHander(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "post article bad request body")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	authdUserName := common.GetUserName(req.Context())
	if reqArticle.UserName != authdUserName {
		err := apperrors.NotMatchUser.Wrap(errors.New("does note match reqBody user and idtoken user"), "invalid paramater")
		apperrors.ErrorHandler(w, req, err)
		return
	}

	article, err := c.service.PostArticleService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ResponseBodyEncodeFailed.Wrap(err, "failed to encode article response")
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

}

// /article/listのハンドラ
func (c *ArticleContoller) ListArticleHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	// クエリパラメータpageを取得
	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			err = apperrors.BadParam.Wrap(err, "article list queryparam must be number")
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}

	articleList, err := c.service.GetArticleListService(page)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err := json.NewEncoder(w).Encode(articleList); err != nil {
		err = apperrors.ResponseBodyEncodeFailed.Wrap(err, "failed to encode article list response")
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}

}

// /article/1のハンドラ
func (c *ArticleContoller) ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		err = apperrors.BadParam.Wrap(err, "article id queryparam must be number")
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	article, err := c.service.GetArticleService(articleID)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ResponseBodyEncodeFailed.Wrap(err, "failed to encode article response")
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
}

// /article/niceのハンドラ
func (c *ArticleContoller) PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article

	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		err = apperrors.ReqBodyDecodeFailed.Wrap(err, "post article bad request body")
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
		return
	}

	article, err := c.service.PostNiceService(reqArticle)
	if err != nil {
		apperrors.ErrorHandler(w, req, err)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		err = apperrors.ResponseBodyEncodeFailed.Wrap(err, "failed to encode article response")
		http.Error(w, "fail to encode json\n", http.StatusInternalServerError)
		return
	}
}
