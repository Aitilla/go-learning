package tasks

import (
	"context"
	"go-learning/db"
	"strings"
	"fmt"
	"bufio"
)

func CreateTask(ctx context.Context, reader *bufio.Reader) error {
	fmt.Print("Enter wanted task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	fmt.Print("What type of task: ")
	taskType, _ := reader.ReadString('\n')
	taskType = strings.TrimSpace(taskType)

	_, err := db.Conn.Exec(ctx,
		"INSERT INTO todo (task, type) VALUES ($1, $2)",
		task, taskType)
	if err != nil {
		return err
	}

	fmt.Println("Task created succesfully")
	return nil
}