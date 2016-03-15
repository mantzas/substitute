package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/mantzas/adaptlog"
	"github.com/mantzas/substitute/http/mux"
	"github.com/mantzas/substitute/log"
)

func main() {
	port := flag.Int("port", 8080, "port of the substitution service")
	portMgmt := flag.Int("mgmtport", 8081, "port of the substitution management service")
	flag.Parse()

	if *port == 0 || *portMgmt == 0 {
		flag.Usage()
		return
	}

	fmt.Printf("Service starting on port %d with management port %d.", *port, *portMgmt)
	fmt.Println()

	adaptlog.Configure(new(log.Logger), adaptlog.AnyLevel)

	go func() {

		fmt.Println("Starting management service.")
		adaptlog.Fatal(http.ListenAndServe(":8081", mux.GetMgmtServerMux()))
	}()

	fmt.Println("Starting service.")
	adaptlog.Fatal(http.ListenAndServe(":8080", mux.GetServerMux()))
}
