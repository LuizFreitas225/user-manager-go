package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type PostgresDatabase struct {
}

func (pd *PostgresDatabase) Open() (*sql.DB, error) {
	dataSourceName, driver := pd.getDataSource()
	db, err := sql.Open(driver, dataSourceName)

	if err != nil {
		return &sql.DB{}, err
	}

	return db, nil
}

func (db *PostgresDatabase) getDataSource() (string, string) {
	// Carregar .env
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: .env não encontrado, usando variáveis de ambiente existentes")
	}

	driver := os.Getenv("POSTGRES_DRIVER")
	user := os.Getenv("POSTGRES_USER")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")
	dbname := os.Getenv("POSTGRES_DB")

	// Montar connection string
	dataSourceName := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	return dataSourceName, driver
}
