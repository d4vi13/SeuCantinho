package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	httpSwagger "github.com/swaggo/http-swagger"
	_ "github.com/d4vi13/SeuCantinho/server/docs"

	"github.com/d4vi13/SeuCantinho/server/internal/database"
	"github.com/d4vi13/SeuCantinho/server/internal/routes"
)

// @tile API SeuCantinho
// @description API que gerencia os usuários, as reservas, os espaços e os pagamentos
// @version 1.0
// @contact Equipe Seu Cantinho
// @host localhost:8080
func main() {

	addr := flag.String("addr", ":"+os.Getenv("PORT"), "HTTP Network Address")
	flag.Parse()

	mux := http.NewServeMux()

	routes.RegisterRoutes(mux)

	//Registrar rota de documentação SWAGGER
    mux.Handle("/docs/", httpSwagger.WrapHandler)


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
