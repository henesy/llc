package main

import (
	"os"
	"fmt"
)

// If there is an error, call fatal in a pleasant format
func efatal(err error, s ...interface{}) {
	if err != nil {
		s = append([]interface{}{"err:"}, s...)
		s = append(s, "- " + err.Error())
	
		fatal(s...)
	}
}

// End the program while printing an error
func fatal(s ...interface{}) {
	fmt.Fprintln(os.Stderr, s...)
	os.Exit(1)
}
