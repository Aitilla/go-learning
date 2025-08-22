package tasks

import (
	"context"
	"go-learning/db"
	"log"
	"fmt"
)

func ShowTasks(ctx context.Context) {
	rows, err := db.Conn.Query(ctx, "SELECT id, task, type, status FROM todo")
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var task, taskType string
		var status bool

		err := rows.Scan(&id, &task, &taskType, &status)
		if err != nil {
			log.Fatal("something went wrong", err)
		}
		fmt.Println("ID:", id, "| Task:", task, "| Type:", taskType, "| Completed:", status)
	}
}
