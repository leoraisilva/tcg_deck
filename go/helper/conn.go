package helper

import (
	"database/sql"
	"fmt"
)

const (
	host   = "localhost"
	port   = "5432"
	user   = "postgres"
	pass   = "postgres"
	dbname = "tcg_db"
)

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s, port=%s, user=%s, pass=%s, dbname=%s",
		host, port, user, pass, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Printf("Erro ao Conectar ao Banco: %v\n", err)
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		fmt.Printf("Erro ao Ping no Banco: %v\n", err)
		return nil, err
	}

	fmt.Printf("Conexão estabelecida com Sucesso")
	return db, nil
}
