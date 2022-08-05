package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/mnlprz/go/repositorio/models"
	"github.com/mnlprz/go/repositorio/services"
)

func handlerEntradaID(w http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case "GET":
		idParam := req.URL.Query().Get("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		entrada, err := services.GetEntrada(id)
		if err != nil {
			log.Fatal(err)
		}
		entradaJson, err := json.Marshal(entrada)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(entradaJson)
		if err != nil {
			log.Fatal(err)
		}

	case "POST":
		var entrada models.Entrada
		dec := json.NewDecoder(req.Body)
		err := dec.Decode(&entrada)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(entrada)
		err = services.InsertEntrada(entrada)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write([]byte("Insertado correctamente."))
		if err != nil {
			log.Fatal(err)
		}

	case "PUT":
		var entrada models.Entrada
		dec := json.NewDecoder(req.Body)
		err := dec.Decode(&entrada)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(entrada)
		err = services.UpdateEntrada(entrada)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write([]byte("Actualizado correctamente."))
		if err != nil {
			log.Fatal(err)
		}

	case "DELETE":
		idParam := req.URL.Query().Get("id")
		id, err := strconv.ParseInt(idParam, 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		err = services.DeleteEntrada(id)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write([]byte("Eliminado correctamente."))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func handlerGetEntradasID(w http.ResponseWriter, req *http.Request) {

	http.HandleFunc("/getentradas/{campo1}", func(w http.ResponseWriter, req *http.Request) {
		campo1 := req.URL.Query().Get("campo1")
		entradas, err := services.GetEntradas(campo1)
		if err != nil {
			log.Fatal(err)
		}
		entradaJson, err := json.Marshal(entradas)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(entradaJson)
		if err != nil {
			log.Fatal(err)
		}
	})
}
