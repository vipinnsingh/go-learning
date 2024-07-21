package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
)

type server struct {
	Address string
	Proxy   *httputil.ReverseProxy
}

type Server interface {
	IsAlive()
	GetAddress()
	Serve(rw http.ResponseWriter, r *http.Request)
}

func (server *server) IsAlive() bool {
	return true
}

func (server *server) GetAddress() string {
	return server.Address
}

func (server *server) Serve(rw http.ResponseWriter, req *http.Request) {
	server.Proxy.ServeHTTP(rw, req)
}

func newServer(address string) *server {

	serverURL, err := url.Parse(address)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		handleErr(err)
	}
	fmt.Printf("serverURL: %v\n", serverURL)
	server := server{
		Address: address,
		Proxy:   httputil.NewSingleHostReverseProxy(serverURL),
	}

	return &server
}

type loadBalancer struct {
	Port            string
	Servers         []server
	RoundRobinCount int
}

func (lb *loadBalancer) getNextAvailableServer() server {

	server := lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	for !server.IsAlive() {
		lb.RoundRobinCount++
		server = lb.Servers[lb.RoundRobinCount%len(lb.Servers)]
	}
	lb.RoundRobinCount++

	return server
}

func (lb *loadBalancer) serveProxy(rw http.ResponseWriter, req *http.Request) {

	targetServer := lb.getNextAvailableServer()

	fmt.Printf("forwarding request to targetServer with address : %v\n", targetServer.Address)

	targetServer.Serve(rw, req)
}

func newLoadBalancer(port string, servers []server) *loadBalancer {

	return &loadBalancer{
		Port:            port,
		Servers:         servers,
		RoundRobinCount: 0,
	}
}

func main() {

	fmt.Println("Started main function")

	servers := []server{
		*newServer("https://www.swiggy.com"),
		*newServer("https://www.zomato.com"),
		*newServer("https://www.duckduckgo.com"),
	}

	fmt.Printf("servers: %v\n", servers)

	loadBalancer := newLoadBalancer("3000", servers)
	fmt.Printf("loadBalancer: %v\n", loadBalancer)

	handleRedirect := func(rw http.ResponseWriter, req *http.Request) {
		loadBalancer.serveProxy(rw, req)
	}

	http.HandleFunc("/", handleRedirect)

	fmt.Println("Server is running at localhost:", loadBalancer.Port)
	http.ListenAndServe(":"+loadBalancer.Port, nil)
}

func handleErr(err error) {
	if err != nil {
		fmt.Printf("err: %v\n", err)
		os.Exit(1)
	}
}
