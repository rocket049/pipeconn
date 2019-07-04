package main

import (
	"net/rpc"

	"gitee.com/rocket049/pipeconn"
)

import "errors"

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}

	quo.Rem = args.A % args.B
	quo.Quo = args.A / args.B

	return nil
}

func main() {
	arith := new(Arith)
	server := rpc.NewServer()
	server.Register(arith)
	conn := pipeconn.NewServerPipeConn()
	server.ServeConn(conn)
}
