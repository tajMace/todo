// taj-mace
// function to mark a task as completed

package commands

import (
	"errors"
	"slices"

	"github.com/tajMace/todo/src/helpers"
	"github.com/tajMace/todo/types"
)

var ErrIdNotInUse = errors.New("ID not associated with any incomplete tasks")

func Done(tm types.TaskManager, id int) (types.TaskManager, error) {
	res := slices.Clone(tm.Tasks)

	idx := helpers.GetValidTaskIndex(tm, id)
	if idx == -1 {
		return tm, ErrIdNotInUse
	}

	res[idx].MarkDone()
	tm.Tasks = res
	return tm, nil
}
