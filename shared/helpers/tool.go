package helpers

import (
	"strconv"
)

// ParseStringToUint64 pretty self explanatory.
// Especially used for query param conversion.
func ParseStringToUint64(str string) uint64 {
	var err error
	var id uint64
	id, err = strconv.ParseUint(str, 10, 32)
	if err != nil {
		return 0
	}
	return id
}
