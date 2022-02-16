package banco

import (
	"api/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//Abre conexao com o bd e retorna
func Conectar() (*sql.DB, error) {
	db, erro := sql.Open("mysql", config.StringConecaoBanco)
	if erro != nil {
		return nil, erro
	}

	if erro = db.Ping(); erro != nil {
		db.Close()
		return nil, erro
	}
	return db, nil
}
