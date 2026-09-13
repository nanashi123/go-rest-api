package main

import (
	"context"
	"fmt"
	"os"

	"google.golang.org/api/idtoken"
)

func main() {
	googleClientID := os.Getenv("GOOGLE_CLIENTID")
	idToken := os.Getenv("ID_TOKEN")

	// Validatorの生成
	tokenValidator, err := idtoken.NewValidator(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}

	// Validateメソッドによる検証作業
	payload, err := tokenValidator.Validate(context.Background(), idToken, googleClientID)
	if err != nil {
		fmt.Println("validate err: ", err)
		return
	}

	// 検証成功の場合にはpayloadにペイロードの情報が格納
	fmt.Println(payload.Claims["name"])
}
