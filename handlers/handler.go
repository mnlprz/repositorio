package handlers

import (
	"net/http"
)

func SetHandlers() {

	http.HandleFunc("/cargatablaofertas", func(w http.ResponseWriter, req *http.Request) {
		setHandlersTablaOfertas(w, req)
	})

	http.HandleFunc("/cargatablacontratos", func(w http.ResponseWriter, req *http.Request) {
		setHandlersTablaContratos(w, req)
	})

	http.HandleFunc("/contrato/{contrato}", func(w http.ResponseWriter, req *http.Request) {
		setHandlersContrato(w, req)
	})

	http.HandleFunc("/contratos/{nup}", func(w http.ResponseWriter, req *http.Request) {
		setHandlersContratos(w, req)
	})
}
