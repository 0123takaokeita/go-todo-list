package main

import (
	"fmt"
	"time"
)

type Todo struct {
	ID        int       `json:"id"`
	Title     string    `json:"title"`
	Priority  int       `json:"priority"`
	Deadline  time.Time `json:"deadline"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}


func add(title string) {
	id := len(todos) + 1
	todo := Todo{
		ID:        id,
		Title:     title,
		Priority:  1,
		Deadline:  time.Now(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	todos[id] = todo
	fmt.Printf("タスクを追加しました。%s\n", title)
}

func all() {
	fmt.Println("一覧表示します。")
	for k, todo := range todos {
		fmt.Printf("%d: %s\n", k, todo.Title)
	}
}
