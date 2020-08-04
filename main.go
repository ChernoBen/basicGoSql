package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	//criando conexão
	db, err := sql.Open("mysql", "root:Sapiencia@10@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Println("conectado com sucesso !")
	}
	//criando base de dados
	_, err = db.Exec("CREATE DATABASE testDB")
	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Println("OK")
	}
	//usando database testDB
	_, err = db.Exec("USE testDB")
	if err != nil {
		fmt.Println(err.Error())

	} else {
		fmt.Println("testDB selecionado com sucesso!")

	}

	//create table
	stmt, err := db.Prepare("CREATE TABLE employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50),last_name varchar(30),PRIMARY KEY (id));")
	if err != nil {
		fmt.Println(err.Error())
	}

	//executando comando
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("tabela criada com sucesso!")
	}
	defer db.Close()
}
