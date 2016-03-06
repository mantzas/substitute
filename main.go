package main

import (
	"flag"
	"fmt"
	"github.com/mantzas/substitute/http/middleware"
	"log"
	"net/http"
)

func handler1() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("handler1"))
	})
}

func handler2() http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

        w.WriteHeader(http.StatusOK)
        w.Write([]byte("handler2"))
	})
}

func main() {
	port := flag.Int("port", 8080, "port of the substitution service")
	portMgmt := flag.Int("mgmtport", 8081, "port of the substitution service")
	flag.Parse()

	if *port == 0 || *portMgmt == 0 {
		flag.Usage()
		return
	}

	fmt.Printf("Service starting on port %d with management port %d.", *port, *portMgmt)
	fmt.Println()

	go func() {

		fmt.Println("Starting management service.")
		log.Fatal(http.ListenAndServe(":8081", getMgmtServerMux()))
	}()

	fmt.Println("Starting service.")
	log.Fatal(http.ListenAndServe(":8080", getServerMux()))
}

func getMgmtServerMux() *http.ServeMux {

	serverMux := http.NewServeMux()
	serverMux.Handle("/", middleware.DefaultGetMiddleware(handler1()))
	return serverMux
}

func getServerMux() *http.ServeMux {

	serverMux := http.NewServeMux()
	serverMux.Handle("/", middleware.DefaultGetMiddleware(handler2()))
	return serverMux
}
