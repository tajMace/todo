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

		tasks, err := Add(tasks, "todo", "")
		if err != nil {
			t.Fatalf("got an error, but didn't expect one")
		}

		if len(tasks.Tasks) != 1 {
			t.Fatalf("number of tasks wrong: got %d, wanted %d", len(tasks.Tasks), 1)
		}

		wantTask := types.Task{
			ID:   0,
			Text: "todo",
			Done: false,
			Due:  "",
		}

		want := types.TaskManager{
			NextID: 1,
			Tasks:  []types.Task{wantTask},
		}

		if tasks.NextID != want.NextID {
			t.Errorf("NextID wrong: got %d, wanted %d", tasks.NextID, want.NextID)
		}

		if !reflect.DeepEqual(tasks.Tasks, want.Tasks) {
			t.Errorf("Tasks wrong: got %+v, wanted %+v", tasks.Tasks, want.Tasks)
		}
	})

	t.Run("addition to a non-empty list", func(t *testing.T) {
		wantTask1 := types.Task{
			ID:   0,
			Text: "todo",
			Done: false,
			Due:  "",
		}
		wantTask2 := types.Task{
			ID:   1,
			Text: "todo2",
			Done: false,
			Due:  "2026-07-26",
		}

		want := types.TaskManager{
			NextID: 2,
			Tasks:  []types.Task{wantTask1, wantTask2},
		}

		tasks, err := Add(types.TaskManager{}, "todo", "")
		if err != nil {
			t.Fatalf("got an error, but didn't expect one")
		}

		tasks, err = Add(tasks, "todo2", "2026-07-26")
		if err != nil {
			t.Fatalf("got an error, but didn't expect one")
		}

		if len(tasks.Tasks) != 2 {
			t.Logf("got %+v, wanted %+v", tasks, want)
			t.Fatalf("number of tasks wrong: got %d, wanted %d", len(tasks.Tasks), 2)
		}

		if tasks.NextID != want.NextID {
			t.Errorf("NextID wrong: got %d, wanted %d", tasks.NextID, want.NextID)
		}

		if !reflect.DeepEqual(tasks.Tasks, want.Tasks) {
			t.Errorf("Tasks wrong: got %+v, wanted %+v", tasks.Tasks, want.Tasks)
		}
	})

	t.Run("incorrectly formatted date", func(t *testing.T) {
		_, err := Add(types.TaskManager{}, "todo", "not-a-date")
		want := ErrIncorrectDateFormat

		if err == nil {
			t.Fatal("Expected an error, but didn't get one")
		}

		if err != want {
			t.Errorf("Error wrong: got %+v, wanted %+v", err, want)
		}
	})
}
