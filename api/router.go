package api

import (
	"database/sql"
	"net/http"

	"github.com/TaisukeFujise/blog_api/api/middlewares"
	"github.com/TaisukeFujise/blog_api/controllers"
	"github.com/TaisukeFujise/blog_api/services"
	"github.com/gorilla/mux"
)

/*
- service構造体の生成（ser）
- controller構造体の生成（aCon, cCon）
- Routerの登録
*/
func NewRouter(db *sql.DB) *mux.Router {
	ser := services.NewMyAppService(db)
	aCon := controllers.NewArticleController(ser)
	cCon := controllers.NewCommentController(ser)

	r := mux.NewRouter()

	r.HandleFunc("/hello", aCon.HelloHandler).Methods(http.MethodGet)

	r.HandleFunc("/article", aCon.PostArticleHandler).Methods(http.MethodPost)
	r.HandleFunc("/article/list", aCon.ArticleListHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{id:[0-9]+}", aCon.ArticleDetailHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/nice", aCon.PostNiceHandler).Methods(http.MethodPost)

	r.HandleFunc("/comment", cCon.PostCommentHandler).Methods(http.MethodPost)

	r.Use(middlewares.LoggingMiddleware)
	return r
}
