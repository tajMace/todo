// taj-mace
// tests for addition of todo items

package add

import (
	"testing"

	"github.com/tajMace/todo/types"
)

func TestAdd(t *testing.T) {
	t.Run("basic success: adds a todo item", func(t *testing.T) {
		tasks := types.TaskManager{}
		newTask := types.Task{
			ID: tasks.NextID,
			Text: "todo",
			Done: false,
			Due: nil,
		}

		newTasks = Add(tasks, newTask)
		if len(newTasks) == 0 {
			t.Fatal("expected item to be added, but wasn't")
		}

		want := types.TaskManager{
			NextID: 1,
			Tasks: []types.Task{newTask}
		}
		if newTasks != want {
			t.Error("item appended incorrectly")
		}
	})
}
