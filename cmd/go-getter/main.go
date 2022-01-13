package main

import (
	"log"
	"net/url"

	"github.com/hashicorp/go-getter"
)

func main() {
	var u url.URL
	u.Scheme = "http"
	u.Host = "127.0.0.1:9090"
	u.Path = "/gitrepository/default/openapis/42c153c36cd5226e299a2dd74e507aa288c13e8f.tar.gz"
	u.RawQuery = "checksum=0f916722e794c757c90da055f6f4b59bae86b16152a9f5c9cae604743b515ddf"

	dst := "/home/kurtis/wrkspc/github.com/angelokurtis/go-laboratory/bin/42c153c36cd5226e299a2dd74e507aa288c13e8f"
	log.Printf("getting from %s to %s", u.String(), dst)

	err := getter.GetAny(dst, u.String())
	dieIfNotNil(err)
}

func dieIfNotNil(err error) {
	if err != nil {
		panic(err)
	}
}
