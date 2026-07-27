// taj-mace
// command to order existing todo tasks

package commands

import (
	"slices"
	"strings"

	"github.com/tajMace/todo/types"
)

func List(tm types.TaskManager) []types.Task {
	tasks := slices.Clone(tm.Tasks)
	tasks = slices.DeleteFunc(tasks, func(a types.Task) bool {
		return a.Done == true
	})

	slices.SortFunc(tasks, func(a, b types.Task) int {
		if a.Due != b.Due {
			return dateCompare(a.Due, b.Due)
		}

		return a.ID - b.ID
	})

	return tasks
}

func dateCompare(a, b string) int {
	if a == b {
		return 0
	}

	if a == "" {
		return 1
	}

	if b == "" {
		return -1
	}

	return strings.Compare(a, b)
}
