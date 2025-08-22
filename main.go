package main

import (
	"bufio"
	"context"
	"fmt"
	"go-learning/db"
	"log"
	"os"
	"strings"
	"go-learning/tasks"
)

func main() {
	db.Init()
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("What would you like todo?")
		fmt.Println("1. Check tasks \n2. Create a task \n3. Update a task \n4. Delete a task \n5. Exit")
		fmt.Println("Type the number of the option")
		choice, _ := reader.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			fmt.Println("Check tasks")
			tasks.ShowTasks(ctx)
		case "2":
			if err := tasks.CreateTask(ctx, reader); err != nil {
				log.Fatal("failed to create task:", err)
			}
		case "3":
			if err := tasks.UpdateTask(ctx, reader); err != nil {
				log.Fatal("failed to update task", err)
			}
		case "4":
			if err := tasks.DeleteTask(ctx, reader); err != nil {
				log.Fatal("failed to delete task", err)
			}
		case "5":
			fmt.Println("Program exited")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice")
		}
	}
}
