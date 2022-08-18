package services

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"

	"github.com/mnlprz/go/repositorio/database"
	"github.com/mnlprz/go/repositorio/models"
)

func GetContrato(id int64) (models.Contrato, error) {

	const query = "SELECT * FROM contratos WHERE id = $1"
	var contrato models.Contrato

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := db.QueryRow(query, id).Scan(&contrato.Id, &contrato.Nup, &contrato.CodEntidad, &contrato.CodCentro, &contrato.NumContrato, &contrato.CodProd, &contrato.CodSubProd, &contrato.Moneda, &contrato.Saldo, &contrato.CreatedAt, &contrato.UpdatedAt)
	if r == sql.ErrNoRows {
		return contrato, errors.New("No se encontro ID")
	}
	return contrato, nil
}

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
