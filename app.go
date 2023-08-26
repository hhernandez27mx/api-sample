package main

import (
	"fmt"
	"net/http"

	"azt.com/api-sample/infra/controller"
)

func init() {

}
func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Página de inicio")
	})
	http.HandleFunc("/contact", controller.HomeContact)
	http.HandleFunc("/contact/add", controller.AdddContactHandler)
	port := ":8080"
	http.ListenAndServe(port, nil)
}
