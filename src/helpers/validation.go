// taj-mace
// shared validation helpers

package helpers

import "time"

func IsValidDateFormat(date string) bool {
	_, err := time.Parse("1970-01-01", date)
	return err == nil
}
