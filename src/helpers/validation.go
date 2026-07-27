// taj-mace
// shared validation helpers

package helpers

import (
	"time"

	"github.com/tajMace/todo/types"
)

func IsValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}

func GetValidTaskIndex(tm types.TaskManager, id int) int {
	for i, task := range tm.Tasks {
		if task.ID == id {
			return i
		}
	}
	return -1
}
