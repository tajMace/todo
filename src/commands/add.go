// taj-mace

package commands

import (
	"errors"

	"github.com/tajMace/todo/src/helpers"
	"github.com/tajMace/todo/types"
)

var ErrIncorrectDateFormat = errors.New("due date must be in YYYY-MM-DD")

/*
 * Adds a task to the given task list
 */
func Add(tasks types.TaskManager, text, due string) (types.TaskManager, error) {
	if due != "" && !helpers.IsValidDateFormat(due) {
		return tasks, ErrIncorrectDateFormat
	}

	newTask := types.NewTask(tasks.NextID, text, due)

	return types.TaskManager{
		NextID: tasks.NextID + 1,
		Tasks:  append(tasks.Tasks, newTask),
	}, nil
}
