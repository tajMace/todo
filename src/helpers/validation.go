// taj-mace
// shared validation helpers

package helpers

import "time"

func IsValidDateFormat(date string) bool {
	_, err := time.Parse("2006-01-02", date)
	return err == nil
}
