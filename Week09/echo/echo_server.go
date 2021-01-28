package echo

import (
	"bufio"
	"context"
	"log"
	"net"
	"sync"
)

type Server struct {
	lis net.Listener
	addr string
}

func NewServer(addr string) *Server {
	s := new(Server)
	s.addr = addr
	return s
}

func (s *Server) Start(ctx context.Context) error {
	var err error
	s.lis, err = net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	go func() {
		for {
			conn, err := s.lis.Accept()
			if err != nil {
				log.Printf("accept error: %v", err)
				return
			}

			log.Printf("accept client: %s", conn.RemoteAddr().String())
			go s.handleConn(conn)
		}
	}()
	return nil
}

func (s *Server) Stop(ctx context.Context) error {
	return s.lis.Close()
}

func (s *Server) handleConn(conn net.Conn) {
	var wg sync.WaitGroup
	ch := make(chan string)
	rd := bufio.NewReader(conn)
	wd := bufio.NewWriter(conn)
	defer conn.Close()

	wg.Add(2)
	go s.handleRead(&wg, rd, ch)
	go s.handleWrite(&wg, wd, ch)
	wg.Wait()
}

func (s *Server) handleRead(wg *sync.WaitGroup, rd *bufio.Reader, ch chan <- string) {
	for {
		msg, err := rd.ReadString('\n')
		if err != nil {
			log.Printf("read error: %v", err)
			break
		}

		log.Printf("read msg: %s", msg)
		ch <- msg
	}
	close(ch)
	wg.Done()
}

func (s *Server) handleWrite(wg *sync.WaitGroup, wd *bufio.Writer, ch <- chan string) {
	for msg := range ch {
		_, err := wd.WriteString(msg)
		if err != nil {
			log.Printf("write error: %v", err)
			break
		}

		err = wd.Flush()
		if err != nil {
			log.Printf("write flush error: %v", err)
			break
		}
	}
	wg.Done()
}