package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/tClown11/distributed-demo/log"
	"github.com/tClown11/distributed-demo/service"
)

func main() {
	log.Run("./distributed.log")
	host, port := "localhost", "4000"
	ctx, err := service.Start(
		context.Background(),
		"Log Service",
		host,
		port,
		log.RegisterHandlers,
	)
	if err != nil {
		stlog.Fatalln(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down log service")
}
