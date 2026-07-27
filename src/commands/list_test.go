// taj-mace
// tests the ordering of a given list of Tasks
package commands

import (
	"reflect"
	"testing"

	"github.com/tajMace/todo/types"
)

func TestList(t *testing.T) {
	t.Run("correctly orders by ID (no dates)", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("test2", "")
		tm.NewTaskOnManager("test", "")

		want := []types.Task{
			{ID: 0, Text: "test2", Due: ""},
			{ID: 1, Text: "test", Due: ""},
		}

		got := List(tm)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("ID order wrong: got %+v, wanted %+v", got, want)
		}
	})

	t.Run("correctly orders by date", func(t *testing.T) {
		tm := types.NewTaskManager(nil)
		tm.NewTaskOnManager("test", "2026-07-27")
		tm.NewTaskOnManager("test2", "2026-07-26")

		want := []types.Task{
			types.NewTask(1, "test2", "2026-07-26"), // earlier date, sorts first
			types.NewTask(0, "test", "2026-07-27"),
		}

		got := List(tm)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("date order wrong: got %+v, wanted %+v", got, want)
		}
	})

	t.Run("undated tasks sort after dated tasks", func(t *testing.T) {
		tm := types.TaskManager{
			NextID: 2,
			Tasks: []types.Task{
				{ID: 0, Text: "no date", Due: ""},
				{ID: 1, Text: "has date", Due: "2026-07-27"},
			},
		}
		want := []types.Task{
			{ID: 1, Text: "has date", Due: "2026-07-27"},
			{ID: 0, Text: "no date", Due: ""},
		}

		got := List(tm)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("undated ordering wrong: got %+v, wanted %+v", got, want)
		}
	})

	t.Run("multiple undated tasks fall back to ID order", func(t *testing.T) {
		tm := types.TaskManager{
			NextID: 2,
			Tasks: []types.Task{
				{ID: 1, Text: "second", Due: ""},
				{ID: 0, Text: "first", Due: ""},
			},
		}
		want := []types.Task{
			{ID: 0, Text: "first", Due: ""},
			{ID: 1, Text: "second", Due: ""},
		}

		got := List(tm)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("undated tiebreak wrong: got %+v, wanted %+v", got, want)
		}
	})

	t.Run("correctly filters doneness", func(t *testing.T) {
		want := []types.Task{
			{ID: 0, Text: "test", Due: "2026-07-27", Done: false},
		}
		tm := types.TaskManager{
			NextID: 2,
			Tasks: []types.Task{
				{ID: 0, Text: "test", Due: "2026-07-27", Done: false},
				{ID: 1, Text: "test2", Due: "2026-07-26", Done: true},
			},
		}

		got := List(tm)
		if !reflect.DeepEqual(got, want) {
			t.Errorf("doneness filter wrong: got %+v, wanted %+v", got, want)
		}
	})

	t.Run("all tasks done returns empty result", func(t *testing.T) {
		tm := types.TaskManager{
			NextID: 2,
			Tasks: []types.Task{
				{ID: 0, Text: "a", Done: true},
				{ID: 1, Text: "b", Done: true},
			},
		}

		got := List(tm)
		if len(got) != 0 {
			t.Errorf("expected no tasks, got %+v", got)
		}
	})

	t.Run("doesn't fail on an empty list", func(t *testing.T) {
		want := []types.Task(nil)
		got := List(types.NewTaskManager(nil))
		if !reflect.DeepEqual(got, want) {
			t.Errorf("failed on empty list: got %+v, wanted %+v", got, want)
		}
	})

	t.Run("does not mutate the input TaskManager", func(t *testing.T) {
		original := types.TaskManager{
			NextID: 2,
			Tasks: []types.Task{
				{ID: 0, Text: "b", Due: "2026-07-28"},
				{ID: 1, Text: "a", Due: "2026-07-26"},
			},
		}
		inputCopy := types.TaskManager{
			NextID: original.NextID,
			Tasks:  append([]types.Task{}, original.Tasks...),
		}

		_ = List(original)

		if !reflect.DeepEqual(original, inputCopy) {
			t.Errorf("List mutated its input: got %+v, wanted unchanged %+v", original, inputCopy)
		}
	})
}
