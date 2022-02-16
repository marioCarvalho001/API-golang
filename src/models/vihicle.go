package models

import (
	"errors"
)

// Estrutura de chamado puxar a tabela de frota
type Vehicle struct {
	ID        uint64  `json:"id"`
	Fleet_id  uint64  `json:"fleet_id"`
	Name      string  `json:"name"`
	Max_speed float64 `json:"max_speed"`
}

//Chamar os metodos e valida
func (vehicle *Vehicle) Preparar() error {
	if erro := vehicle.validar(); erro != nil {
		return erro
	}
	return nil
}
func (vehicle *Vehicle) validar() error {
	if vehicle.Max_speed == 0 {
		return errors.New("Max_vm é obrigatorio")
	}
	return nil
}
