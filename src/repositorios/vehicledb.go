package repositorios

import (
	"api/src/models"
	"database/sql"
)

//Representa um repositorio
type vehicle struct {
	db *sql.DB
}

//Criar um novo repositorio de usuarios
func NovoRepositorioVehicle(db *sql.DB) *vehicle {
	return &vehicle{db}
}

//inseri uma nova frota no banco de dados

func (i vehicle) BuscarVehicle() ([]models.Vehicle, error) {
	statement, erro := i.db.Query("select * from vehicle")
	if erro != nil {
		return nil, erro
	}
	defer statement.Close()
	var vehicles []models.Vehicle

	for statement.Next() {
		var vehicle models.Vehicle

		if erro := statement.Scan(&vehicle.ID, &vehicle.Fleet_id, &vehicle.Name, &vehicle.Max_speed); erro != nil {
			return nil, erro
		}

		vehicles = append(vehicles, vehicle)
	}
	return vehicles, nil
}

func (i vehicle) CriarVehicle(vehicle models.Vehicle) (uint64, error) {

	statement, erro := i.db.Prepare("insert into vehicle (name, fleet_id,max_speed) values (?,?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(vehicle.Name, vehicle.Fleet_id, vehicle.Max_speed)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return uint64(ultimoIDInserido), nil

}
