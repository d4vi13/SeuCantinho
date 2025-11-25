package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/d4vi13/SeuCantinho/server/internal/database"
	"github.com/d4vi13/SeuCantinho/server/internal/routes"
)

func main() {

	addr := flag.String("addr", ":"+os.Getenv("PORT"), "HTTP Network Address")
	flag.Parse()

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Starting server %s\n", *addr)
	err = http.ListenAndServe(*addr, mux)
	if err != nil {
		log.Fatal(err)
	}
	os.Exit(1)

}
