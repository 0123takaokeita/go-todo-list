package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-tty"
)

func main() {
	// 起動したら入力街状態にする。
	// 1. タスクの追加
	// 2. タスクの一覧
	// 3. タスクの削除
	// 4. タスクの編集
	// 5. 終了
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}

	defer tty.Close()

	for {
		fmt.Print("a. 追加, l. 一覧, d. 削除, e. 編集, x. 終了\n")

		r, err := tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		if r != 'x' {
			continue
		}

		if r == 'x' {
			fmt.Println("終了します。")
			break
		}
	}
}
