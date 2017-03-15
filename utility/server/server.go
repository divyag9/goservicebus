package server

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/divyag9/goservicebus/utility/packages/servicebus"
)

// ServiceBusCaller interface for calling ServiceBus
type ServiceBusCaller interface {
	CallServiceBus(requestBytes []byte) (*servicebus.Response, error)
}

// Server servicebus
type Server struct {
	ServiceBusCaller
}

// Caller interface for servicebus
type Caller struct {
	serviceBusEndPoint string
}

// CallServiceBus makes a request tot he servicebus endpoint
func (c *Caller) CallServiceBus(requestBytes []byte) (*servicebus.Response, error) {
	req, err := http.NewRequest("POST", c.serviceBusEndPoint, bytes.NewBuffer(requestBytes))
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

// NewServer creates new servicebus  server
func NewServer(serviceBusEndPoint string) *Server {
	return &Server{ServiceBusCaller: &Caller{serviceBusEndPoint: serviceBusEndPoint}}
}
