package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
)

var serverIP = "unknown"
var serverPort = flag.Int("port", 8080, "server port")

func myHandler(w http.ResponseWriter, r *http.Request) {
	s := struct {
		Host       string
		ServerIP   string
		ServerPort int
		RequestURI string
		Header     interface{}
	}{
		Host:       r.Host,
		ServerIP:   serverIP,
		ServerPort: *serverPort,
		RequestURI: r.RequestURI,
		Header:     r.Header,
	}
	str, _ := json.Marshal(s)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, string(str))
}

func init() {
	flag.Parse()
	initServerIP()
}

func initServerIP() {
	addrs, _ := net.InterfaceAddrs()
	for _, addr := range addrs {
		if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				serverIP = ipnet.IP.String()
				break
			}
		}
	}
}

func main() {
	http.HandleFunc("/", myHandler)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *serverPort), nil))
}
