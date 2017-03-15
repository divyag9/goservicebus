package server

import (
	"errors"
	"reflect"
	"testing"

	"github.com/divyag9/goservicebus/utility/packages/servicebus"
)

type FakeServer struct {
	Response *servicebus.Response
	Err      error
}

func (f *FakeServer) CallServiceBus(requestBytes []byte) (*servicebus.Response, error) {
	if f.Err != nil {
		return nil, f.Err
	}
	return f.Response, nil
}

var cases = []struct {
	server           *Server
	fakeServer       *FakeServer
	requestBytes     []byte
	expectedResponse *servicebus.Response
	expectedErr      error
}{
	{
		server: NewServer("http://servicebus.sgtec.io/Execute.svc/Execute"),
		fakeServer: &FakeServer{
			Response: &servicebus.Response{Jsonrpc: "2.0",
				Result: &servicebus.Result{PingResponse: "Utility.Ping - Waited 1 ms"},
			},
			Err: nil,
		},
		requestBytes: nil,
		expectedResponse: &servicebus.Response{Jsonrpc: "2.0",
			Result: &servicebus.Result{PingResponse: "Utility.Ping - Waited 1 ms"},
		},
		expectedErr: nil,
	},
	{
		server: NewServer("http://servicebus.sgtec.io/Execute.svc/Execute"),
		fakeServer: &FakeServer{
			Response: &servicebus.Response{},
			Err:      errors.New("Fake Error"),
		},
		requestBytes:     nil,
		expectedResponse: nil,
		expectedErr:      errors.New("Fake Error"),
	},
}

func TestCallServiceBus(t *testing.T) {
	for _, c := range cases {
		c.server.ServiceBusCaller = c.fakeServer
		response, err := c.server.CallServiceBus(c.requestBytes)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}

		if !reflect.DeepEqual(c.expectedResponse, response) {
			t.Errorf("Expected %q but got %q", c.expectedResponse, response)
		}
	}
}
