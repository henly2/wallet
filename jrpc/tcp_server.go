package jrpc

import (
	"net/rpc"
	"net"
	"log"
	"fmt"
)

// port like: ":8080"
func StartTcpServer(port string) error{
	log.Println("Tcp jrpc server start to create...")

	addr, err := net.ResolveTCPAddr("tcp", port)
	if err != nil {
		log.Printf("Tcp jrpc server failed to ResolveTCPAddr ", err.Error())
		return err;
	}

	listener, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Printf("Tcp jrpc server failed to start ", err.Error())
		return err;
	}

	log.Println("Tcp jrpc server successful listen...", addr)
	for{
		conn, err := listener.Accept();
		if err != nil {
			log.Printf("Fatal error: ", err.Error())
			continue
		}

		go rpc.ServeConn(conn)
	}

	return nil;
}

func StartTcpClient(addr string, method string, params string, res *string) error {
	client, err := rpc.Dial("tcp", addr)
	if err != nil {
		log.Printf("Fatal error: ", err.Error())
		return err
	}
	defer client.Close()

	err = client.Call(method, params, res)
	if err != nil {
		log.Printf("call failed: ", err.Error())
		return err
	}

	fmt.Printf("Reply: %s\n", *res)
	return nil
}