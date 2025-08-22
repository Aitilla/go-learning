package tasks
import (
	"context"
	"go-learning/db"
	"strconv"
	"fmt"
	"strings"
	"bufio"
)

func UpdateTask(ctx context.Context, reader *bufio.Reader) error {
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