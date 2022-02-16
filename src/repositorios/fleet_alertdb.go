package repositorios

import (
	"api/src/models"
	"database/sql"
)

//Representa um repositorio
type fleet_alert struct {
	db *sql.DB
}

//Criar um novo repositorio de usuarios
func NovoRepositorioFleet_alert(db *sql.DB) *fleet_alert {
	return &fleet_alert{db}
}

//Buscar
func (i fleet_alert) BuscarFleet_alert(ID uint64) ([]models.Fleet_alert, error) {

	statement, erro := i.db.Query("select * from fleet_alert where fleet_id =?", ID)
	if erro != nil {
		return nil, erro
	}
	defer statement.Close()

	var fleet_alerts []models.Fleet_alert

	for statement.Next() {
		var fleet_alert models.Fleet_alert

		if erro := statement.Scan(&fleet_alert.ID, &fleet_alert.Fleet_id, &fleet_alert.Webhook); erro != nil {
			return nil, erro
		}

		fleet_alerts = append(fleet_alerts, fleet_alert)
	}
	return fleet_alerts, nil
}

var n uint64

func Urlcd(rsc uint64) {
	n = rsc
}

func (i fleet_alert) CriarFleet_alert(fleet_alert models.Fleet_alert) (uint64, error) {
	statement, erro := i.db.Prepare("insert into fleet_alert (webhook,fleet_id) values (?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(fleet_alert.Webhook, n)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}

	//Retorna o
	return uint64(ultimoIDInserido), nil
}
