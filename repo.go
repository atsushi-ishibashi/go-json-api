package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db *sql.DB
)

func init() {
	var err error
	db, err = sql.Open("mysql", "root:@/gotest?parseTime=true")
	if err != nil {
		panic(err)
	}
}

//RepoFindTodo ....
func RepoFindTodo(id int) Todo {
	var t Todo
	if err := t.findByID(id); err != nil {
		fmt.Printf("Failed QueryRow: %s", err)
		return Todo{}
	}
	return t
}

//RepoCreateTodo ....
func RepoCreateTodo(t Todo) Todo {
	if err := t.insertTodo(); err != nil {
		fmt.Printf("Failed RepoCreateTodo: %s", err)
		return Todo{}
	}
	return t
}

//RepoDestroyTodo ...
func RepoDestroyTodo(id int) error {
	var t Todo
	if err := t.destroyByID(id); err != nil {
		fmt.Printf("Failed destroyByID: %s", err)
		return err
	}
	return nil
}
