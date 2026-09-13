// handlersパッケージのハンドラを利用して、サーバーを起動する

package main

import (
	"encoding/json"
	"fmt"
	"time"
)

type Comment struct {
	CommentID int       `json:"comment_id"`
	ArticleID int       `json:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
}

type Article struct {
	ID          int       `json:"article_id"`
	Title       string    `json:"title"`
	Contents    string    `json:"contens"`
	UserName    string    `json:"user_name"`
	NiceNum     int       `json:"nice_num"`
	CommentList []Comment `json:"comment_list"`
	CreatedAt   time.Time `json:"created_at"`
}

func main() {
	comment1 := Comment{
		CommentID: 1,
		ArticleID: 1,
		Message:   "test comment1",
		CreatedAt: time.Now(),
	}
	comment2 := Comment{
		CommentID: 2,
		ArticleID: 1,
		Message:   "test comment2",
		CreatedAt: time.Now(),
	}
	article := Article{
		ID:          1,
		Title:       "test title",
		Contents:    "test contents",
		UserName:    "test user",
		NiceNum:     1,
		CommentList: []Comment{comment1, comment2},
		CreatedAt:   time.Now(),
	}

	jsonData, err := json.Marshal(article)
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	jsonIndentData, err := json.MarshalIndent(article, "", "  ")
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	fmt.Printf("%s\n", jsonData)
	fmt.Printf("%s\n", jsonIndentData)
}
