package main

import (
	"net/http"

	"github.com/silassilverio/Loja/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
