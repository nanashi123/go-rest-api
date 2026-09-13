package main_test

import (
	"fmt"
	"testing"
)

func TestA(t *testing.T) {
	// TestAにおける後処理の定義
	t.Cleanup(func() {
		// 後処理内容
		fmt.Println("cleanup")
	})

	// テストの実施
	fmt.Println("testA")
}
