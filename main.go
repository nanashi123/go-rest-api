package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Nanashi123/go-api-practice/api"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbUser     = os.Getenv("DB_USER")
	dbPassword = os.Getenv("DB_PASSWORD")
	dbDatabase = os.Getenv("DB_NAME")
	dbConn     = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPassword, dbDatabase)
)

func main() {
	// 1. サーバー全体で使用するsql.DB型を一つ生成す
	db, err := sql.Open("mysql", dbConn)
	if err != nil {
		log.Println("failtoconnectDB")
		return
	}

	r := api.NewRouter(db)

	log.Println("server start at port 8080")
	fmt.Println("http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
