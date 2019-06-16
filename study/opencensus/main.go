package main

import (
	"os"
	"log"
	"github.com/micro/go-micro"
	"fmt"
	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
	"contrib.go.opencensus.io/exporter/zipkin"
	"go.opencensus.io/trace"
	wrapperTrace "github.com/micro/go-plugins/wrapper/trace/opencensus"
	"study/proto"
	"context"
)


type Greeter struct{}

func (g *Greeter) Hello(ctx context.Context, req *proto.HelloRequest, rsp *proto.HelloResponse) error {
	rsp.Greeting = "Hello " + req.Name
	return nil
}


func main1(){

	apiUrl:="http://localhost:9411/api/v2/spans"
	hostPort,_:=os.Hostname()
	serviceName:="go.micro.srv.greeter"

	loaclEndpoint,err:=openzipkin.NewEndpoint(serviceName,hostPort)
	if err!=nil {
		log.Fatal("failed to create the local zipkinEndpoint:%v",err)
	}
	reporter:=zipkinHTTP.NewReporter(apiUrl)
	ze:=zipkin.NewExporter(reporter,loaclEndpoint)
	trace.RegisterExporter(ze)
	trace.ApplyConfig(trace.Config{DefaultSampler:trace.AlwaysSample()})


	service := micro.NewService(
		micro.Name("go.micro.srv.greeter"),
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


func main() {
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