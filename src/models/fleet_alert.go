package models

import (
	"errors"
)

// Estrutura de chamado puxar a tabela de frota
type Fleet_alert struct {
	ID       uint64 `json:"id"`
	Fleet_id uint64 `json:"fleet_id"`
	Webhook  string `json:"webhook"`
}

//Chamar os metodos e valida
func (fleet_alert *Fleet_alert) Preparar() error {
	if erro := fleet_alert.validar(); erro != nil {
		return erro
	}
	return nil
}
func (fleet_alert *Fleet_alert) validar() error {
	if fleet_alert.Webhook == "" {
		return errors.New("Nome obrigatorio")
		//validar a url no get
	}
	return nil
}
