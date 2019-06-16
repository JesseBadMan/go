package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opentracing"
	"github.com/opentracing/opentracing-go"
	zipkinopentracing "github.com/openzipkin/zipkin-go-opentracing"
	"log"
	"os"
	"study/proto"
	"time"
)



type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}

func main() {
	// Create a new service. Optionally include some options here.
	apiURL := "http://localhost:9411/api/v2/spans"
	hostPort,_ := os.Hostname()
	serviceName := "go.micro.srv.greeter"

	collector, err := zipkinopentracing.NewHTTPCollector(apiURL)
	if err != nil {
		log.Fatalf("unable to create Zipkin HTTP collector: %v", err)
		return
	}
	tracer, err := zipkinopentracing.NewTracer(
		zipkinopentracing.NewRecorder(collector, false, hostPort, serviceName),
	)
	if err != nil {
		log.Fatalf("unable to create Zipkin tracer: %v", err)
		return
	}
	opentracing.InitGlobalTracer(tracer)

	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
		micro.RegisterTTL(time.Duration(10)*time.Second),
		micro.WrapHandler(wrapperTrace.NewHandlerWrapper()),
		micro.WrapClient(wrapperTrace.NewClientWrapper()),

	)

	// Init will parse the command line flags.
	service.Init()

	// Register handler
	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	// Run the server
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}


func main1() {
	// Create a new service. Optionally include some options here.
	service := micro.NewService(micro.Name("go.micro.srv.greeter"))
	service.Init()

	// Create new greeter client
	greeter := proto.NewGreeterClient("go.micro.srv.greeter", service.Client())

	// Call the greeter
	rsp, err := greeter.Hello(context.TODO(), &proto.HelloRequest{Name: "John"})
	if err != nil {
		fmt.Println(err)
	}

	// Print response
	fmt.Println(rsp.Greeting)
}