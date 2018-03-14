package jrpc

import (
	"log"
	"net/rpc"
	"fmt"
)

// Call a JRPC to Http server
// @parameter: addr string, like "127.0.0.1:8080"
// @parameter: method string
// @parameter: params string
// @parameter: res *string
// @return: error
func CallJRPCToHttpServer(addr string, method string, params string, res *string) error {
	log.Println("Call JRPC to Http server...", addr)

	client, err := rpc.DialHTTP("tcp", addr)
	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}
	defer client.Close()

	err = client.Call(method, params, res)
	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	fmt.Println("Params: ", params)
	fmt.Println("Reply: ", *res)
	return nil
}
