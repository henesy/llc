# llc

** This is a work in progress. **

Less Lossy Calculator.

Calculator for big maths using polish notation for ease of implementation. 

Reads from stdin by default. 

Writes to stdout by default.

## Features

### Planned

Comments as `#`.

Arithmetic:

	* / - + !

Binary (might not do):

	^ & | ~

Functions:

	pow(x, to)		# Raise `x` to the `to` power
	log(b [n], x)	# Log base `b` of `x` with `n` as a base specifying natural log
	flatten(x)		# Given a radical `x`, emit as float-like value
	let(x, be)		# Set `x` equal to the result of `be`

Constants:

	π or Π			# 3.14 …
	e				# 2.71 …

Done:

☺

## Build

	go build

## Run

	./llc

## Install

	go install

## Usage

Utf-8 is the default. 

Parentheses are optional for singleton expressions. 

An example session:

	; llc
	(let σ (* 2 3))
	σ
	6				# Output
	/ σ 2
	3				# Output

## Contributing

Please do.
