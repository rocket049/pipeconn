/*	golang 标准 rpc 返回结构体时，如果某个结构体成员为零值，那么返回时该成员变量会被自动忽略。
	因此必须注意：客户端用来接受返回值的变量必须在使用前清零！
*/
package main

import (
	"fmt"
	"log"
	"net/rpc"

	"gitee.com/rocket049/pipeconn"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	var client *rpc.Client
	var err error
	conn, err := pipeconn.NewClientPipeConn("./pipe-server")
	if err != nil {
		return
	}
	client = rpc.NewClient(conn)
	defer client.Close()
	args := &Args{1, 3}
	var ret2 Quotient
	var reply int
	for i := 1; i < 10; i++ {
		args.A = i
		err = client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			log.Fatal("arith error:", err)
		} else {
			fmt.Printf("Arith: \n\t%d*%d=%d\n", args.A, args.B, reply)
		}
		ret2.Quo, ret2.Rem = 0, 0
		err = client.Call("Arith.Divide", args, &ret2)
		if err != nil {
			log.Fatal("arith error:", err)
		} else {
			fmt.Printf("\t%d/%d=%d -- %d\n", args.A, args.B, ret2.Quo, ret2.Rem)
		}
	}

}
