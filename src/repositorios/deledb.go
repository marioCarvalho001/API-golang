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

func (delet delete) Deletar() error {
	statement, erro := delet.db.Prepare("delete from alert ")
	if erro != nil {
		return erro
	}
	defer statement.Close()
	if _, erro = statement.Exec(); erro != nil {
		return erro
	}
	return nil
}
