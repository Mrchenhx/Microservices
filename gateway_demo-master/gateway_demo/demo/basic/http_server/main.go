// @Author: Wilson 
// @Create: 2022/05/01 14:16:00
// @Description: TODO

package main

import (
	"log"
	"net/http"
	"time"
)

var (
	Addr = ":1210"
)

func main() {
	// create the route
	mux := http.NewServeMux()
	// set the route rule
	mux.HandleFunc("/bye", sayBye)
	//  create the server
	server := &http.Server{
		Addr:         Addr,
		WriteTimeout: time.Second * 3,
		Handler:      mux,
	}
	// listen the port and provide serve
	log.Println("Starting httpserver at" + Addr)
	log.Fatal(server.ListenAndServe())
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(1 * time.Second)
	w.Write([]byte("Bye bye, this is httpServer"))
}
