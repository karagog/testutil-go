// Package fakehttpserver implements a fake HTTP service for unit testing.
package fakehttpserver

import (
	"net/http"
	"net/url"
)

// FakeHTTPServer is a fake service that just returns whatever you tell it.
// Use it as follows:
//
// fakeServer := FakeHTTPServer{
//   ResponseJSON: <arbitrary JSON string>
//   ResponseHTTPStatus: <optional non-ok status>
// }
// s := httptest.NewServer(http.HandlerFunc(fakeServer.Handle))
// defer s.Close()
//
// Then you can create a HTTP client and it will hit our test service
// with the address 's.Listener.Addr().String()'.
type FakeHTTPServer struct {
	// This response will be given automatically regardless of the request.
	ResponseJSON       string
	ResponseHTTPStatus int

	// When a request is received, the URL that was requested will be stored here for inspection.
	LastURLReceived *url.URL
}

func (f *FakeHTTPServer) Handle(w http.ResponseWriter, r *http.Request) {
	f.LastURLReceived = r.URL
	if f.ResponseHTTPStatus != http.StatusOK && f.ResponseHTTPStatus != 0 {
		http.Error(w, "Test-injected error", f.ResponseHTTPStatus)
		return
	}
	status := http.StatusOK
	if f.ResponseHTTPStatus != 0 {
		status = f.ResponseHTTPStatus
	}
	w.WriteHeader(status)
	w.Write([]byte(f.ResponseJSON))
}
