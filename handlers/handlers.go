package handlers

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/TaisukeFujise/blog_api/models"
	"github.com/gorilla/mux"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "Hello world\n")
}

// POST /article: リクエストボディで受け取った記事を投稿する
func PostArticleHandler(w http.ResponseWriter, req *http.Request) {
	var reqArticle models.Article
	if err := json.NewDecoder(req.Body).Decode(&reqArticle); err != nil {
		http.Error(w, "fail to decode json\n", http.StatusBadRequest)
	}
	article := reqArticle
	json.NewEncoder(w).Encode(article)
}

// GET /article/list: クエリパラメータpageで指定されたページ（1ページ5個の記事を表示）に表示するための記事一覧を取得
func ArticleListHandler(w http.ResponseWriter, req *http.Request) {
	queryMap := req.URL.Query()

	var page int
	if p, ok := queryMap["page"]; ok && len(p) > 0 {
		var err error
		page, err = strconv.Atoi(p[0])
		if err != nil {
			http.Error(w, "Invalid query parameter", http.StatusBadRequest)
			return
		}
	} else {
		page = 1
	}
	log.Println(page)

	articleList := []models.Article{models.Article1, models.Article2}
	json.NewEncoder(w).Encode(articleList)
}

// GET /article/{id}: 指定IDの記事を取得する
func ArticleDetailHandler(w http.ResponseWriter, req *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(req)["id"])
	if err != nil {
		http.Error(w, "Invalid query parameter", http.StatusBadRequest)
		return
	}

	log.Println(articleID)

	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

// POST /article/nice: 記事にいいねをつける
func PostNiceHandler(w http.ResponseWriter, req *http.Request) {
	article := models.Article1
	json.NewEncoder(w).Encode(article)
}

// POST /comment: リクエストボディで受け取ったコメントを投稿する
func PostCommentHandler(w http.ResponseWriter, req *http.Request) {
	comment := models.Comment1
	json.NewEncoder(w).Encode(comment)
}
