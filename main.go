package main

import (
	"net/http"

	"github.com/ProductStore/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
