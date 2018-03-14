package jrpc

import (
	"log"
	"net"
	"net/http"
)

// Start a JRPC Http Server2
// @parameter: port string, like ":8080"
// @return: error
func StartJRPCHttpServer2(port string) error {
	log.Println("Start JRPC Http server...", port)

	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	log.Println("Start JRPC Http server successfully, listen on port: ", port)

	err = http.Serve(listener, nil)
	if err != nil {
		log.Println("Error: ", err.Error())
		return err
	}

	return nil
}
