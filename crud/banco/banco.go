package banco

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql" //Driver de conexao com mysql
)

// Abre Conexao
func Conectar() (*sql.DB, error) {
	var strConnection = "root:123456@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", strConnection)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}
	return db, nil

}
