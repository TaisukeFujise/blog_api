package repositories_test

import (
	"testing"

	"github.com/TaisukeFujise/blog_api/models"
	"github.com/TaisukeFujise/blog_api/repositories"
	"github.com/TaisukeFujise/blog_api/repositories/testdata"

	_ "github.com/go-sql-driver/mysql"
)

func TestSelectArticleList(t *testing.T) {
	expectedNum := len(testdata.ArticleTestData)

	repo := repositories.NewArticleRepository(testDB)
	ctx := t.Context()
	got, err := repo.SelectArticleList(ctx, 1)
	if err != nil {
		t.Fatal(err)
	}

	if num := len(got); num != expectedNum {
		t.Errorf("want %d but got %d articles\n", expectedNum, num)
	}
}

func TestSelectArticleDetail(t *testing.T) {
	tests := []struct {
		testTitle string
		expected  models.Article
	}{
		{
			testTitle: "subtest1",
			expected:  testdata.ArticleTestData[0],
		}, {
			testTitle: "subtest2",
			expected:  testdata.ArticleTestData[1],
		},
	}

	repo := repositories.NewArticleRepository(testDB)
	ctx := t.Context()
	for _, test := range tests {
		t.Run(test.testTitle, func(t *testing.T) {
			got, err := repo.SelectArticleDetail(ctx, test.expected.ID)
			if err != nil {
				t.Fatal(err)
			}

			if got.ID != test.expected.ID {
				t.Errorf("ID: get %d but want %d\n", got.ID, test.expected.ID)
			}
			if got.Title != test.expected.Title {
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
}

func TestInsertArticle(t *testing.T) {
	article := models.Article{
		Title:    "insertTest",
		Contents: "testest",
		UserName: "saki",
	}
	expectedArticleNum := 3

	repo := repositories.NewArticleRepository(testDB)
	ctx := t.Context()
	newArticle, err := repo.InsertArticle(ctx, article)
	if err != nil {
		t.Fatal(err)
	}
	if newArticle.ID != expectedArticleNum {
		t.Errorf("new article id is expected %d but got %d\n", expectedArticleNum, newArticle.ID)
	}

	t.Cleanup(func() {
		const sqlStr = `
			delete from articles
			where title = ? and contents = ? and username = ?
		`
		testDB.Exec(sqlStr, article.Title, article.Contents, article.UserName)
	})
}

func TestUpdateNiceNum(t *testing.T) {
	articleID := 1
	repo := repositories.NewArticleRepository(testDB)
	ctx := t.Context()

	before, err := repo.SelectArticleDetail(ctx, articleID)
	if err != nil {
		t.Fatal("fail to get before data")
	}

	err = repo.UpdateNiceNum(ctx, articleID)
	if err != nil {
		t.Fatal(err)
	}

	after, err := repo.SelectArticleDetail(ctx, articleID)
	if err != nil {
		t.Fatal("fail to get after data")
	}
	if after.NiceNum-before.NiceNum != 1 {
		t.Error("fail to update nice num")
	}
}
