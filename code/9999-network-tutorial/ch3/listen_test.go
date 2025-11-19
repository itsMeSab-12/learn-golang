package ch3

import (
	"fmt"
	"net"
	"testing"
)

func Listen() error {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return err
	}

	defer func() { _ = listener.Close() }()
	fmt.Printf("bound to %q", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go func(c net.Conn) {
			defer c.Close()

			//Logic
		}(conn)
	}
}

func TestListener(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		t.Fatal(err)
	}

	defer func() { _ = listener.Close() }()
	t.Logf("bound to %q", listener.Addr())

	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Fatal(err)
		}

		go func(c net.Conn) {
			defer c.Close()

			//Logic
		}(conn)
	}
}
