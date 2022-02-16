package repositorios

import "database/sql"

//Representa um repositorio
type delete struct {
	db *sql.DB
}

//Criar um novo repositorio de usuarios
func NovoRepotorioReset(db *sql.DB) *delete {
	return &delete{db}
}

func (d *delete) Deletar() error {
	statement, erro := d.db.Prepare("DROP DATABASE apidb")
	if erro != nil {
		return erro
	}
	defer statement.Close()

	if _, erro = statement.Exec(); erro != nil {
		return erro
	}
	return nil
}
