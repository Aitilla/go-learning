package main

import (
	"bufio"
	"context"
	"fmt"
	"go-learning/db"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	db.Init()
	ctx := context.Background()
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("What would you like todo?")
	fmt.Println("1. Check tasks \n2. Create a task \n3. Update a task \n4. Delete a task \n5. Exit")
	fmt.Println("Type the number of the option")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	switch choice {
	case "1":
		fmt.Println("Check tasks")
		showTasks(ctx)
	case "2":
		if err := createTask(ctx, reader); err != nil {
			log.Fatal("failed to create task:", err)
		}
	case "3":
		if err := updateTask(ctx, reader); err != nil {
			log.Fatal("failed to update task", err)
		}
	case "4":
		if err := deleteTask(ctx, reader); err != nil {
			log.Fatal("failed to delete task", err)
		}
	case "5":
		fmt.Println("Program exited")
		os.Exit(0)
	default:
		fmt.Println("Invalid choice")
	}
}

func showTasks(ctx context.Context) {
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

func updateTask(ctx context.Context, reader *bufio.Reader) error {
	fmt.Print("Enter the task id: ")
	edit, _ := reader.ReadString('\n')
	edit = strings.TrimSpace(edit)
	id, err := strconv.Atoi(edit)
	if err != nil {
		return fmt.Errorf("invalid id: %v", err)
	}

	fmt.Print("Mark as completed? (y/n): ")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	var newStatus bool
	if choice == "y" {
		newStatus = true
	} else {
		newStatus = false
	}

	_, err = db.Conn.Exec(ctx, "UPDATE todo SET status = $1 WHERE id = $2", newStatus, id)
	if err != nil {
		return fmt.Errorf("failed to update task: %v", err)
	}

	fmt.Println("Task updated successfully")
	return nil
}

func deleteTask(ctx context.Context, reader *bufio.Reader) error {
	
	fmt.Println("Enter task id: ")
	delete, _ := reader.ReadString('\n')
	delete = strings.TrimSpace(delete)
	id, err := strconv.Atoi(delete)
	if err != nil {
		return fmt.Errorf("invalid id: %v", err)
	}

	fmt.Print("Are you sure? (y/n)")
	choice, _ := reader.ReadString('\n')
	choice = strings.TrimSpace(choice)

	if choice == "n"{
		os.Exit(0)
	}

	_, err = db.Conn.Exec(ctx, "DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		return err
	}

	fmt.Println("Task deleted succesfully")
	return nil
}