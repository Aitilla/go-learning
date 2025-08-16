package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"log"
	"go-learning/db"
	"context"
)

func main() {


	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What would you like todo?")
	fmt.Println("1. Check tasks \n2. Create a task \n3. Exit")
	fmt.Println("Type the number of the option")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Check tasks")
		showTasks()
	case "2":
		task, taskType := createTask(reader)
		fmt.Println("Task:", task, "with type", taskType, "has been created")
	case "3":
		fmt.Println("Program exited")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}

func showTasks(){
	db.Init()
	ctx := context.Background()

	rows, err := db .Conn.Query(ctx, "SELECT task, type, status FROM todo")
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

func createTask(reader *bufio.Reader) (string, string) {
	fmt.Print("Enter wanted task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	fmt.Print("What type of task: ")
	taskType, _ := reader.ReadString('\n')
	taskType = strings.TrimSpace(taskType)

	return task, taskType
}
