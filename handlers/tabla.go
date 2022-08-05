package handlers

import (
	"log"
	"net/http"

	"github.com/mnlprz/go/repositorio/services"
)

func setHandlersTabla() {

	http.HandleFunc("/cargatablaofertas", func(w http.ResponseWriter, req *http.Request) {
		err := services.CargaTablaOfertas()
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write([]byte("Tabla OFERTAS cargada exitosamente."))
		if err != nil {
			log.Fatal(err)
		}
	})

	http.HandleFunc("/cargatablacontratos", func(w http.ResponseWriter, req *http.Request) {
		err := services.CargaTablaContratos()
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write([]byte("Tabla CONTRATOS cargada exitosamente."))
		if err != nil {
			log.Fatal(err)
		}
	})
}
