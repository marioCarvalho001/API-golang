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

//Chamar os campos e valida
func (vehicle *Vehicle) Preparar() error {
	if erro := vehicle.validar(); erro != nil {
		return erro
	}
	return nil
}
func (vehicle *Vehicle) validar() error {
	if vehicle.Fleet_id == 0 {
		return errors.New("Fleet_id name invalido")
	}
	if vehicle.Name == "" {
		return errors.New("campo name invalido")
	}
	return nil
}
