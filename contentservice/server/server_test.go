package server

import (
	"errors"
	"reflect"
	"testing"

	"github.com/divyag9/goservicebus/packages/servicebus"
)

type FakeServer struct {
	Response *servicebus.Response
	Err      error
}

func (f *FakeServer) mockCallServiceBus(requestBytes []byte, serviceBusEndPoint string) (*servicebus.Response, error) {
	if f.Err != nil {
		return nil, f.Err
	}
	return f.Response, nil
}

var cases = []struct {
	f                *FakeServer
	requestBytes     []byte
	endPoint         string
	expectedResponse *servicebus.Response
	expectedErr      error
}{
	{
		f: &FakeServer{
			Response: &servicebus.Response{Jsonrpc: "2.0",
				Result: &servicebus.Result{Contractorid: 72494,
					Releasedate:   "2015-08-06T15:09:30",
					Scandate:      "2017-03-09T10:33:09",
					Imagetype:     1,
					Imagewidth:    100,
					Imageheight:   100,
					Deptcode:      "01",
					Descprefix:    "test",
					Desctext:      "test",
					Category:      "test",
					Ordernumber:   600016555,
					Archived:      "N",
					Datecreated:   "2017-03-09T10:33:09",
					Datemodefied:  "2017-03-09T10:33:09",
					Filesize:      180,
					ID:            1810448062,
					Imagefilename: "\\\\filer\\QA01\\ImageStore\\ServiceBus\\600\\016\\555\\da00563b-bb38-49b1-b3ef-29dbce63fbed.png",
					Thumbnailsize: 0,
					Webfilename:   "QA01/ImageStore/ServiceBus/600/016/555/da00563b-bb38-49b1-b3ef-29dbce63fbed.png",
				}},
			Err: nil,
		},
		requestBytes: nil,
		endPoint:     "",
		expectedResponse: &servicebus.Response{Jsonrpc: "2.0",
			Result: &servicebus.Result{Contractorid: 72494,
				Releasedate:   "2015-08-06T15:09:30",
				Scandate:      "2017-03-09T10:33:09",
				Imagetype:     1,
				Imagewidth:    100,
				Imageheight:   100,
				Deptcode:      "01",
				Descprefix:    "test",
				Desctext:      "test",
				Category:      "test",
				Ordernumber:   600016555,
				Archived:      "N",
				Datecreated:   "2017-03-09T10:33:09",
				Datemodefied:  "2017-03-09T10:33:09",
				Filesize:      180,
				ID:            1810448062,
				Imagefilename: "\\\\filer\\QA01\\ImageStore\\ServiceBus\\600\\016\\555\\da00563b-bb38-49b1-b3ef-29dbce63fbed.png",
				Thumbnailsize: 0,
				Webfilename:   "QA01/ImageStore/ServiceBus/600/016/555/da00563b-bb38-49b1-b3ef-29dbce63fbed.png",
			}},
		expectedErr: nil,
	},
	{
		f: &FakeServer{
			Response: &servicebus.Response{},
			Err:      errors.New("Fake Error"),
		},
		requestBytes:     nil,
		endPoint:         "",
		expectedResponse: nil,
		expectedErr:      errors.New("Fake Error"),
	},
}

func TestCallServiceBus(t *testing.T) {
	for _, c := range cases {
		busServer := NewServer(c.f.mockCallServiceBus)
		response, err := busServer.CallServiceBus(c.requestBytes, c.endPoint)
		if !reflect.DeepEqual(err, c.expectedErr) {
			t.Errorf("Expected err to be %q but it was %q", c.expectedErr, err)
		}

		if !reflect.DeepEqual(c.expectedResponse, response) {
			t.Errorf("Expected %q but got %q", c.expectedResponse, response)
		}
	}
}
