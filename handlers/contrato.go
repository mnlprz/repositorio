package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/mnlprz/go/repositorio/models"
	"github.com/mnlprz/go/repositorio/services"
)

func setHandlersContrato(w http.ResponseWriter, req *http.Request) {

	nup := req.URL.Query().Get("nup")
	contratos, err := services.GetContratos(nup)
	if err != nil {
		log.Fatal(err)
	}
	contratosJson, err := json.Marshal(contratos)
	if err != nil {
		log.Fatal(err)
	}
	_, err = w.Write(contratosJson)
	if err != nil {
		log.Fatal(err)
	}
}

func setHandlersContratoID(w http.ResponseWriter, req *http.Request) {

	switch req.Method {

	case "GET":
		id := req.URL.Query().Get("id")
		contratos, err := services.GetContratos(id)
		if err != nil {
			log.Fatal(err)
		}
		contratosJson, err := json.Marshal(contratos)
		if err != nil {
			log.Fatal(err)
		}
		_, err = w.Write(contratosJson)
		if err != nil {
			log.Fatal(err)
		}
	case "POST":

		var contrato models.Contrato
		err := json.NewDecoder(req.Body).Decode(&contrato)
		if err != nil {
			log.Fatal(err)
		}
		data := url.Values{}
		data.Add("nup", contrato.Nup)
		data.Add("codEntidad", contrato.CodEntidad)
		data.Add("codCentro", contrato.CodCentro)
		data.Add("numContrato", contrato.NumContrato)
		data.Add("codProd", contrato.CodProd)
		data.Add("codSubProd", contrato.CodSubProd)
		data.Add("moneda", contrato.Moneda)
		data.Add("saldo", contrato.Saldo)
		err = services.PostContrato(data)
		if err != nil {
			log.Fatal(err)
		}
	}

}
