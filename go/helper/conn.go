package helper

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "tcg_db"
)

func GetConnection() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
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
