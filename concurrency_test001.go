package main

import (
	"fmt"
	"sync"
)

func RunGreet() {

	// 構造体を定義
	type Greets struct {
		greet string
	}

	// 挨拶のリストを定義
	greets := []Greets{
		{"Hello from Goroutine 1"},
		{"Hello from Goroutine 2"},
	}

	// WaitGroupを初期化
	var wg sync.WaitGroup

	// リストの要素数だけ待機列へ追加する
	wg.Add(len(greets))

	for n := range greets {

		// 無名関数
		go func(msg string) {

			// deferは最後に実行したい処理を予約する
			// wgの待機列から１つ消す
			defer wg.Done()

			// msgを出力
			fmt.Println(msg)

		}(greets[n].greet) //匿名関数の即時実行
	}

	// wgが終了するまで待機
	wg.Wait()
}
