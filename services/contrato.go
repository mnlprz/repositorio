package services

import (
	"database/sql"
	"errors"
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
		return contrato, errors.New("no se encontro id")
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
	_, err := http.PostForm(URL, data)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

func PutContrato(contrato models.Contrato) error {

	const query = "UPDATE contratos SET NUP = $1, COD_ENTIDAD = $2, COD_CENTRO = $3, NUM_CONTRATO = $4, COD_PROD = $5, COD_SUBPROD = $6, MONEDA = $7, SALDO = $8 WHERE ID = $9"

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r, err := db.Exec(query, contrato.Nup, contrato.CodEntidad, contrato.CodCentro, contrato.NumContrato, contrato.CodProd, contrato.CodSubProd, contrato.Moneda, contrato.Saldo, contrato.Id)
	if err != nil {
		log.Fatal(err)
	}
	n, _ := r.RowsAffected()
	if n == 0 {
		return errors.New("no se encontro el id a modificar")
	}
	return nil
}

func DeleteContrato(id int64) error {

	const query = "DELETE FROM contratos WHERE ID = $1"

	db, err := database.GetConnection()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r, err := db.Exec(query, id)
	if err != nil {
		log.Fatal(err)
	}
	n, _ := r.RowsAffected()
	if n == 0 {
		return errors.New("no se encontro el id a eliminar")
	}
	return nil
}
