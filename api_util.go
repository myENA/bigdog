package bigdog

import (
	"net"
	"net/http"
	"runtime"
	"time"
)

// defaultHTTPClient returns the http client used when an implementor does not provide their own
func defaultHTTPClient() *http.Client {
	// pooled transport config shamelessly borrowed from
	// https://github.com/hashicorp/go-cleanhttp/blob/v0.5.1/cleanhttp.go
	return &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		},
	}
}
