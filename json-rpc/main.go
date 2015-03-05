package main

import (
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"

	"golang.org/x/net/websocket"
)

// MultiplyArgs .
type MultiplyArgs struct {
	A int
	B int
}

// AverageNums .
type AverageNums []int

// Arith .
type Arith struct{}

// Multiply returns one number times another number
func (t *Arith) Multiply(args *MultiplyArgs, reply *int) error {
	*reply = args.A * args.B
	return nil
}

// Average calculates the average of a slice of ints, returning a float64
func (t *Arith) Average(nums *AverageNums, reply *float64) error {
	total := 0
	for _, num := range *nums {
		total += num
	}
	*reply = float64(total) / float64(len(*nums))
	return nil
}

func main() {
	rpc.Register(&Arith{})

	http.Handle("/conn", websocket.Handler(serve))
	http.ListenAndServe("localhost:7000", nil)
}

func serve(ws *websocket.Conn) {
	jsonrpc.ServeConn(ws)
}
