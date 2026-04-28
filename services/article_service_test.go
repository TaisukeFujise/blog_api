package services_test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"

	"github.com/TaisukeFujise/blog_api/repositories"
	"github.com/TaisukeFujise/blog_api/services"
	_ "github.com/go-sql-driver/mysql"
)

var aSer *services.ArticleService

func TestMain(m *testing.M) {
	dbUser := "docker"
	dbPassword := "docker"
	dbDatabase := "sampledb"
	dbConn := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	articleRepo := repositories.NewArticleRepository(db)
	commentRepo := repositories.NewCommentRepository(db)
	aSer = services.NewArticleService(articleRepo, commentRepo)
	m.Run()
}

func BenchmarkGetArticleService(b *testing.B) {
	articleID := 1

	ctx := b.Context()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := aSer.GetArticleService(ctx, articleID)
		if err != nil {
			b.Error(err)
			break
		}
	}
}
