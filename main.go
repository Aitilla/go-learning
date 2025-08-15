package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
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
