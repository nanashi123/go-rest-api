package repositories_test

import (
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"testing"

	_ "github.com/go-sql-driver/mysql"
)

// テスト全体で共有する sql.DB 型
var testDB *sql.DB

var (
	dbUser     = "docker"
	dbPassword = "docker"
	dbDatabase = "sampledb"
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		fmt.Println("setup error: ", err)
		os.Exit(1)
	}

	// ユニットテストをすべて実行
	resultCode := m.Run()
	teardown()
	if err := teardown(); err != nil {
		fmt.Println("teardown error:", err)
		resultCode = 1
	}
	os.Exit(resultCode)
}

// 全テスト共通の前処理を書く
func setup() error {
	if err := connectDB(); err != nil {
		return err
	}
	if err := cleanupDB(); err != nil {
		fmt.Println("cleanupDB error:", err)
		return err
	}
	if err := setupTestData(); err != nil {
		fmt.Println("setupTestData error:", err)
		return err
	}
	return nil
}

// 全テスト処理の後処理を書く
func teardown() error {
	err := cleanupDB()
	err2 := testDB.Close()
	if err != nil {
		return err
	}
	if err2 != nil {
		return err2
	}
	return nil
}

func connectDB() error {
	var err error
	testDB, err = sql.Open("mysql", dbConn)
	if err != nil {
		return err
	}
	return nil
}

func setupTestData() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/setupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func cleanupDB() error {
	cmd := exec.Command("mysql", "-h", "127.0.0.1", "-u", "docker", "sampledb", "--password=docker", "-e", "source ./testdata/cleanupDB.sql")
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}
