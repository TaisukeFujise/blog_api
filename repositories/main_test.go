package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"
)

// テストパッケージ repositories_test 内全体で使える変数 testDB を用意
var testDB *sql.DB

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}

func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanup", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setup")
		return err
	}
	return nil
}

func teardown() {
	cleanupDB()
	testDB.Close()
}

func setupTestData() error {
	f, err := os.Open("./testdata/setupDB.sql")
	if err != nil {
		return err
	}
	defer f.Close()
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "--password=docker", "sampledb")
	cmd.Stdin = f
	return cmd.Run()
}

func cleanupDB() error {
	f, err := os.Open("./testdata/cleanupDB.sql")
	if err != nil {
		return err
	}
	defer f.Close()
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "--password=docker", "sampledb")
	cmd.Stdin = f
	return cmd.Run()
}
