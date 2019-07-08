// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package testserver

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
)

// TestServer wraps a http.Server that runs on a separate goroutine
// and is intended for unit testing.
type TestServer struct {
	Port   int
	server *http.Server
}

// CreateAndStart creates and returns a TestServer that waits for web
// requests on the provided port using the provided multiplexer on a
// different goroutine.
//
// If the port is 0, a port number is dynamically chosen. This port
// can be retrieved by checking the Port field of the TestServer struct.
func CreateAndStart(port int, mux *http.ServeMux) (*TestServer, error) {
	ln, err := net.Listen("tcp", fmt.Sprintf("127.0.0.1:%d", port))
	if err != nil {
		return nil, err
	}

	port, err = strconv.Atoi(strings.Split(ln.Addr().String(), ":")[1])
	if err != nil {
		return nil, err
	}
	server := &http.Server{Handler: mux}

	go func() {
		_ = server.Serve(ln)
	}()

	return &TestServer{
		Port:   port,
		server: server,
	}, nil
}

// Stop immediately closes all connections and shuts down the server.
func (ms TestServer) Stop() {
	_ = ms.server.Close()
}
