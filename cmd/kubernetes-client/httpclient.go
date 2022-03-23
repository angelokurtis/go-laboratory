package main

import (
	"context"
	"net/http"

	"k8s.io/client-go/rest"
	"moul.io/http2curl"

	"github.com/angelokurtis/go-laboratory/internal/log"
)

func newHttpClient(c *rest.Config) (*http.Client, error) {
	config := *c

	transport, err := rest.TransportFor(&config)
	if err != nil {
		return nil, err
	}
	var httpClient *http.Client
	if transport != http.DefaultTransport || config.Timeout > 0 {
		httpClient = &http.Client{
			Transport: &Interceptor{RoundTripper: transport, config: &config},
			Timeout:   config.Timeout,
		}
	} else {
		httpClient = http.DefaultClient
	}

	return httpClient, nil
}

type Interceptor struct {
	http.RoundTripper
	config *rest.Config
}

func (i *Interceptor) RoundTrip(request *http.Request) (*http.Response, error) {
	//path, err := os.Getwd()
	//if err == nil {
	//	_ = os.WriteFile(path+"/ca.crt", i.config.CAData, 0644)
	//}

	response, err := i.RoundTripper.RoundTrip(request)

	req2 := request.Clone(context.TODO())
	delete(req2.Header, "Authorization")
	delete(req2.Header, "Accept")
	command, _ := http2curl.GetCurlCommand(req2)
	log.Debugf(">>> %s", command.String())
	if response != nil {
		log.Debugf("<<< %s %s", response.Proto, response.Status)
	}
	//fmt.Println(command.String() + " --cacert '" + path + "/ca.crt'")
	return response, err
}
