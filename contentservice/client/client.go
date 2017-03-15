package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/divyag9/goservicebus/packages/servicebus"
	"github.com/divyag9/goservicebus/server"
)

var (
	serviceBusEndpoint = flag.String("servicebus_endpoint", "http://servicebus.qa01.local/Execute.svc/Execute", "The servicebus execute endpoint")
	contractorID       = flag.Int64("contractorid", 72494, "Contractor Id for the PUT call")
	orderNumber        = flag.Int64("ordernumber", 600016555, "OrderNumber for the PUT call")
	imageType          = flag.Int("imagetype", 1, "Imagetype for the PUT call")
	fileName           = flag.String("filename", "../testdata/e3e0f976-79a5-4059-ac23-d44386a6d4da.png", "Filename for the PUT call")
	imageWidth         = flag.Int("imagewidth", 100, "Imagewidth for the PUT call")
	imageHeight        = flag.Int("imageheight", 100, "Imageheight for the PUT call")
	releaseDate        = flag.String("releasedate", "2015-08-06", "Releasedate for the PUT call")
	deptCode           = flag.String("deptcode", "01", "Department code for the PUT call")
)

func main() {
	flag.Parse()
	request := getRequest()
	count := 0
	var elapsed time.Duration
	for count < 20 {
		start := time.Now()
		response, err := callServiceBus(request)
		if err != nil {
			log.Fatalf("Error making servicebuscall: %v", err)
		}
		if response.Result != nil {
			fmt.Println("put id: ", response.Result.ID)
		} else {
			fmt.Println("put error: ", response.Error)
		}
		//elapsed = elapsed + timeTrack(time.Now(), "servicebus call")
		elapsed = elapsed + time.Since(start)
		count++

	}
	average := elapsed / time.Duration(count)
	fmt.Println("Average elapsed: ", average)

}

func getFileContents(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("Error opening file: %s", err)
	}
	defer file.Close()

	stats, err := file.Stat()
	if err != nil {
		return nil, err
	}
	size := stats.Size()
	fileBytes := make([]byte, size)
	reader := bufio.NewReader(file)
	_, err = reader.Read(fileBytes)
	if err != nil {
		return nil, err
	}

	return fileBytes, nil
}

func getRequest() *servicebus.Request {
	// Create Params
	p := &servicebus.Params{}
	p.Contractorid = *contractorID
	p.Deptcode = *deptCode
	p.Filename = *fileName
	p.Imageheight = int32(*imageHeight)
	p.Imagetype = int32(*imageType)
	p.Imagewidth = int32(*imageWidth)
	p.Ordernumber = *orderNumber
	p.Releasedate = *releaseDate
	fileContents, err := getFileContents(*fileName)
	if err != nil {
		log.Fatalf("Error retrieving the file contents: %v", err)
	}
	p.Filecontents = fileContents

	request := &servicebus.Request{}
	request.Jsonrpc = "2.0"
	request.Method = "CONTENTSERVICE.PUT"
	request.Params = p

	return request
}

func callServiceBus(request *servicebus.Request) (*servicebus.Response, error) {
	start := time.Now()
	requestBytes, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	elapsed := time.Since(start)
	fmt.Println("Elapsed Marshal: ", elapsed)
	// Contact the server and print out its response.
	busServer := server.NewServer(server.CallServiceBus)
	response, err := busServer.CallServiceBus(requestBytes, *serviceBusEndpoint)
	if err != nil {
		return nil, err
	}

	return response, nil
}
