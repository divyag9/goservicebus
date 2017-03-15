package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/divyag9/goservicebus/packages/servicebus"
)

type ServiceBusCaller func(requestBytes []byte, serviceBusEndPoint string) (*servicebus.Response, error)

//BusServer struct with the caller
type BusServer struct {
	CallServiceBus ServiceBusCaller
}

// CallServiceBus makes a request tot he servicebus endpoint
func CallServiceBus(requestBytes []byte, serviceBusEndPoint string) (*servicebus.Response, error) {
	req, err := http.NewRequest("POST", serviceBusEndPoint, bytes.NewBuffer(requestBytes))
	if err != nil {
		return nil, err
	}
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	body, _ := ioutil.ReadAll(resp.Body)
	body = bytes.TrimPrefix(body, []byte("\xef\xbb\xbf"))

	response := &servicebus.Response{}
	err = json.Unmarshal(body, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func NewServer(sbc ServiceBusCaller) *BusServer {
	return &BusServer{CallServiceBus: sbc}
}
