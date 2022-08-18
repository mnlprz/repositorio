package models

import (
	"errors"
	"time"
)

type Oferta struct {
	NumPersona string
	Oferta     string
	Nombre     string
	Apellido   string
	SitIva     string
	Digital    string
	Fecha      string
	Created_at time.Time
	Updated_at time.Time
}

func (oferta *Oferta) Validate() error {

	if oferta.NumPersona == "" {
		return errors.New("NUM_PERSONA invalido")
	}
	if oferta.Oferta == "" {
		return errors.New("OFERTA invalido")
	}
	if oferta.Nombre == "" {
		return errors.New("NOMBRE invalido")
	}
	if oferta.Apellido == "" {
		return errors.New("APELLIDO invalido")
	}
	if oferta.SitIva == "" {
		return errors.New("SIT_IVA invalido")
	}
	if oferta.Digital == "" {
		return errors.New("DIGITAL invalido")
	}
	if oferta.Fecha == "" {
		return errors.New("FECHA invalido")
	}
	return nil
}
