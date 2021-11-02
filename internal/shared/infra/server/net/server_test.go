package server

import (
	"bufio"
	"context"
	"feeder-service/internal/products/application/count"
	"feeder-service/internal/products/application/create"
	"feeder-service/internal/products/application/increase"
	"feeder-service/internal/products/infra/server/net/handler"
	"feeder-service/internal/products/infra/storage/fs"
	"feeder-service/internal/products/infra/storage/in_memory"
	"feeder-service/internal/shared/domain/config"
	"feeder-service/internal/shared/infra/counter"
	"fmt"
	"net"
	"testing"
	"time"
)

func setupTest() (context.Context, Server, config.ServerConfig) {
	ctx := context.Background()
	serverConfig := config.ServerConfig{
		Port:      4000,
		ConnLimit: 2,
		Timeout:   3,
	}

	persistRepository := fs.NewProductRepository()
	tempRepository := in_memory.NewProductRepository()

	invalidCounter := &counter.Counter{}
	duplicatedCounter := &counter.Counter{}

	createUC := create.NewCreateProductUseCase(persistRepository, tempRepository)
	increaseInvalidUC := increase.NewIncreaseInvalidProductUseCase(invalidCounter)
	increaseDuplicatedUC := increase.NewIncreaseDuplicatedProductUseCase(duplicatedCounter)
	countUC := count.NewCountProductsUseCase(tempRepository, invalidCounter, duplicatedCounter)

	createHandler := handler.NewCreateHandler(createUC, increaseInvalidUC, increaseDuplicatedUC)
	reportHandler := handler.NewReportHandler(countUC)

	ctx, srv := New(ctx, serverConfig, createHandler, reportHandler)

	return ctx, srv, serverConfig
}

func TestShutdownWithTimeout(t *testing.T) {
	ctx, srv, serverConfig := setupTest()

	defer srv.Shutdown()
	start := time.Now()
	err := srv.Run(ctx)
	elapsed := time.Since(start) / time.Second

	if err.Error() != "Server shutdown!" {
		t.Errorf("Expected: %v, got: %v", "Server shutdown!", err.Error())
	}
	if elapsed != time.Duration(serverConfig.Timeout) {
		t.Errorf("Expected to end: %v, got: %v", serverConfig.Timeout, elapsed)
	}
}

func TestShutdownWithTerminateCommand(t *testing.T) {
	ctx, srv, serverConfig := setupTest()

	go func() {
		conn, _ := net.Dial("tcp", ":"+fmt.Sprint((serverConfig.Port)))

		fmt.Fprintf(conn, "terminate\n")
	}()

	defer srv.Shutdown()
	start := time.Now()
	err := srv.Run(ctx)
	elapsed := time.Since(start) / time.Second

	if err.Error() != "Server shutdown!" {
		t.Errorf("Expected: %v, got: %v", "Server shutdown!", err.Error())
	}
	if elapsed >= time.Duration(serverConfig.Timeout) {
		t.Errorf("Expected to end immediately, got: %v", elapsed)
	}
}

func TestConnectionLimits(t *testing.T) {
	ctx, srv, serverConfig := setupTest()

	errorMsg := "There are already " + fmt.Sprint(serverConfig.ConnLimit) + " clients connected."

	for i := 0; i <= int(serverConfig.ConnLimit); i++ {
		go func(i int) {
			conn, _ := net.Dial("tcp", ":"+fmt.Sprint((serverConfig.Port)))

			message, _ := bufio.NewReader(conn).ReadString('\n')
			if i == int(serverConfig.ConnLimit) {
				if message != errorMsg {
					t.Errorf("Expected: %v, got: %v", errorMsg, message)
				} else {
					t.Logf("Properly message received: %v", message)
				}
			}
		}(i)
	}

	defer srv.Shutdown()
	start := time.Now()
	err := srv.Run(ctx)
	elapsed := time.Since(start) / time.Second

	if err.Error() != "Server shutdown!" {
		t.Errorf("Expected: %v, got: %v", "Server shutdown!", err.Error())
	}
	if elapsed != time.Duration(serverConfig.Timeout) {
		t.Errorf("Expected to end: %v, got: %v", serverConfig.Timeout, elapsed)
	}

}
