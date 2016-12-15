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

	draw(d, theX)
	print("\n")
	draw(d, stripes)
	print("\n")
	draw(d, diamond)
}

var (
	newline = byte(10)
	dot     = byte(46)
	pipe    = byte(124)
	x       = byte(120)
)

func draw(d int, formatter func(line []byte, i, d int)) {
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
		formatter(line, i, d)
		print(string(line))
	}
}

func theX(line []byte, i, d int) {
	line[2*i+1] = x
	line[2*d-2*i-1] = x
}

func stripes(line []byte, i, d int) {
	for i := 0; i < d/2+d%2; i++ {
		line[4*i+1] = x
	}
}

func diamond(line []byte, i, d int) {
	i = (i + d/2) % d
	line[2*i+1] = x
	line[2*d-2*i-1] = x
}
