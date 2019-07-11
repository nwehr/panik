[![Build Status](https://travis-ci.org/nwehr/panik.svg?branch=master)](https://travis-ci.org/nwehr/panik)

Introduction
------------

Panik provides a one-liner solution break out of a function early when you encounter errors. 

Install
-------

```
$ go get github.com/nwehr/panik
```

Usage
-----

```go
package main

import (
	"github.com/nwehr/panik"
	"errors"
)

func main() {
	// we can only recover a panic in a deferred function
	defer func() {
		if err := recover(); err != nil {
			println(err.(error).Error())
		}
	}()
	
	str, err := stringOrError() 

	// we will only panic if err is not nil
	panik.OnError(err) 
	
	// we never get this far if we panic
	println(str)
}

func stringOrError() (string, error) {
	return "", errors.New("Some error")
}
```