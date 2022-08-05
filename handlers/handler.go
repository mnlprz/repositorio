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

	http.HandleFunc("/entrada/{id}", func(w http.ResponseWriter, req *http.Request) {
		handlerEntradaID(w, req)
	})

	http.HandleFunc("/getentradas/{campo1}", func(w http.ResponseWriter, req *http.Request) {
		handlerGetEntradasID(w, req)
	})

	http.HandleFunc("/getcontratos/{nup}", func(w http.ResponseWriter, req *http.Request) {
		setHandlersContrato(w, req)
	})

	http.HandleFunc("/contrato/{id}", func(w http.ResponseWriter, req *http.Request) {
		setHandlersContratoID(w, req)
	})
}
