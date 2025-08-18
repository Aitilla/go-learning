package main

import (
	"bufio"
	"context"
	"fmt"
	"go-learning/db"
	"log"
	"os"
	"strings"
)

func main() {
	db.Init()
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What would you like todo?")
	fmt.Println("1. Check tasks \n2. Create a task \n3. Exit")
	fmt.Println("Type the number of the option")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Check tasks")
		showTasks(ctx)
	case "2":
		if err := createTask(ctx, reader); err != nil {
			log.Fatal("Failed to create task:", err)
		}
	case "3":
		fmt.Println("Program exited")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}

func showTasks(ctx context.Context) {
	rows, err := db.Conn.Query(ctx, "SELECT task, type, status FROM todo")
	if err != nil {
		log.Fatal("Something went wrong", err)
	}
	defer rows.Close()

	for rows.Next() {
		var task, taskType string
		var status bool

		err := rows.Scan(&task, &taskType, &status)
		if err != nil {
			log.Fatal("Something went wrong", err)
		}
		fmt.Println("Task:", task, "| Type:", taskType, "| Completed:", status)
	}
}

func createTask(ctx context.Context, reader *bufio.Reader) error {
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
