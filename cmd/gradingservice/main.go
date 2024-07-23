package main

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/tClown11/distributed-demo/grades"
	"github.com/tClown11/distributed-demo/log"
	"github.com/tClown11/distributed-demo/registry"
	"github.com/tClown11/distributed-demo/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName:      registry.GradingService,
		ServiceURL:       serviceAddress,
		RequiredServices: []registry.ServiceName{registry.LogService},
		ServiceUpdateURL: serviceAddress + "/services",
		HeartBeatURL:     serviceAddress + "/heartbeat",
	}

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.Registerhandlers)
	if err != nil {
		stlog.Fatal(err)
	}

	if logProvider, err := registry.GetProvider(registry.LogService); err == nil {
		fmt.Printf("Logging service found at: %v \n", logProvider)
		log.SetClientLogger(logProvider, r.ServiceName)
	}

	<-ctx.Done()

	fmt.Println("Shutting down grading service")
}
