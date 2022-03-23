package main

import (
	"crypto/tls"
	"fmt"
	"net/http"
	"os"

	"github.com/hashicorp/go-cleanhttp"
	"k8s.io/client-go/rest"
	"moul.io/http2curl"
)

const defaultHTTPTransportMaxIdleConnsPerHost = 10

func newHttpClient(c *rest.Config) (*http.Client, error) {

	httpClient := cleanhttp.DefaultPooledClient()
	transport := httpClient.Transport.(*http.Transport)

	transport.MaxIdleConnsPerHost = defaultHTTPTransportMaxIdleConnsPerHost

	tlsConfig := transport.TLSClientConfig
	if tlsConfig == nil {
		tlsConfig = &tls.Config{}
		transport.TLSClientConfig = tlsConfig
	}
	tlsConfig.MinVersion = tls.VersionTLS12

	if c.Insecure {
		tlsConfig.InsecureSkipVerify = true
	}

	return httpClient, nil
}

type Interceptor struct {
	http.RoundTripper
	config *rest.Config
}

func (i *Interceptor) RoundTrip(request *http.Request) (*http.Response, error) {
	path, err := os.Getwd()
	if err == nil {
		_ = os.WriteFile(path+"/ca.crt", i.config.CAData, 0644)
	}

	response, err := i.RoundTripper.RoundTrip(request)

	command, _ := http2curl.GetCurlCommand(request)
	fmt.Println(command.String() + " --cacert '" + path + "/ca.crt'")
	return response, err
}
