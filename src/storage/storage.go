// taj-mace
// functions to assist with storage and recall of data

package storage

import (
	"encoding/json"
	"os"

	"github.com/tajMace/todo/types"
)

func LoadTasks(path string) (types.TaskManager, error) {
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return types.NewTaskManager(nil), nil
	} else if err != nil {
		return types.NewTaskManager(nil), err
	}

	var tm types.TaskManager
	err = json.Unmarshal(data, &tm)
	if err != nil {
		return types.NewTaskManager(nil), err
	}

	return tm, nil
}

func SaveTasks(path string, tm types.TaskManager) error {
	data, err := json.MarshalIndent(tm, "", "\t")
	if err != nil {
		return err
	}

	err = os.WriteFile(path, data, 0644)
	if err != nil {
		return err
	}

	return nil
}
