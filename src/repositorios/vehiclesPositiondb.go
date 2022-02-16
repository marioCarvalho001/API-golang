package repositorios

import (
	"api/src/models"
	"database/sql"
)

type vposition struct {
	db *sql.DB
}

func NovoRepositorioVPosition(db *sql.DB) *vposition {
	return &vposition{db}
}

func (i vposition) BuscarPosition(ID uint64) ([]models.Position, error) {

	statement, erro := i.db.Query("select * from vehicle_position where vehicle_id =?", ID)
	if erro != nil {
		return nil, erro
	}
	defer statement.Close()
	var positions []models.Position

	for statement.Next() {
		var position models.Position

		if erro := statement.Scan(&position.ID, &position.Veiculo_id, &position.Data, &position.Latitude, &position.Longitude, &position.Velocidade_atual, &position.Max_speed); erro != nil {
			return nil, erro
		}

		positions = append(positions, position)
	}
	return positions, nil
}

var idv uint64

func UrlID(i uint64) {
	idv = i
}

func (i vposition) CriarPositionAtual(position models.Position) (uint64, error) {
	statement, erro := i.db.Prepare("insert into vehicle_position ( times_stamp, latitude, longitude, current_speed, max_speed, vehicle_id) values (?,?,?,?,?,?) ")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()
	resultado, erro := statement.Exec(position.Data, position.Latitude, position.Longitude, position.Velocidade_atual, position.Max_speed, idv)
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
