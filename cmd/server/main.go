package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/d4vi13/SeuCantinho/internal/routes"
)

func main() {

	addr := flag.String("addr", ":4000", "HTTP Network Address")
	flag.Parse()

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	fmt.Printf("Starting server %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(1)

}
