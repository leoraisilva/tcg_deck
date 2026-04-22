package helper

import (
	"database/sql"
	"fmt"
	"os"

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
	err = ConnMigration(db)
	if err != nil {
		return nil, err
	}

	fmt.Printf("Conexão estabelecida com Sucesso")
	return db, nil
}

func ConnMigration(db *sql.DB) error {
	sqlByte, err := os.ReadFile("../resource/migration.sql")
	if err != nil {
		fmt.Printf("Erro ao adicionar migration: %v\n", err)
		return err
	}

	_, err = db.Exec(string(sqlByte))
	if err != nil {
		fmt.Printf("Erro ao criar migration: %v\n", err)
		return err
	}
	fmt.Printf("Migration adicionado com Sucesso!!\n")
	return nil
}
