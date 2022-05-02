// @Author: Wilson 
// @Create: 2022/05/01 14:16:00
// @Description: TODO

package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func main() {
	// create the connection pool
	transport := &http.Transport{
		DialContext: (&net.Dialer{
			Timeout: 30 * time.Second, // connection time
			KeepAlive: 30 * time.Second, // alive time
		}).DialContext,
		MaxIdleConns: 100, // the max Idle connections
		IdleConnTimeout: 90*time.Second, // Idle connection timeout
		TLSHandshakeTimeout: time.Second, // tls handshake timeout
		ExpectContinueTimeout: time.Second, // 100-continue status code timeout
	}
	// create the client
	client := &http.Client{
		Timeout: time.Second * 30,
		Transport: transport,
	}
	// request the url
	res, err := client.Get("http://127.0.0.1:1210/bye")
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	// read the content
	bds, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bds))
}
