// taj-mace
// entry point of the app

package main

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/tajMace/todo/src/commands"
	"github.com/tajMace/todo/src/storage"
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
		}
	default:
		{
			fmt.Println("Usage: todo <title> [arguments]")
		}
	}

	err = storage.SaveTasks(path, result)
	handleErrorReturn(err)
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
