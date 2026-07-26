// taj-mace
// tests for addition of todo items

package add

import (
	"reflect"
	"testing"

	"github.com/tajMace/todo/types"
)

func TestAdd(t *testing.T) {
	t.Run("addition to an empty list", func(t *testing.T) {
		tasks := types.TaskManager{}

		newTasks := Add(tasks, "todo", "")
		if len(newTasks.Tasks) == 0 {
			t.Fatalf("expected item to be added, but wasn't")
		}

		wantTask := types.Task{
			ID:   tasks.NextID,
			Text: "todo",
			Done: false,
			Due:  "",
		}

		want := types.TaskManager{
			NextID: 1,
			Tasks:  []types.Task{wantTask},
		}

		if newTasks.NextID != want.NextID {
			t.Errorf("NextID wrong: got %d, wanted %d", newTasks.NextID, want.NextID)
		}

		if !reflect.DeepEqual(newTasks.Tasks, want.Tasks) {
			t.Errorf("Tasks wrong: got %+v, wanted %+v", newTasks.Tasks, want.Tasks)
		}
	})
}
