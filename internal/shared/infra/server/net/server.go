package server

import (
	"context"
	"errors"
	product_handlers "feeder-service/internal/products/infra/server/net/handler"
	"feeder-service/internal/shared/domain/config"
	"feeder-service/internal/shared/infra/counter"
	"feeder-service/internal/shared/infra/server/net/handler"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"time"
)

type server struct {
	config        config.ServerConfig
	listener      net.TCPListener
	wg            sync.WaitGroup
	limiter       *counter.Counter
	handler       handler.Handler
	reportHandler product_handlers.ReportHandler
}

// Server is a TCP server that limits connections
type Server interface {
	Run(ctx context.Context) error
	Shutdown()
}

// New creates a new server and return it with context
func New(ctx context.Context, serverConfig config.ServerConfig, handler handler.Handler, reportHandler product_handlers.ReportHandler) (context.Context, Server) {
	srv := &server{
		config:        serverConfig,
		limiter:       &counter.Counter{},
		handler:       handler,
		reportHandler: reportHandler,
	}

	srv.wg.Add(1)

	return serverContext(ctx, serverConfig.Timeout), srv
}

// Run launches the server and accepts connections
func (s *server) Run(ctx context.Context) error {
	defer s.wg.Done()

	portStr := strconv.Itoa(s.config.Port)

	l, err := net.ListenTCP(s.config.Network, &net.TCPAddr{
		Port: s.config.Port,
	})
	if err != nil {
		return err
	}
	s.listener = *l

	log.Println("Server running on port", portStr)

	go func() {
		for {
			c, err := l.Accept()
			if err != nil {
				fmt.Println("Error connecting:", err)
				break
			}
			if !isMaxConnectionsLimit(s, c) {
				handleConnection(ctx, s, c)
			}
		}
	}()

	<-ctx.Done()
	return errors.New("Server shutdown by timeout.")
}

// Shutdown closes listener and wait that the rest of connections were closed
func (s *server) Shutdown() {
	fmt.Println("Shutdown")
	s.listener.Close()
	s.wg.Wait()
	s.reportHandler.Handle()
}

// serverContext prepares context for a graceful shutdown of server after a timeout or SIGINT signal
func serverContext(ctx context.Context, timeout uint8) context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	ctx, cancel := context.WithTimeout(ctx, time.Duration(timeout)*time.Second)
	go func() {
		<-c
		cancel()
	}()

	return ctx
}

// isMaxConnectionsLimit checks if maximum number of connections is exceeded
func isMaxConnectionsLimit(s *server, c net.Conn) bool {
	if s.limiter.Value() == uint64(s.config.ConnLimit) {
		_, err := c.Write([]byte("There are already 5 clients connected."))
		if err != nil {
			fmt.Println("Error writting:", err.Error())
		}
		c.Close()
		return true
	}
	return false
}

// handleConnection manages all states of a valid connection
func handleConnection(ctx context.Context, s *server, c net.Conn) {
	s.wg.Add(1)
	s.limiter.Inc()

	fmt.Println("Client " + c.RemoteAddr().String() + " connected")

	go func() {
		ctxCancel, cancel := context.WithCancel(ctx)

		go func() {
			s.handler.Handle(c)
			cancel()
		}()

		<-ctxCancel.Done()
		closeConn(s, c)
	}()
}

// closeConn closes connection leaving a slot to another one
func closeConn(s *server, c net.Conn) {
	c.Close()
	fmt.Println("Client " + c.RemoteAddr().String() + " left.")

	s.limiter.Dec()
	s.wg.Done()
}
