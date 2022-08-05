package models

import "time"

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
