package main

import (
	"log"
	"os"
	"strconv"
)

func main() {
	d, err := strconv.Atoi(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	for _, f := range []formatter{
		theX,
		stripes,
		diamond,
	} {
		draw(d, f)
		print("\n")
	}
}

var (
	newline = byte(10)
	dot     = byte(46)
	pipe    = byte(124)
	x       = byte(120)
)

type formatter func(line []byte, i, d int)

func draw(d int, f formatter) {
	tmpl := make([]byte, d*2+2)
	tmpl[len(tmpl)-1] = newline
	tmpl[len(tmpl)-2] = pipe
	for i := 0; i < d; i++ {
		tmpl[2*i] = pipe
		tmpl[2*i+1] = dot
	}

	for i := 0; i < d; i++ {
		line := make([]byte, len(tmpl))
		copy(line, tmpl)
		f(line, i, d)
		print(string(line))
	}
}

func theX(line []byte, i, d int) {
	line[offset(i)] = x
	line[offset(d-i-1)] = x
}

func stripes(line []byte, i, d int) {
	for i := 0; i < d/2+d%2; i++ {
		line[offset(2*i)] = x
	}
}

func diamond(line []byte, i, d int) {
	r := i + d/2
	line[offset(r%d+r/d)] = x
	line[offset((d-i+d/2)%d-r/d)] = x
}

func offset(col int) int {
	return 2*col + 1
}
