// taj-mace
// function to remove a task

package commands

import (
	"slices"

	"github.com/tajMace/todo/src/helpers"
	"github.com/tajMace/todo/types"
)

func Remove(tm types.TaskManager, id int) (types.TaskManager, error) {
	idx := helpers.GetValidTaskIndex(tm, id)
	if idx < 0 {
		return tm, ErrIdNotInUse
	}

	tasks := slices.Clone(tm.Tasks)
	tasks = slices.Delete(tasks, idx, idx+1)
	tm.Tasks = tasks

	return tm, nil
}
