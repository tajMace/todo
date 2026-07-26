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
		wantTask := types.NewTask(0, "todo", "")
		want := types.NewTaskManager([]types.Task{wantTask})

		got := addAndAssertNoErrors(t, types.TaskManager{}, "todo", "")

		compareTasks(t, got, want)
	})

	t.Run("addition to a non-empty list", func(t *testing.T) {
		wantTask1 := types.NewTask(0, "todo", "")
		wantTask2 := types.NewTask(1, "todo2", "2026-07-26")
		want := types.NewTaskManager([]types.Task{wantTask1, wantTask2})

		got := types.TaskManager{}
		got = addAndAssertNoErrors(t, got, "todo", "")
		got = addAndAssertNoErrors(t, got, "todo2", "2026-07-26")

		compareTasks(t, got, want)
	})

	t.Run("incorrectly formatted date", func(t *testing.T) {
		want := ErrIncorrectDateFormat

		err := addAndAssertError(t, types.TaskManager{}, "todo", "not-a-date")

		compareErrors(t, err, want)
	})
}

/* ========== helpers ========== */
func addAndAssertNoErrors(t testing.TB, tasks types.TaskManager, text, due string) types.TaskManager {
	t.Helper()

	got, err := Add(tasks, text, due)
	if err != nil {
		t.Fatalf("got an error, but didn't expect one")
	}

	return got
}

func addAndAssertError(t testing.TB, tasks types.TaskManager, text, due string) error {
	t.Helper()

	_, err := Add(tasks, text, due)
	if err == nil {
		t.Fatal("Expected an error, but didn't get one")
	}

	return err
}

func compareTasks(t testing.TB, got, want types.TaskManager) {
	t.Helper()

	if len(got.Tasks) != len(want.Tasks) {
		t.Fatalf("number of tasks wrong: got %d, wanted %d", len(got.Tasks), len(want.Tasks))
	}

	if got.NextID != want.NextID {
		t.Errorf("NextID wrong: got %d, wanted %d", got.NextID, want.NextID)
	}

	if !reflect.DeepEqual(got.Tasks, want.Tasks) {
		t.Errorf("Tasks wrong: got %+v, wanted %+v", got.Tasks, want.Tasks)
	}
}

func compareErrors(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("Error wrong: got %+v, wanted %+v", got, want)
	}
}
