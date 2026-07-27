// taj-mace
// tests for removing a task
package commands

import (
	"reflect"
	"testing"

	"github.com/tajMace/todo/types"
)

func TestRemove(t *testing.T) {
	t.Run("removes an existing task", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task", "")

		got, err := Remove(tm, 0)
		if err != nil {
			t.Fatal("got an error, but wasn't expecting one")
		}

		if len(got.Tasks) != 0 {
			t.Errorf("expected list to be empty, got %+v", got.Tasks)
		}
	})

	t.Run("other tasks remain unchanged after removal", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task1", "")
		tm.NewTaskOnManager("task2", "")

		got, err := Remove(tm, 0)
		if err != nil {
			t.Fatal("got an error, but wasn't expecting one")
		}

		if len(got.Tasks) != 1 {
			t.Fatalf("expected 1 task remaining, got %d: %+v", len(got.Tasks), got.Tasks)
		}
		if got.Tasks[0].ID != 1 || got.Tasks[0].Text != "task2" {
			t.Errorf("wrong task remained: got %+v", got.Tasks[0])
		}
	})

	t.Run("nonexistent id returns error", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task", "")

		_, err := Remove(tm, 99)
		if err != ErrIdNotInUse {
			t.Errorf("got error %v, wanted %v", err, ErrIdNotInUse)
		}
	})

	t.Run("empty list returns error", func(t *testing.T) {
		tm := types.NewTaskManager(nil)

		_, err := Remove(tm, 0)
		if err != ErrIdNotInUse {
			t.Errorf("got error %v, wanted %v", err, ErrIdNotInUse)
		}
	})

	t.Run("does not mutate the input TaskManager", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task1", "")
		tm.NewTaskOnManager("task2", "")

		before := types.TaskManager{
			NextID: tm.NextID,
			Tasks:  append([]types.Task{}, tm.Tasks...),
		}

		_, err := Remove(tm, 0)
		if err != nil {
			t.Fatal("got an error, but wasn't expecting one")
		}

		if !reflect.DeepEqual(tm, before) {
			t.Errorf("Remove mutated its input: got %+v, wanted unchanged %+v", tm, before)
		}
	})

	t.Run("removing then adding avoids ID collision", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("task1", "") // ID 0
		tm.NewTaskOnManager("task2", "") // ID 1

		tm, err := Remove(tm, 1)
		if err != nil {
			t.Fatal("got an error, but wasn't expecting one")
		}

		tm, err = Add(tm, "task3", "")
		if err != nil {
			t.Fatal("got an error, but wasn't expecting one")
		}

		seen := map[int]bool{}
		for _, task := range tm.Tasks {
			if seen[task.ID] {
				t.Fatalf("duplicate ID %d found after remove+add: %+v", task.ID, tm.Tasks)
			}
			seen[task.ID] = true
		}
	})
}
