// taj-mace
// entry point of the app

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/tajMace/todo/src/commands"
	"github.com/tajMace/todo/src/storage"
	"github.com/tajMace/todo/types"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		earlyReturn()
	}

	home, err := os.UserHomeDir()
	handleErrorReturn(err)

	path := filepath.Join(home, ".todo.json")

	tm, err := storage.LoadTasks(path)
	handleErrorReturn(err)

	result := tm
	mutation := false

	switch args[0] {
	case "add":
		{
			if len(args) < 2 {
				earlyReturn()
			}
			text := args[1]
			date := ""
			if len(args) > 3 && (args[2] == "--due" || args[2] == "--d") {
				date = args[3]
			}

			result, err = commands.Add(tm, text, date)
			handleErrorReturn(err)

			mutation = true
		}
	case "list":
		{

			data := commands.List(result)
			if len(data) == 0 {
				fmt.Println("No items todo! Well done!")
				break
			}
			printFormatter(data)
		}
	case "done":
		{
			if len(args) < 2 {
				earlyReturn()
			}

			id, err := strconv.Atoi(args[1])
			if err != nil {
				earlyReturn()
			}

			result, err = commands.Done(tm, id)

			mutation = true
		}
	default:
		{
			fmt.Println("Usage: todo <title> [arguments]")
		}
	}

	if mutation {
		err = storage.SaveTasks(path, result)
		handleErrorReturn(err)
	}
}

/* ========== HELPERS ========== */
func handleErrorReturn(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

func earlyReturn() {
	fmt.Println("Usage: todo <title> [arguments]")
	os.Exit(1)
}

func printFormatter(tasks []types.Task) {
	index := 0

	fmt.Print("\033[H\033[2J") // terminal clear code

	fmt.Println(" ================================ ")
	fmt.Println(" ========== TASKS TODO ==========")
	fmt.Println(" ================================ ")

	fmt.Println("")
	fmt.Println(" ---------- Dated Tasks ----------")
	for _, task := range tasks {
		if task.Due == "" {
			break
		}

		fmt.Printf("[ %d ] \t %s - %s\n", task.ID, task.Text, task.Due)
		index++
	}

	fmt.Println("")
	fmt.Println(" --------- Undated Tasks ---------")
	for _, task := range tasks[index:] {
		fmt.Printf("[ %d ] \t %s\n", task.ID, task.Text)
	}

	fmt.Println("")
}
