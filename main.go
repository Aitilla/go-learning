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

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What would you like todo?")
	fmt.Println("1. Check tasks \na2. Create a task")
	fmt.Println("Type the number of the option")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)


	fmt.Print("Enter wanted task: ")
	task, _ := reader.ReadString('\n')
	task = strings.TrimSpace(task)

	fmt.Print("What type of task: ")
	taskType, _ := reader.ReadString('\n')
	taskType = strings.TrimSpace(taskType)

	fmt.Println(task, "with type", taskType, "has been created")
}
