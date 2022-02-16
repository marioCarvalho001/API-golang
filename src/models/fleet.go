package models

import (
	"errors"
)

// Estrutura de chamado puxar a tabela de frota
type Fleet struct {
	ID        uint64  `json:"id"`
	Name      string  `json:"name"`
	Max_speed float64 `json:"max_speed"`
}

//Chama os valores passado via post os metodos e valida
func (fleet *Fleet) Preparar() error {
	if erro := fleet.validar(); erro != nil {

		return erro

	}
	return nil
}
func (fleet *Fleet) validar() error {
	if fleet.Name == "" {
		return errors.New("Nome obrigatorio")

	}
	if fleet.Max_speed == 0 {
		return errors.New("Max_vm é obrigatorio")
	}
	return nil
}
