package main

import (
	"fmt"
	"log"

	"github.com/mattn/go-tty"
)

var todos map[int]Todo

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

	todos = make(map[int]Todo)


	for {
		fmt.Print("a. 追加, l. 一覧, d. 削除, e. 編集, x. 終了\n")

		r, err := tty.ReadRune()

		if err != nil {
			log.Fatal(err)
		}

		switch r {
		case 'a':
			fmt.Print("title: ")
			// 入力を受ける
			title, err := tty.ReadString()
			if err != nil {
				log.Fatal(err)
			}
			add(title)
		case 'l':
			all()
		case 'd':
			fmt.Println("削除します。")
		case 'e':
			fmt.Println("編集します。")
		case 'x':
			fmt.Println("終了します。")
			return
		default:
			fmt.Println("無効な値です。")
		}
	}
}
