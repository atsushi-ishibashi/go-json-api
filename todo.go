package main

import (
	"fmt"
	"time"
)

//Todo ...
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

func (t *Todo) findByID(id int) error {
	if err := db.QueryRow("SELECT id, name, completed, due FROM todos WHERE id=?", id).Scan(&t.ID, &t.Name, &t.Completed, &t.Due); err != nil {
		return err
	}
	return nil
}

func (t *Todo) insertTodo() error {
	stmt, err := db.Prepare("INSERT todos SET name=?,completed=?,due=?")
	if err != nil {
		fmt.Printf("Failed Prepare: %s", err)
		return err
	}
	res, err := stmt.Exec(t.Name, t.Completed, t.Due)
	if err != nil {
		fmt.Printf("Failed Exec: %s", err)
		return err
	}
	id, err := res.LastInsertId()
	if err != nil {
		fmt.Printf("Failed LastInsertId: %s", err)
		return err
	}
	t.ID = int(id)
	return nil
}

func (t *Todo) destroyByID(id int) error {
	stmt, err := db.Prepare("DELETE FROM todos WHERE id=?")
	if err != nil {
		fmt.Printf("Failed Prepare: %s", err)
		return err
	}
	res, err := stmt.Exec(id)
	if err != nil {
		fmt.Printf("Failed Exec: %s", err)
		return err
	}
	if _, err := res.RowsAffected(); err != nil {
		fmt.Printf("Failed RowsAffected: %s", err)
		return err
	}
	return nil
}

//Todos ...
type Todos []Todo

func (ts *Todos) selectAll() error {
	rows, err := db.Query("SELECT id, name, completed, due FROM todos")
	if err != nil {
		fmt.Printf("Failed selectAll: %s", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Name, &t.Completed, &t.Due); err != nil {
			fmt.Printf("Failed Scan: %s", err)
			return err
		}
		*ts = append(*ts, t)
	}
	return nil
}
