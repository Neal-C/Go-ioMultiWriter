//lint:file-ignore ST1006 because

package main

import (
	"bytes"
	"fmt"
	"io"
)

type Conn struct {
	io.Writer
}

func (self *Conn) Write(p []byte) (int, error){
	fmt.Println("writing to the connection", string(p));
	return self.Writer.Write(p);
}
func NewConn() *Conn {
	return &Conn{
		Writer: new(bytes.Buffer),
	}
}

type Server struct {
	//map[net.Conn]bool
	peers map[*Conn]bool
}

func NewServer() (server *Server) {
	server = &Server{
		peers: make(map[*Conn]bool),
	}

	for i := 0; i < 10; i++ {
		server.peers[NewConn()] = true;
	}

	return
}

func (self *Server) broadcast(message []byte) error {
	peers := []io.Writer{};

	for peer := range self.peers {
		peers = append(peers, peer);
	}

	multiWriter := io.MultiWriter(peers...);
	_ , err := multiWriter.Write(message);
	return err;
}

func main() {
	server := NewServer();
	server.broadcast([]byte("foo"));
}