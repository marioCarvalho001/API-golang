package repositorios

import (
	"api/src/models"
	"database/sql"
)

//Representa um repositorio
type fleet struct {
	db *sql.DB
}

//Criar um novo repositorio de valores
func NovoRepositorioFleet(db *sql.DB) *fleet {
	return &fleet{db}
}

//inseri uma nova frota no banco de dados
func (i fleet) CriarFleets(fleet models.Fleet) (uint64, error) {

	statement, erro := i.db.Prepare("insert into fleet (name,max_speed) values (?,?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(fleet.Name, fleet.Max_speed)
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

//busca uma nova frota no banco de dados
func (i fleet) BuscarFleet() ([]models.Fleet, error) {
	statement, erro := i.db.Query("select * from fleet")
	if erro != nil {
		return nil, erro
	}
	defer statement.Close()

	var fleets []models.Fleet

	for statement.Next() {
		var fleet models.Fleet

		if erro := statement.Scan(&fleet.ID, &fleet.Name, &fleet.Max_speed); erro != nil {
			return nil, erro
		}

		fleets = append(fleets, fleet)
	}
	return fleets, nil
}
