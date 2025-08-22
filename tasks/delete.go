package tasks
import (
	"context"
	"go-learning/db"
	"strings"
	"fmt"
	"strconv"
	"os"
	"bufio"
)

func DeleteTask(ctx context.Context, reader *bufio.Reader) error {

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

	if choice == "n" {
		os.Exit(0)
	}

	_, err = db.Conn.Exec(ctx, "DELETE FROM todo WHERE id = $1", id)
	if err != nil {
		return err
	}

	fmt.Println("Task deleted succesfully")
	return nil
}
