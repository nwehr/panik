package panik

import (
	"errors"
	"fmt"
	"runtime"
)

// OnError panics if err != nil
func OnError(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		panic(errors.New(err.Error() + fmt.Sprintf("; %s:%d", file, line)))
	}
}
