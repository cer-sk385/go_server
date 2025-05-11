package main

import (
	"fmt"
)

func RunGreet() {
	// 挨拶のリストを定義
	greets := []string{
		"Hello from Goroutine 1",
		"Hello from Goroutine 2",
	}

	// 完了通知用のチャネルを作成
	done := make(chan bool)

	for _, msg := range greets {
		// 無名関数
		go func(msg string) {
			// msgを出力
			fmt.Println(msg)
			// 処理完了を通知
			done <- true
		}(msg) //匿名関数の即時実行
	}

	// 全てのゴルーチンの完了を待機
	for range greets {
		<-done
	}
}
