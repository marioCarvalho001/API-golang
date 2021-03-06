package models

import (
	"errors"
)

// Estrutura de chamado puxar a tabela de frota

type Position struct {
	ID               uint64  `json:"id"`
	Veiculo_id       string  `json:"vehicle_id"`
	Data             int     `json:"times_stamp"`
	Latitude         float64 `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	Velocidade_atual int     `json:"current_speed"`
	Max_speed        int     `json:"max_speed"`
}

//Chamar os metodos e valida
func (position *Position) Preparar() error {
	if erro := position.validar(); erro != nil {
		return erro
	}
	return nil
}
func (position *Position) validar() error {
	if position.Data == 0 {
		position.Data = 00000000
	}
	if position.Latitude == 0 {
		return errors.New("Latitude invalida")
	}
	if position.Longitude == 0 {
		return errors.New("Latitude invalida")
	}

	if position.Max_speed == 0 {
		return errors.New("Max_speed invalida")
	}
	return nil
}
