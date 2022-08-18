package models

import (
	"errors"
	"time"
)

type Contrato struct {
	Id          int64
	Nup         string
	CodEntidad  string
	CodCentro   string
	NumContrato string
	CodProd     string
	CodSubProd  string
	Moneda      string
	Saldo       string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

const PESOS, DOLARES = "ARS", "USD"

func (contrato *Contrato) Validate() error {

	if contrato.Nup == "" {
		return errors.New("Campo NUP invalido")
	}

	if contrato.CodEntidad == "" {
		return errors.New("Campo COD_ENTIDAD invalido")
	}

	if contrato.CodCentro == "" {
		return errors.New("Campo COD_CENTRO invalido")
	}

	if contrato.NumContrato == "" {
		return errors.New("Campo NUM_CONTRATO invalido")
	}

	if contrato.CodProd == "" {
		return errors.New("Campo COD_PROD invalido")
	}

	if contrato.CodSubProd == "" {
		return errors.New("Campo COD_SUBPROD invalido")
	}

	if contrato.Moneda != PESOS && contrato.Moneda != DOLARES {
		return errors.New("Campo MONEDA invalido")
	}

	if contrato.Saldo == "" {
		return errors.New("Campo SALDO invalido")
	}

	return nil
}
