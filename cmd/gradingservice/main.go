package gradingservice

import (
	"context"
	"fmt"
	stlog "log"

	"github.com/tClown11/distributed-demo/grades"
	"github.com/tClown11/distributed-demo/registry"
	"github.com/tClown11/distributed-demo/service"
)

func main() {
	host, port := "localhost", "6000"
	serviceAddress := fmt.Sprintf("http://%s:%s", host, port)

	r := registry.Registration{
		ServiceName: registry.GradingService,
		ServiceURL:  serviceAddress,
	}

	ctx, err := service.Start(context.Background(),
		host,
		port,
		r,
		grades.Registerhandlers)
	if err != nil {
		stlog.Fatal(err)
	}
	<-ctx.Done()

	fmt.Println("Shutting down grading service")
}
