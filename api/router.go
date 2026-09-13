package api

import (
	"database/sql"
	"net/http"

	"github.com/Nanashi123/go-api-practice/api/middlewares"
	"github.com/Nanashi123/go-api-practice/controllers"
	"github.com/Nanashi123/go-api-practice/services"
	"github.com/gorilla/mux"
)

func NewRouter(db *sql.DB) *mux.Router {
	// sql.DB型をもとに、サーバー全体で使用するサービス構造体MyAppServiceを一つ生成する
	ser := services.NewMyAppService(db)
	// MyAppService 型をもとに、サーバー全体で使用するコントローラ構造体MyAppControllerを一つ生成する
	aCon := controllers.NewArticleContoller(ser)
	cCon := controllers.NewCommentContoller(ser)

	r := mux.NewRouter()
	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article", aCon.PostArticleHander).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ListArticleHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)
	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	// ルータrに登録されているハンドラの前処理・後処理としてLoggingMiddlewareが使われるようになる
	r.Use(middlewares.LoggingMiddleware)
	r.Use(middlewares.AuthMiddleware)

	return r
}
