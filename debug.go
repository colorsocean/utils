package utils

import (
	"fmt"
	"runtime"
)

func FileLine(errors ...error) string {
	_, file, line, ok := runtime.Caller(1)

	if ok {
		return fmt.Sprintf("[%s:%d]% v", file, line, errors)
	} else {
		return "unknown"
	}
}
