package services

import (
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/mnlprz/go/repositorio/database"
	"github.com/mnlprz/go/repositorio/models"
)

func GetContratos(nup string) ([]models.Contrato, error) {

	const query = "SELECT * FROM contratos WHERE nup = $1"
	var contratos []models.Contrato

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r, err := db.Query(query, nup)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()

	for r.Next() {
		var contrato models.Contrato
		err = r.Scan(&contrato.Id, &contrato.Nup, &contrato.CodEntidad, &contrato.CodCentro, &contrato.NumContrato, &contrato.CodProd, &contrato.CodSubProd, &contrato.Moneda, &contrato.Saldo, &contrato.CreatedAt, &contrato.UpdatedAt)
		if err != nil {
			log.Fatal(err)
		}
		contratos = append(contratos, contrato)
	}
	if r.Err() != nil {
		log.Fatal(err)
	}

	return contratos, nil
}

func PostContrato(data url.Values) error {

	const URL = "http://localhost:5556/postcontrato/"
	res, err := http.PostForm(URL, data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)
	return nil
}
