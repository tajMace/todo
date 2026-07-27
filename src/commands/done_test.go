// taj-mace
// tests for marking a task as done
package commands

import (
	"reflect"
	"testing"

	"github.com/tajMace/todo/types"
)

func TestDone(t *testing.T) {
	t.Run("successful mark", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task", "")

		got, err := Done(tm, 0)
		if err != nil {
			t.Fatal("got an error, but wasn't expecting one")
		}

		if !got.Tasks[0].Done {
			t.Error("expected task to be done, but wasn't")
		}
	})

	t.Run("other tasks remain unchanged", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task1", "")
		tm.NewTaskOnManager("task2", "")

		got, err := Done(tm, 0)
		if err != nil {
			t.Fatal("got an error, but wasn't expecting one")
		}

		if got.Tasks[1].Done {
			t.Error("task 1 should not have been affected by marking task 0 done")
		}
		if got.Tasks[1].Text != "task2" {
			t.Errorf("task 1 text changed unexpectedly: got %q", got.Tasks[1].Text)
		}
	})

	t.Run("nonexistent id returns error", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task", "")

		_, err := Done(tm, 99)
		if err != ErrIdNotInUse {
			t.Errorf("got error %v, wanted %v", err, ErrIdNotInUse)
		}
	})

	t.Run("empty list returns error", func(t *testing.T) {
		tm := types.NewTaskManager(nil)

		_, err := Done(tm, 0)
		if err != ErrIdNotInUse {
			t.Errorf("got error %v, wanted %v", err, ErrIdNotInUse)
		}
	})

	t.Run("marking an already-done task again is a no-op success", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task", "")

		tm, err := Done(tm, 0)
		if err != nil {
			t.Fatal("got an error, but wasn't expecting one")
		}

		got, err := Done(tm, 0)
		if err != nil {
			t.Fatal("got an error marking an already-done task, but wasn't expecting one")
		}
		if !got.Tasks[0].Done {
			t.Error("expected task to still be done")
		}
	})

	t.Run("does not mutate the input TaskManager", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task", "")

		before := types.TaskManager{
			NextID: tm.NextID,
			Tasks:  append([]types.Task{}, tm.Tasks...),
		}

		_, err := Done(tm, 0)
		if err != nil {
			t.Fatal("got an error, but wasn't expecting one")
		}

		if !reflect.DeepEqual(tm, before) {
			t.Errorf("Done mutated its input: got %+v, wanted unchanged %+v", tm, before)
		}
	})
}
