package main

import (
	"io"
	"os"
	"bufio"
	"flag"
)

var (
	interp	bool	// Interpreted mode - print all assignments
	noLazy	bool	// Lazy mode - evaluate as needed
)


// Less lossy calculator
func main() {
	flag.BoolVar(&interp, "i", false, "Enable explicit interpreted mode, emits all assignment values")
	flag.BoolVar(&noLazy, "l", false, "Disable lazy evaluation of variable assignments")
	flag.Parse()

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	
	// Read in expressions
	for {
		r, _, err := in.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}

			fatal("err: could not read rune - ", err)
		}
		
		// TODO - lex/parse into s-expressions
		
		// Output
		_, err = out.WriteRune(r)
		efatal(err, "could not write rune")
		out.Flush()
	}
	
	out.Flush()
	
}
