package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	var strConnection = "root:123456@/devbook?charset=utf8&parseTime=True&loc=Local"
	db, err := sql.Open("mysql", strConnection)

	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Conexao aberta")

	linhas, err := db.Query("select * from usuarios")

	if err != nil {
		log.Fatal(err)

	}

	defer linhas.Close()

	fmt.Println(linhas)
}
