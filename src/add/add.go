// taj-mace

package add

import (
	"github.com/tajMace/todo/types"
)

/*
 * Adds a task to the given task list
 */
func Add(tasks types.TaskManager, text, due string) types.TaskManager {
	newTask := types.Task{
		ID:   tasks.NextID,
		Text: text,
		Due:  due,
		Done: false,
	}
	return types.TaskManager{
		NextID: tasks.NextID + 1,
		Tasks:  append(tasks.Tasks, newTask),
	}
}
