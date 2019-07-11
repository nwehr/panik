package panik

import (
	"errors"
	"fmt"
	"runtime"
)

var fnPanic func(v interface{})

// OnError panics if err != nil
func OnError(err error) {
	if err != nil {
		_, file, line, _ := runtime.Caller(1)
		fnPanic(errors.New(err.Error() + fmt.Sprintf("; %s:%d", file, line)))
	}
}

func init() {
	fnPanic = func(v interface{}) {
		panic(v)
	}
}
