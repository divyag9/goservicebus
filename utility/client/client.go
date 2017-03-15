package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/divyag9/goservicebus/utility/packages/servicebus"
	"github.com/divyag9/goservicebus/utility/server"
)

func main() {
	delayMilliSeconds := flag.Int("delayMilliSeconds", 0, "Number of seconds to sleep")
	serviceBusEndPoint := flag.String("servicebus_endpoint", "http://servicebus.sgtec.io/Execute.svc/Execute", "The servicebus execute endpoint")

	flag.Parse()

	params := &servicebus.Params{}
	params.DelayMilliSeconds = int32(*delayMilliSeconds)
	request := &servicebus.Request{}
	request.Jsonrpc = "2.0"
	request.Method = "UTILITY.PING"
	request.Params = params

	count := 0
	var elapsed time.Duration
	for count < 20 {
		start := time.Now()
		response, err := callServiceBus(request, *serviceBusEndPoint)
		if err != nil {
			log.Fatalf("Error making servicebuscall: %v", err)
		}
		if response.Result != nil {
			fmt.Println("PingResponse: ", response.Result.PingResponse)
		} else {
			fmt.Println("ping error: ", response.Error)
		}
		elapsed = elapsed + time.Since(start)
		count++

	}
	average := elapsed / time.Duration(count)
	fmt.Println("Average elapsed: ", average)

}

func callServiceBus(request *servicebus.Request, serviceBusEndpoint string) (*servicebus.Response, error) {
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	// Contact the server and print out its response.
	busServer := server.NewServer(serviceBusEndpoint)
	response, err := busServer.CallServiceBus(requestBytes)
	if err != nil {
		return nil, err
	}

	return response, nil
}
