package trace

import (
	"os"
	"log"
	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
	zipkin "contrib.go.opencensus.io/exporter/zipkin"
	"go.opencensus.io/trace"
	"github.com/opentracing/opentracing-go"
	zipkinopentracing "github.com/openzipkin/zipkin-go-opentracing"

)

func TraceBoot(){
	apiUrl:="http://localhost:9410/api/spans"
	hostPort,_:=os.Hostname()
	serviceName:="go.micro.srv.greeter"
	loaclEndpoint,err:=openzipkin.NewEndpoint(serviceName,hostPort)
	if err!=nil {
		log.Fatal("failed to create the local zipkinEndpoint:%v",err)
	}
	reporter:=zipkinHTTP.NewReporter(apiUrl)
	ze:=zipkin.NewExporter(reporter,loaclEndpoint )
	trace.RegisterExporter(ze)
	trace.ApplyConfig(trace.Config{DefaultSampler:trace.AlwaysSample()})
	return
}

func TraceBoot1() {
	apiURL := "http://localhost:9411/api/v1/spans"
	hostPort,_ := os.Hostname()
	serviceName := "go.micro.srv.order"

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
	return
}

