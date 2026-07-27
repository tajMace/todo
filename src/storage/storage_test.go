// taj-mace
// small tests for load and save functionality

package storage

import (
	"reflect"
	"testing"

	"github.com/tajMace/todo/types"
)

func TestSaveAndLoadRoundTrip(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/todo.json"

	want := types.NewTaskManager([]types.Task{
		types.NewTask(0, "todo", ""),
		types.NewTask(1, "todo2", "2026-07-26"),
	})

	err := SaveTasks(path, want)
	if err != nil {
		t.Fatalf("SaveTasks returned an error: %v", err)
	}

	got, err := LoadTasks(path)
	if err != nil {
		t.Fatalf("LoadTasks returned an error: %v", err)
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("round trip mismatch: got %+v, want %+v", got, want)
	}
}

func TestLoadTasksMissingFile(t *testing.T) {
	dir := t.TempDir()
	path := dir + "/does-not-exist.json"

	got, err := LoadTasks(path)
	if err != nil {
		t.Fatalf("expected no error for missing file, got: %v", err)
	}

	want := types.NewTaskManager(nil)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}
