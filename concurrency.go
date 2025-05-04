// Package main は、Goroutineを使用した並行処理のサンプルを提供します。
package main

import (
	"fmt"
	"sync"
	"time"
)

// Task は並行処理で実行するタスクの構造体
type Task struct {
	ID        int
	Duration  time.Duration
	Completed bool
}

// ProcessTask は個々のタスクを処理する関数
func ProcessTask(task *Task, wg *sync.WaitGroup) {
	defer wg.Done() // タスク完了時にWaitGroupのカウンタをデクリメント

	fmt.Printf("タスク %d を開始します\n", task.ID)
	time.Sleep(task.Duration) // タスクの処理時間をシミュレート
	task.Completed = true
	fmt.Printf("タスク %d が完了しました\n", task.ID)
}

// RunConcurrentTasks は複数のタスクを並行処理で実行する関数
func RunConcurrentTasks() {
	// タスクの作成
	tasks := []Task{
		{ID: 1, Duration: 2 * time.Second},
		{ID: 2, Duration: 1 * time.Second},
		{ID: 3, Duration: 3 * time.Second},
	}

	// WaitGroupの初期化
	var wg sync.WaitGroup
	wg.Add(len(tasks)) // 実行するタスクの数を設定

	// 各タスクをGoroutineで実行
	for i := range tasks {
		go ProcessTask(&tasks[i], &wg)
	}

	// すべてのタスクが完了するまで待機
	wg.Wait()
	fmt.Println("すべてのタスクが完了しました")
}
