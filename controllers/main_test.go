package controllers_test

import (
	"testing"

	_ "github.com/go-sql-driver/mysql"

	"github.com/TaisukeFujise/blog_api/controllers"
	"github.com/TaisukeFujise/blog_api/controllers/testdata"
)

var aCon *controllers.ArticleController

func TestMain(m *testing.M) {
	ser := testdata.NewServiceMock()
	aCon = controllers.NewArticleController(ser)
	m.Run()
}
